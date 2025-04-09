package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"log/slog"

	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:     "send",
	Aliases: []string{"s"}, // L'alias pour la commande
	Short:   "Envoyer un Mobitag",
	Long:    `Envoi d'un Mobitag à un numéro de téléphone.`,
	Example: `mobitag send --to <destinataire> --message <message> --from <expéditeur>
mobitag send --to 123456 --message "Hello, world!"
mobitag send -t 123456 -m "Hello, world!" -f 654321`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("OPTNC_MOBITAGNC_API_KEY") == "" {
			slog.Error("La clé API 'OPTNC_MOBITAGNC_API_KEY' n'est pas définie dans les variables d'environnement. Veuillez définir cette clé ou utiliser la commande 'mobitag web' en attendant d'avoir une clé.")
			os.Exit(1)
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		to, _ := cmd.Flags().GetString("to")
		message, _ := cmd.Flags().GetString("message")
		from, _ := cmd.Flags().GetString("from")
		cut, _ := cmd.Flags().GetBool("cut")
		verbose, _ := cmd.Flags().GetBool("verbose")

		SendSMS(to, message, from, cut, verbose)
	},
}

// sendSMS sends an SMS to the specified receiver mobile number
// receiverMobile: the mobile number of the receiver, like 654321
// message: the message to send
func SendSMS(receiverMobile string, message string, senderMobile string, cut bool, verbose bool) {
	// Replace all newline characters with spaces
	message = strings.ReplaceAll(message, "\n", " ")

	// Check if message exceeds 160 characters
	if len(message) > 160 {
		if !cut {
			slog.Error("Le message dépasse la limite de 160 caractères length=" + fmt.Sprint(len(message)))
			os.Exit(1) // Exit the program
		}
		slog.Warn("Le message dépasse la limite de 160 caractères et sera coupé length=" + fmt.Sprint(len(message)))
		message = message[:155] + "[...]"
	}

	// Get the Mobitag API key from the environment
	mobitagAPIKey := os.Getenv("OPTNC_MOBITAGNC_API_KEY")
	apiURL := "https://api.opt.nc/mobitag/sendSms"

	client := &http.Client{}
	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		slog.Error("Une erreur s'est produite lors de la création de la requête erreur=" + err.Error())
	}

	// log all parameters
	if verbose {
		if senderMobile != "" {
			slog.Info("Expéditeur=" + senderMobile)
		}
		slog.Info("Destinataire=" + receiverMobile)
	}

	slog.Info("Message envoyé=" + message)

	// set request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-apikey", mobitagAPIKey)

	// set request payload with receiver mobile number, message, and optionally sender mobile number
	var reqBody strings.Builder
	reqBody.WriteString(`{"to":"`)
	reqBody.WriteString(receiverMobile)
	reqBody.WriteString(`","message":"`)
	reqBody.WriteString(message)
	if senderMobile != "" {
		reqBody.WriteString(`","from":"`)
		reqBody.WriteString(senderMobile)
	}
	reqBody.WriteString(`"}`)
	req.Body = io.NopCloser(strings.NewReader(reqBody.String()))

	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Une erreur s'est produite lors de l'envoi de la requête erreur=" + err.Error())
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("Une erreur s'est produite lors de la fermeture du corps de la réponse erreur=" + err.Error())
		}
	}()

	if resp.StatusCode == 443 {
		slog.Error("La clé API est invalide. Veuillez demander une nouvelle clé ou utiliser la commande 'mobitag web' en attendant.")
		os.Exit(1)
	}

	slog.Info("Accusé réception=" + resp.Status)
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("to", "t", "", "Numéro de téléphone du destinataire")
	sendCmd.Flags().StringP("message", "m", "", "Message à envoyer")
	sendCmd.Flags().StringP("from", "f", "", "Numéro de téléphone de l'expéditeur")
	err := sendCmd.MarkFlagRequired("to")
	if err != nil {
		slog.Error("Erreur lors du marquage du flag 'to' comme requis error=" + err.Error())
		os.Exit(1)
	}

	err = sendCmd.MarkFlagRequired("message")
	if err != nil {
		slog.Error("Erreur lors du marquage du flag 'message' comme requis error=" + err.Error())
		os.Exit(1)
	}

	sendCmd.Flags().BoolP("cut", "c", false, "Couper le message à 160 caractères afin de ne pas excéder la limite")
	sendCmd.Flags().BoolP("verbose", "v", false, "Afficher les détails de l'envoi")

}
