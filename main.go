package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type loker struct {
	Name   string
	Number string
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Init jumlah loker : ")

	scanner.Scan()

	input, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Terdapat pilihan :")
	fmt.Println("1. Input data")
	fmt.Println("2. Tampilkan loker")
	fmt.Println("3. Hapus loker")
	fmt.Println("4. Cari loker")

	fmt.Print("Pilih nomor : ")
	scanner.Scan()

	inputNo, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	switch inputNo {
	case 1:
		inputData(input, scanner)
	case 2:
		showLoker()
	case 3:
		removeLoker()
	case 4:
		searchLoker()
	}

	fmt.Println("Selesai")
}

func inputData(input int, scanner *bufio.Scanner) {
	fmt.Print("Input loker : ")
	scanner.Scan()

	// dataInput := strings.Split(scanner.Text(), " ")

	// data := make([]loker, dataInput)
}

func removeLoker() {
	fmt.Println("remove loker")
}

func showLoker() {
	fmt.Println("show loker")
}

func searchLoker() {
	fmt.Println("search loker")
}

// masukkan jumlah loker = 3
// pilihan
// - input
// - tampilkan loker
// - hapus loker
// - cari loker

// jika input
