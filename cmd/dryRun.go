package cmd

import (
	"os"

	"log/slog"

	"github.com/spf13/cobra"
)

// dryRunCmd represents the dryRun command
var dryRunCmd = &cobra.Command{
	Use:     "dryRun",
	Aliases: []string{"dr"}, // L'alias pour la commande
	Short:   "Vérification de la configuration",
	Long:    `Vérification de la présence de la clé API: OPTNC_MOBITAGNC_API_KEY`,

	Run: func(cmd *cobra.Command, args []string) {
		var mobitagAPIKey = os.Getenv("OPTNC_MOBITAGNC_API_KEY")
		if mobitagAPIKey != "" {
			slog.Info("Dry run test passed: OPTNC_MOBITAGNC_API_KEY environment variable is set")
		} else {
			slog.Error("Dry run test failed: OPTNC_MOBITAGNC_API_KEY environment variable is not set")
		}
	},
}

func init() {
	rootCmd.AddCommand(dryRunCmd)
}
