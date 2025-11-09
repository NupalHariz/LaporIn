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
)
