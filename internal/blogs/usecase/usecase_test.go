package usecase

import (
	"context"
	"testing"

	"github.com/realtemirov/task-for-dell/internal/blogs/mock"
	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/pkg/logger"
	"github.com/realtemirov/task-for-dell/pkg/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestBlofUC_Create(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of blog
	logger := logger.NewApiLogger(nil)
	mockBlogRepo := mock.NewMockRepository(ctrl)
	blogUC := NewBlogUseCase(nil, mockBlogRepo, logger)

	// model of blog
	blog := models.Blog{}

	// context
	ctx := context.Background()

	// mock the Create method of the repository
	mockBlogRepo.EXPECT().Create(
		ctx,
		gomock.Eq(&blog),
	).Return(&blog, nil)

	// call the Create method of the usecase
	createdBlog, err := blogUC.Create(context.Background(), &blog)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, createdBlog)
}

func TestBlofUC_Update(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of blog
	logger := logger.NewApiLogger(nil)
	mockBlogRepo := mock.NewMockRepository(ctrl)
	blogUC := NewBlogUseCase(nil, mockBlogRepo, logger)

	// model of blog
	blog := models.Blog{
		ID:    1,
		Title: "update-title",
	}

	// mock the Update method of the repository
	mockBlogRepo.EXPECT().Update(
		context.Background(),
		gomock.Eq(&blog),
	).Return(&blog, nil)

	// call the Update method of the usecase
	updatedBlog, err := blogUC.Update(context.Background(), &blog)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, updatedBlog)
}

func TestBlofUC_Delete(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of blog
	logger := logger.NewApiLogger(nil)
	mockBlogRepo := mock.NewMockRepository(ctrl)
	blogUC := NewBlogUseCase(nil, mockBlogRepo, logger)

	// blog id
	blogID := int64(1)

	// mock the Delete method of the repository
	mockBlogRepo.EXPECT().Delete(
		context.Background(),
		gomock.Eq(blogID),
	).Return(nil)

	// call the Delete method of the usecase
	err := blogUC.Delete(context.Background(), blogID)

	// check the result
	require.NoError(t, err)
}

func TestBlofUC_GetByID(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of blog
	logger := logger.NewApiLogger(nil)
	mockBlogRepo := mock.NewMockRepository(ctrl)
	blogUC := NewBlogUseCase(nil, mockBlogRepo, logger)

	// blog id
	blogID := int64(1)

	// mock the GetByID method of the repository
	mockBlogRepo.EXPECT().GetByID(
		context.Background(),
		gomock.Eq(blogID),
	).Return(&models.Blog{}, nil)

	// call the GetByID method of the usecase
	blog, err := blogUC.GetByID(context.Background(), blogID)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, blog)
}

func TestBlofUC_GetAll(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of blog
	logger := logger.NewApiLogger(nil)
	mockBlogRepo := mock.NewMockRepository(ctrl)
	blogUC := NewBlogUseCase(nil, mockBlogRepo, logger)

	// entity of blog list, context, query
	entity := models.BlogList{}
	ctx := context.Background()
	query := utils.Query{
		Page:   1,
		Limit:  10,
		Search: "",
		Sort:   "",
	}

	// mock the GetAll method of the repository
	mockBlogRepo.EXPECT().GetAll(
		ctx,
		&query,
	).Return(&entity, nil)

	// call the GetAll method of the usecase
	blogList, err := blogUC.GetAll(ctx, &query)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, blogList)
}
