package index

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/sdk"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/cli/pkg/utils"
)

func validateIndex(index string) (*sdk.Client, error) {
	client := session.Connect(config.Token())
	indicesResponse, err := client.GetIndices()
	if err != nil {
		return nil, err
	}

	// Create a map of indices to compare against
	var indexNames []string
	available := make(map[string]bool)
	for _, idx := range indicesResponse.GetData() {
		available[idx.Name] = true
		indexNames = append(indexNames, idx.Name)
	}

	// If the index is not present in the map, print output and suggest
	// options for what the argument may have intended
	if !available[index] {
		var msg strings.Builder
		fmt.Fprintf(&msg, "index '%s' does not exist", index)
		if suggestions := utils.SuggestFor(index, indexNames); len(suggestions) > 0 {
			msg.WriteString("\n\nDid you mean this?\n")
			for _, s := range suggestions {
				fmt.Fprintf(&msg, "\t%s\n", s)
			}
		}
		return nil, fmt.Errorf("%s", msg.String())
	}
	return client, nil
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

			client, err := validateIndex(args[0])
			if err != nil {
				return err
			}
			response, err := client.GetIndex(args[0], queryParameters)
			if err != nil {
				return err
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

			client, err := validateIndex(args[0])
			if err != nil {
				return err
			}
			response, err := client.GetIndex(args[0], queryParameters)
			if err != nil {
				return err
			}

			var viewportOutput interface{}
			viewportOutput = response.GetData()
			if opts.Full {
				viewportOutput = response
			}
			ui.Viewport(args[0], viewportOutput)

			return nil
		},
	}

	cmdList.Flags().BoolVarP(&opts.Full, "full", "f", false, i18n.C.IndexFlagFullResponse)
	cmdBrowse.Flags().BoolVarP(&opts.Full, "full", "f", false, i18n.C.IndexFlagFullResponse)

	cmd.AddCommand(cmdList)
	cmd.AddCommand(cmdBrowse)

	return cmd
}
