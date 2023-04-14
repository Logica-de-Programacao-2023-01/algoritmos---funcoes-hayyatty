package main

import (
	"fmt"
)

func main() {
	var (
		x    float64
		soma float64
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
	fmt.Println("A lista é ", list)
	media, err := media(list, soma)
	if err != nil {
		fmt.Println("Ocorreu um erro em fazer a media, a lista esta vazia")
		return
	}
	fmt.Print("A media da lista é", media)
}
func media(list []float64, soma float64) (float64, error) {
	y := float64(len(list))
	med := soma / y
	if len(list) == 0 {
		return 0, fmt.Errorf("A lista esta vazia")
	}

	for _, i := range list {
		soma += i
	}
	return med, nil
	return 0, nil
}
