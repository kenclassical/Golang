package main

func fline_max(array []int) int {
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
	sum := fline_max(number)
	print(sum)
}
