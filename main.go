package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/mraron/njudge/cmd"
)

var profile = os.Getenv("NJUDGE_PROFILE")

func main() {
	if profile == "true" {
		fmt.Printf("Creating file %q\n", "cpuprofile")
		f, err := os.Create("cpuprofile")
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(f)

		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				pprof.StopCPUProfile()
				os.Exit(0)
			}
		}()
	}

	cmd.Execute()
}
