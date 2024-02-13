package main

import (
	"fmt"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/infra"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/config"
	initSentry "github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/sentry"
)

func main() {
	// Config
	c, cErr := config.New()

	// Sentry
	if err := initSentry.SetUp(c); err != nil {
		sentry.CaptureException(fmt.Errorf("initSentry err: %w", err))
	}

	// DB
	_, err := infra.NewDBConnector(c)
	if err != nil {
		sentry.CaptureException(fmt.Errorf("initDB err: %w", err))
	}

	defer func() {
		if cErr != nil {
			sentry.CaptureException(fmt.Errorf("config err: %w", cErr))
		}
		sentry.Recover()
		sentry.Flush(2 * time.Second)
	}()
}
