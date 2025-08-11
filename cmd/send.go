package cmd

import (
	"os"

	"log/slog"

	"github.com/spf13/cobra"
)

var sendCmd = &cobra.Command{
	Use:     "send",
	Aliases: []string{"s"}, // L'alias pour la commande
	Short:   "Envoyer un Mobitag",
	Long:    `Envoi d'un Mobitag à un numéro de téléphone.`,
	Example: `mobitag send --to <destinataire> --message <message> --from <expéditeur>
mobitag send --to 123456 --message "Hello, world!" --from 654321
mobitag s -t 123456 -m "Hello, world!" -f 654321`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("OPTNC_MOBITAGNC_API_KEY") == "" {
			slog.Error("La clé API 'OPTNC_MOBITAGNC_API_KEY' n'est pas définie dans les variables d'environnement. Veuillez définir cette clé ou utiliser la commande 'mobitag web' en attendant d'avoir une clé.")
			os.Exit(1)
		}

		// Configuration du niveau de journalisation en fonction du flag verbose
		verboseLevel, _ := cmd.Flags().GetString("verbose")
		switch verboseLevel {
		case "warn":
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn})))
		case "info":
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))
		case "debug":
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))
		default:
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError})))
		}

		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		to, _ := cmd.Flags().GetString("to")
		message, _ := cmd.Flags().GetString("message")
		from, _ := cmd.Flags().GetString("from")
		cut, _ := cmd.Flags().GetBool("cut")

		return SendSMS(to, message, from, cut)
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

	sendCmd.Flags().StringP("to", "t", "", "Numéro de téléphone du destinataire (obligatoire)")
	sendCmd.Flags().StringP("message", "m", "", "Message à envoyer (obligatoire)")
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
	sendCmd.Flags().StringP("verbose", "v", "info", "Niveau de journalisation (warn, info, debug)")
}
