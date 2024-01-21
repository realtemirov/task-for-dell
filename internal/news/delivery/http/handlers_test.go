package http

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/realtemirov/task-for-dell/config"

	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/internal/news/mock"
	"github.com/realtemirov/task-for-dell/internal/news/usecase"
	"github.com/realtemirov/task-for-dell/pkg/logger"
	"github.com/realtemirov/task-for-dell/pkg/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestNewsHandlers_Create(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.LoadConfig("./../../../../config/config-local")
	require.NoError(t, err)

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()

	mockNewsUC := mock.NewMockUseCase(ctrl)
	newsUC := usecase.NewNewsUseCase(cfg, mockNewsUC, logger)
	newsHandler := NewNewsHandlers(cfg, newsUC, logger)
	handler := newsHandler.Create()

	t.Run("Create succes case", func(t *testing.T) {
		new := models.New{
			Title:   "title-test",
			Content: "content-test",
		}

		bufferData, err := utils.AnyToBytesBuffer(new)
		require.NoError(t, err)
		require.NotNil(t, bufferData)

		request := httptest.NewRequest(http.MethodPost, "/v1/news", strings.NewReader(bufferData.String()))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)

		mockNewsUC.EXPECT().Create(gomock.Any(), &new).Return(&new, nil)

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("Create Validate error case", func(t *testing.T) {

		new := models.New{
			Title:   "",
			Content: "",
		}

		bufferData, err := utils.AnyToBytesBuffer(new)
		require.NoError(t, err)
		require.NotNil(t, bufferData)

		request := httptest.NewRequest(http.MethodPost, "/v1/news", strings.NewReader(bufferData.String()))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		err = handler(echoCtx)
		require.NoError(t, err)
	})
}

func TestNewsHandlers_Update(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.LoadConfig("./../../../../config/config-local")
	require.NoError(t, err)

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()

	mockNewsUC := mock.NewMockUseCase(ctrl)
	newsUC := usecase.NewNewsUseCase(cfg, mockNewsUC, logger)
	newsHandler := NewNewsHandlers(cfg, newsUC, logger)
	handler := newsHandler.Update()

	t.Run("Update succes case", func(t *testing.T) {
		new := models.New{
			ID:      1,
			Title:   "title-test",
			Content: "content-test",
		}

		bufferData, err := utils.AnyToBytesBuffer(new)
		require.NoError(t, err)
		require.NotNil(t, bufferData)

		request := httptest.NewRequest(http.MethodPost, "/v1/news/1", strings.NewReader(bufferData.String()))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		mockNewsUC.EXPECT().Update(gomock.Any(), &new).Return(&new, nil)

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("Update Validate error case", func(t *testing.T) {
		new := models.New{
			ID:      1,
			Title:   "",
			Content: "",
		}

		bufferData, err := utils.AnyToBytesBuffer(new)
		require.NoError(t, err)
		require.NotNil(t, bufferData)

		request := httptest.NewRequest(http.MethodPost, "/v1/news/1", strings.NewReader(bufferData.String()))
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
		new := models.New{
			ID:      1,
			Title:   "title-test",
			Content: "content-test",
		}

		bufferData, err := utils.AnyToBytesBuffer(new)
		require.NoError(t, err)
		require.NotNil(t, bufferData)

		request := httptest.NewRequest(http.MethodPost, "/v1/news/1", strings.NewReader(bufferData.String()))
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

func TestNewsHandlers_GetByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.LoadConfig("./../../../../config/config-local")
	require.NoError(t, err)

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()
	mockNewsUC := mock.NewMockUseCase(ctrl)
	newsUC := usecase.NewNewsUseCase(cfg, mockNewsUC, logger)
	newsHandler := NewNewsHandlers(cfg, newsUC, logger)
	handler := newsHandler.GetByID()

	t.Run("GetByID succes case", func(t *testing.T) {
		new := models.New{
			ID:      1,
			Title:   "title-test",
			Content: "content-test",
		}

		request := httptest.NewRequest(http.MethodGet, "/v1/news/1", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		mockNewsUC.EXPECT().GetByID(gomock.Any(), int64(1)).Return(&new, nil)

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("GetByID ID Param error case", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodGet, "/v1/news/1", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("test")

		err = handler(echoCtx)
		require.NoError(t, err)
	})

}

func TestNewsHandlers_Delete(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.LoadConfig("./../../../../config/config-local")
	require.NoError(t, err)

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()
	mockNewsUC := mock.NewMockUseCase(ctrl)
	newsUC := usecase.NewNewsUseCase(cfg, mockNewsUC, logger)
	newsHandler := NewNewsHandlers(cfg, newsUC, logger)
	handler := newsHandler.Delete()

	t.Run("Delete succes case", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodDelete, "/v1/news/1", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("1")

		mockNewsUC.EXPECT().Delete(gomock.Any(), int64(1)).Return(nil)

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("Delete ID Param error case", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodDelete, "/v1/news/1", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)
		echoCtx.SetParamNames("id")
		echoCtx.SetParamValues("test")

		err = handler(echoCtx)
		require.NoError(t, err)
	})

}

func TestNewsHandlers_GetAll(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	cfg, err := config.LoadConfig("./../../../../config/config-local")
	require.NoError(t, err)

	logger := logger.NewApiLogger(cfg)
	logger.InitLogger()
	mockNewsUC := mock.NewMockUseCase(ctrl)
	newsUC := usecase.NewNewsUseCase(cfg, mockNewsUC, logger)
	newsHandler := NewNewsHandlers(cfg, newsUC, logger)
	handler := newsHandler.GetAll()

	t.Run("GetAll succes case", func(t *testing.T) {
		news := []*models.New{
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

		request := httptest.NewRequest(http.MethodGet, "/v1/news?limit=10&page=1&search=&sort=asc", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)

		query := utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "asc",
		}

		mockNewsUC.EXPECT().GetAll(gomock.Any(), &query).Return(&models.NewsList{
			TotalCount: 2,
			TotalPage:  1,
			Page:       1,
			Limit:      10,
			HasMore:    false,
			News:       news,
		}, nil)

		err = handler(echoCtx)
		require.NoError(t, err)
	})

	t.Run("GetAll QueryLimit error case", func(t *testing.T) {

		request := httptest.NewRequest(http.MethodGet, "/v1/news?limit=test&page=1&search=&sort=asc", nil)
		response := httptest.NewRecorder()

		e := echo.New()
		echoCtx := e.NewContext(request, response)

		err = handler(echoCtx)
		require.NoError(t, err)
	})
}
