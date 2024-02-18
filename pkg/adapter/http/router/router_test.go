package router

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/adapter/http/handler"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/domain/model"
	repository "github.com/progsystem926/bbs-nextjs-go-back/pkg/lib/mock"
	"github.com/progsystem926/bbs-nextjs-go-back/pkg/usecase"
	"github.com/stretchr/testify/assert"
)

func TestQueryRouter(t *testing.T) {
	ctrlPost := gomock.NewController(t)
	defer ctrlPost.Finish()
	mpr := repository.NewMockPost(ctrlPost)
	pRes := []*model.Post{
		{ID: 1, Text: "testPost1", UserID: "1", CreatedAt: ""},
	}
	mpr.EXPECT().GetPosts().Return(pRes, nil)

	ctrlUser := gomock.NewController(t)
	defer ctrlUser.Finish()
	mur := repository.NewMockUser(ctrlUser)
	uRes := &model.User{
		ID:       1,
		Name:     "testUser1",
		Email:    "testUser1@example.com",
		Password: "testUser1password",
	}
	mur.EXPECT().GetUserById("1").Return(uRes, nil)

	pu := usecase.NewPostUseCase(mpr)
	uu := usecase.NewUserUseCase(mur)
	gh := handler.NewGraphHandler(pu, uu)

	e := echo.New()
	e.POST("/query", gh.QueryHandler())

	reqFile := "testdata/ok_req.json"
	reqBody := loadFile(t, reqFile)
	body := strings.NewReader(string(reqBody))
	req := httptest.NewRequest(http.MethodPost, "/query", body)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	rspFile := "testdata/ok_res.json"
	wnt := loadFile(t, rspFile)
	assert.JSONEq(t, string(wnt), rec.Body.String())
}

func loadFile(t *testing.T, path string) []byte {
	t.Helper()

	bt, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read from %q: %v", path, err)
	}
	return bt
}
