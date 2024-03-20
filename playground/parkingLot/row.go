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
	if r.isFull() {
		return true
	}

	switch v {
	case motorcycle:
		return false
	case compactCar:
		fallthrough
	case regularCar:
		return !r.hasEmptySpotForCarOrVan()
	case van:
		return !r.hasEmptySpotForVan()
	default:
		return true
	}
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

func (r *row) hasEmptySpotForCarOrVan() bool {
	if len(r.spots) == 0 {
		return false
	}

	empty := atomic.Bool{}
	empty.Store(false)

	lop.ForEach(r.spots, func(s *spot, _ int) {
		if s.isEmpty() && !s.isMotorcycleSpot() {
			empty.Store(true)
		}
	})

	return empty.Load()
}

func (r *row) hasEmptySpotForVan() bool {
	if len(r.spots) == 0 {
		return false
	}

	consecutiveEmptyCarSpots := 0
	for _, s := range r.spots {
		if s.isEmpty() {
			if s.isVanSpot() {
				return true
			}

			if s.isCarSpot() {
				consecutiveEmptyCarSpots++
				if consecutiveEmptyCarSpots == vanFitsInCarSpots {
					return true
				}
			} else {
				consecutiveEmptyCarSpots = 0
			}
		} else {
			consecutiveEmptyCarSpots = 0
		}
	}

	return false
}
