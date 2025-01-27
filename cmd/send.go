/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:     "send",
	Aliases: []string{"s"}, // L'alias pour la commande
	Short:   "Envoyer un Mobitag",
	Long:    `Envoi d'un Mobitag à un numéro de téléphone.`,
	//Usage:   "mobitag-cli send --to XXXXXX --message 'Hello world' --from XXXXXX",
	Run: func(cmd *cobra.Command, args []string) {
		to, _ := cmd.Flags().GetString("to")
		message, _ := cmd.Flags().GetString("message")
		from, _ := cmd.Flags().GetString("from")
		sendSMS(to, message, from)
	},
}

// sendSMS sends an SMS to the specified receiver mobile number
// receiverMobile: the mobile number of the receiver, like 654321
// message: the message to send
func sendSMS(receiverMobile string, message string, senderMobile string) {
	// Get the Mobitag API key from the environment
	mobitagAPIKey := os.Getenv("OPTNC_MOBITAGNC_API_KEY")
	apiURL := "https://api.opt.nc/mobitag/sendSms"

	messageToSend := message + "\nFrom: " + senderMobile

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		fmt.Printf("An error occurred while creating the request: %v\n", err)
		return
	}

	// set request headers
	req.Header.Set("Content-Type", "application/json")
	// Authenticate the request
	req.Header.Set("x-apikey", mobitagAPIKey)

	// set request paylod with receiver mobile number and message
	reqBody := fmt.Sprintf(`{"to":"%s","message":"%s"}`, receiverMobile, messageToSend)
	req.Body = io.NopCloser(strings.NewReader(reqBody))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("❗An error occurred while sending the request: %v\n", err)
		return
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
	sendCmd.MarkFlagRequired("to")
	sendCmd.MarkFlagRequired("message")
}
