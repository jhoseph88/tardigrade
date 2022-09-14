package cmd

import (
	"github.com/spf13/cobra"
)

var db, transformation string

// backfillCmd represents the backfill command
var backfillCmd = &cobra.Command{
	Use:   "backfill",
	Short: "Applies user-defined function to backfill or transform data in the content-engine database.",
	Long:  `tardigrade is intended for routine tasks like backfilling data for new columns added in a migration.`,
	Run: func(cmd *cobra.Command, args []string) {
		RunTransformation(db, transformation)
	},
}

func init() {
	rootCmd.AddCommand(backfillCmd)

	rootCmd.Flags().StringVarP(
		&transformation,
		"transformation",
		"t",
		"BackfillLogoFileType",
		"Function defining the transformation that should take place.",
	)

	rootCmd.Flags().StringVarP(
		&db,
		"database",
		"d",
		"projects/rsg-sawa-dev-branches/instances/sawa-dev-branches/databases/cameron-dev",
		"Database to invoke the transformation on",
	)
}
