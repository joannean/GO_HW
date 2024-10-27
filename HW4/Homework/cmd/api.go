/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/joannean/GO_HW/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Path     string `mapstructure:"path"`
		JsonFile string `mapstructure:"jsonFile"`
	} `mapstructure:"database"`
}

// apiCmd represents the api command
var apiCmd = &cobra.Command{
	Use:   "api",
	Short: "Run API Server",
	Long:  `create todolist, get todolist, update todolist, delete todolist`,
	Run: func(cmd *cobra.Command, args []string) {
		cfgFile, _ := cmd.Flags().GetString("config")

		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("Error reading config file: %s\n", err)
			os.Exit(1)
		}

		var config Config
		if err := viper.Unmarshal(&config); err != nil {
			fmt.Printf("Unable to decode config into struct: %s\n", err)
			os.Exit(1)
		}
		todo.Run_api()
	},
}

func init() {
	rootCmd.AddCommand(apiCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// apiCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// apiCmd.Flags().StringP("port", "p", "8080", "Todo Server Port")
	apiCmd.Flags().StringP("config", "p", "config.yaml", "Path to the config file")
}
