package taskmodule

const (
	addTaskQuery = `
	INSERT INTO task (
		title,
		completed,
		hour,
		minute
	) VALUES (
		$1,
		$2,
		$3,
		$4
	) returning id
`
	getTaskQuery = `
	SELECT
		title,
		completed,
		hour,
		minute
	FROM
		task
	WHERE
		id=$1
`

	getTaskAllQuery = `
	SELECT
		*
	FROM
		task
`

	updateTaskQuery = `
	UPDATE
		task
	SET
		title=COALESCE($1, title),
		completed=COALESCE($2, completed),
		hour=COALESCE($3, hour),
		minute=COALESCE($4, minute)
	WHERE
		id=$5
	returning id	
	
`

	removeTaskQuery = `
	
	DELETE FROM
		task
	WHERE
		id=$1
`
	toggleTaskQuery = `
	
	UPDATE
		task
	SET
		completed=NOT(completed)
	WHERE
		id=$1
	returning id
`
)
