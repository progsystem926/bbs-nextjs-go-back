package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/labstack/echo/v4"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/adapter/http/resolver"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/graph/generated"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/usecase"
)

type Graph interface {
	QueryHandler() echo.HandlerFunc
}

type GraphHandler struct {
	PostUserCase usecase.Post
	UserUseCase  usecase.User
}

func NewGraphHandler(pu usecase.Post, uu usecase.User) Graph {
	GraphHandler := GraphHandler{
		PostUserCase: pu,
		UserUseCase:  uu,
	}
	return &GraphHandler
}

func (g *GraphHandler) QueryHandler() echo.HandlerFunc {
	rslvr := resolver.Resolver{
		PostUseCase: g.PostUserCase,
		UserUseCase: g.UserUseCase,
	}

	srv := handler.NewDefaultServer(
		generated.NewExecutableSchema(generated.Config{Resolvers: &rslvr}),
	)

	return func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	}
}
