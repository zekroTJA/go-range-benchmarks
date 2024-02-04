package main

func SimpleSum(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}

	return sum
}
