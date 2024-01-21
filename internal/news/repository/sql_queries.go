package repository

import "fmt"

var (

	// list of fields from news table.
	fieldsOfNewsTable = `id, title, content, created_at`

	// query for create new news.
	createQuery = fmt.Sprintf(`
	INSERT INTO news
	(
		title,
		content
	)
	VALUES ($1, $2)
	RETURNING %s`, fieldsOfNewsTable)

	// query for update news.
	updateQuery = fmt.Sprintf(`
	UPDATE news SET
		title = $1,
		content = $2
	WHERE id = $3
	RETURNING %s`, fieldsOfNewsTable)

	// query for delete news.
	deleteQuery = `DELETE FROM news WHERE id = $1`

	// query for get news by id.
	getByIDQuery = fmt.Sprintf(`
	SELECT 
		%s 
	FROM news
	WHERE
		id = $1
	`, fieldsOfNewsTable)

	// query for get total count of newsList.
	getTotalCountQuery = `SELECT COUNT(id) FROM news WHERE 1=1 `

	// query for get all newsList.
	getAllQuery = fmt.Sprintf(`
	SELECT 
		%s 
	FROM news
	WHERE
		1=1
	`, fieldsOfNewsTable)
)
