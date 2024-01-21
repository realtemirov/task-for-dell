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

// TestNewRepo_Create tests Create method.
func TestNewRepo_Create(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// new's repository
	repo := NewNewsRepository(sqlxDB)

	// Create new success case
	t.Run("Create", func(t *testing.T) {

		// temprorary new
		new := &models.New{
			ID:      1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			new.ID,
			new.Title,
			new.Content,
		)

		// mock query with args and return rows
		mock.ExpectQuery(createQuery).WithArgs(
			new.Title,
			new.Content,
		).WillReturnRows(rows)

		// call Create method
		createdNew, err := repo.Create(context.Background(), new)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, createdNew)
		require.Equal(t, new.ID, createdNew.ID)
		require.Equal(t, new.Title, createdNew.Title)
		require.Equal(t, new.Content, createdNew.Content)
	})

	// Create New error case
	t.Run("Create Error", func(t *testing.T) {

		// temprorary new
		new := &models.New{
			ID:      1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock query with args and return error
		mock.ExpectQuery(createQuery).WithArgs(
			new.ID,
			new.Title,
			new.Content,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Create method
		createdNew, err := repo.Create(context.Background(), new)

		// check error and result
		require.Error(t, err)
		require.Nil(t, createdNew)
	})
}

// TestNewRepo_Update tests Update method.
func TestNewRepo_Update(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// new repository
	repo := NewNewsRepository(sqlxDB)

	// Update new success case
	t.Run("Update", func(t *testing.T) {

		// temprorary new
		new := &models.New{
			ID:      1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			new.ID,
			new.Title,
			new.Content,
		)

		// mock query with args and return rows
		mock.ExpectQuery(updateQuery).WithArgs(
			new.Title,
			new.Content,
			new.ID,
		).WillReturnRows(rows)

		// call Update method
		updatedNew, err := repo.Update(context.Background(), new)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, updatedNew)
		require.Equal(t, new.ID, updatedNew.ID)
		require.Equal(t, new.Title, updatedNew.Title)
		require.Equal(t, new.Content, updatedNew.Content)
	})

	// Update New error case
	t.Run("Update Error", func(t *testing.T) {

		// temprorary New
		new := &models.New{
			ID:      1,
			Title:   "test-title",
			Content: "test-content",
		}

		// mock query with args and return error
		mock.ExpectQuery(updateQuery).WithArgs(
			new.Title,
			new.Content,
			new.ID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Update method
		updatedNew, err := repo.Update(context.Background(), new)

		// check error and result
		require.Error(t, err)
		require.Nil(t, updatedNew)
	})
}

// TestNewRepo_Delete tests Delete method.
func TestNewRepo_Delete(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// new repository
	repo := NewNewsRepository(sqlxDB)

	// Delete newNew success case
	t.Run("Delete", func(t *testing.T) {

		// delete new id
		newID := int64(1)

		// mock query with args and return result
		mock.ExpectExec(deleteQuery).WithArgs(
			newID,
		).WillReturnResult(sqlmock.NewResult(1, 1))

		// call Delete method
		err := repo.Delete(context.Background(), newID)

		// check error
		require.NoError(t, err)
	})

	// Delete new error case
	t.Run("Delete Error", func(t *testing.T) {

		// delete new id
		newID := int64(1)

		// mock query with args and return error
		mock.ExpectExec(deleteQuery).WithArgs(
			newID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call Delete method
		err := repo.Delete(context.Background(), newID)

		// check error
		require.Error(t, err)
	})

	// Delete new RowsAffected equal to zero case
	t.Run("Delete RowsAffected equal to zero", func(t *testing.T) {

		// delete new id
		newID := int64(1)

		// mock query with args and return result, but rows affected equal to zero
		mock.ExpectExec(deleteQuery).WithArgs(
			newID,
		).WillReturnResult(sqlmock.NewResult(1, 0))

		// call Delete method
		err := repo.Delete(context.Background(), newID)

		// check error
		require.Error(t, err)
	})

	// Delete new RowsAffected error case
	t.Run("Delete RowsAffected Error", func(t *testing.T) {

		// delete new id
		newID := int64(1)

		// mock query with args and return error which rows affected
		mock.ExpectExec(deleteQuery).WithArgs(
			newID,
		).WillReturnResult(sqlmock.NewErrorResult(fmt.Errorf("rows affected error")))

		// call Delete method
		err := repo.Delete(context.Background(), newID)

		// check error
		require.Error(t, err)
	})
}

// TestNewRepo_GetByID tests GetByID method.
func TestNewRepo_GetByID(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// new repository
	repo := NewNewsRepository(sqlxDB)

	// GetByID success case
	t.Run("GetByID", func(t *testing.T) {

		// new id
		newID := int64(1)

		// mock rows
		rows := sqlmock.NewRows(
			[]string{"id", "title", "content"},
		).AddRow(
			newID,
			"test-title",
			"test-content",
		)

		// mock query with args and return rows
		mock.ExpectQuery(getByIDQuery).WithArgs(
			newID,
		).WillReturnRows(rows)

		// call GetByID method
		new, err := repo.GetByID(context.Background(), newID)

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, new)
		require.Equal(t, newID, new.ID)
		require.Equal(t, "test-title", new.Title)
		require.Equal(t, "test-content", new.Content)
	})

	// GetByID error case
	t.Run("GetByID Error", func(t *testing.T) {

		// new id
		newID := int64(1)

		// mock query with args and return error
		mock.ExpectQuery(getByIDQuery).WithArgs(
			newID,
		).WillReturnError(sqlmock.ErrCancelled)

		// call GetByID method
		new, err := repo.GetByID(context.Background(), newID)

		// check error and result
		require.Error(t, err)
		require.Nil(t, new)
	})
}

// TestNewRepo_GetAll tests GetAll method.
func TestNewRepo_GetAll(t *testing.T) {
	t.Parallel()

	// create mock db
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	require.NoError(t, err)
	defer db.Close()

	// create sqlx db with mock db
	sqlxDB := sqlx.NewDb(db, "sqlmock")
	defer sqlxDB.Close()

	// new repository
	repo := NewNewsRepository(sqlxDB)

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
			fmt.Sprintf("%s ORDER BY created_at %s OFFSET $1 LIMIT $2", getAllQuery, query.GetSort()),
		).WithArgs(
			query.GetOffset(),
			query.GetLimit(),
		).WillReturnRows(rows)

		// call GetAll method
		news, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "",
		})

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, news)
		require.Len(t, news.News, 2)
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
				"%s ORDER BY created_at %s OFFSET $1 LIMIT $2",
				fmt.Sprintf(`%s%s`, getAllQuery, " AND title LIKE '%"+query.Search+"%' "), query.GetSort()),
		).WithArgs(
			query.GetOffset(),
			query.GetLimit(),
		).WillReturnRows(rows)

		// call GetAll method
		News, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "test",
			Sort:   "",
		})

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, News)
		require.Len(t, News.News, 2)
	})

	// GetAll TotalCount equal to zero case
	t.Run("GetAll TotalCount equal to zero", func(t *testing.T) {

		// mock query for get total count zero
		mock.ExpectQuery(getTotalCountQuery).WillReturnRows(
			sqlmock.NewRows([]string{"count"}).AddRow(0),
		)

		// call GetAll method
		news, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "",
		})

		// check error and result
		require.NoError(t, err)
		require.NotNil(t, news)
		require.Len(t, news.News, 0)

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
		news, err := repo.GetAll(context.Background(), &query)

		// check error and result
		require.Error(t, err)
		require.Nil(t, news)
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
			fmt.Sprintf("%s ORDER BY created_at %s OFFSET $1 LIMIT $2", getAllQuery, query.GetSort()),
		).WithArgs(
			query.GetOffset(),
			query.GetLimit(),
		).WillReturnRows(rows)

		// call GetAll method
		news, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "",
		})

		// check error and result
		require.Error(t, err)
		require.Nil(t, news)

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
			fmt.Sprintf("%s ORDER BY created_at %s OFFSET $1 LIMIT $2", getAllQuery, query.GetSort()),
		).WithArgs(
			query.GetOffset(),
			query.GetLimit(),
		).WillReturnRows(rows)

		// call GetAll method
		news, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "",
		})

		// check error and result
		require.Error(t, err)
		require.Nil(t, news)

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
			fmt.Sprintf("%s ORDER BY created_at %s OFFSET $1 LIMIT $2", getAllQuery, query.GetSort()),
		).WithArgs(
			query.GetOffset(),
			query.GetLimit(),
		).WillReturnError(sqlmock.ErrCancelled)

		// call GetAll method
		new, err := repo.GetAll(context.Background(), &utils.Query{
			Limit:  10,
			Page:   1,
			Search: "",
			Sort:   "",
		})

		// check error and result
		require.Error(t, err)
		require.Nil(t, new)

	})
}
