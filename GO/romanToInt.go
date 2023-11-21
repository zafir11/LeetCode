package main

func romanToInt(s string) int {

	roman := make(map[byte]int)
	roman['I'] = 1
	roman['V'] = 5
	roman['X'] = 10
	roman['L'] = 50
	roman['C'] = 100
	roman['D'] = 500
	roman['M'] = 1000

	var nextval, result int
	for i := len(s) - 1; i >= 0; i-- {
		currentval := roman[s[i]]
		// check for subtration
		if currentval < nextval {
			result -= currentval
		} else {
			result += currentval
		}
		nextval = currentval
	}

	return result
}
