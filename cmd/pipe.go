package cmd

import (
	"bufio"
	"io"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

// pipeCmd represents the pipe command
var pipeCmd = &cobra.Command{
	Use:     "pipe",
	Aliases: []string{"p"}, // L'alias pour la commande
	Short:   "Envoyer un Mobitag depuis un pipe",
	Long:    `Envoi d'un Mobitag de moins de 160 caractères à un numéro de téléphone depuis un pipe.`,
	Example://`<sortie d'une commande> | mobitag pipe --to <destinataire> --from <expéditeur>
	// pwd | mobitag pipe --to 123456 --from 654321
	`echo "Hello c'est $(whoami) : alors on se le fait ce café ?" | mobitag p -t 123456 -f 654321`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if os.Getenv("OPTNC_MOBITAGNC_API_KEY") == "" {
			slog.Error("La clé API 'OPTNC_MOBITAGNC_API_KEY' n'est pas définie dans les variables d'environnement. Veuillez définir cette clé ou utiliser la commande 'mobitag web' en attendant d'avoir une clé.")
			os.Exit(1)
		}
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		to, _ := cmd.Flags().GetString("to")
		from, _ := cmd.Flags().GetString("from")

		var message string
		// Check if there is input from stdin
		if stat, _ := os.Stdin.Stat(); (stat.Mode() & os.ModeCharDevice) == 0 {
			reader := bufio.NewReader(os.Stdin)
			input, _ := io.ReadAll(reader)
			message = string(input)
		} else {
			slog.Error("Aucune entrée n'a été trouvée. Veuillez utiliser un pipe pour envoyer un message.")
			os.Exit(1)
		}

		cut, _ := cmd.Flags().GetBool("cut")
		logLevel, _ := cmd.Flags().GetString("log-level")

		// Configuration du niveau de journalisation en fonction du flag log-level
		switch logLevel {
		case "warn":
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelWarn})))
		case "info":
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo})))
		case "debug":
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})))
		default:
			slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelError})))
		}

		return SendSMS(to, message, from, cut)
	},
}

func init() {
	rootCmd.AddCommand(pipeCmd)

	pipeCmd.Flags().StringP("to", "t", "", "Numéro de téléphone du destinataire (obligatoire)")
	// pipeCmd.Flags().StringP("from", "f", "", "Numéro de téléphone de l'expéditeur")
	err := pipeCmd.MarkFlagRequired("to")
	if err != nil {
		slog.Error("Erreur lors du marquage du flag 'to' comme requis error=" + err.Error())
		os.Exit(1)
	}

	pipeCmd.Flags().BoolP("cut", "c", false, "Couper le message si sa taille dépasse 160 caractères afin de ne pas excéder la limite")
	pipeCmd.Flags().StringP("log-level", "l", "info", "Niveau de journalisation (warn, info, debug)")
}
