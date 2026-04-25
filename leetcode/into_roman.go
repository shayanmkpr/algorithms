package main

func intToRoman(num int) string {
	if num > 3999 { // this should not happen
		return "What?"
	}

	var roman string
	var copy int
	var rem int
	var tenCntr int

	rem = 0
	copy = num
	tenCntr = 1 // Start at 1 instead of 10

	romanMap := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	for copy != 0 {
		rem = copy % 10
		copy = copy / 10

		switch {
		case rem == 9:
			roman = romanMap[tenCntr] + romanMap[tenCntr*10] + roman
		case rem == 4:
			roman = romanMap[tenCntr] + romanMap[tenCntr*5] + roman
		case rem >= 5:
			for range rem - 5 {
				roman = romanMap[tenCntr] + roman
			}
			roman = romanMap[tenCntr*5] + roman
		case rem > 0:
			for range rem {
				roman = romanMap[tenCntr] + roman
			}
		}

		tenCntr *= 10 // Multiply by 10 to move to next decimal place
	}

	return roman
}
