package parkingLot

import (
	lop "github.com/samber/lo/parallel"
	"sync/atomic"
)

// row is a contiguous row of parking spots
type row struct {
	spots []*spot
}

func (r *row) isEmpty() bool {
	if len(r.spots) == 0 {
		return false
	}

	empty := atomic.Bool{}
	empty.Store(true)

	lop.ForEach(r.spots, func(s *spot, _ int) {
		if s.isOccupied() {
			empty.Store(false)
		}
	})

	return empty.Load()
}

func (r *row) isFull() bool {
	if len(r.spots) == 0 {
		return true
	}

	full := atomic.Bool{}
	full.Store(true)

	lop.ForEach(r.spots, func(s *spot, _ int) {
		if s.isEmpty() {
			full.Store(false)
		}
	})

	return full.Load()
}

func (r *row) isFullFor(v vehicleType) bool {
	//TODO implement me
	panic("implement me")
}

func (r *row) spotCount() int {
	return len(r.spots)
}

func (r *row) availableSpots() int {
	if len(r.spots) == 0 {
		return 0
	}

	available := atomic.Int32{}
	available.Store(0)

	lop.ForEach(r.spots, func(s *spot, _ int) {
		if s.isEmpty() {
			available.Add(1)
		}
	})

	return int(available.Load())
}

func (r *row) spotsTakenBy(v vehicleType) int {
	if len(r.spots) == 0 {
		return 0
	}

	takenBy := atomic.Int32{}
	takenBy.Store(0)

	lop.ForEach(r.spots, func(s *spot, _ int) {
		if s.isOccupied() && *s.occupiedBy == v {
			takenBy.Add(1)
		}
	})

	return int(takenBy.Load())
}
