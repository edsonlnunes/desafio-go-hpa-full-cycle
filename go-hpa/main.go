package main


import (
	"math"
	"fmt"
	"net/http"
)

func makeCalc() float64{

	var valor float64 = 0.000002;

	for i := 1; i < 100000000; i++ {
		valor += math.Sqrt(valor);
	}

	return valor;
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		result := makeCalc();
		fmt.Fprintf(w, "Result: %f", result);
	})

	http.ListenAndServe(":8000", nil);
}