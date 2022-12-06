package models

import (
	"airline/database"
	models "airline/models/queries"
	"fmt"
	"log"
	"time"
)

type Flight struct {
	FlightId         uint
	AircraftId       uint
	DepartureAirport string
	DepartureTime    time.Time
	ArrivalAirport   string
	ArrivalTime      time.Time
	FlightPrice      float32
}

func GetAllFlights() (*[]Flight, error) {
	rows, err := database.GetDB().Query(models.FindFlightsQuery)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var flights []Flight

	for rows.Next() {
		f := Flight{}
		err := rows.Scan(
			&f.FlightId,
			&f.AircraftId,
			&f.DepartureAirport,
			&f.ArrivalAirport,
			&f.DepartureTime,
			&f.ArrivalTime,
			&f.FlightPrice,
		)
		if err != nil {
			fmt.Println(err)
			continue
		}
		flights = append(flights, f)
	}

	return &flights, nil
}
