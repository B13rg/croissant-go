// Test a croissant file
//
//nolint:gochecknoinits,gochecknoglobals
package cmd

import (
	"github.com/b13rg/croissant-go/pkg/croissant"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

// Test a croissant file.
var TestCmd = &cobra.Command{
	Use:   "test path",
	Short: "Test loading a Croissant file",
	Long:  `Test loading a Croissant file`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 || args[0] == "" {
			log.Logger.Fatal().Msg("invalid path param")
		}

		ds, err := croissant.NewDataSetFromPath(args[0])
		if err != nil {
			log.Logger.Fatal().AnErr("DatasetFromPath", err).Msg("error loading file")
		}

		log.Logger.Info().Str("name", ds.Name).Msg("Dataset name")
	},
}

func init() {
	RootCmd.AddCommand(TestCmd)
}
