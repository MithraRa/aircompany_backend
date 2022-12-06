package models

const InsertUserQuery = `
	INSERT INTO 
	    passengers (name, email, password) 
	VALUES ($1, $2, $3)`

const FindUserByEmailQuery = `SELECT 
	COALESCE(passenger_id, 0),
	COALESCE(name, ''),
	COALESCE(lastname, ''),
	COALESCE(patronymic, ''),
	COALESCE(document_data, ''),
	COALESCE(phone, ''),
	COALESCE(password, ''),
	COALESCE(email, '')
FROM 
	passengers
WHERE
    email = $1`

const FindUserByIdQuery = `SELECT 
	COALESCE(passenger_id, 0),
	COALESCE(name, ''),
	COALESCE(lastname, ''),
	COALESCE(patronymic, ''),
	COALESCE(document_data, ''),
	COALESCE(phone, ''),
	COALESCE(password, ''),
	COALESCE(email, '')
FROM 
	passengers
WHERE
    passenger_id = $1`

const UpdateUserQuery = `
UPDATE 
	passengers
SET
	name = $1,
	lastname = $2,
	patronymic = $3,
	document_data = $4,
	phone = $5
WHERE 
	passenger_id = $6`

const DeleteUserIdQuery = `
DELETE 
	FROM passengers     
WHERE 
    passenger_id = $1`

const DeleteUserEmailQuery = `
DELETE 
	FROM passengers     
WHERE 
    email = $1`
