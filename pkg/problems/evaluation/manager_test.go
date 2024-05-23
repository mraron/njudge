package evaluation_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/mraron/njudge/pkg/problems/evaluation"
	"testing"
	"time"
)

func TestManager(t *testing.T) {
	m := evaluation.NewManager()
	_ = m.AddStep("compilation", func(ctx context.Context) error {
		fmt.Println("compile", ctx.Err())
		time.Sleep(1 * time.Second)
		return nil
	}, nil)
	_ = m.AddStep("test1", func(ctx context.Context) error {
		fmt.Println("test1")
		time.Sleep(1 * time.Second)
		return errors.New("dafuq")
	}, []string{"compilation"})
	_ = m.AddStep("test2", func(ctx context.Context) error {
		fmt.Println("test2")
		time.Sleep(1 * time.Second)
		return nil
	}, []string{"compilation", "test1"})
	_ = m.AddStep("test3", func(ctx context.Context) error {
		fmt.Println("test3")
		time.Sleep(1 * time.Second)
		return nil
	}, []string{"test1"})
	_ = m.Run(context.Background())
}
