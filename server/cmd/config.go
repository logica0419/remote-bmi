package cmd

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfg = &Config{}

type Config struct {
	Address string `mapstructure:"address" json:"address"`
	MySQL   struct {
		Hostname string `mapstructure:"hostname" json:"hostname,omitempty"`
		Port     int    `mapstructure:"port" json:"port,omitempty"`
		Username string `mapstructure:"username" json:"username,omitempty"`
		Password string `mapstructure:"password" json:"password,omitempty"`
		Database string `mapstructure:"database" json:"database,omitempty"`
	} `mapstructure:"mysql" json:"mysql"`
}

func loadConfig(cfgFile string) error {
	viper.SetDefault("address", ":3000")
	viper.SetDefault("mysql.hostname", "127.0.0.1")
	viper.SetDefault("mysql.port", 3306)
	viper.SetDefault("mysql.username", "isucon")
	viper.SetDefault("mysql.password", "isucon")
	viper.SetDefault("mysql.database", "Remote-BMI")

	viper.AutomaticEnv()

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("config")
		viper.SetConfigType("json")
	}

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print("Unable to find config file, default settings or environmental variables are to be used.")
		} else {
			return fmt.Errorf("Error: failed to load config file - %s ", err)
		}
	}

	err := viper.Unmarshal(cfg)
	if err != nil {
		return fmt.Errorf("Error: failed to parse configs - %s ", err)
	}

	return nil
}

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Print current configurations to stdout",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Current Configurations:")

		scs := spew.ConfigState{
			Indent:                  "\t",
			DisablePointerAddresses: true,
		}
		scs.Dump(cfg)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
