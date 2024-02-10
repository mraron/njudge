package sandbox_test

import (
	"context"
	"flag"
	"github.com/mraron/njudge/pkg/language/sandbox"
	"github.com/stretchr/testify/assert"
	"log/slog"
	"testing"
	"time"
)

var verbose = flag.Bool("verbose", false, "log sandbox")

func TestIsolate_Run(t *testing.T) {
	testcases := []struct {
		name        string
		config      sandbox.RunConfig
		command     []string
		wantVerdict sandbox.Verdict
	}{
		{
			name: "sh_echo",
			config: sandbox.RunConfig{
				RunID: "sh_echo",
			},
			command:     []string{"/bin/sh", "-c", "echo \"nigga\""},
			wantVerdict: sandbox.VerdictOK,
		},
		{
			name: "sh_yes",
			config: sandbox.RunConfig{
				RunID:     "sh_yes",
				TimeLimit: 50 * time.Millisecond,
			},
			command:     []string{"/bin/sh", "-c", "yes"},
			wantVerdict: sandbox.VerdictTL,
		},
	}

	var opts []sandbox.IsolateOption
	if *verbose {
		opts = append(opts, sandbox.IsolateOptionUseLogger(slog.Default()))
	}
	s, err := sandbox.NewIsolate(558, opts...)
	assert.Nil(t, err)
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.Init(context.Background())
			assert.Nil(t, err)

			st, err := s.Run(context.Background(), tc.config, tc.command[0], tc.command[1:]...)
			assert.Nil(t, err)

			assert.Equal(t, tc.wantVerdict, st.Verdict)
			err = s.Cleanup(context.Background())
			assert.Nil(t, err)
		})
	}
}
