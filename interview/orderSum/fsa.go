package orderSum

import "strings"

type sumByResult map[string]int

type sumByResultByQuarter map[quarter]sumByResult

func sumByFsaByQuarter(orders orderList) sumByResultByQuarter {
	result := make(sumByResultByQuarter, 4)

	result[Q1] = make(sumByResult, len(orders))
	result[Q2] = make(sumByResult, len(orders))
	result[Q3] = make(sumByResult, len(orders))
	result[Q4] = make(sumByResult, len(orders))

	for _, order := range orders {
		quarter := monthToQuarter[strings.Split(order.timestamp, "-")[1]]
		resultByFsa := result[quarter]
		fsa := strings.Split(order.postalCode, " ")[0]
		resultByFsa[fsa] += order.unitVolume
	}

	return result
}

func sumByFsa(orders orderList) sumByResult {
	result := make(sumByResult, len(orders))

	for _, order := range orders {
		fsa := strings.Split(order.postalCode, " ")[0]
		result[fsa] += order.unitVolume
	}

	return result
}
