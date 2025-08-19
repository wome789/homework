package main

import "fmt"

func main() {
	//fmt.Println(romanToInt("MCMXCIV"))
	//fmt.Println(intToRoman(999))
	//fmt.Println((10 % 1000 % 100) / 10)
	fmt.Println(intToRoman2(1994))
}

func romanToInt(s string) int {
	romanMap := make(map[rune]int)
	romanMap['I'] = 1
	romanMap['V'] = 5
	romanMap['X'] = 10
	romanMap['L'] = 50
	romanMap['C'] = 100
	romanMap['D'] = 500
	romanMap['M'] = 1000

	var romanRunes = []rune(s)

	var total int
	for i := 0; i < len(romanRunes); i++ {
		currentValue := romanMap[romanRunes[i]]
		var nextValue int
		if len(romanRunes)-1 != i {
			nextValue = romanMap[romanRunes[i+1]]
		}

		if currentValue < nextValue {
			total += nextValue - currentValue
			i++
		} else {
			total += currentValue
		}
	}

	return total
}

func intToRoman(intValue int) string {
	romanString := ""
	thousand := intValue / 1000

	// 千位
	if thousand > 0 {
		for i := 1; i <= thousand; i++ {
			romanString += "M"
		}
	}

	// 百位
	thousandRemainder := (intValue % 1000) / 100
	if thousandRemainder > 0 {
		if thousandRemainder == 9 {
			romanString += "CM"
		} else if thousandRemainder == 4 {
			romanString += "CD"
		} else if thousandRemainder == 5 {
			romanString += "D"
		} else if thousandRemainder > 5 {
			romanString += "D"
			for i := 1; i <= thousandRemainder-5; i++ {
				romanString += "C"
			}
		} else {
			for i := 1; i <= thousandRemainder; i++ {
				romanString += "C"
			}
		}
	}

	// 十位
	tenRemainder := (intValue % 1000 % 100) / 10
	if tenRemainder > 0 {
		if tenRemainder == 9 {
			romanString += "XC"
		} else if tenRemainder == 4 {
			romanString += "XL"
		} else if tenRemainder == 5 {
			romanString += "L"
		} else if tenRemainder > 5 {
			romanString += "L"
			for i := 1; i <= tenRemainder-5; i++ {
				romanString += "X"
			}
		} else {
			for i := 1; i <= tenRemainder; i++ {
				romanString += "X"
			}
		}
	}

	// 个位
	onesRemainder := intValue % 1000 % 100 % 10
	if onesRemainder > 0 {
		if onesRemainder == 9 {
			romanString += "IX"
		} else if onesRemainder == 4 {
			romanString += "IV"
		} else if onesRemainder == 5 {
			romanString += "V"
		} else if onesRemainder > 5 {
			romanString += "V"
			for i := 1; i <= onesRemainder-5; i++ {
				romanString += "I"
			}
		} else {
			for i := 1; i <= onesRemainder; i++ {
				romanString += "I"
			}
		}
	}

	return romanString
}

func intToRoman2(intValue int) string {
	romanString := ""
	romans := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; i < len(values); i++ {
		for {
			if intValue >= values[i] {
				romanString += romans[i]
				intValue -= values[i]
			} else {
				break
			}
		}
	}
	return romanString
}
