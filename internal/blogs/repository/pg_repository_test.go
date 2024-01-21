package repository

import (
	"context"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/realtemirov/task-for-dell/internal/models"
	"github.com/realtemirov/task-for-dell/pkg/utils"
	"github.com/stretchr/testify/require"
)

// TestBlogRepo_Create tests Create method.
func TestBlogRepo_Create(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// blog repository
	repo := NewBlogsRepository(sqlxDB)

	// Create blog success case
	t.Run("Create", func(t *testing.T) {

		// temprorary blog
		blog := &models.Blog{
			ID:      1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			blog.ID,
			blog.Title,
			blog.Content,
		)

		// mock query with args and return rows
		mock.ExpectQuery(createQuery).WithArgs(
			blog.Title,
			blog.Content,
		).WillReturnRows(rows)

		// call Create method
		createdBlog, err := repo.Create(context.Background(), blog)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, createdBlog)
		require.Equal(t, blog.ID, createdBlog.ID)
		require.Equal(t, blog.Title, createdBlog.Title)
		require.Equal(t, blog.Content, createdBlog.Content)
	})

	// Create blog error case
	t.Run("Create Error", func(t *testing.T) {

		// temprorary blog
		blog := &models.Blog{
			ID:      1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock query with args and return error
		mock.ExpectQuery(createQuery).WithArgs(
			blog.ID,
			blog.Title,
			blog.Content,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Create method
		createdBlog, err := repo.Create(context.Background(), blog)

		// check error and result
		require.Error(t, err)
		require.Nil(t, createdBlog)
	})
}

// TestBlogRepo_Update tests Update method.
func TestBlogRepo_Update(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// blog repository
	repo := NewBlogsRepository(sqlxDB)

	// Update blog success case
	t.Run("Update", func(t *testing.T) {

		// temprorary blog
		blog := &models.Blog{
			ID:      1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			blog.ID,
			blog.Title,
			blog.Content,
		)

		// mock query with args and return rows
		mock.ExpectQuery(updateQuery).WithArgs(
			blog.Title,
			blog.Content,
			blog.ID,
		).WillReturnRows(rows)

		// call Update method
		updatedBlog, err := repo.Update(context.Background(), blog)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, updatedBlog)
		require.Equal(t, blog.ID, updatedBlog.ID)
		require.Equal(t, blog.Title, updatedBlog.Title)
		require.Equal(t, blog.Content, updatedBlog.Content)
	})

	// Update blog error case
	t.Run("Update Error", func(t *testing.T) {

		// temprorary blog
		blog := &models.Blog{
			ID:      1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock query with args and return error
		mock.ExpectQuery(updateQuery).WithArgs(
			blog.Title,
			blog.Content,
			blog.ID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Update method
		updatedBlog, err := repo.Update(context.Background(), blog)

		// check error and result
		require.Error(t, err)
		require.Nil(t, updatedBlog)
	})
}

// TestBlogRepo_Delete tests Delete method.
func TestBlogRepo_Delete(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// blog repository
	repo := NewBlogsRepository(sqlxDB)

	// Delete blog success case
	t.Run("Delete", func(t *testing.T) {

		// delete blog id
		blogID := int64(1)

		// mock query with args and return result
		mock.ExpectExec(deleteQuery).WithArgs(
			blogID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		// call Delete method
		err := repo.Delete(context.Background(), blogID)

		// check error
		require.NoError(t, err)
	})

	// Delete blog error case
	t.Run("Delete Error", func(t *testing.T) {

		// delete blog id
		blogID := int64(1)

		// mock query with args and return error
		mock.ExpectExec(deleteQuery).WithArgs(
			blogID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Delete method
		err := repo.Delete(context.Background(), blogID)

		// check error
		require.Error(t, err)
	})

	// Delete blog RowsAffected equal to zero case
	t.Run("Delete RowsAffected equal to zero", func(t *testing.T) {

		// delete blog id
		blogID := int64(1)

		// mock query with args and return result, but rows affected equal to zero
		mock.ExpectExec(deleteQuery).WithArgs(
			blogID,
		).WillReturnResult(sqlmock.NewResult(1, 0))

		// call Delete method
		err := repo.Delete(context.Background(), blogID)

		// check error
		require.Error(t, err)
	})

	// Delete blog RowsAffected error case
	t.Run("Delete RowsAffected Error", func(t *testing.T) {

		// delete blog id
		blogID := int64(1)

		// mock query with args and return error which rows affected
		mock.ExpectExec(deleteQuery).WithArgs(
			blogID,
		).WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("rows affected error")))

		// call Delete method
		err := repo.Delete(context.Background(), blogID)

		// check error
		require.Error(t, err)
	})
}

// TestBlogRepo_GetByID tests GetByID method.
func TestBlogRepo_GetByID(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// blog repository
	repo := NewBlogsRepository(sqlxDB)

	// GetByID success case
	t.Run("GetByID", func(t *testing.T) {

		// blog id
		blogID := int64(1)

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			blogID,
			"test-title",
			"test-content",
		)

		// mock query with args and return rows
		mock.ExpectQuery(getByIDQuery).WithArgs(
			blogID,
		).WillReturnRows(rows)

		// call GetByID method
		blog, err := repo.GetByID(context.Background(), blogID)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, blog)
		require.Equal(t, blogID, blog.ID)
		require.Equal(t, "test-title", blog.Title)
		require.Equal(t, "test-content", blog.Content)
	})

	// GetByID error case
	t.Run("GetByID Error", func(t *testing.T) {

		// blog id
		blogID := int64(1)

		// mock query with args and return error
		mock.ExpectQuery(getByIDQuery).WithArgs(
			blogID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call GetByID method
		blog, err := repo.GetByID(context.Background(), blogID)

		// check error and result
		require.Error(t, err)
		require.Nil(t, blog)
	})
}

// TestBlogRepo_GetAll tests GetAll method.
func TestBlogRepo_GetAll(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// blog repository
	repo := NewBlogsRepository(sqlxDB)

	// GetAll success case, without search
	t.Run("GetAll", func(t *testing.T) {

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			int64(1),
			"test-title",
			"test-content",
		).AddRow(
			int64(1),
			"test-title",
			"test-content",
		)
		query := utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
		}

		// mock query for get total count
		mock.ExpectQuery(getTotalCountQuery).
			WillReturnRows(
				sqlmock.NewRows([]string{"count"}).AddRow(2),
			)

		// mock query with args and return rows
		mock.ExpectQuery(
			fmt.Sprintf("%s ORDER BY created_by $1 OFFSET $2 LIMIT $3", getAllQuery),
		).WithArgs(
			query.GetSort(),
			query.GetOffset(),
			query.GetLimit(),
		).WillReturnRows(rows)

		// call GetAll method
		blogs, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "",
		})

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, blogs)
		require.Len(t, blogs.Blogs, 2)
	})

	// GetAll success case, with search
	t.Run("GetAll Search", func(t *testing.T) {

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			int64(1),
			"test-title",
			"test-content",
		).AddRow(
			int64(1),
			"test-title",
			"test-content",
		)

		// mock query
		query := utils.Query{
			Limit:  10,
			Page:   1,
			Search: "test",
		}

		// mock query for get total count, with search
		mock.ExpectQuery(
			fmt.Sprintf("%s%s", getTotalCountQuery, " AND title LIKE '%"+query.Search+"%' "),
		).WillReturnRows(
			sqlmock.NewRows([]string{"count"}).AddRow(2),
		)

		// mock query with args and return rows, with search
		mock.ExpectQuery(
			fmt.Sprintf(
				"%s ORDER BY created_by $1 OFFSET $2 LIMIT $3",
				fmt.Sprintf(`%s%s`, getAllQuery, " AND title LIKE '%"+query.Search+"%' ")),
		).WithArgs(
			query.GetSort(),
			query.GetOffset(),
			query.GetLimit(),
		).WillReturnRows(rows)

		// call GetAll method
		blogs, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "test",
			Sort:   "",
		})

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, blogs)
		require.Len(t, blogs.Blogs, 2)
	})

	// GetAll TotalCount equal to zero case
	t.Run("GetAll TotalCount equal to zero", func(t *testing.T) {

		// mock query for get total count zero
		mock.ExpectQuery(getTotalCountQuery).WillReturnRows(
			sqlmock.NewRows([]string{"count"}).AddRow(0),
		)

		// call GetAll method
		blogs, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "",
		})

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, blogs)
		require.Len(t, blogs.Blogs, 0)

	})

	// GetAll TotalCount error case
	t.Run("GetAll TotalCount Error", func(t *testing.T) {

		// mock query for get total count error
		query := utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
		}

		// mock query for get total count error
		mock.ExpectQuery(getTotalCountQuery).WillReturnError(sqlmock.ErrCancelled)

		// call GetAll method
		blogs, err := repo.GetAll(context.Background(), &query)

		// check error and result
		require.Error(t, err)
		require.Nil(t, blogs)
	})

	// GetAll rows error case
	t.Run("GetAll rows Error", func(t *testing.T) {

		// mock rows and error
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			int64(1),
			"test-title",
			"test-content",
		).AddRow(
			int64(1),
			"test-title",
			"test-content",
		).RowError(1, fmt.Errorf("row error"))

		// mock query for get total count
		query := utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
		}
		mock.ExpectQuery(getTotalCountQuery).
			WillReturnRows(
				sqlmock.NewRows([]string{"count"}).AddRow(2),
			)

		// mock query with args and return rows
		mock.ExpectQuery(
			fmt.Sprintf("%s ORDER BY created_by $1 OFFSET $2 LIMIT $3", getAllQuery),
		).WithArgs(
			query.GetSort(),
			query.GetOffset(),
			query.GetLimit(),
		).WillReturnRows(rows)

		// call GetAll method
		blogs, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "",
		})

		// check error and result
		require.Error(t, err)
		require.Nil(t, blogs)

	})

	// GetAll rows.Scan error case
	t.Run("GetAll rows.Scan Error", func(t *testing.T) {

		// mock rows and error
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(interface{}(nil), interface{}(nil), interface{}(nil))

		// mock query for get total count
		query := utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
		}
		mock.ExpectQuery(getTotalCountQuery).
			WillReturnRows(
				sqlmock.NewRows([]string{"count"}).AddRow(2),
			)

		// mock query with args and return rows
		mock.ExpectQuery(
			fmt.Sprintf("%s ORDER BY created_by $1 OFFSET $2 LIMIT $3", getAllQuery),
		).WithArgs(
			query.GetSort(),
			query.GetOffset(),
			query.GetLimit(),
		).WillReturnRows(rows)

		// call GetAll method
		blogs, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "",
		})

		// check error and result
		require.Error(t, err)
		require.Nil(t, blogs)

	})

	// GetAll error case
	t.Run("GetAll Error", func(t *testing.T) {

		// mock query for get total count
		query := utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
		}
		mock.ExpectQuery(getTotalCountQuery).
			WillReturnRows(
				sqlmock.NewRows([]string{"count"}).AddRow(2),
			)

		// mock query with args and return error
		mock.ExpectQuery(
			fmt.Sprintf("%s ORDER BY created_by $1 OFFSET $2 LIMIT $3", getAllQuery),
		).WithArgs(
			query.GetSort(),
			query.GetOffset(),
			query.GetLimit(),
		).WillReturnError(sqlmock.ErrCancelled)

		// call GetAll method
		blogs, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "",
		})

		// check error and result
		require.Error(t, err)
		require.Nil(t, blogs)

	})
}
