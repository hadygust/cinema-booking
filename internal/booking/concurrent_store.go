package booking

import (
	"sync"
)

type ConcurrentStore struct {
	bookings map[string]Booking
	sync.RWMutex
}

func NewConcurrentStore() *ConcurrentStore {
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

func (m *ConcurrentStore) Book(b Booking) error {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.bookings[b.SeatID]; ok {
		return ErrSeatAlreadyBooked
	}

	m.bookings[b.SeatID] = b
	return nil
}

func (m *ConcurrentStore) ListBookings(movieID string) []Booking {
	m.RLock()
	defer m.RUnlock()

	var bookings []Booking

	for _, booking := range m.bookings {
		if booking.MovieID == movieID {
			bookings = append(bookings, booking)
		}
	}

	return bookings
}
