package cmd

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strings"

	"log/slog"
)

// buildRequestBody construit le corps JSON pour l'envoi du SMS
func buildRequestBody(receiverMobile, encodedMessage, senderMobile string) string {
	var reqBody strings.Builder
	reqBody.WriteString(`{"to":"`)
	reqBody.WriteString(receiverMobile)
	reqBody.WriteString(`","message":"`)
	reqBody.WriteString(encodedMessage)
	if senderMobile != "" {
		reqBody.WriteString(`","from":"`)
		reqBody.WriteString(senderMobile)
	}
	reqBody.WriteString(`"}`)
	return reqBody.String()
}

// handleHTTPError gère les codes d'erreur HTTP de l'API Mobitag
func handleHTTPError(resp *http.Response) error {
	switch resp.StatusCode {
	case 443:
		slog.Error("La clé API est invalide. Veuillez demander une nouvelle clé ou utiliser la commande 'mobitag web' en attendant.")
		return fmt.Errorf("clé API invalide (code 443)")
	case 401:
		slog.Error("Accès non autorisé. Veuillez vérifier votre clé API ou utiliser la commande 'mobitag web' pour obtenir une nouvelle clé.")
		return fmt.Errorf("accès non autorisé (code 401)")
	case 400:
		slog.Error("Requête invalide. Veuillez vérifier les paramètres envoyés.")
		return fmt.Errorf("requête invalide (code 400)")
	case 202:
		return nil
	default:
		slog.Error("Une erreur s'est produite lors de l'envoi du message. Code d'erreur=" + fmt.Sprint(resp.StatusCode) + " message=" + resp.Status)
		return fmt.Errorf("erreur lors de l'envoi du message: code %d, %s", resp.StatusCode, resp.Status)
	}
}

// sendSMS sends an SMS to the specified receiver mobile number
// receiverMobile: the mobile number of the receiver, like 654321
// message: the message to send
func SendSMS(receiverMobile string, message string, senderMobile string, cut bool) error {
	// Nettoyage et découpage du message
	msg := strings.ReplaceAll(message, "\n", " ")
	if len(msg) > 160 {
		if !cut {
			slog.Error("Le message dépasse la limite de 160 caractères length=" + fmt.Sprint(len(msg)))
			return fmt.Errorf("le message dépasse la limite de 160 caractères length=%d", len(msg))
		}
		slog.Warn("Le message dépasse la limite de 160 caractères et sera coupé length=" + fmt.Sprint(len(msg)))
		msg = msg[:155] + "[...]"
	}

	// Encodage Base64
	encodedMessage := base64.StdEncoding.EncodeToString([]byte(msg))
	slog.Debug("Message encodé en Base64=" + encodedMessage)

	// Récupération de la clé API
	mobitagAPIKey := os.Getenv("OPTNC_MOBITAGNC_API_KEY")
	apiURL := "https://api.opt.nc/mobitag/sendSms"

	// Construction de la requête HTTP
	client := &http.Client{}
	reqBody := buildRequestBody(receiverMobile, encodedMessage, senderMobile)
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(reqBody))
	if err != nil {
		slog.Error("Une erreur s'est produite lors de la création de la requête erreur=" + err.Error())
		return fmt.Errorf("erreur lors de la création de la requête: %w", err)
	}

	// Ajout des headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-apikey", mobitagAPIKey)

	// Logs
	if senderMobile != "" {
		slog.Debug("Expéditeur=" + senderMobile)
	}
	slog.Debug("Destinataire=" + receiverMobile)
	slog.Info("Message envoyé=" + msg)

	// Envoi de la requête
	resp, err := client.Do(req)
	if err != nil {
		slog.Error("Une erreur s'est produite lors de l'envoi de la requête erreur=" + err.Error())
		return fmt.Errorf("erreur lors de l'envoi de la requête: %w", err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			slog.Warn("Une erreur s'est produite lors de la fermeture du corps de la réponse erreur=" + err.Error())
		}
	}()

	// Gestion des erreurs HTTP
	if err := handleHTTPError(resp); err != nil {
		return err
	}

	slog.Info("Accusé réception=" + resp.Status)
	return nil
}
