package main

import "fmt"

func generateMatrix(n int) [][]int {

	matrix := make([][]int, n)
	for l := 0; l < n; l++ {
		matrix[l] = make([]int, n)
	}

	indexArray := [2][4]int{{1, 0, -1, 0}, {0, 1, 0, -1}}

	count := 0
	i := 0
	j := -1
	v := 1

	for {
		orginI := i
		orginJ := j
		for k := 0; k < 4; k++ {
			j += indexArray[0][k]
			i += indexArray[1][k]
			if i >= 0 && i < n && j >= 0 && j < n && matrix[i][j] == 0 {
				matrix[i][j] = v
				v++
				count = 0
				break
			} else {
				count++
				i = orginI
				j = orginJ
			}
		}
		if count == 4 {
			break
		}
	}

	return matrix
}


func main() {
	fmt.Println(generateMatrix(3))
}