package orderSum

import "strings"

type order struct {
	timestamp    string
	postalCode   string
	customerName string
	unitVolume   int
}

type orderList []order

type sumByResult map[string]int

const (
	Q1 = "Q1"
	Q2 = "Q2"
	Q3 = "Q3"
	Q4 = "Q4"
)

func makeMonthToQuarter() map[string]quarter {
	result := make(map[string]quarter, 12)
	result["01"] = Q1
	result["02"] = Q1
	result["03"] = Q1
	result["04"] = Q2
	result["05"] = Q2
	result["06"] = Q2
	result["07"] = Q3
	result["08"] = Q3
	result["09"] = Q3
	result["10"] = Q4
	result["11"] = Q4
	result["12"] = Q4
	return result
}

var monthToQuarter map[string]quarter

func init() {
	monthToQuarter = makeMonthToQuarter()
}

type quarter string

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
