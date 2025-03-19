package cmd

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

// Variables injectées au build avec ldflags
var (
	Commit       = "none"
	Date         = "unknown"
	BuiltBy      = "GoReleaser"
	GoVersion    = runtime.Version()
	GitTreeState = "unknown"
)

// aboutCmd represents the about command
var aboutCmd = &cobra.Command{
	Use:     "about",
	Aliases: []string{"a"},
	Short:   "Informations sur ce CLI",
	Long:    `Ce cli d'innovation disruptive (périmètre fonctionnel, stack, Open Source) est la suite de tout un processus d'innovations et d'un contexte : il convient que l'utilisateur final puisse en prorendre connaissance...depuis le terminal.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Définition des couleurs avec les codes hexadécimaux de l'entreprise
		titleColor := color.New(color.FgHiYellow).Add(color.BgBlue, color.Bold) // Yellow text on dark blue background
		sloganColor := color.New(color.FgHiWhite, color.Italic)                 // #FFFFFF (white) in italic
		sectionColor := color.New(color.FgHiYellow)                             // #FBBC09 (yellow)
		labelColor := color.New(color.FgHiWhite)                                // #FFFFFF (white) pour les labels
		valueColor := color.New(color.FgHiWhite)                                // #FFFFFF (white) pour les valeurs

		// Générer l'ASCII art en jaune sur fond bleu
		myFigure := figure.NewFigure("mobitag", "starwars", true)
		asciiArtLines := myFigure.Slicify()
		maxLength := 0
		for _, line := range asciiArtLines {
			if len(line) > maxLength {
				maxLength = len(line)
			}
		}
		for _, line := range asciiArtLines {
			paddedLine := line + strings.Repeat(" ", maxLength-len(line))
			titleColor.Add(color.BgBlue).Println(paddedLine)
		}

		// Affichage du slogan en blanc et italique
		sloganColor.Println("\"Le plus court chemin entre le terminal et un mobitag.\"")

		// Affichage des métadonnées injectées
		fmt.Println()
		sectionColor.Println("--- Build Info ---")
		labelColor.Printf("%-20s ", "GitVersion:")
		valueColor.Println(Version)

		labelColor.Printf("%-20s ", "Git Commit:")
		valueColor.Println(Commit)

		labelColor.Printf("%-20s ", "GitTreeState:")
		valueColor.Println("clean") // Ajuste selon ton besoin

		labelColor.Printf("%-20s ", "BuildDate:")
		valueColor.Println(Date)

		labelColor.Printf("%-20s ", "BuiltBy:")
		valueColor.Println(BuiltBy)

		labelColor.Printf("%-20s ", "GoVersion:")
		valueColor.Println(GoVersion)

		labelColor.Printf("%-20s ", "Compiler:")
		valueColor.Println(runtime.Compiler)

		labelColor.Printf("%-20s ", "Platform:")
		valueColor.Printf("%s/%s\n", runtime.GOOS, runtime.GOARCH)

		// Affichage des ressources
		fmt.Println()
		sectionColor.Println("--- Ressources ---")
		labelColor.Printf("%-20s ", "Licence:")
		valueColor.Println("AGPLv3")

		labelColor.Printf("%-20s ", "Code:")
		valueColor.Println("https://github.com/opt-nc/mobitag-cli")

		labelColor.Printf("%-20s ", "Roadmap:")
		valueColor.Println("https://github.com/orgs/opt-nc/projects/24")

		labelColor.Printf("%-20s ", "Site Web:")
		valueColor.Println("http://www.mobitag.nc")

	},
}

func init() {
	rootCmd.AddCommand(aboutCmd)
}
