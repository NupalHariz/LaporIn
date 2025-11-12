package report

const (
	insertReport = `
	INSERT INTO reports
	(
		title,
		description,
		category,
		location,
		photo_url,
		ticket_code
	) VALUES
	(
		:title,
		:description,
		:category,
		:location,
		:photo_url,
		:ticket_code
	)
	`

	getReport = `
	SELECT 
		id,
		title,
		description,
		category,
		location,
		photo_url,
		ticket_code,
		status,
		status_desc,
		status_proof_url,
		created_at,
		updated_at
	FROM
		reports
	`

	updateReport=`
	UPDATE
		reports
	`

	countReport=`
		SELECT
			COUNT(*)
		FROM
			reports
	`
)
