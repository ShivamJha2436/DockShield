package config

import (
	"github.com/spf13/viper"
	"log"
)

type AppConfig struct {
	MinSeverity string `mapstructure:"min_severity"`
	OutputStyle string `mapstructure:"output_style"`
	SaveReport  bool   `mapstructure:"save_report"`
}

var Cfg AppConfig

func LoadConfig() {
	viper.SetConfigName("dockshield")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// Defaults
	viper.SetDefault("min_severity", "HIGH")
	viper.SetDefault("output_style", "table")
	viper.SetDefault("save_report", false)

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("⚠️ No config file found. Using default settings.")
	}

	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("Config parse error: %v", err)
	}
}
