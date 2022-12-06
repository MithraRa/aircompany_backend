package models

const GetFreeSeatsByFlightIdQuery = `SELECT * FROM getfreeseatsbyflightid($1)`

const InsertNewTicketQuery = `
INSERT INTO 
	tickets (passenger_id, total_amount, seat_code) 
VALUES 
    ($1, $2, $3)`

const InsertNewTicketFlightQuery = `
INSERT INTO 
	ticket_flight
VALUES 
    (LASTVAL(), $1)`
