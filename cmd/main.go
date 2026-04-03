package main

import (
	"fmt"

	"github.com/bloodgroup-cplusplus/gigw_ddd_hexagon/internal/adapters/core/arithmetic"
	//"github.com/bloodgroup-cplusplus/gigw_ddd_hexagon/internal/ports"
)

func main() {
	// ports

	//var core ports.ArithmeticPort

	//core = arithmetic.NewAdapter()
	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(1, 3)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(result)

}
