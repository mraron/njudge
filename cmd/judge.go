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
		BindEnvs(judge.Server{})
		viper.SetEnvPrefix("njudge")

		viper.SetConfigName("judge")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		return viper.MergeInConfig()
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		s := judge.NewServer()
		if err := viper.Unmarshal(&s); err != nil {
			return err
		}

		return s.Run()
	},
}

func init() {
	RootCmd.AddCommand(JudgeCmd)
}
