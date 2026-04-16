package booking

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

const defaultHoldTTL = time.Minute * 2

type RedisStore struct {
	rdb *redis.Client
}

func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{
		rdb: rdb,
	}
}

func sessionKey(id string) string {
	return fmt.Sprintf("session:%s", id)
}

func (m *RedisStore) Book(b Booking) error {
	session, err := m.hold(b)

	if err != nil {
		return err
	}

	log.Printf("session booked: %v", session)

	return nil
}

func (m *RedisStore) ListBookings(movieID string) []Booking {

	return nil
}

func (m *RedisStore) hold(b Booking) (Booking, error) {
	id := uuid.New().String()
	now := time.Now()
	ctx := context.Background()
	key := fmt.Sprintf("seat:%s:%s", b.MovieID, b.SeatID)

	b.ID = id
	val, _ := json.Marshal(b)

	res := m.rdb.SetArgs(ctx, key, val, redis.SetArgs{
		Mode: "NX",
		TTL:  defaultHoldTTL,
	})

	ok := res.Val() == "OK"

	if !ok {
		log.Printf("failed for: %s", b.ID)
		return Booking{}, ErrSeatAlreadyBooked
	}

	m.rdb.Set(ctx, sessionKey(id), key, defaultHoldTTL)

	return Booking{
		ID:        id,
		MovieID:   b.MovieID,
		SeatID:    b.SeatID,
		UserID:    b.UserID,
		Status:    "held",
		ExpiresAt: now.Add(defaultHoldTTL),
	}, nil
}
