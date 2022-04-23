package cmd

import (
	"fmt"
	"github.com/mraron/njudge/web"
	"github.com/mraron/njudge/web/helpers/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var WebCmd = &cobra.Command{
	Use: "web",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		viper.SetEnvPrefix("njudge")

		viper.SetConfigName("web")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		return viper.MergeInConfig()
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Server{}

		err := viper.Unmarshal(&cfg)
		fmt.Println(cfg)
		if err != nil {
			return err
		}

		s := web.Server{Server: cfg}
		s.Run()

		return nil
	},
}

func init() {
	RootCmd.AddCommand(WebCmd)
}
