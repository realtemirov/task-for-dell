package usecase

import (
	"context"

	"github.com/realtemirov/task-for-dell/config"
	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/internal/news"
	"github.com/realtemirov/task-for-dell/pkg/logger"
	"github.com/realtemirov/task-for-dell/pkg/utils"
)

// News Usecase
type newsUC struct {
	cfg  *config.Config
	repo news.Repository
	log  logger.Logger
}

// News UseCase contructor
func NewNewsUseCase(cfg *config.Config, repo news.Repository, log logger.Logger) news.UseCase {
	return &newsUC{
		cfg:  cfg,
		repo: repo,
		log:  log,
	}
}

// Create implements news.UseCase.
func (u *newsUC) Create(ctx context.Context, news *models.New) (*models.New, error) {
	return u.repo.Create(ctx, news)
}

// Update implements news.UseCase.
func (u *newsUC) Update(ctx context.Context, news *models.New) (*models.New, error) {
	return u.repo.Update(ctx, news)
}

// Delete implements news.UseCase.
func (u *newsUC) Delete(ctx context.Context, newsID int64) error {
	return u.repo.Delete(ctx, newsID)
}

// GetAll implements news.UseCase.
func (u *newsUC) GetAll(ctx context.Context, query *utils.Query) (*models.NewsList, error) {
	return u.repo.GetAll(ctx, query)
}

// GetByID implements news.UseCase.
func (u *newsUC) GetByID(ctx context.Context, newsID int64) (*models.New, error) {
	return u.repo.GetByID(ctx, newsID)
}
