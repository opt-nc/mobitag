/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// sendPipeCmd represents the sendPipe command
var sendPipeCmd = &cobra.Command{
	Use:     "echo \"message\" | sendPipe --to <destinataire>",
	Aliases: []string{"sp"}, // L'alias pour la commande
	Short:   "Envoyer un Mobitag depuis un pipe",
	Long:    `Envoi d'un Mobitag à un numéro de téléphone depuis un pipe.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("OPTNC_MOBITAGNC_API_KEY") == "" {
			log.Fatalf("❗ La clé API 'OPTNC_MOBITAGNC_API_KEY' n'est pas définie dans les variables d'environnement.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		to, _ := cmd.Flags().GetString("to")
		from, _ := cmd.Flags().GetString("from")

		var message string
		// Check if there is input from stdin
		if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
			reader := bufio.NewReader(os.Stdin)
			input, _ := io.ReadAll(reader)
			message = string(input)
		} else {
			log.Fatalf("❗ Aucune entrée n'a été trouvée. Veuillez utiliser un pipe pour envoyer un message.")
		}

		SendSMS(to, message, from)
	},
}

func init() {
	rootCmd.AddCommand(sendPipeCmd)

	sendPipeCmd.Flags().StringP("to", "t", "", "Numéro de téléphone du destinataire")
	sendPipeCmd.Flags().StringP("from", "f", "", "Numéro de téléphone de l'expéditeur")
	sendPipeCmd.MarkFlagRequired("to")
}
