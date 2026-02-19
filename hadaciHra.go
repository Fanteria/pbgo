package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

func HadaciHra() {

	fmt.Println("hadej moje čislo")
	var hadaneCislo string

	cisloKUhadnuti := rand.Int64N(20)
	pocetPokusu := 0
	hranicePoctuPokusu := 10
	for {
		pocetPokusu = pocetPokusu + 1
		fmt.Scan(&hadaneCislo)
		cislo, chyba := strconv.ParseInt(hadaneCislo, 10, 0)
		if chyba == nil {
			if cislo > cisloKUhadnuti {
				fmt.Println("uber")
			}
			if cislo < cisloKUhadnuti {
				fmt.Println("přidej")
			}
			if (cislo) == cisloKUhadnuti {
				fmt.Println("Hura,uhadl jsi to!!!!!!")
				break
			}

		} else {
			fmt.Println("Špatně, chlape")
		}
		if pocetPokusu > hranicePoctuPokusu {
			fmt.Println("číslo které hádáš je", cisloKUhadnuti-cislo, "od čísla které jsi zadal")
		}
	}
}
