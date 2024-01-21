package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/realtemirov/task-for-dell/internal/blogs"
	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/pkg/utils"
)

type blogsRepo struct {
	db *sqlx.DB
}

// NewBlogsRepository constructor
func NewBlogsRepository(db *sqlx.DB) blogs.Repository {
	return &blogsRepo{db: db}
}

// Create implements blogs.Repository.
func (r *blogsRepo) Create(ctx context.Context, blog *models.Blog) (*models.Blog, error) {

	// result for response
	result := models.Blog{}

	// insert entitiy and scan result
	if err := r.db.QueryRowxContext(
		ctx,
		createQuery,
		&blog.Title,
		&blog.Content,
	).StructScan(&result); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.Create.StructScan")
	}

	// if no error, return result
	return &result, nil
}

// Update implements blogs.Repository.
func (r *blogsRepo) Update(ctx context.Context, blog *models.Blog) (*models.Blog, error) {

	// response result
	result := models.Blog{}

	// update entity and scan result
	if err := r.db.QueryRowxContext(
		ctx,
		updateQuery,
		&blog.Title,
		&blog.Content,
		&blog.ID,
	).StructScan(&result); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.Update.StructScan")
	}

	// if no error, return result
	return &result, nil
}

// Delete implements blogs.Repository.
func (r *blogsRepo) Delete(ctx context.Context, blogID int64) error {

	// delete entity and return result
	result, err := r.db.ExecContext(ctx, deleteQuery, blogID)
	if err != nil {
		return errors.Wrap(err, "blogsRepo.Delete.ExecContext")
	}

	// if didn't rows affected, return error
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "blogsRepo.Delete.RowsAffected")
	}

	if rowsAffected == 0 {
		return errors.Wrap(sql.ErrNoRows, "blogsRepo.Delete.RowsAffected")
	}

	return nil
}

// GetByID implements blogs.Repository.
func (r *blogsRepo) GetByID(ctx context.Context, blogID int64) (*models.Blog, error) {

	result := models.Blog{}

	// get entity by id and scan result
	if err := r.db.QueryRowxContext(
		ctx,
		getByIDQuery,
		blogID,
	).StructScan(&result); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetByID.StructScan")
	}

	// if no error, return result
	return &result, nil
}

// GetAll implements blogs.Repository.
func (r *blogsRepo) GetAll(ctx context.Context, query *utils.Query) (*models.BlogList, error) {

	var (

		// total blogs count
		totalCount      int
		totalCountQuery string = getTotalCountQuery
		allBlogsQuery   string = getAllQuery
	)

	// if search query is empty, get all blogs
	if query.Search != "" {

		// change query for get total count
		totalCountQuery = fmt.Sprintf("%s%s", getTotalCountQuery, " AND title LIKE '%"+query.Search+"%' ")

		// change query for get all blogs
		allBlogsQuery = fmt.Sprintf(`%s %s`, getAllQuery, " AND title LIKE '%"+query.Search+"%' ")
	}

	// change query for get all blogs, sort by created_at, add offset and limit
	allBlogsQuery += fmt.Sprintf(" ORDER BY created_at %s OFFSET $1 LIMIT $2", query.GetSort())

	// get total count and scan result
	if err := r.db.QueryRowContext(
		ctx,
		totalCountQuery,
	).Scan(&totalCount); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetAll.QueryRowContext.Scan")
	}

	// if total count is 0, return empty list
	if totalCount == 0 {
		return &models.BlogList{
			TotalCount: totalCount,
			TotalPage:  utils.GetTotalPages(totalCount, query.GetLimit()),
			Page:       query.GetPage(),
			Limit:      query.GetLimit(),
			HasMore:    false,
			Blogs:      make([]*models.Blog, 0),
		}, nil
	}

	// get all blogs
	rows, err := r.db.QueryxContext(
		ctx,
		allBlogsQuery,
		query.GetOffset(),
		query.GetLimit(),
	)
	if err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetAll.QueryxContext")
	}
	defer rows.Close()

	// blogs list for response
	blogs := make([]*models.Blog, 0, query.GetLimit())

	// scan rows
	for rows.Next() {
		blog := models.Blog{}
		if err := rows.StructScan(&blog); err != nil {
			return nil, errors.Wrap(err, "blogsRepo.GetAll.StructScan")
		}

		blogs = append(blogs, &blog)
	}

	// if error, return error
	if err = rows.Err(); err != nil {
		return nil, errors.Wrap(err, "blogsRepo.GetAll.rows.Err")
	}

	// if no error, return result
	return &models.BlogList{
		TotalCount: totalCount,
		TotalPage:  utils.GetTotalPages(totalCount, query.GetLimit()),
		Page:       query.GetPage(),
		Limit:      query.GetLimit(),
		HasMore:    utils.GetHasMore(query.GetPage(), totalCount, query.GetLimit()),
		Blogs:      blogs,
	}, nil
}
