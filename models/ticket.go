package models

import (
	"airline/database"
	models "airline/models/queries"
	"log"
	"time"
)

type TicketFlight struct {
	FlightId uint
	TicketId uint
}

type Ticket struct {
	TicketId    uint
	PassengerId uint
	Total       float32
	SeatCode    string
}

type TicketSeat struct {
	SeatCode string
	Price    float32
}

type PersonalTicket struct {
	TicketId         uint
	PassengerId      uint
	Total            float32
	SeatCode         string
	FlightId         int
	DepartureAirport string
	ArrivalAirport   string
	DepartureTime    time.Time
	ArrivalTime      time.Time
}

func GetFreeSeats(flightId int) (*[]TicketSeat, error) {
	rows, err := database.GetDB().Query(models.GetFreeSeatsByFlightIdQuery, flightId)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var freeSeats []TicketSeat

	for rows.Next() {
		f := TicketSeat{}
		err := rows.Scan(
			&f.SeatCode,
			&f.Price,
		)
		if err != nil {
			continue
		}
		freeSeats = append(freeSeats, f)
	}

	return &freeSeats, nil
}

func CreateTicket(flightId, userId int, price float32, seatCode string) error {
	tx, err := database.GetDB().Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(models.InsertNewTicketQuery, userId, price, seatCode)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(models.InsertNewTicketFlightQuery, flightId)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func GetAllPersonalTickets(userId int) ([]PersonalTicket, error) {
	rows, err := database.GetDB().Query(models.FindAllUserTicketsByIdQuery, userId)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var tickets []PersonalTicket

	for rows.Next() {
		f := PersonalTicket{}
		err := rows.Scan(
			&f.TicketId,
			&f.PassengerId,
			&f.Total,
			&f.SeatCode,
			&f.FlightId,
			&f.DepartureAirport,
			&f.ArrivalAirport,
			&f.DepartureTime,
			&f.ArrivalTime,
		)
		if err != nil {
			continue
		}
		tickets = append(tickets, f)
	}

	return tickets, nil
}
