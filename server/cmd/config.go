package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/davecgh/go-spew/spew"
	"github.com/logica0419/remote-bmi/server/benchmark"
	"github.com/logica0419/remote-bmi/server/repository"
	"github.com/logica0419/remote-bmi/server/router"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var c = &Config{}

type Config struct {
	Address string `mapstructure:"address" json:"address" yaml:"address"`
	MySQL   struct {
		Hostname string `mapstructure:"hostname" json:"hostname" yaml:"hostname"`
		Port     int    `mapstructure:"port" json:"port" yaml:"port"`
		Username string `mapstructure:"username" json:"username" yaml:"username"`
		Password string `mapstructure:"password" json:"password" yaml:"password"`
		Database string `mapstructure:"database" json:"database" yaml:"database"`
	} `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Version string `mapstructure:"version" json:"version" yaml:"version"`
	BenchIP string `mapstructure:"bench_ip" json:"bench_ip" yaml:"bench_ip"`
}

func newRouterConfig(c *Config) *router.Config {
	return &router.Config{
		Address: c.Address,
		Version: c.Version,
	}
}

func newRepositoryConfig(c *Config) *repository.Config {
	return &repository.Config{
		Hostname: c.MySQL.Hostname,
		Port:     c.MySQL.Port,
		Username: c.MySQL.Username,
		Password: c.MySQL.Password,
		Database: c.MySQL.Database,
	}
}

func newBenchmarkerConfig(c *Config) *benchmark.Config {
	return &benchmark.Config{
		Version: c.Version,
		BenchIP: c.BenchIP,
	}
}

func loadConfig(configFile string) error {
	viper.SetDefault("address", ":3000")
	viper.SetDefault("mysql.hostname", "127.0.0.1")
	viper.SetDefault("mysql.port", 3306)
	viper.SetDefault("mysql.username", "isucon")
	viper.SetDefault("mysql.password", "isucon")
	viper.SetDefault("mysql.database", "Remote-BMI")
	viper.SetDefault("version", "isucon-test")
	viper.SetDefault("bench_ip", "127.0.0.1")

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Print("Unable to find config file, default settings or environmental variables are to be used.")
		} else {
			return fmt.Errorf("Error: failed to load config file - %s ", err)
		}
	}

	err := viper.Unmarshal(c)
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
		scs.Dump(c)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
