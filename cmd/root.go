package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "TempLI",
	Short: "An app made to get weather information through the terminal",
	Long: `You can use this app to get weather information through the terminal.
it can also be used with the --raw flag to get the response in json format.`,
	Version: "0.0.1",
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("version", "v", false, "Print the version number")
	rootCmd.SetVersionTemplate("Hey there, I'm version {{.Version}}")

}
