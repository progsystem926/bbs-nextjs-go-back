package sentry

import (
	"strings"

	"github.com/getsentry/sentry-go"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/config"
	"golang.org/x/xerrors"
)

func SetUp(c *config.Config) error {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:              c.SentryDsn,
		Environment:      c.Env,
		Debug:            true,
		AttachStacktrace: true,
		TracesSampleRate: 1.0,
		BeforeSend: func(event *sentry.Event, hint *sentry.EventHint) *sentry.Event {
			for i := range event.Exception {
				exception := &event.Exception[i]
				if !strings.Contains(exception.Type, "wrapError") {
					continue
				}
				sp := strings.SplitN(exception.Value, ":", 2)
				if len(sp) != 2 {
					continue
				}
				exception.Type, exception.Value = sp[0], sp[1]
			}
			return event
		},
	}); err != nil {
		return xerrors.Errorf("fail to init sentry: %w", err)
	}

	return nil
}
