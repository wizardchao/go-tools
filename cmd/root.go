package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "",
	Short: "",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(m3u8Cmd)
}

func Execute() error {
	return rootCmd.Execute()
}
