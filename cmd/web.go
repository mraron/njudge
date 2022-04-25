package cmd

import (
	"github.com/mraron/njudge/web"
	"github.com/mraron/njudge/web/helpers/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

var WebCmd = &cobra.Command{
	Use:   "web",
	Short: "manage web related parts, for example start webserver, run migrations etc.",
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
		if err != nil {
			return err
		}

		s := web.Server{Server: cfg}
		s.Run()

		return nil
	},
}

var (
	File       string
	Problemset string
	Problem    string
	Language   string
	User       int
)

var SubmitCmd = &cobra.Command{
	Use: "submit",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Server{}

		err := viper.Unmarshal(&cfg)
		if err != nil {
			return err
		}

		s := web.Server{Server: cfg}
		src, err := ioutil.ReadFile(File)
		if err != nil {
			return err
		}

		s.SetupEnvironment()
		if err := s.ProblemStore.UpdateProblem(Problem); err != nil {
			return err
		}

		id, err := s.Submit(User, Problemset, Problem, Language, src)
		if err != nil {
			return err
		}

		log.Print("submission received with id ", id)
		return nil
	},
}

func init() {
	RootCmd.AddCommand(WebCmd)

	SubmitCmd.Flags().IntVar(&User, "user", 0, "ID of user on behalf we make the submission")
	SubmitCmd.Flags().StringVar(&Problemset, "problemset", "main", "Problemset of problem")
	SubmitCmd.Flags().StringVar(&Problem, "problem", "", "Problem")
	SubmitCmd.Flags().StringVar(&Language, "language", "cpp14", "Language")
	SubmitCmd.Flags().StringVar(&File, "file", "", "file to submit")

	SubmitCmd.MarkFlagRequired("user")
	SubmitCmd.MarkFlagRequired("problem")
	SubmitCmd.MarkFlagRequired("file")

	WebCmd.AddCommand(SubmitCmd)
}
