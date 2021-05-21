package cmd

import (
	"github.com/spf13/cobra"
	"template/adapters/bot"
)

func init() {
	rootCmd.AddCommand(botCmd)
}

var botCmd = &cobra.Command{
	Use:   "start-bot",
	Short: "Starting bot server",
	Run: func(cmd *cobra.Command, args []string) {
		bot.StartBot()
	},
}

