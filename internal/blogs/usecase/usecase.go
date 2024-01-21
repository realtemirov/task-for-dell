package usecase

import (
	"context"

	"github.com/realtemirov/task-for-dell/config"
	"github.com/realtemirov/task-for-dell/internal/blogs"
	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/pkg/logger"
	"github.com/realtemirov/task-for-dell/pkg/utils"
)

// Blogs Usecase
type blogUC struct {
	cfg  *config.Config
	repo blogs.Repository
	log  logger.Logger
}

// Blogs UseCase contructor
func NewBlogUseCase(cfg *config.Config, repo blogs.Repository, log logger.Logger) blogs.UseCase {
	return &blogUC{
		cfg:  cfg,
		repo: repo,
		log:  log,
	}
}

// Create implements blogs.UseCase.
func (u *blogUC) Create(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	return u.repo.Create(ctx, blog)
}

// Update implements blogs.UseCase.
func (u *blogUC) Update(ctx context.Context, blog *models.Blog) (*models.Blog, error) {
	return u.repo.Update(ctx, blog)
}

// Delete implements blogs.UseCase.
func (u *blogUC) Delete(ctx context.Context, blogID int64) error {
	return u.repo.Delete(ctx, blogID)
}

// GetAll implements blogs.UseCase.
func (u *blogUC) GetAll(ctx context.Context, query *utils.Query) (*models.BlogList, error) {
	return u.repo.GetAll(ctx, query)
}

// GetByID implements blogs.UseCase.
func (u *blogUC) GetByID(ctx context.Context, blogID int64) (*models.Blog, error) {
	return u.repo.GetByID(ctx, blogID)
}
