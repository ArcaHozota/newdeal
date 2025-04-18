package tools

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Snowflake ID generator implemented in Go.
// Ported and extended from the original Java classes SnowflakeIdGenerator and SnowflakeUtils.
//   Epoch:           2015‑01‑01 00:00:00 (UTC)
//   Layout:          41 bits time | 5 bits datacenter | 5 bits worker | 12 bits sequence
//   Time granularity: milliseconds.
//
// Usage examples:
//   // 1. Create a persistent generator (preferred for high throughput)
//   gen, err := utils.NewSnowflakeIDGenerator(1, 1)
//   if err != nil { log.Fatal(err) }
//   id, _ := gen.NextID()
//   fmt.Println(id)
//
//   // 2. One‑shot ID similar to Java SnowflakeUtils.snowflakeId()
//   id, _ := utils.SnowflakeID()
//   fmt.Println(id)

const (
	// custom epoch (2015‑01‑01) in milliseconds
	timeEpoch int64 = 1420041600000

	workerIDBits     uint8 = 5
	datacenterIDBits uint8 = 5
	sequenceBits     uint8 = 12

	maxWorkerID     int64 = -1 ^ (-1 << workerIDBits)     // 31
	maxDatacenterID int64 = -1 ^ (-1 << datacenterIDBits) // 31

	workerIDShift            = sequenceBits
	datacenterIDShift        = sequenceBits + workerIDBits
	timestampLeftShift       = sequenceBits + workerIDBits + datacenterIDBits
	sequenceMask       int64 = -1 ^ (-1 << sequenceBits)
)

// init seeds the default PRNG used for SnowflakeID().
func init() {
	rand.Seed(time.Now().UnixNano())
}

// SnowflakeIDGenerator provides unique 64‑bit IDs.
// All exported methods are safe for concurrent use.
//
// The internal ID layout (most‑significant -> least):
//   41 bits  timestamp offset from epoch (ms)
//   5  bits  datacenter ID
//   5  bits  worker ID
//   12 bits  sequence within the same millisecond
//
// When system clock moves backwards this implementation will return an error.
// You could choose to block instead by replacing the error with a sleep/retry.

type SnowflakeIDGenerator struct {
	mu           sync.Mutex
	workerID     int64
	datacenterID int64
	sequence     int64
	lastTS       int64
}

// NewSnowflakeIDGenerator validates IDs and returns a generator.
func NewSnowflakeIDGenerator(workerID, datacenterID int64) (*SnowflakeIDGenerator, error) {
	if workerID < 0 || workerID > maxWorkerID {
		return nil, fmt.Errorf("worker ID must be between 0 and %d", maxWorkerID)
	}
	if datacenterID < 0 || datacenterID > maxDatacenterID {
		return nil, fmt.Errorf("datacenter ID must be between 0 and %d", maxDatacenterID)
	}
	return &SnowflakeIDGenerator{
		workerID:     workerID,
		datacenterID: datacenterID,
		sequence:     0,
		lastTS:       -1,
	}, nil
}

// NextID returns the next unique ID. Thread‑safe.
func (s *SnowflakeIDGenerator) NextID() (uint64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ts := currentMillis()

	if ts < s.lastTS {
		return 0, fmt.Errorf("clock moved backwards: refusing to generate id for %d ms", s.lastTS-ts)
	}

	if ts == s.lastTS {
		s.sequence = (s.sequence + 1) & sequenceMask
		if s.sequence == 0 {
			ts = s.waitNextMillis()
		}
	} else {
		s.sequence = 0
	}

	s.lastTS = ts

	id := ((ts - timeEpoch) << timestampLeftShift) |
		(s.datacenterID << datacenterIDShift) |
		(s.workerID << workerIDShift) |
		s.sequence

	return uint64(id), nil
}

// SnowflakeID is a convenience wrapper mirroring Java's SnowflakeUtils.snowflakeId().
// It picks random worker/datacenter IDs (0‑31) each call and returns a unique ID.
func SnowflakeID() (uint64, error) {
	worker := rand.Int63n(maxWorkerID + 1)
	dc := rand.Int63n(maxDatacenterID + 1)

	gen, err := NewSnowflakeIDGenerator(worker, dc)
	if err != nil {
		return 0, err
	}
	return gen.NextID()
}

func (s *SnowflakeIDGenerator) waitNextMillis() int64 {
	ts := currentMillis()
	for ts <= s.lastTS {
		ts = currentMillis()
	}
	return ts
}

func currentMillis() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
