/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"createfiles/io"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	_ "io"
	"os"
	"strconv"
	"strings"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "createfiles",
	Short: "Generate multiple files with different sizes",
	Long:  `This application is used to save your time by generating multiple files for you with different sizes`,
	Run: func(cmd *cobra.Command, args []string) {
		size, err := cmd.Flags().GetString("size")
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Errorf("Error in the size file flag!")
			os.Exit(1)
		}
		unit := size[len(size)-2:]
		unitSize, err := strconv.Atoi(size[:len(size)-2])
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Errorf("Error in the size flag!")
			os.Exit(1)
		}
		upperUnit := strings.ToUpper(unit)
		unitSizeInKB := MapToKb(upperUnit, unitSize)
		count, err := cmd.Flags().GetInt("count")
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Errorf("Error in count file flag!")
			os.Exit(1)
		}
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.WithFields(log.Fields{
				"error": err,
			}).Errorf("Error in count file flag!")
			os.Exit(1)
		}
		io.CreateBatch(path, count, unitSizeInKB)
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
	rootCmd.Flags().StringP("path", "p", "data", "The path of the directory to generate file on")
	rootCmd.Flags().IntP("count", "c", 1, "The count of files to be generated")
}

func MapToKb(unit string, unitsize int) int {
	switch unit {
	case "KB":
		unitsize = unitsize * 1
		break
	case "MB":
		unitsize = unitsize * 1000
		break
	case "GB":
		unitsize = unitsize * 1000000
		break
	default:
		log.Errorf("Unsupported size unit: %s", unit)
		os.Exit(1)
	}
	return unitsize
}
