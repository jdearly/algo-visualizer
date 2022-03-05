package algo

import (
	"fmt"
	"time"

	"github.com/jdearly/algo-visualizer/pkg/gotui"
)

func InsertionSort(items []int, output bool) ([]int, int) {
	var j int
	var ops int

	for i := 1; i < len(items); i++ {
		j = i
		for j > 0 && (items[j] < items[j-1]) {
			items[j], items[j-1] = items[j-1], items[j]
			j = j - 1
			ops++
			// print current order
			// TODO: this is temporary until real tui is built
			if output {
				for i, v := range items {
					if i == j {
						fmt.Println(gotui.PrintStars(v) + " <- j")
					} else if i == j-1 {
						fmt.Println(gotui.PrintStars(v) + " <- j-1")
					} else {
						fmt.Println(gotui.PrintStars(v))
					}
				}
				time.Sleep(500 * time.Millisecond)
			}
		}
	}

	return items, ops
}
