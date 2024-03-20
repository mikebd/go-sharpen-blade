package parkingLot

type vehicleType int8

const (
	motorcycle vehicleType = iota
	compactCar
	regularCar
	van
)

const vanFitsInCarSpots = 3

type api interface {
	// isEmpty returns true if the parking lot is empty
	isEmpty() bool

	// isFull returns true if the parking lot is full
	isFull() bool

	// isFullFor returns true if the parking lot is full for the given vehicle type
	isFullFor(vehicleType) bool

	// spotCount returns the number of spots in the parking lot
	spotCount() int

	// availableSpots returns the number of available spots in the parking lot
	availableSpots() int

	// spotsTakenBy returns the number of spots taken by the given vehicle type
	spotsTakenBy(vehicleType) int
}

/*

// parkingSpots is a contiguous row of spots
type parkingSpots struct {

}

*/
