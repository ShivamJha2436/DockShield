package config

import (
	"log"

	"github.com/spf13/viper"
)

type AppConfig struct {
	MinSeverity string `mapstructure:"min_severity"`
	OutputStyle string `mapstructure:"output_style"`
	SaveReport  bool   `mapstructure:"save_report"`
}

var Cfg AppConfig

func LoadConfig() {
	viper.SetConfigName("dockshield") // loads dockshield.yaml
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".") // look in current directory

	// Set defaults
	viper.SetDefault("min_severity", "HIGH")
	viper.SetDefault("output_style", "table")
	viper.SetDefault("save_report", false)

	if err := viper.ReadInConfig(); err != nil {
		log.Println("⚠️ Config not found. Using defaults.")
	}

	if err := viper.Unmarshal(&Cfg); err != nil {
		log.Fatalf("❌ Config parse error: %v", err)
	}
}
