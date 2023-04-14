package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		x float64
	)
	list := []float64{}
	for {
		fmt.Print("Digite um numero para a lista")
		fmt.Println("Digite zero para parar")
		fmt.Scan(&x)
		if x != 0 {
			list = append(list, x)
		} else {
			break
		}
	}
	menor, err := ordenar(list)
	if err != nil {
		fmt.Println("Ocorreu um erro e a lista ta vazia")
		return
	}
	fmt.Print("O menor numero da lista é", menor)
}
func ordenar(list []float64) (float64, error) {
	sort.Float64s(list)
	if len(list) == 0 {
		return 0, fmt.Errorf("A lista está vazia")
	}
	y := list[0]
	return y, nil
}
