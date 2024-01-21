package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/realtemirov/task-for-dell/config"
	"github.com/realtemirov/task-for-dell/internal/blogs/mock"
	"github.com/realtemirov/task-for-dell/internal/blogs/usecase"
	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/pkg/logger"
	"github.com/realtemirov/task-for-dell/pkg/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestBlogHandlers_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.LoadConfig("./../../../../config/config-local")
	require.NoError(t, err)

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()

	mockBlogUC := mock.NewMockUseCase(ctrl)
	blogUC := usecase.NewBlogUseCase(cfg, mockBlogUC, logger)
	blogHandler := NewBlogsHandlers(cfg, blogUC, logger)
	handler := blogHandler.Create()

	t.Run("Create succes case", func(t *testing.T) {
		blog := models.Blog{
			Title:   "title-test",
			Content: "content-test",
		}

		bufferData, err := utils.AnyToBytesBuffer(blog)
		require.NoError(t, err)
		require.NotNil(t, bufferData)

		request := httptest.NewRequest(http.MethodPost, "/v1/blogs", strings.NewReader(bufferData.String()))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)

		mockBlogUC.EXPECT().Create(gomock.Any(), &blog).Return(&blog, nil)

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("Create Validate error case", func(t *testing.T) {

		blog := models.Blog{
			Title:   "",
			Content: "",
		}

		bufferData, err := utils.AnyToBytesBuffer(blog)
		require.NoError(t, err)
		require.NotNil(t, bufferData)

		request := httptest.NewRequest(http.MethodPost, "/v1/blogs", strings.NewReader(bufferData.String()))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		err = handler(echoCtx)
		require.NoError(t, err)
	})
}

func TestBlogHandlers_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.LoadConfig("./../../../../config/config-local")
	require.NoError(t, err)

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()

	mockBlogUC := mock.NewMockUseCase(ctrl)
	blogUC := usecase.NewBlogUseCase(cfg, mockBlogUC, logger)
	blogHandler := NewBlogsHandlers(cfg, blogUC, logger)
	handler := blogHandler.Update()

	t.Run("Update succes case", func(t *testing.T) {
		blog := models.Blog{
			ID:      1,
			Title:   "title-test",
			Content: "content-test",
		}

		bufferData, err := utils.AnyToBytesBuffer(blog)
		require.NoError(t, err)
		require.NotNil(t, bufferData)

		request := httptest.NewRequest(http.MethodPost, "/v1/blogs/1", strings.NewReader(bufferData.String()))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		mockBlogUC.EXPECT().Update(gomock.Any(), &blog).Return(&blog, nil)

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("Update Validate error case", func(t *testing.T) {
		blog := models.Blog{
			ID:      1,
			Title:   "",
			Content: "",
		}

		bufferData, err := utils.AnyToBytesBuffer(blog)
		require.NoError(t, err)
		require.NotNil(t, bufferData)

		request := httptest.NewRequest(http.MethodPost, "/v1/blogs/1", strings.NewReader(bufferData.String()))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("Update ID Param error case", func(t *testing.T) {
		blog := models.Blog{
			ID:      1,
			Title:   "title-test",
			Content: "content-test",
		}

		bufferData, err := utils.AnyToBytesBuffer(blog)
		require.NoError(t, err)
		require.NotNil(t, bufferData)

		request := httptest.NewRequest(http.MethodPost, "/v1/blogs/1", strings.NewReader(bufferData.String()))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("test")

		err = handler(echoCtx)
		require.NoError(t, err)
	})

}

func TestBlogHandlers_GetByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.LoadConfig("./../../../../config/config-local")
	require.NoError(t, err)

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()
	mockBlogUC := mock.NewMockUseCase(ctrl)
	blogUC := usecase.NewBlogUseCase(cfg, mockBlogUC, logger)
	blogHandler := NewBlogsHandlers(cfg, blogUC, logger)
	handler := blogHandler.GetByID()

	t.Run("GetByID succes case", func(t *testing.T) {
		blog := models.Blog{
			ID:      1,
			Title:   "title-test",
			Content: "content-test",
		}

		request := httptest.NewRequest(http.MethodGet, "/v1/blogs/1", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		mockBlogUC.EXPECT().GetByID(gomock.Any(), int64(1)).Return(&blog, nil)

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("GetByID ID Param error case", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodGet, "/v1/blogs/1", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("test")

		err = handler(echoCtx)
		require.NoError(t, err)
	})

}

func TestBlogHandlers_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.LoadConfig("./../../../../config/config-local")
	require.NoError(t, err)

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()
	mockBlogUC := mock.NewMockUseCase(ctrl)
	blogUC := usecase.NewBlogUseCase(cfg, mockBlogUC, logger)
	blogHandler := NewBlogsHandlers(cfg, blogUC, logger)
	handler := blogHandler.Delete()

	t.Run("Delete succes case", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/v1/blogs/1", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		mockBlogUC.EXPECT().Delete(gomock.Any(), int64(1)).Return(nil)

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("Delete ID Param error case", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodDelete, "/v1/blogs/1", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("test")

		err = handler(echoCtx)
		require.NoError(t, err)
	})

}

func TestBlogHandlers_GetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.LoadConfig("./../../../../config/config-local")
	require.NoError(t, err)

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()
	mockBlogUC := mock.NewMockUseCase(ctrl)
	blogUC := usecase.NewBlogUseCase(cfg, mockBlogUC, logger)
	blogHandler := NewBlogsHandlers(cfg, blogUC, logger)
	handler := blogHandler.GetAll()

	t.Run("GetAll succes case", func(t *testing.T) {
		blogs := []*models.Blog{
			{
				ID:      1,
				Title:   "title-test",
				Content: "content-test",
			},
			{
				ID:      2,
				Title:   "title-test",
				Content: "content-test",
			},
		}

		request := httptest.NewRequest(http.MethodGet, "/v1/blogs?limit=10&page=1&search=&sort=asc", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)

		query := utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "asc",
		}

		mockBlogUC.EXPECT().GetAll(gomock.Any(), &query).Return(&models.BlogList{
			TotalCount: 2,
			TotalPage:  1,
			Page:       1,
			Limit:      10,
			HasMore:    false,
			Blogs:      blogs,
		}, nil)

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("GetAll QueryLimit error case", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodGet, "/v1/blogs?limit=test&page=1&search=&sort=asc", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)

		err = handler(echoCtx)
		require.NoError(t, err)
	})
}
