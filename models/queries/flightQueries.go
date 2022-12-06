package models

const FindFlightsQuery = `SELECT 
	flight_id AS FlightId,
	aircraft_id AS AircraftId,
	(SELECT CONCAT(airport_name, ' ', city) FROM airports WHERE airport_id = departure_airport) AS DepartureAirport,
	(SELECT CONCAT(airport_name, ' ', city) FROM airports WHERE airport_id = arrival_airport) AS ArrivalAirport,
	departure_time AS DepartureTime,
	arrival_time AS ArrivalTime,
	flight_price AS FlightPrice
FROM
	flights`

const FindAllUserTicketsByIdQuery = `
SELECT 
 	tickets.ticket_id as TicketId,
	tickets.passenger_id AS UserId,
	tickets.total_amount AS Total,
	tickets.seat_code AS SeatCode,
	flights.flight_id AS FlightId,
	(SELECT CONCAT(airport_name, ' ', city) FROM airports WHERE airport_id = departure_airport) AS DepartureAirport,
	(SELECT CONCAT(airport_name, ' ', city) FROM airports WHERE airport_id = arrival_airport) AS ArrivalAirport,
	departure_time AS DepartureTime,
	arrival_time AS ArrivalTime
FROM 
	tickets
JOIN ticket_flight ON tickets.ticket_id = ticket_flight.ticket_id
JOIN flights ON flights.flight_id = ticket_flight.flight_id
WHERE tickets.passenger_id = $1`
