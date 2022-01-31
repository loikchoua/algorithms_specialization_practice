package merge

func Sort(input []int) []int {
	switch len(input) {
	case 1:
		return input
	default:
		half := len(input) / 2
		firstHalf := Sort(input[0:half])
		secondHalf := Sort(input[half:])

		return merge(firstHalf, secondHalf)
	}
}

func merge(a []int, b []int) []int {
	response := make([]int, len(a)+len(b))
	i := 0
	j := 0
	for index := 0; index < len(response); index++ {

		switch {
		case i < len(a) && j < len(b):
			if a[i] <= b[j] {
				response[index] = a[i]
				i += 1
			} else {
				response[index] = b[j]
				j += 1
			}
		case i < len(a):
			response[index] = a[i]
			i += 1
		case j < len(b):
			response[index] = b[j]
			j += 1
		}
	}
	return response
}

func SortAndCountInversions(a []int) (sorted []int, numberOfInversions int) {
	switch len(a) {
	case 1:
		sorted = a
		return
	default:
		x, b := SortAndCountInversions(a[:len(a)/2])
		y, c := SortAndCountInversions(a[len(a)/2:])
		z, d := mergeAndCountInversions(x, y)
		return z, b + c + d
	}
}

func mergeAndCountInversions(a []int, b []int) ([]int, int) {
	response := make([]int, len(a)+len(b))
	i := 0
	j := 0
	numberOfInversions := 0
	for index := 0; index < len(response); index++ {

		switch {
		case i < len(a) && j < len(b):
			if a[i] <= b[j] {
				response[index] = a[i]
				i += 1
			} else {
				response[index] = b[j]
				numberOfInversions += len(a) - i
				j += 1
			}
		case i < len(a):
			response[index] = a[i]
			i += 1
		case j < len(b):
			response[index] = b[j]
			j += 1
		}
	}
	return response, numberOfInversions
}
