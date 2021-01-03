package reservation

func NewReservation(vertical, reservationDate string) Reservation {
	switch(vertical) {
	case "flight":
		return FlightReservationImplementation{reservationDate,}
	case "hotel":
		return HotelReservationImplementation{reservationDate,}
	default:
		return nil
	}
}