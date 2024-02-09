package evaluation_test

import (
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"golang.org/x/net/context"
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	m := evaluation.NewManager()
	m.AddStep("compilation", func(ctx context.Context) error {
		fmt.Println("compile", ctx.Err())
		time.Sleep(1 * time.Second)
		return nil
	}, nil)
	m.AddStep("test1", func(ctx context.Context) error {
		fmt.Println("test1")
		time.Sleep(1 * time.Second)
		return errors.New("dafuq")
	}, []string{"compilation"})
	m.AddStep("test2", func(ctx context.Context) error {
		fmt.Println("test2")
		time.Sleep(1 * time.Second)
		return nil
	}, []string{"compilation", "test1"})
	m.AddStep("test3", func(ctx context.Context) error {
		fmt.Println("test3")
		time.Sleep(1 * time.Second)
		return nil
	}, []string{"test1"})
	m.Run(context.Background())
}
