package usecase

import (
	"context"
	"testing"

	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/internal/news/mock"
	"github.com/realtemirov/task-for-dell/pkg/logger"
	"github.com/realtemirov/task-for-dell/pkg/utils"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"
)

func TestNewUC_Create(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of new
	logger := logger.NewApiLogger(nil)
	mockNewRepo := mock.NewMockRepository(ctrl)
	newUC := NewNewsUseCase(nil, mockNewRepo, logger)

	// model of new
	new := models.New{}

	// context
	ctx := context.Background()

	// mock the Create method of the repository
	mockNewRepo.EXPECT().Create(
		ctx,
		gomock.Eq(&new),
	).Return(&new, nil)

	// call the Create method of the usecase
	createdNew, err := newUC.Create(context.Background(), &new)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, createdNew)
}

func TestNewUC_Update(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of new
	logger := logger.NewApiLogger(nil)
	mockNewRepo := mock.NewMockRepository(ctrl)
	newUC := NewNewsUseCase(nil, mockNewRepo, logger)

	// model of new
	new := models.New{
		ID:    1,
		Title: "update-title",
	}

	// mock the Update method of the repository
	mockNewRepo.EXPECT().Update(
		context.Background(),
		gomock.Eq(&new),
	).Return(&new, nil)

	// call the Update method of the usecase
	updatedNew, err := newUC.Update(context.Background(), &new)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, updatedNew)
}

func TestNewUC_Delete(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of new
	logger := logger.NewApiLogger(nil)
	mockNewRepo := mock.NewMockRepository(ctrl)
	newUC := NewNewsUseCase(nil, mockNewRepo, logger)

	// new id
	newID := int64(1)

	// mock the Delete method of the repository
	mockNewRepo.EXPECT().Delete(
		context.Background(),
		gomock.Eq(newID),
	).Return(nil)

	// call the Delete method of the usecase
	err := newUC.Delete(context.Background(), newID)

	// check the result
	require.NoError(t, err)
}

func TestNewUC_GetByID(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of new
	logger := logger.NewApiLogger(nil)
	mockNewRepo := mock.NewMockRepository(ctrl)
	newUC := NewNewsUseCase(nil, mockNewRepo, logger)

	// new id
	newID := int64(1)

	// mock the GetByID method of the repository
	mockNewRepo.EXPECT().GetByID(
		context.Background(),
		gomock.Eq(newID),
	).Return(&models.New{}, nil)

	// call the GetByID method of the usecase
	new, err := newUC.GetByID(context.Background(), newID)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, new)
}

func TestNewUC_GetAll(t *testing.T) {
	t.Parallel()

	// Create a new instance of the gomock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// logger, repository, usecase of new
	logger := logger.NewApiLogger(nil)
	mockNewRepo := mock.NewMockRepository(ctrl)
	newUC := NewNewsUseCase(nil, mockNewRepo, logger)

	// entity of NEW list, context, query
	entity := models.NewsList{}
	ctx := context.Background()
	query := utils.Query{
		Page:   1,
		Limit:  10,
		Search: "",
		Sort:   "",
	}

	// mock the GetAll method of the repository
	mockNewRepo.EXPECT().GetAll(
		ctx,
		&query,
	).Return(&entity, nil)

	// call the GetAll method of the usecase
	newList, err := newUC.GetAll(ctx, &query)

	// check the result
	require.NoError(t, err)
	require.NotNil(t, newList)
}
