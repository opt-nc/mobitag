package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/charmbracelet/fang"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "mobitag",
	Short: "Envoyer des sms avec Mobitag",
	Long:  `CLI permettant d'envoyer des sms gratuits en contactant l'API de Mobitag. Afin de vérifier la configuration, utilisez la commande 'mobitag check-apikey'.`,
	Example: `mobitag send --to <destinataire> --message <message>
whoami | mobitag pipe --to 123456`,

	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Help for this command")
	if err := rootCmd.PersistentFlags().MarkHidden("help"); err != nil {
		// Option 1: log or handle gracefully
		fmt.Fprintf(os.Stderr, "failed to hide help flag: %v\n", err)
		os.Exit(1)
	}

	if err := fang.Execute(context.Background(), rootCmd, fang.WithoutVersion()); err != nil {
		os.Exit(1)
	}

}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.mobitag-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
