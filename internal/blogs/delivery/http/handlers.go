package http

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/realtemirov/task-for-dell/config"
	"github.com/realtemirov/task-for-dell/internal/blogs"
	"github.com/realtemirov/task-for-dell/internal/models"

	"github.com/realtemirov/task-for-dell/pkg/httpErrors"
	"github.com/realtemirov/task-for-dell/pkg/logger"
	"github.com/realtemirov/task-for-dell/pkg/utils"
)

type blogsHandlers struct {
	cfg     *config.Config
	blogsUC blogs.UseCase
	logger  logger.Logger
}

// NewBlogsHandlers constructs a new blogsHandlers.
func NewBlogsHandlers(cfg *config.Config, blogsUC blogs.UseCase, logger logger.Logger) blogs.Handlers {
	return &blogsHandlers{
		cfg:     cfg,
		blogsUC: blogsUC,
		logger:  logger,
	}
}

// Create
// @Summary Create blog
// @Description Create blog
// @Tags Blogs
// @Accept  json
// @Produce  json
// @Param body body models.BlogSwagger true "blog"
// @Success 201 {object} models.Blog
// @Failure 400 {object} httpErrors.ErrorMessage
// @Failure 500 {object} httpErrors.ErrorMessage
// @Router /blogs [POST]
func (h *blogsHandlers) Create() echo.HandlerFunc {
	return func(c echo.Context) error {

		// model of blog for create
		var (
			err         error
			blog        *models.Blog = &models.Blog{}
			createdBlog *models.Blog = &models.Blog{}
		)

		// bind request body to blog
		if err = c.Bind(blog); err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// validate blog
		if err = utils.ValidateStruct(c.Request().Context(), blog); err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// create blog
		createdBlog, err = h.blogsUC.Create(c.Request().Context(), blog)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// return created blog
		return c.JSON(http.StatusCreated, createdBlog)
	}
}

// Update
// @Summary Update
// @Description Update blog
// @Tags Blogs
// @Accept  json
// @Produce  json
// @Param id path int true "blog_id"
// @Param body body models.BlogSwagger true "body"
// @Success 200 {object} models.Blog
// @Failure 400 {object} httpErrors.ErrorMessage
// @Failure 404 {object} httpErrors.ErrorMessage
// @Failure 500 {object} httpErrors.ErrorMessage
// @Router /blogs/{id} [PUT]
func (h *blogsHandlers) Update() echo.HandlerFunc {
	return func(c echo.Context) error {

		var (
			err         error
			blogID      int64
			blog        *models.Blog = &models.Blog{}
			updatedBlog *models.Blog = &models.Blog{}
		)

		// get blog id from url
		blogID, err = utils.StringToInt64(c.Param("id"))
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// bind request body to blog
		if err = c.Bind(blog); err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// validate blog
		if err = utils.ValidateStruct(c.Request().Context(), blog); err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}
		blog.ID = blogID

		// update blog
		updatedBlog, err = h.blogsUC.Update(c.Request().Context(), blog)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		// return updated blog
		return c.JSON(http.StatusOK, updatedBlog)
	}
}

// Delete
// @Summary Delete
// @Description Delete blog
// @Tags Blogs
// @Accept  json
// @Produce  json
// @Param id path int true "blog_id"
// @Success 204 "No Content"
// @Failure 400 {object} httpErrors.ErrorMessage
// @Failure 404 {object} httpErrors.ErrorMessage
// @Failure 500 {object} httpErrors.ErrorMessage
// @Router /blogs/{id} [DELETE]
func (h *blogsHandlers) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {

		var (
			err    error
			blogID int64
		)
		blogID, err = utils.StringToInt64(c.Param("id"))
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		err = h.blogsUC.Delete(c.Request().Context(), blogID)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		return c.NoContent(http.StatusNoContent)
	}
}

// GetByID
// @Summary GetByID
// @Description Getting blog by id
// @Tags Blogs
// @Accept  json
// @Produce  json
// @Param id path int true "blog_id"
// @Success 200 {object} models.Blog
// @Failure 400 {object} httpErrors.ErrorMessage
// @Failure 404 {object} httpErrors.ErrorMessage
// @Failure 500 {object} httpErrors.ErrorMessage
// @Router /blogs/{id} [GET]
func (h *blogsHandlers) GetByID() echo.HandlerFunc {
	return func(c echo.Context) error {

		var (
			err    error
			blogID int64
			blog   *models.Blog = &models.Blog{}
		)

		blogID, err = utils.StringToInt64(c.Param("id"))
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		blog, err = h.blogsUC.GetByID(c.Request().Context(), blogID)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		return c.JSON(http.StatusOK, blog)
	}
}

// GetAll
// @Summary GetAll
// @Description Get all blogs with pagination and search
// @Tags Blogs
// @Accept  json
// @Produce  json
// @Param query query utils.Query true "query"
// @Success 200 {object} models.BlogList
// @Failure 400 {object} httpErrors.ErrorMessage
// @Failure 404 {object} httpErrors.ErrorMessage
// @Failure 500 {object} httpErrors.ErrorMessage
// @Router /blogs [GET]
func (h *blogsHandlers) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {

		var (
			err      error
			query    *utils.Query     = &utils.Query{}
			blogList *models.BlogList = &models.BlogList{}
		)

		query, err = utils.GetPaginationFromCtx(c)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		blogList, err = h.blogsUC.GetAll(c.Request().Context(), query)
		if err != nil {
			return httpErrors.ErrResponseWithLog(c, h.logger, err)
		}

		return c.JSON(http.StatusOK, blogList)
	}
}
