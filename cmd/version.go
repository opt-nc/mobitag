/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v"}, // L'alias pour la commande
	Short:   "Affiche la version de l'application",
	Long:    `Affiche la version de l'application`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
