package cmd

import (
	"io/ioutil"
	"log"

	"github.com/mraron/njudge/web"
	"github.com/mraron/njudge/web/helpers/config"
	"github.com/mraron/njudge/web/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

var SubmitCmdArgs struct {
	File       string
	Problemset string
	Problem    string
	Language   string
	User       int
}

var SubmitCmd = &cobra.Command{
	Use: "submit",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Server{}

		err := viper.Unmarshal(&cfg)
		if err != nil {
			return err
		}

		s := web.Server{Server: cfg}
		src, err := ioutil.ReadFile(SubmitCmdArgs.File)
		if err != nil {
			return err
		}

		s.SetupEnvironment()
		if err := s.ProblemStore.Update(); err != nil {
			return err
		}

		id, err := s.Submit(SubmitCmdArgs.User, SubmitCmdArgs.Problemset, SubmitCmdArgs.Problem, SubmitCmdArgs.Language, src)
		if err != nil {
			return err
		}

		log.Print("submission received with id ", id)
		return nil
	},
}

var ActivateCmdArgs struct {
	Name string
}

var ActivateCmd = &cobra.Command{
	Use: "activate",
	RunE: func(cmd *cobra.Command, args []string) error {
		cfg := config.Server{}

		err := viper.Unmarshal(&cfg)
		if err != nil {
			return err
		}

		s := web.Server{Server: cfg}

		s.SetupEnvironment()
		s.ConnectToDB()
		_, err = models.Users(qm.Where("name = ?", ActivateCmdArgs.Name)).UpdateAll(s.DB, models.M{"activation_key": nil})

		return err
	},
}

func init() {
	RootCmd.AddCommand(WebCmd)

	SubmitCmd.Flags().IntVar(&SubmitCmdArgs.User, "user", 0, "ID of user on behalf we make the submission")
	SubmitCmd.Flags().StringVar(&SubmitCmdArgs.Problemset, "problemset", "main", "Problemset of problem")
	SubmitCmd.Flags().StringVar(&SubmitCmdArgs.Problem, "problem", "", "Problem")
	SubmitCmd.Flags().StringVar(&SubmitCmdArgs.Language, "language", "cpp14", "Language")
	SubmitCmd.Flags().StringVar(&SubmitCmdArgs.File, "file", "", "file to submit")

	SubmitCmd.MarkFlagRequired("user")
	SubmitCmd.MarkFlagRequired("problem")
	SubmitCmd.MarkFlagRequired("file")

	WebCmd.AddCommand(SubmitCmd)

	ActivateCmd.Flags().StringVar(&ActivateCmdArgs.Name, "name", "", "name os the user to activate")
	ActivateCmd.MarkFlagRequired("name")

	WebCmd.AddCommand(ActivateCmd)
}
