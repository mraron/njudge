package cmd

import (
	"github.com/mraron/njudge/pkg/glue"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GlueCmd = &cobra.Command{
	Use:   "glue",
	Short: "",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		BindEnvs(glue.Config{})
		viper.SetEnvPrefix("njudge")

		viper.SetConfigName("glue")
		viper.AddConfigPath(".")
		viper.AutomaticEnv()
		return viper.MergeInConfig()
	},

	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := glue.Config{}

		err := viper.Unmarshal(&cfg)
		if err != nil {
			return err
		}

		s := glue.Server{Config: cfg}
		s.Run()

		return nil
	},
}

func init() {
	RootCmd.AddCommand(GlueCmd)
}
