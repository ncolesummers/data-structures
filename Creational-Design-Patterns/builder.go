package reservation

type ReservationBuilder interface {
	Vertical(string) ReservationBuilder
	ReservationDate(string) ReservationBuilder
	Build() Reservation
}

type reservationBuilder struct {
	vertical string
	rdate string
}

func (r *reservationBuilder) Vertical(v string) ReservationBuilder {
	r.vertical = v
	return r
}

func (r *reservationBuilder) ReservationDate(date string) ReservationBuilder {
	r.rdate = date
	return r
}

func (r *reservationBuilder) Build() Reservation {
	var builtReservation Reservation

	switch r.vertical {
	case "flight":
		builtReservation = FlightReservationImplementation{r.date}
	case "hotel":
		builtReservation = HotelReservationImplementation{r.date}
	}

	return builtReservation
}

func NewReservationBuilder() ReservationBuilder {
	return &reservationBuilder{}
}