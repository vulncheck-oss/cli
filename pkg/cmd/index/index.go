package index

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

// validateIndex checks whether index exists. If it does not but close matches
// are found, an interactive select is presented so the user can pick one.
// Returns the confirmed index name, or an error if the name is unrecognised.
func validateIndex(index string) (string, error) {
	indicesResponse, err := session.Connect(config.Token()).GetIndices()
	if err != nil {
		return "", err
	}

	// Create a map of indices to compare against
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

	// If the index is not present in the map but close matches exist, present
	// an interactive select so the user can choose the intended index
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

type Options struct {
	Full bool
}

func Command() *cobra.Command {

	opts := &Options{
		Full: false,
	}

	cmd := &cobra.Command{
		Use:   "index <command>",
		Short: i18n.C.IndexShort,
	}

	// Define flags for index commands
	keys := reflect.TypeOf(sdk.IndexQueryParameters{})

	// Dynamically add flags for index commands (list and browse)
	for i := 0; i < keys.NumField(); i++ {
		flag := keys.Field(i).Tag.Get("json") // Get the json tag value which is the correct API field
		name := keys.Field(i).Name
		cmd.PersistentFlags().String(flag, "", name)
	}

	cmdList := &cobra.Command{
		Use:   "list <index>",
		Short: i18n.C.IndexListShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.IndexErrorRequired)
			}

			// Create a new IndexQueryParameters struct and set the values from the flags
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
							fmt.Println(err)
							continue
						}
						field.SetInt(int64(intValue))
					}
				}
			}

			index := args[0]
			client := session.Connect(config.Token())
			response, err := client.GetIndex(index, queryParameters)

			// If GetIndex fails due to a HTTP request error fallback
			// and attempt to validate the index name argument provided if
			// there was a typo/spelling error it will suggest similar names
			if err != nil {
				if _, ok := err.(sdk.ReqError); ok {
					corrected, validationErr := validateIndex(index)
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

			var terminalOutput interface{}
			terminalOutput = response.GetData()
			if opts.Full {
				terminalOutput = response
			}

			ui.Json(terminalOutput)

			return nil
		},
	}

	cmdBrowse := &cobra.Command{
		Use:   "browse <index>",
		Short: i18n.C.IndexBrowseShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return ui.Error(i18n.C.IndexErrorRequired)
			}

			// Create a new IndexQueryParameters struct and set the values from the flags
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
							fmt.Println(err)
							continue
						}
						field.SetInt(int64(intValue))
					}
				}
			}

			index := args[0]
			client := session.Connect(config.Token())
			response, err := client.GetIndex(index, queryParameters)

			// If GetIndex fails due to a HTTP request error fallback
			// and attempt to validate the index name argument provided if
			// there was a typo/spelling error it will suggest similar names
			if err != nil {
				if _, ok := err.(sdk.ReqError); ok {
					corrected, validationErr := validateIndex(index)
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

			var viewportOutput interface{}
			viewportOutput = response.GetData()
			if opts.Full {
				viewportOutput = response
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
