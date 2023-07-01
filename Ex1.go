package main

func find_max(array []int) int {
	max := array[0]
	for i := range array {
		if array[i] > max {
			max = array[i]
		}
	}
	return max
}

func main() {
	number := []int{9, 6, 2, 4, 7, 1, 11}
	sum := find_max(number)
	print(sum)
}
