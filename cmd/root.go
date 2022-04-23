package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var RootCmd = &cobra.Command{
	Use:     "@todo1",
	Version: "v0.1.0",
	Short:   "@todo2 short",
	Long:    "@todo3 long",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	fmt.Println("??")

	viper.SetConfigName("judge")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.MergeInConfig()
}
