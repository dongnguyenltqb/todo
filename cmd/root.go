package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"todo/pkg/infra"
	"todo/pkg/utils"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

var rootCmd = &cobra.Command{
	Use: "todo",
	Run: func(cmd *cobra.Command, args []string) {
		utils.ClearScreen()
		Cmd := exec.Command("todo", "--help")
		Cmd.Stdout = os.Stdout
		Cmd.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	infra.Setup()
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("./")
		viper.SetConfigName(".config")
		viper.SetConfigType("yaml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
