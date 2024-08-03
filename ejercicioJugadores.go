package main

import "fmt"

func main() {
	mapaJugadores := map[string]int{
		"John Smith":          1001,
		"John Michael Garcia": 1002,
		"Sarah Williams":      1003,
		"Daniel Martinez":     1004,
		"Emily Brown":         1005,
		"James Rodriguez":     1006,
		"Sophia Lee":          1007,
	}
	nuevosJugadores := []string{"Franco Lee", "Francisco Herrera", "Emiliano Peña", "Carlos Rodriguez", "Agustín Martinez", "Julián Aponte", "Guillermo Rodriguez"}
	i := 0
	emptyMap := make(map[string]int)
	for k, v := range mapaJugadores {
		k = nuevosJugadores[i]
		// fmt.Println(k)
		emptyMap[k] = v
		i++
		fmt.Println(k,v)
	}
	
	// fmt.Println(emptyMap)

}
