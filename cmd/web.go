package cmd

import (
	"log/slog"
	"os"

	"github.com/pkg/browser"
	"github.com/spf13/cobra"
)

// webCmd représente la commande web
var webCmd = &cobra.Command{
	Use:   "web",
	Short: "Ouvre le navigateur pour afficher la version web de Mobitag",
	Long: `Cette commande lance le navigateur par défaut de l'utilisateur 
et affiche la version web de Mobitag accessible à l'adresse suivante : 
http://www.mobitag.nc.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := "http://www.mobitag.nc"
		err := browser.OpenURL(url)
		if err != nil {
			slog.Error("Échec de l'ouverture du navigateur", "erreur", err)
			os.Exit(1)
		}
	},
}

func init() {
	// Ajoute la commande web à la commande racine
	rootCmd.AddCommand(webCmd)
}
