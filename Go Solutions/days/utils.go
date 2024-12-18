package days

import "strconv"

func convertToBase(numberToConvert int, base int) string {
	var convertedNumber string
	if numberToConvert == 0 {
		return "0"
	}
	for numberToConvert > 0 {
		remainder := numberToConvert % base
		numberToConvert /= base
		convertedNumber = strconv.Itoa(remainder) + convertedNumber
	}
	return convertedNumber
}
