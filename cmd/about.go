package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:     "about",
	Aliases: []string{"a"},
	Short:   "Informations sur ce CLIi",
	Long:    `Ce cli d'innovation disruptive (périmètre fonctionnel, stack, Open Source) est la suite de tout un processus d'innovations et d'un contexte : il convient que l'utilisateur final puisse en prorendre connaissance...depuis le terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`mobitag-cli : le plus court chemin entre le terminal et un mobitag.

Ce cli est l'aboutissement de tout un cheminement autour de mobitag et de sa digitalisation : 

1. Célébration des 25 ans de mobitag avec l'API sur APIGEE : https://bit.ly/3WiwffU
2. Un premier hackathon 2024-06-22 : https://bit.ly/4gPqzSN
3. Intégration à un assistant IA via Open Interpreter  : https://bit.ly/3D6iGKW

📑 Ressources

- Licence : [AGPLv3](https://www.gnu.org/licenses/agpl-3.0.fr.html#license-text)
- Code : https://github.com/opt-nc/mobitag-cli
- Roadmap : https://github.com/orgs/opt-nc/projects/24`)
	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
