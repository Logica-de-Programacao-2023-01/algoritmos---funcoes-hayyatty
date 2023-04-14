package main

import "fmt"

func main() {
	var (
		x, y float32
	)
	fmt.Println("Digite 2 numeros para serem divididos")
	fmt.Scan(&x)
	fmt.Scan(&y)
	resultado, err := dividir(x, y)
	if err != nil {
		fmt.Println("Ocorreu um erro ao dividir", x, "e", y)
		return
	}
	fmt.Print(resultado)
}
func dividir(x float32, y float32) (float32, error) {
	if y == 0 {
		return 0, fmt.Errorf("Nao pode dividir por zero")
	}
	return x / y, nil
}
