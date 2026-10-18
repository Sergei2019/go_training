package main

import (
	"fmt"
	"sync"
)

func main() {
	var lst []int = []int{}
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		lst = append(lst, i)
	}
	fmt.Println(lst)
	{
		wg.Add(1)
		var lst1 []int = make([]int, len(lst))
		go func() {
			defer wg.Done()
			copy(lst1, lst)
			fmt.Println(lst1)
		}()

	}
	{
		for i := 0; i <= 10; i++ {
			for j := 0; j <= 100; j++ {
				fmt.Print(i*j, "\t")
			}
		}
	}
	wg.Wait()
	fmt.Println("все горутины завершены")

}
