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
	"github.com/vulncheck-oss/cli/pkg/utils"
	"github.com/vulncheck-oss/sdk"
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
		flag := utils.NormalizeString(keys.Field(i).Name)
		cmd.PersistentFlags().String(flag, "", keys.Field(i).Name)
	}

	cmdList := &cobra.Command{
		Use:   "list <index>",
		Short: i18n.C.IndexListShort,
		RunE: func(cmd *cobra.Command, args []string) error {
			response, err := getIndexResponse(cmd, args)
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
			response, err := getIndexResponse(cmd, args)
			if err != nil {
				return err
			}

			loadPageFunc := func(index string, page int) (*sdk.IndexResponse, error) {
				queryParams := buildQueryParameters(cmd)
				queryParams.Page = page
				return session.Connect(config.Token()).GetIndex(index, queryParams)
			}

			ui.ViewportPaginated(args[0], response.GetData(), response.Meta, loadPageFunc)
			return nil
		},
	}

	cmd.AddCommand(cmdList)
	cmd.AddCommand(cmdBrowse)

	return cmd
}

func getIndexResponse(cmd *cobra.Command, args []string) (*sdk.IndexResponse, error) {
	if len(args) != 1 {
		return nil, ui.Error(i18n.C.IndexErrorRequired)
	}

	queryParameters := buildQueryParameters(cmd)

	return session.Connect(config.Token()).GetIndex(args[0], queryParameters)
}

func buildQueryParameters(cmd *cobra.Command) sdk.IndexQueryParameters {
	queryParameters := sdk.IndexQueryParameters{}
	keys := reflect.TypeOf(sdk.IndexQueryParameters{})

	for i := 0; i < keys.NumField(); i++ {
		flag := utils.NormalizeString(keys.Field(i).Name)
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

	return queryParameters
}
