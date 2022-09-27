package main

import (
	"fmt"
	"os"
	"strconv"
)

type DataPerson struct {
	name    string
	address string
	job     string
	reason  string
}

func main2() {

	// fmt.Println(datas[1].name)

	datas := []DataPerson{
		{name: "Giva", address: "Jakarta", job: "Developer", reason: "Belajar"},
		{name: "Iqbal", address: "Jakarta", job: "Developer", reason: "Belajar"},
		{name: "Ucup", address: "Jakarta", job: "Developer", reason: "Belajar"},
		{name: "Nugraha", address: "Jakarta", job: "Developer", reason: "Belajar"},
		{name: "Bayu", address: "Jakarta", job: "Developer", reason: "Belajar"},
		{name: "Putu", address: "Jakarta", job: "Developer", reason: "Belajar"},
		{name: "Satrio", address: "Jakarta", job: "Developer", reason: "Belajar"},
		{name: "Agus", address: "Jakarta", job: "Developer", reason: "Belajar"},
		{name: "Barru", address: "Jakarta", job: "Developer", reason: "Belajar"},
		{name: "Hasan", address: "Jakarta", job: "Developer", reason: "Belajar"},
	}

	printDatas := func(a string) {
		index, _ := strconv.Atoi(a)

		if index == 0 || index > len(datas) {
			fmt.Println("Data Tidak Ditemukan")
			return
		}

		fmt.Println("Nama : ", datas[index-1].name)
		fmt.Println("Alamat : ", datas[index-1].address)
		fmt.Println("Pekerjaan :", datas[index-1].job)
		fmt.Println("Alasan : ", datas[index-1].reason)

		return

	}

	printDatas(os.Args[1])

}
