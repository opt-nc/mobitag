package cmd

import (
	"fmt"
	"runtime"

	"github.com/common-nighthawk/go-figure"
	"github.com/spf13/cobra"
)

// Variables injectées au build avec ldflags
var (
	Commit    = "none"
	Date      = "unknown"
	BuiltBy   = "GoReleaser"
	GoVersion = runtime.Version()
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:     "about",
	Aliases: []string{"a"},
	Short:   "Informations sur ce CLI",
	Long:    `Ce cli d'innovation disruptive (périmètre fonctionnel, stack, Open Source) est la suite de tout un processus d'innovations et d'un contexte : il convient que l'utilisateur final puisse en prorendre connaissance...depuis le terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Générer l'ASCII art
		myFigure := figure.NewFigure("mobitag-cli", "doom", true)
		myFigure.Print()

		fmt.Println("\"Le plus court chemin entre le terminal et un mobitag.\"")

		// Affichage des métadonnées injectées
		fmt.Println("\n--- Build Info ---")
		fmt.Printf("%-20s %s\n", "GitVersion:", Version)
		fmt.Printf("%-20s %s\n", "Git Commit:", Commit)
		fmt.Printf("%-20s %s\n", "GitTreeState:", "clean") // Ajuste selon ton besoin
		fmt.Printf("%-20s %s\n", "BuildDate:", Date)
		fmt.Printf("%-20s %s\n", "BuiltBy:", BuiltBy)
		fmt.Printf("%-20s %s\n", "GoVersion:", GoVersion)
		fmt.Printf("%-20s %s\n", "Compiler:", runtime.Compiler)
		fmt.Printf("%-20s %s/%s\n", "Platform:", runtime.GOOS, runtime.GOARCH)

		fmt.Println("\n--- Ressources ---")
		fmt.Printf("%-20s %s\n", "Licence:", "AGPLv3")
		fmt.Printf("%-20s %s\n", "Code:", "https://github.com/opt-nc/mobitag-cli")
		fmt.Printf("%-20s %s\n", "Roadmap:", "https://github.com/orgs/opt-nc/projects/24")
		fmt.Printf("%-20s %s\n", "Site Web:", "http://www.mobitag.nc")

	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
