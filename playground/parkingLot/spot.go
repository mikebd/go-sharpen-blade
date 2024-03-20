package parkingLot

// spot is a parking spot that can fit a vehicle up to maxVehicleSize
// and may or may not be occupied by a vehicle
type spot struct {
	maxVehicleSize vehicleType
	occupiedBy     *vehicleType
}

func (s *spot) isMotorcycleSpot() bool {
	return s.maxVehicleSize == motorcycle
}

func (s *spot) isCarSpot() bool {
	return s.maxVehicleSize == compactCar || s.maxVehicleSize == regularCar
}

func (s *spot) isVanSpot() bool {
	return s.maxVehicleSize == van
}

func (s *spot) isEmpty() bool {
	return s.occupiedBy == nil
}

func (s *spot) isOccupied() bool {
	return s.occupiedBy != nil
}
