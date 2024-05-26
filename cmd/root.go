package cmd

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func BindEnvs(iface interface{}, parts ...string) {
	ifv := reflect.ValueOf(iface)
	ift := reflect.TypeOf(iface)
	for i := 0; i < ift.NumField(); i++ {
		v := ifv.Field(i)
		t := ift.Field(i)
		tv, ok := t.Tag.Lookup("mapstructure")
		if !ok {
			continue
		}
		squash := strings.Contains(tv, ",squash")
		tv = strings.Split(tv, ",")[0]
		switch v.Kind() {
		case reflect.Struct:
			if squash {
				BindEnvs(v.Interface(), parts...)
			} else {
				BindEnvs(v.Interface(), append(parts, tv)...)
			}
		default:
			viper.MustBindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}

var RootCmd = &cobra.Command{
	Use:     "njudge",
	Version: "v0.5.0",
	Long:    "cli tool to manage njudge instance",
}

func Execute() {
	RootCmd.AddCommand(NewWebCmd(viper.GetViper()))
	RootCmd.AddCommand(NewJudgeCmd(viper.GetViper()))
	RootCmd.AddCommand(NewGlueCmd(viper.GetViper()))
	RootCmd.AddCommand(NewTestProblemCmd())

	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
