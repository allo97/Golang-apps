package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Magazyn musi być stworzony zanim się zacznie pobieranie produktów
	var magazyn = generateMagazyn(50, 100)

	babka1 := babka{"babka1", "znicze"}
	babka2 := babka{"babka2", "znicze"}
	babka3 := babka{"babka3", "wiazanki"}
	babka4 := babka{"babka4", "wiazanki"}

	koszNaZnicze := make(chan string, 10)
	koszNaWiazanki := make(chan string, 10)

	poslancy := []poslaniec{{"poslaniec1"}, {"poslaniec2"}, {"poslaniec3"}, {"poslaniec4"}, {"poslaniec5"}}

	var wg sync.WaitGroup

	fmt.Println("Uruchamiam babki")

	wg.Add(4)
	go babka1.pobierzProduktZMagazynu(magazyn.znicze, koszNaZnicze, &wg)
	go babka2.pobierzProduktZMagazynu(magazyn.znicze, koszNaZnicze, &wg)
	go babka3.pobierzProduktZMagazynu(magazyn.wiazanki, koszNaWiazanki, &wg)
	go babka4.pobierzProduktZMagazynu(magazyn.wiazanki, koszNaWiazanki, &wg)

	fmt.Println("Uruchamiam poslancow")

	for _, poslaniec := range poslancy {
		wg.Add(1)
		go poslaniec.pobierzProduktyZKosza(koszNaZnicze, koszNaWiazanki, magazyn, &wg)
	}

	wg.Wait()
}

func generateMagazyn(iloscWiazanek int, iloscZniczy int) magazyn {
	magazyn := magazyn{znicze: make(chan string, iloscZniczy), wiazanki: make(chan string, iloscWiazanek)}

	for i := 0; i < iloscWiazanek; i++ {
		magazyn.wiazanki <- "wiazanka" + strconv.Itoa(i+1)
	}

	for i := 0; i < iloscZniczy; i++ {
		magazyn.znicze <- "znicz" + strconv.Itoa(i+1)
	}

	fmt.Println("Magazyn został stworzony")
	return magazyn
}

func (babka babka) pobierzProduktZMagazynu(magazyn chan string, kosz chan string, wg *sync.WaitGroup) {
	for len(magazyn) != 0 {
		if len(kosz) >= 10 {
			fmt.Printf("Pełny kosz! %s musi poczekać aż posłaniec odbierze produkt!\n", babka.nazwa)
		} else {
			produkt := <-magazyn
			kosz <- produkt
			fmt.Printf("%s pobiera %s i dodaje %s do kosza\n", babka.nazwa, babka.typ, produkt)
		}
		time.Sleep(time.Second)

	}
	defer wg.Done()
}

func (poslaniec poslaniec) pobierzProduktyZKosza(kosz_na_znicze chan string, kosz_na_wiazanki chan string, magazyn magazyn, wg *sync.WaitGroup) {
	for len(magazyn.znicze) != 0 || len(magazyn.wiazanki) != 0 {
		if len(kosz_na_znicze) > 1 && len(kosz_na_wiazanki) > 0 {
			znicz1 := <-kosz_na_znicze
			znicz2 := <-kosz_na_znicze
			wiazanka := <-kosz_na_wiazanki

			fmt.Printf("%s pobiera %s i %s i %s\n", poslaniec.nazwa, znicz1, znicz2, wiazanka)
		} else {
			fmt.Printf("Pusty kosz! %s musi poczekać!\n", poslaniec.nazwa)
		}

		time.Sleep(time.Second)

	}
	defer wg.Done()
}

type magazyn struct {
	znicze   chan string
	wiazanki chan string
}

type babka struct {
	nazwa string
	typ   string
}

type poslaniec struct {
	nazwa string
}
