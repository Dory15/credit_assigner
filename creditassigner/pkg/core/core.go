package core

import (
	"fmt"
)

type ICreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type CreditAssigner struct{}

func (creditAssigner CreditAssigner) Assign(investment int32) (int32, int32, int32, error) {

	max300 := investment / 300
	max500 := investment / 500
	max700 := investment / 700

	fmt.Println(max300, max500, max700)

	for i := max300; i >= 0; i-- {
		for j := max500; j >= 0; j-- {
			for k := max700; k >= 0; k-- {
				if (i*300 + j*500 + k*700) == investment {
					return i, j, k, nil
				}
			}
		}
	}
	return 0, 0, 0, fmt.Errorf("No se puede asignar el crédito")

}
