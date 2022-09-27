package session2

import (
	"fmt"
)

func main2() {
	var name string

	for i := 0; i <= 10; i++ {
		if i%2 == 1 {
			name = "ganjil"
		} else {
			name = "genap"
		}

		fmt.Println(i, name)

	}
}
