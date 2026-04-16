package booking

type MemoryStore struct {
	bookings map[string]Booking
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		bookings: map[string]Booking{},
	}
}

func (m *MemoryStore) Book(b Booking) error {
	if _, ok := m.bookings[b.SeatID]; ok {
		return ErrSeatAlreadyBooked
	}

	m.bookings[b.SeatID] = b
	return nil
}

func (m *MemoryStore) ListBookings(movieID string) []Booking {
	var bookings []Booking

	for _, booking := range m.bookings {
		if booking.MovieID == movieID {
			bookings = append(bookings, booking)
		}
	}

	return bookings
}
