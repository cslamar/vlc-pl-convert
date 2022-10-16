/*
Copyright © 2022 Chris Slamar <chris@slamar.com>
*/
package cmd

import (
	"fmt"
	playlist "github.com/cslamar/vlc-pl-convert/pkg"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "vlc-pl-convert",
	Short: "Create new VLC playlist with pause tracks for recording",
	Run: func(cmd *cobra.Command, args []string) {
		inputFile, err := cmd.Flags().GetString("input")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		outputFile, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		pauseSeconds, err := cmd.Flags().GetInt("pause-seconds")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		if err := playlist.Generate(inputFile, outputFile, pauseSeconds, verbose); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("input", "i", "", "input VLC xspf playlist file")
	rootCmd.Flags().StringP("output", "o", "output.xspf", "output VLC xspf playlist file")
	rootCmd.Flags().IntP("pause-seconds", "p", 2, "pause delay between tracks")
	rootCmd.Flags().BoolP("verbose", "v", false, "enable verbose debugging output")
	rootCmd.MarkFlagRequired("input")
}
