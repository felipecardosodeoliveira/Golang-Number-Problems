package main

import "fmt"

/**
	FAZER UMA CONTAGEM DE 1 A 100
	SEMRE QUE APARACER UM MÚLTIPLO DE 3 MPRIMIR "PIN".
	E NOS MÚLTIPLOS DE 5 FALAR "PAN"
**/

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("PIN\n")
		} else if i%5 == 0 {
			fmt.Printf("PAN\n")
		}
	}
}
