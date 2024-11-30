package main

import "fmt"

/**
	EXIBIR TODOS OS NÚMEROS DE 1 A 100 QUE SÃO DIVISIVEIS POR 3
**/

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d\n", i)
		}
	}
}
