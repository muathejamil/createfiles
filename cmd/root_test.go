package cmd

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func BenchmarkExecute(b *testing.B) {
	viper.AddConfigPath("C:\\Users\\Lenovo\\GolandProjects\\createfiles\\config")
	viper.SetConfigName("config")
	err := viper.ReadInConfig()
	if err != nil {
		log.Error("Error in reading config file")
		os.Exit(3)
	}
	for n := 0; n < b.N; n++ {
		Execute()
	}
}
