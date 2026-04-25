package main

func myAtoi(s string) int {
	var result int64 = 0
	var negative bool
	var numberStart bool = false

	for i, element := range s {

		if string(element) == " " {
			continue
		}

		if element <= '9' && element >= '0' {
			digit := int64(element - '0')
			result = result*10 + digit

			if i != 0 && string(s[i-1]) == "-" && numberStart == false {
				negative = true
			} else if i != 0 && string(s[i-1]) == "+" && numberStart == false {
				negative = false
			}
			numberStart = true

		} else if numberStart {
			break
		}

	}

	if negative {
		result = -1 * result
	}

	// rounding:
	if result < -2147483648 {
		return -2147483648
	} else if result > 2147483647 {
		return 2147483647
	}

	return int(result)
}
