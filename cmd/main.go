package main

import (
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/getsentry/sentry-go"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/adapter/http/handler"
	authMiddleware "github.com/progsystem926/bbs-nextjs-go-back/pkg/adapter/http/middleware"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/adapter/http/router"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/infra"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/config"
	initSentry "github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/sentry"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/usecase"
)

func main() {
	// Config
	c, cErr := config.New()

	// Sentry
	if err := initSentry.SetUp(c); err != nil {
		sentry.CaptureException(fmt.Errorf("initSentry err: %w", err))
	}

	// DB
	db, err := infra.NewDBConnector(c)
	if err != nil {
		sentry.CaptureException(fmt.Errorf("initDb err: %w", err))
	}

	// DI
	mr := infra.NewPostRepository(db)
	ur := infra.NewUserRepository(db, c)
	au := usecase.NewAuthUseCase(ur, c)
	mu := usecase.NewPostUseCase(mr)
	uu := usecase.NewUserUseCase(ur, c)
	ch := handler.NewCsrfHandler()
	lh := handler.NewLoginHandler(au)
	gh := handler.NewGraphHandler(mu, uu)
	sh := handler.NewSignUpHandler(uu)
	ph := playground.Handler("GraphQL", "/query")
	am := authMiddleware.NewAuthMiddleware(au)

	// Rooting
	r := router.NewInitRouter(ch, lh, gh, sh, ph, am)
	_, err = r.InitRouting(c)
	if err != nil {
		sentry.CaptureException(fmt.Errorf("InitRouting at NewInitRoute err: %w", err))
	}

	defer func() {
		if cErr != nil {
			sentry.CaptureException(fmt.Errorf("config err: %w", cErr))
		}
		sentry.Recover()
		sentry.Flush(2 * time.Second)
	}()
}
