package cmd

import (
	"github.com/mraron/njudge/internal/judge"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var JudgeCmd = &cobra.Command{
	Use:   "judge",
	Short: "start judge server",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		BindEnvs(judge.HTTPServer{})
		viper.SetEnvPrefix("njudge")

		viper.SetConfigName("judge")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		return viper.MergeInConfig()
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := judge.ServerConfig{}
		if err := viper.Unmarshal(&cfg); err != nil {
			return err
		}

		s, err := judge.NewServer(cfg)
		if err != nil {
			return err
		}

		s.Run()
		return nil
	},
}

func init() {
	RootCmd.AddCommand(JudgeCmd)
}
