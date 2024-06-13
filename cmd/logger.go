package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

type WebhookHandler struct {
	next slog.Handler
	url  string
}

func (w WebhookHandler) Enabled(ctx context.Context, level slog.Level) bool {
	return true
}

func (w WebhookHandler) Handle(ctx context.Context, record slog.Record) error {
	if record.Level == slog.LevelError {
		r := record.Clone()
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
		go func() {
			defer cancel()
			attrs := make([]byte, 0, 1024)
			r.Attrs(func(attr slog.Attr) bool {
				attrs = append(attrs, []byte("- ")...)
				attrs = append(attrs, []byte(attr.Key)...)
				attrs = append(attrs, '=')
				attrs = append(attrs, []byte(attr.Value.String())...)
				attrs = append(attrs, '\n')
				return true
			})
			msg, err := json.Marshal(struct {
				Content string `json:"content"`
			}{fmt.Sprintf("<@&1244740023677616189> %s **%s**\n```%s```", r.Time.Format(time.DateTime), r.Message, string(attrs))})
			if err != nil {
				return
			}
			req, _ := http.NewRequestWithContext(ctx, "POST", w.url, bytes.NewBuffer(msg))
			req.Header.Set("Content-Type", "application/json")
			_, _ = http.DefaultClient.Do(req)
		}()
	}
	return w.next.Handle(ctx, record)
}

func (w WebhookHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return WebhookHandler{
		next: w.next.WithAttrs(attrs),
		url:  w.url,
	}
}

func (w WebhookHandler) WithGroup(name string) slog.Handler {
	return WebhookHandler{
		next: w.next.WithGroup(name),
		url:  w.url,
	}
}

func getHookedLogger(discordWebhookURL string) *slog.Logger {
	var (
		logger       *slog.Logger
		innerHandler slog.Handler = slog.NewTextHandler(os.Stdout, nil)
	)
	if discordWebhookURL != "" {
		logger = slog.New(WebhookHandler{
			next: innerHandler,
			url:  discordWebhookURL,
		})
	} else {
		logger = slog.New(innerHandler)
	}
	return logger
}
