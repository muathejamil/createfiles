/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"createfiles/io"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	_ "io"
	"os"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "createfiles",
	Short: "Generate multiple files with different sizes",
	Long:  `This application is used to save your time by generating multiple files for you with different sizes`,
	Run: func(cmd *cobra.Command, args []string) {
		size, sizeErr := cmd.Flags().GetString("size")
		if sizeErr != nil {
			log.WithFields(log.Fields{
				"error": sizeErr,
			}).Errorf("Error in size file flag!")
			os.Exit(1)
		}

		fmt.Println(size)
		count, countErr := cmd.Flags().GetInt32("count")
		if countErr != nil {
			log.WithFields(log.Fields{
				"error": sizeErr,
			}).Errorf("Error in count file flag!")
			os.Exit(1)
		}
		file := io.NewFile("file1.txt", "./data/", 1)
		io.PopulateTheFile("./data/file1.txt", 1)
		fmt.Println(file)
		fmt.Println(count)
	},
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	Version: "1.0.0",
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
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.createfiles.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().StringP("size", "s", "1kb", "Size of required files")
	rootCmd.Flags().Int32P("count", "c", 1, "The count of files to be generated")
}
