package main


import (
	"math"
	"fmt"
)

func makeCalc() float64{

	var valor float64 = 0.000002;

	for i := 1; i < 100000000; i++ {
		valor += math.Sqrt(valor);
	}

	return valor;
}

func main() {
	result := makeCalc();

	fmt.Printf("%f", result);
	fmt.Printf("Code.Education Hocks");
}