package index

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vulncheck-oss/cli/pkg/config"
	"github.com/vulncheck-oss/cli/pkg/i18n"
	"github.com/vulncheck-oss/cli/pkg/session"
	"github.com/vulncheck-oss/cli/pkg/ui"
	"github.com/vulncheck-oss/sdk-go"
)

func Command() *cobra.Command {

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

			response, err := session.Connect(config.Token()).GetIndex(args[0], queryParameters)
			if err != nil {
				return err
			}
			ui.Json(response.GetData())
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

			response, err := session.Connect(config.Token()).GetIndex(args[0], queryParameters)
			if err != nil {
				return err
			}
			ui.Viewport(args[0], response.GetData())
			return nil
		},
	}

	cmd.AddCommand(cmdList)
	cmd.AddCommand(cmdBrowse)

	return cmd
}
