package utils

import (
	crand "crypto/rand"
	rand "math/rand"

	"encoding/binary"
	"log"
)

const idSource = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// GetRandomNumber generates a cryptographically secure number between 0 and the specified number
func GetRandomNumber(max int) int {
	var src cryptoSource
	rnd := rand.New(src)
	return rnd.Intn(max)
}

// GenerateRandomID generates a truely random ID for a new user
func GenerateRandomID() string {
	b := make([]byte, 8)

	for i := range b {
		b[i] = idSource[GetRandomNumber(len(idSource))]
	}

	return string(b)
}

type cryptoSource struct{}

func (s cryptoSource) Seed(seed int64) {}

func (s cryptoSource) Int63() int64 {
	return int64(s.Uint64() & ^uint64(1<<63))
}

func (s cryptoSource) Uint64() (v uint64) {
	err := binary.Read(crand.Reader, binary.BigEndian, &v)
	if err != nil {
		log.Fatal(err)
	}
	return v
}
