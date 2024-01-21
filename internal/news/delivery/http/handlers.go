package http

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/realtemirov/task-for-dell/config"
	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/internal/news"

	"github.com/realtemirov/task-for-dell/pkg/httpErrors"
	"github.com/realtemirov/task-for-dell/pkg/logger"
	"github.com/realtemirov/task-for-dell/pkg/utils"
)

type newsHandlers struct {
	cfg    *config.Config
	newsUC news.UseCase
	logger logger.Logger
}

// NewNewsHandlers constructs a new newsHandlers.
func NewNewsHandlers(cfg *config.Config, newsUC news.UseCase, logger logger.Logger) news.Handlers {
	return &newsHandlers{
		cfg:    cfg,
		newsUC: newsUC,
		logger: logger,
	}
}

// Create
// @Summary Create new
// @Description Create new
// @Tags News
// @Accept  json
// @Produce  json
// @Param body body models.NewSwagger true "new"
// @Success 201 {object} models.New
// @Failure 400 {object} httpErrors.ErrorMessage
// @Failure 500 {object} httpErrors.ErrorMessage
// @Router /news [POST]
func (h *newsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		// model of new for create
		var (
			err        error
			new        *models.New = &models.New{}
			createdNew *models.New = &models.New{}
		)

		// bind request body to new
		if err = c.Bind(new); err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// validate new
		if err = utils.ValidateStruct(c.Request().Context(), new); err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// create new
		createdNew, err = h.newsUC.Create(c.Request().Context(), new)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// return created new
		return c.JSON(http.StatusCreated, createdNew)
	}
}

// Update
// @Summary Update
// @Description Update new
// @Tags News
// @Accept  json
// @Produce  json
// @Param id path int true "new_id"
// @Param body body models.NewSwagger true "body"
// @Success 200 {object} models.New
// @Failure 400 {object} httpErrors.ErrorMessage
// @Failure 404 {object} httpErrors.ErrorMessage
// @Failure 500 {object} httpErrors.ErrorMessage
// @Router /news/{id} [PUT]
func (h *newsHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		var (
			err        error
			newID      int64
			new        *models.New = &models.New{}
			updatedNew *models.New = &models.New{}
		)

		// get new's id from url
		newID, err = utils.StringToInt64(c.Param("id"))
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// bind request body to new
		if err = c.Bind(new); err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// validate new
		if err = utils.ValidateStruct(c.Request().Context(), new); err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}
		new.ID = newID

		// update new
		updatedNew, err = h.newsUC.Update(c.Request().Context(), new)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// return updated new
		return c.JSON(http.StatusOK, updatedNew)
	}
}

// Delete
// @Summary Delete
// @Description Delete new
// @Tags News
// @Accept  json
// @Produce  json
// @Param id path int true "new_id"
// @Success 204 "No Content"
// @Failure 400 {object} httpErrors.ErrorMessage
// @Failure 404 {object} httpErrors.ErrorMessage
// @Failure 500 {object} httpErrors.ErrorMessage
// @Router /news/{id} [DELETE]
func (h *newsHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		var (
			err   error
			newID int64
		)
		newID, err = utils.StringToInt64(c.Param("id"))
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		err = h.newsUC.Delete(c.Request().Context(), newID)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		return c.NoContent(http.StatusOK)
	}
}

// GetByID
// @Summary GetByID
// @Description Getting new by id
// @Tags News
// @Accept  json
// @Produce  json
// @Param id path int true "new_id"
// @Success 200 {object} models.New
// @Failure 400 {object} httpErrors.ErrorMessage
// @Failure 404 {object} httpErrors.ErrorMessage
// @Failure 500 {object} httpErrors.ErrorMessage
// @Router /news/{id} [GET]
func (h *newsHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		var (
			err   error
			newID int64
			new   *models.New = &models.New{}
		)

		newID, err = utils.StringToInt64(c.Param("id"))
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		new, err = h.newsUC.GetByID(c.Request().Context(), newID)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		return c.JSON(http.StatusOK, new)
	}
}

// GetAll
// @Summary GetAll
// @Description Get all new with pagination and search
// @Tags News
// @Accept  json
// @Produce  json
// @Param query query utils.Query true "query"
// @Success 200 {object} models.BlogList
// @Failure 400 {object} httpErrors.ErrorMessage
// @Failure 404 {object} httpErrors.ErrorMessage
// @Failure 500 {object} httpErrors.ErrorMessage
// @Router /news [GET]
func (h *newsHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		var (
			err      error
			query    *utils.Query     = &utils.Query{}
			newsList *models.NewsList = &models.NewsList{}
		)

		query, err = utils.GetPaginationFromCtx(c)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		newsList, err = h.newsUC.GetAll(c.Request().Context(), query)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		return c.JSON(http.StatusOK, newsList)
	}
}
