package repository

import "fmt"

var (

	// list of fields from blogs table.
	fieldsOfBlogsTable = `id, title, content, created_at`

	// query for create new blog.
	createQuery = fmt.Sprintf(`
	INSERT INTO blogs
	(
		title,
		content
	)
	VALUES ($1, $2)
	RETURNING %s`, fieldsOfBlogsTable)

	// query for update blog.
	updateQuery = fmt.Sprintf(`
	UPDATE blogs SET
		title = $1,
		content = $2
	WHERE id = $3
	RETURNING %s`, fieldsOfBlogsTable)

	// query for delete blog.
	deleteQuery = `DELETE FROM blogs WHERE id = $1`

	// query for get blog by id.
	getByIDQuery = fmt.Sprintf(`
	SELECT 
		%s 
	FROM blogs
	WHERE
		id = $1
	`, fieldsOfBlogsTable)

	// query for get total count of blogs.
	getTotalCountQuery = `SELECT COUNT(id) FROM blogs WHERE 1=1 `

	// query for get all blogs.
	getAllQuery = fmt.Sprintf(`
	SELECT 
		%s 
	FROM blogs
	WHERE
		1=1`, fieldsOfBlogsTable)
)
