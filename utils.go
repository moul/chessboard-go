package chessboard

import "fmt"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

// FIXME: remove duplicates
func combinations(iterable []int, r int) <-chan []int {
	c := make(chan []int, 1)

	go func() {
		defer close(c)

		pool := iterable
		n := len(pool)

		if r > n {
			return
		}

		indices := make([]int, r)
		for i := range indices {
			indices[i] = i
		}

		result := make([]int, r)
		for i, el := range indices {
			result[i] = pool[el]
		}

		c <- result

		for {
			i := r - 1
			for ; i >= 0 && indices[i] == i+n-r; i -= 1 {
			}

			if i < 0 {
				return
			}

			indices[i] += 1
			for j := i + 1; j < r; j += 1 {
				indices[j] = indices[j-1] + 1
			}

			for ; i < len(indices); i += 1 {
				result[i] = pool[indices[i]]
			}
			c <- result
		}
	}()

	return c
}

func permutations(iterable []int, r int) {
	pool := iterable
	n := len(pool)

	if r > n {
		return
	}

	indices := make([]int, n)
	for i := range indices {
		indices[i] = i
	}

	cycles := make([]int, r)
	for i := range cycles {
		cycles[i] = n - i
	}

	result := make([]int, r)
	for i, el := range indices[:r] {
		result[i] = pool[el]
	}

	fmt.Println(result)

	for n > 0 {
		i := r - 1
		for ; i >= 0; i -= 1 {
			cycles[i] -= 1
			if cycles[i] == 0 {
				index := indices[i]
				for j := i; j < n-1; j += 1 {
					indices[j] = indices[j+1]
				}
				indices[n-1] = index
				cycles[i] = n - i
			} else {
				j := cycles[i]
				indices[i], indices[n-j] = indices[n-j], indices[i]

				for k := i; k < r; k += 1 {
					result[k] = pool[indices[k]]
				}

				fmt.Println(result)

				break
			}
		}

		if i < 0 {
			return
		}
	}
}
