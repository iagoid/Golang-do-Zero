package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	ultimoPrimo := pegaUltimoValorEscritoTxt()

	fmt.Printf("Iniciando do número %d\n", ultimoPrimo)
	proximoPrimo(ultimoPrimo)
}

func pegaUltimoValorEscritoTxt() uint64 {
	f, err := os.Open("start_from.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	lastLine := scanner.Text()

	ultimoPrimo, err := strconv.ParseUint(lastLine, 0, 64)
	if err != nil {
		if strings.Contains(err.Error(), "invalid syntax") {
			return 0
		}
		log.Fatal(err)
	}
	return ultimoPrimo
}

func proximoPrimo(ultimoPrimo uint64) {
	var possivelPrimo uint64

	if ultimoPrimo <= 0 { //txt esta vazio
		SalvarNumeroNoTxt(2) //salva o unico par primo
		possivelPrimo = 3
	} else { //caso tenha conteudo no txt
		possivelPrimo = ultimoPrimo + 1
	}

	for true {
		if EhPrimo(possivelPrimo) {
			SalvarNumeroNoTxt(possivelPrimo)
		}
		possivelPrimo = possivelPrimo + 2
	}
}

func EhPrimo(possivelPrimo uint64) bool {
	var divisor uint64

	if EhPar(possivelPrimo) {
		return false
	}

	for divisor = 3; divisor < possivelPrimo; divisor = divisor + 2 {
		if possivelPrimo%divisor == 0 {
			return false
		}
	}
	return true
}

func EhPar(numero uint64) bool {
	if numero%2 == 0 {
		return true
	}
	return false
}

func SalvarNumeroNoTxt(numero uint64) {
	fmt.Println("Salvando ", numero)

	numeroString := fmt.Sprintf("%d\n", numero)
	f, err := os.OpenFile("numeros_primos.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	if _, err := f.WriteString(numeroString); err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile("start_from.txt", []byte(numeroString), 0)
	if err != nil {
		log.Fatal(err)
	}

}
