/*
Copyright © 2025 Dontknow09
*/
package cmd

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
	"github.com/spf13/cobra"
)

var (
	name          string
	modloader     string
	gameversion   string
	loaderversion string
)

type formModel struct {
	form huh.Form
}

func (m formModel) Init() tea.Cmd {
	return nil
}

func (m formModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	form, cmd := m.form.Update(msg)
	m.form = *form.(*huh.Form)
	return m, cmd
}

func (m formModel) View() string {
	return m.form.View()
}

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		p := tea.NewProgram(formModel{form: *huh.NewForm(
			huh.NewGroup().Title("Name"),
			huh.NewGroup().Title("Create new profile"),
		)})
		if _, err := p.Run(); err != nil {
			fmt.Println("Error starting form:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
