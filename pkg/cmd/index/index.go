package index

import (
	"context"
	"fmt"
	"reflect"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/internal/output"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

func validateIndex(ctx context.Context, index string, interactive bool) (string, error) {
	indicesResponse, err := session.ConnectWithContext(ctx, config.Token()).GetIndices()
	if err != nil {
		return "", err
	}

	var indexNames []string
	available := make(map[string]bool)
	for _, idx := range indicesResponse.GetData() {
		available[idx.Name] = true
		indexNames = append(indexNames, idx.Name)
	}

	if available[index] {
		return index, nil
	}

	suggestions := utils.SuggestFor(index, indexNames)
	if len(suggestions) == 0 {
		return "", fmt.Errorf("index '%s' does not exist", index)
	}

	// Non-interactive callers (--json, --no-interactive, CI, headless)
	// can't see a prompt — return the original error with the suggestions
	// surfaced in the message so the user can pick one and retry.
	if !interactive {
		return "", fmt.Errorf("index '%s' does not exist; did you mean: %s", index, joinSuggestions(suggestions))
	}

	options := make([]huh.Option[string], len(suggestions))
	for i, s := range suggestions {
		options[i] = huh.NewOption(s, s)
	}

	var selected string
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title(fmt.Sprintf("index '%s' does not exist. Did you mean one of these?", index)).
				Options(options...).
				Value(&selected),
		),
	)
	if err := form.Run(); err != nil {
		return "", fmt.Errorf("index '%s' does not exist", index)
	}

	return selected, nil
}

func joinSuggestions(s []string) string {
	if len(s) == 0 {
		return ""
	}
	out := s[0]
	for _, x := range s[1:] {
		out += ", " + x
	}
	return out
}

type Options struct {
	Full bool
}

func Command() *cobra.Command {
	opts := &Options{}

	cmd := &cobra.Command{
		Use:   "index <command>",
		Short: i18n.C.IndexShort,
	}

	keys := reflect.TypeOf(sdk.IndexQueryParameters{})

	for i := 0; i < keys.NumField(); i++ {
		flag := keys.Field(i).Tag.Get("json")
		name := keys.Field(i).Name
		cmd.PersistentFlags().String(flag, "", name)
	}

	cmdList := &cobra.Command{
		Use:   "list <index>",
		Short: i18n.C.IndexListShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) != 1 {
				return ui.Error(i18n.C.IndexErrorRequired)
			}

			queryParameters := sdk.IndexQueryParameters{}
			for i := 0; i < keys.NumField(); i++ {
				flag := keys.Field(i).Tag.Get("json")
				if cmd.Flag(flag).Value.String() != "" {
					field := reflect.ValueOf(&queryParameters).Elem().Field(i)
					switch field.Kind() {
					case reflect.String:
						field.SetString(cmd.Flag(flag).Value.String())
					case reflect.Int:
						intValue, err := strconv.Atoi(cmd.Flag(flag).Value.String())
						if err != nil {
							r.Warn("invalid value for --%s: %v", flag, err)
							continue
						}
						field.SetInt(int64(intValue))
					}
				}
			}

			index := args[0]
			client := session.ConnectWithContext(cmd.Context(), config.Token())
			response, err := client.GetIndex(index, queryParameters)

			if err != nil {
				if _, ok := err.(sdk.ReqError); ok {
					corrected, validationErr := validateIndex(cmd.Context(), index, r.Interactive())
					if validationErr != nil {
						return validationErr
					}
					response, err = client.GetIndex(corrected, queryParameters)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}

			var payload interface{} = response.GetData()
			if opts.Full {
				payload = response
			}

			return r.JSON(payload)
		},
	}

	cmdBrowse := &cobra.Command{
		Use:   "browse <index>",
		Short: i18n.C.IndexBrowseShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			r := output.FromCmd(cmd)
			if len(args) != 1 {
				return ui.Error(i18n.C.IndexErrorRequired)
			}

			queryParameters := sdk.IndexQueryParameters{}
			for i := 0; i < keys.NumField(); i++ {
				flag := keys.Field(i).Tag.Get("json")
				if cmd.Flag(flag).Value.String() != "" {
					field := reflect.ValueOf(&queryParameters).Elem().Field(i)
					switch field.Kind() {
					case reflect.String:
						field.SetString(cmd.Flag(flag).Value.String())
					case reflect.Int:
						intValue, err := strconv.Atoi(cmd.Flag(flag).Value.String())
						if err != nil {
							r.Warn("invalid value for --%s: %v", flag, err)
							continue
						}
						field.SetInt(int64(intValue))
					}
				}
			}

			index := args[0]
			client := session.ConnectWithContext(cmd.Context(), config.Token())
			response, err := client.GetIndex(index, queryParameters)

			if err != nil {
				if _, ok := err.(sdk.ReqError); ok {
					corrected, validationErr := validateIndex(cmd.Context(), index, r.Interactive())
					if validationErr != nil {
						return validationErr
					}
					index = corrected
					response, err = client.GetIndex(index, queryParameters)
					if err != nil {
						return err
					}
				} else {
					return err
				}
			}

			var viewportOutput interface{} = response.GetData()
			if opts.Full {
				viewportOutput = response
			}

			// `browse` is interactive — fall back to JSON output whenever the
			// renderer can't show TUI (--json, --no-interactive, non-TTY, CI).
			if r.IsJSON() || !r.Interactive() {
				return r.JSON(viewportOutput)
			}

			ui.Viewport(index, viewportOutput)
			return nil
		},
	}

	cmdList.Flags().BoolVarP(&opts.Full, "full", "f", false, i18n.C.IndexFlagFullResponse)
	cmdBrowse.Flags().BoolVarP(&opts.Full, "full", "f", false, i18n.C.IndexFlagFullResponse)

	cmd.AddCommand(cmdList)
	cmd.AddCommand(cmdBrowse)

	return cmd
}
