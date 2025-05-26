package cmd

import (
	"os"

	"log/slog"

	"github.com/spf13/cobra"
)

// checkApiKeyCmd represents the check-apikey command
var checkApiKeyCmd = &cobra.Command{
	Use:     "check-apikey",
	Aliases: []string{"cap"}, // Alias éventuel, modifiable
	Short:   "Vérifie la présence de la clé API",
	Long:    `Vérifie si la variable d'environnement OPTNC_MOBITAGNC_API_KEY est définie (non vide).`,

	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPTNC_MOBITAGNC_API_KEY")
		if apiKey != "" {
			slog.Info("Clé API présente : OPTNC_MOBITAGNC_API_KEY est définie")
		} else {
			slog.Error("Clé API absente : OPTNC_MOBITAGNC_API_KEY n'est pas définie")
		}
	},
}

func init() {
	rootCmd.AddCommand(checkApiKeyCmd)
}
