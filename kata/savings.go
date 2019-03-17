package kata

import "math"

func NbMonths(startPriceOld, startPriceNew, savingperMonth int, percentLossByMonth float64) [2]int {
	money := float64(startPriceOld)
	valueOld := float64(startPriceOld)
	price := float64(startPriceNew)
	savings := 0.0
	month := 0

	for money < price {
		month++

		if month%2 == 0 {
			percentLossByMonth += 0.5
		}
		valueOld = valueOld * (1.0 - percentLossByMonth/100)
		price = price * (1.0 - percentLossByMonth/100)
		savings = float64(savingperMonth) + savings
		money = savings + valueOld

	}
	return [2]int{month, int(math.Round(money - price))}
}
