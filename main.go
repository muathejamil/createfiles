/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"createfiles/cmd"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	viper.AddConfigPath("C:\\Users\\Lenovo\\GolandProjects\\createfiles\\config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Error in reading config file")
		os.Exit(3)
	}
	cmd.Execute()
}
