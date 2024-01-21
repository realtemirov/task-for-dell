package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/internal/news"
	"github.com/realtemirov/task-for-dell/pkg/utils"
)

type newsRepo struct {
	db *sqlx.DB
}

// NewNewsRepository constructor
func NewNewsRepository(db *sqlx.DB) news.Repository {
	return &newsRepo{db: db}
}

// Create implements news.Repository.
func (r *newsRepo) Create(ctx context.Context, new *models.New) (*models.New, error) {

	// result for response
	result := models.New{}

	// insert entitiy and scan result
	if err := r.db.QueryRowxContext(
		ctx,
		createQuery,
		&new.Title,
		&new.Content,
	).StructScan(&result); err != nil {
		return nil, errors.Wrap(err, "newsRepo.Create.StructScan")
	}

	// if no error, return result
	return &result, nil
}

// Update implements news.Repository.
func (r *newsRepo) Update(ctx context.Context, new *models.New) (*models.New, error) {

	// response result
	result := models.New{}

	// update entity and scan result
	if err := r.db.QueryRowxContext(
		ctx,
		updateQuery,
		&new.Title,
		&new.Content,
		&new.ID,
	).StructScan(&result); err != nil {
		return nil, errors.Wrap(err, "newsRepo.Update.StructScan")
	}

	// if no error, return result
	return &result, nil
}

// Delete implements news.Repository.
func (r *newsRepo) Delete(ctx context.Context, newID int64) error {

	// delete entity and return result
	result, err := r.db.ExecContext(ctx, deleteQuery, newID)
	if err != nil {
		return errors.Wrap(err, "newsRepo.Delete.ExecContext")
	}

	// if didn't rows affected, return error
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "newsRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "newsRepo.Delete.RowsAffected")
	}

	return nil
}

// GetByID implements news.Repository.
func (r *newsRepo) GetByID(ctx context.Context, newsID int64) (*models.New, error) {

	result := models.New{}

	// get entity by id and scan result
	if err := r.db.QueryRowxContext(
		ctx,
		getByIDQuery,
		newsID,
	).StructScan(&result); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetByID.StructScan")
	}

	// if no error, return result
	return &result, nil
}

// GetAll implements news.Repository.
func (r *newsRepo) GetAll(ctx context.Context, query *utils.Query) (*models.NewsList, error) {

	var (

		// total news count
		totalCount      int
		totalCountQuery string = getTotalCountQuery
		allNewsQuery    string = getAllQuery
	)

	// if search query is empty, get all news
	if query.Search != "" {

		// change query for get total count
		totalCountQuery = fmt.Sprintf("%s%s", getTotalCountQuery, " AND title LIKE '%"+query.Search+"%' ")

		// change query for get all blogs
		allNewsQuery = fmt.Sprintf(`%s %s`, getAllQuery, " AND title LIKE '%"+query.Search+"%' ")
	}

	// change query for get all news, sort by created_at, add offset and limit
	allNewsQuery += fmt.Sprintf(" ORDER BY created_at %s OFFSET $1 LIMIT $2", query.GetSort())

	// get total count and scan result
	if err := r.db.QueryRowContext(
		ctx,
		totalCountQuery,
	).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetAll.QueryRowContext.Scan")
	}

	// if total count is 0, return empty list
	if totalCount == 0 {
		return &models.NewsList{
			TotalCount: totalCount,
			TotalPage:  utils.GetTotalPages(totalCount, query.GetLimit()),
			Page:       query.GetPage(),
			Limit:      query.GetLimit(),
			HasMore:    false,
			News:       make([]*models.New, 0),
		}, nil
	}

	// get all news
	rows, err := r.db.QueryxContext(
		ctx,
		allNewsQuery,
		query.GetOffset(),
		query.GetLimit(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetAll.QueryxContext")
	}
	defer rows.Close()

	// news list for response
	newsList := make([]*models.New, 0, query.GetLimit())

	// scan rows
	for rows.Next() {
		news := models.New{}
		if err := rows.StructScan(&news); err != nil {
			return nil, errors.Wrap(err, "newsRepo.GetAll.StructScan")
		}

		newsList = append(newsList, &news)
	}

	// if error, return error
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "newsRepo.GetAll.rows.Err")
	}

	// if no error, return result
	return &models.NewsList{
		TotalCount: totalCount,
		TotalPage:  utils.GetTotalPages(totalCount, query.GetLimit()),
		Page:       query.GetPage(),
		Limit:      query.GetLimit(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetLimit()),
		News:       newsList,
	}, nil
}
