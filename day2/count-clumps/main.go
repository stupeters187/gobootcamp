package main

import "fmt"

func countClumps(xs []int) int {
	clumps := 0
	for i := 1; i < len(xs); i++ {
		curr, prev := xs[i], xs[i-1]
		if curr == prev {
			clumps++
			for ; i < len(xs); i++ {
				curr, prev := xs[i], xs[i-1]
				if curr != prev {
					break
				}
			}
		}
	}
	return clumps
}

func main()  {

	clumps := []int{1,1,2,2,1,1}
	fmt.Println(countClumps(clumps))




}