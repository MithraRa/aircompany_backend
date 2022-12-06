package models

type Aircraft struct {
	AircraftId int
	Model      string
	Seats      *[]Seat
}

type Seat struct {
	AircraftId int
	SeatCode   string
	Price      float32
}

func GetAircraftById(id int) (*Aircraft, error) {
	return &Aircraft{}, nil
}
