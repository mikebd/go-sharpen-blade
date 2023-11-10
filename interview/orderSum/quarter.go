package orderSum

type quarter string

const (
	Q1 quarter = "Q1"
	Q2 quarter = "Q2"
	Q3 quarter = "Q3"
	Q4 quarter = "Q4"
)

type mapMonthQuarter map[string]quarter

// monthToQuarter maps month number (as a two digit string) to quarter
var monthToQuarter = func() mapMonthQuarter {
	result := make(mapMonthQuarter, 12)
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
}()
