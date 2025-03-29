/*
Copyright © 2025 Dontknow09
*/
package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/spf13/cobra"
)

// launchCmd represents the launch command
var launchCmd = &cobra.Command{
	Use:   "launch",
	Short: "Launch a Modpack",
	Long:  `Launch the client with the specified modpack. This command will launch the client with the specified modpack folder.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Info("Launching modpack...")
	},
}

func init() {
	rootCmd.AddCommand(launchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// launchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// launchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
