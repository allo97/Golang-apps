package main

import (
	"fmt"
	"strconv"
)

func main() {
	var magazyn = generateMagazyn()

	babka1 := babka{nazwa: "babka1", typ: "wiazanka"}
	babka2 := babka{nazwa: "babka2", typ: "wiazanka"}
	babka3 := babka{nazwa: "babka3", typ: "znicz"}
	babka4 := babka{nazwa: "babka4", typ: "znicz"}

	fmt.Println(babka1)
	fmt.Println(babka2)
	fmt.Println(babka3)
	fmt.Println(babka4)

	var kosz_na_znicze []string
	var kosz_na_wiazanki []string

	fmt.Println(magazyn)

	fmt.Println(kosz_na_znicze)
	fmt.Println(kosz_na_wiazanki)
}

func generateMagazyn() magazyn {
	fmt.Println("Tworzę magazyn")

	var magazyn magazyn

	for i := 0; i < 50; i++ {
		magazyn.wiazanki[i] = "wiazanka" + strconv.Itoa(i+1)
	}

	for i := 0; i < 100; i++ {
		magazyn.znicze[i] = "znicz" + strconv.Itoa(i+1)
	}
	return magazyn
}

func pobierzProduktZMagazynu(magazyn magazyn, babka babka, kosz []string) {
	fmt.Println("Babka pobiera produkt")

	var produkt string

	if babka.typ == "wiazanka" {
		produkt = magazyn.wiazanki[len(magazyn.wiazanki)-1]
		magazyn.wiazanki = magazyn.wiazanki[:len(magazyn.wiazanki)-1]
	}

	if babka.typ == "znicz" {
		produkt = magazyn.znicze[len(magazyn.znicze)-1]
		magazyn.znicze = magazyn.znicze[:len(magazyn.znicze)-1]
	}

	fmt.Println("Pobrany produkt: ", produkt)

	fmt.Println("Dodaję produkt do kosza:")

	if len(kosz) >= 10 {
		fmt.Println("Kosz już jest zapełniony! Musisz poczekać aż posłaniec odbierze produkt!")
	} else {
		kosz = append(kosz, produkt)
		fmt.Println("Produkt został dodany do kosza")
	}

}

func pobierzProduktyZKosza(kosz_na_wiazanki []string, kosz_na_znicze []string, poslaniec poslaniec) {
	fmt.Println("Poslaniec pobiera produkty")

	if len(kosz_na_wiazanki) > 0 {
		var wiazanka = kosz_na_wiazanki[len(kosz_na_wiazanki)-1]
		fmt.Println("Pobrano 1 wiazanke")
		fmt.Println(wiazanka)

		kosz_na_wiazanki = kosz_na_wiazanki[:len(kosz_na_wiazanki)-1]
	} else {
		fmt.Println("Nie ma wiazanki do pobrania")
	}

	if len(kosz_na_znicze) > 1 {
		var znicz1 = kosz_na_znicze[len(kosz_na_znicze)-1]
		var znicz2 = kosz_na_znicze[len(kosz_na_znicze)-2]

		fmt.Println("Pobrano 2 znicze")
		fmt.Println(znicz1)
		fmt.Println(znicz2)

		kosz_na_znicze = kosz_na_znicze[:len(kosz_na_znicze)-2]
	} else {
		fmt.Println("Nie ma zniczy do pobrania")
	}

}

type magazyn struct {
	znicze   []string
	wiazanki []string
}

type babka struct {
	nazwa string
	typ   string
}

type poslaniec struct {
	nazwa string
}
