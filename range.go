package main

func SimpleSum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}

	return sum
}

func Strings(s []string) (sum int) {
	for _, v := range s {
		sum += len(v)
	}

	return sum
}

func Struct(s []LargeStruct) (sum int64) {
	for _, v := range s {
		sum += v.F0 + v.F10 + v.F50
	}

	return sum
}
