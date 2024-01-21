package news

import (
	"context"

	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/pkg/utils"
)

type UseCase interface {
	Create(ctx context.Context, news *models.New) (*models.New, error)
	Update(ctx context.Context, news *models.New) (*models.New, error)
	Delete(ctx context.Context, newsID int64) error
	GetByID(ctx context.Context, newsID int64) (*models.New, error)
	GetAll(ctx context.Context, query *utils.Query) (*models.NewsList, error)
}
