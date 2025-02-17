package cmd

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:     "send --to <destinataire> --message <message>",
	Aliases: []string{"s"}, // L'alias pour la commande
	Short:   "Envoyer un Mobitag",
	Long:    `Envoi d'un Mobitag à un numéro de téléphone.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("OPTNC_MOBITAGNC_API_KEY") == "" {
			log.Fatalf("❗ La clé API 'OPTNC_MOBITAGNC_API_KEY' n'est pas définie dans les variables d'environnement.")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		to, _ := cmd.Flags().GetString("to")
		message, _ := cmd.Flags().GetString("message")
		from, _ := cmd.Flags().GetString("from")
		SendSMS(to, message, from)
	},
}

// sendSMS sends an SMS to the specified receiver mobile number
// receiverMobile: the mobile number of the receiver, like 654321
// message: the message to send
func SendSMS(receiverMobile string, message string, senderMobile string) {
	// Replace all newline characters with spaces
	message = strings.ReplaceAll(message, "\n", " ")

	// Get the Mobitag API key from the environment
	mobitagAPIKey := os.Getenv("OPTNC_MOBITAGNC_API_KEY")
	apiURL := "https://api.opt.nc/mobitag/sendSms"

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		log.Fatalf("An error occurred while creating the request: %v\n", err)
	}

	// log all parameters
	fmt.Printf("📞  Destinataire: %s\n", receiverMobile)
	fmt.Printf("📜  Message: %s", message)
	if senderMobile != "" {
		fmt.Printf("📞  Expéditeur: %s\n", senderMobile)
	}

	// set request headers
	req.Header.Set("Content-Type", "application/json")
	// Authenticate the request
	req.Header.Set("x-apikey", mobitagAPIKey)

	// set request payload with receiver mobile number, message, and optionally sender mobile number
	var reqBody string
	if senderMobile != "" {
		reqBody = fmt.Sprintf(`{"to":"%s","message":"%s","from":"%s"}`, receiverMobile, message, senderMobile)
	} else {
		reqBody = fmt.Sprintf(`{"to":"%s","message":"%s"}`, receiverMobile, message)
	}
	req.Body = io.NopCloser(strings.NewReader(reqBody))

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("❗An error occurred while sending the request: %v\n", err)
	}
	defer resp.Body.Close()

	fmt.Printf("ℹ️  Accusé reception: %v\n", resp.Status)
	fmt.Printf("📜  Code retour: %v\n", resp.StatusCode)
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("to", "t", "", "Numéro de téléphone du destinataire")
	sendCmd.Flags().StringP("message", "m", "", "Message à envoyer")
	sendCmd.Flags().StringP("from", "f", "", "Numéro de téléphone de l'expéditeur")
	err := sendCmd.MarkFlagRequired("to")
	if err != nil {
		log.Fatalf("Erreur lors du marquage du flag 'to' comme requis : %v", err)
	}

	err = sendCmd.MarkFlagRequired("message")
	if err != nil {
		log.Fatalf("Erreur lors du marquage du flag 'message' comme requis : %v", err)
	}

}
