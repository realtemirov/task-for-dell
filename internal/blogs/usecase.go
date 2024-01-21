package blogs

import (
	"context"

	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/pkg/utils"
)

type UseCase interface {
	Create(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	Update(ctx context.Context, blog *models.Blog) (*models.Blog, error)
	Delete(ctx context.Context, blogID int64) error
	GetByID(ctx context.Context, blogID int64) (*models.Blog, error)
	GetAll(ctx context.Context, query *utils.Query) (*models.BlogList, error)
}
