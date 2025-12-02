package day_2

func Puzzle1(data string) int {
	total := 0
	productRanges := ParseProductRanges(data)
	for _, productRange := range productRanges {
		for _, product := range NewProductRange(productRange[0], productRange[1]).products {
			if !product.valid {
				total += product.ID
			}
		}
	}
	return total
}
