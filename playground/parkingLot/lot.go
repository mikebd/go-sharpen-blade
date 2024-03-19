package parkingLot

import (
	lop "github.com/samber/lo/parallel"
	"sync/atomic"
)

type lot struct {
	rows []*row
}

func newParkingLot() *lot {
	return &lot{}
}

func (l *lot) isEmpty() bool {
	if len(l.rows) == 0 {
		return false
	}

	empty := atomic.Bool{}
	empty.Store(true)

	lop.ForEach(l.rows, func(r *row, _ int) {
		if !r.isEmpty() {
			empty.Store(false)
		}
	})

	return empty.Load()
}

func (l *lot) isFull() bool {
	if len(l.rows) == 0 {
		return true
	}

	full := atomic.Bool{}
	full.Store(true)

	lop.ForEach(l.rows, func(r *row, _ int) {
		if !r.isFull() {
			full.Store(false)
		}
	})

	return full.Load()
}

func (l *lot) isFullFor(v vehicleType) bool {
	//TODO implement me
	panic("implement me")
}

func (l *lot) spotCount() int {
	//TODO implement me
	panic("implement me")
}

func (l *lot) availableSpots() int {
	//TODO implement me
	panic("implement me")
}

func (l *lot) spotsTakenBy(v vehicleType) int {
	//TODO implement me
	panic("implement me")
}
