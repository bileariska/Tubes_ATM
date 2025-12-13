package main

import "fmt"

func main() {
	var saldo, akhir_saldo, saldo_penarikan int
	var status_kata string
	var pin, pin_user, pilihan int
	var Status_ulang bool

	Status_ulang = true
	pin_user = 12345678
	saldo = 200000

	for Status_ulang {
		fmt.Println("Selamat datang di ATM")

		fmt.Println("Masukan Pin anda")
		fmt.Scan(&pin)
		if pin == pin_user {
			fmt.Println("apakah anda ingin memilih opsi ?")
			fmt.Println("1. Cek saldo")
			fmt.Println("2. Tarik saldo")
			fmt.Println("3. Setor saldo")
			fmt.Scan(&pilihan)

			if pilihan == 1 {
				saldo = saldo + saldo_penarikan
				fmt.Println("saldo anda :", saldo)
			} else if pilihan == 2 {
				fmt.Printf("Saldo anda sisa: %d\n", saldo)
				fmt.Println("Masukkan nominal saldo yang ingin anda tarik? ")
				fmt.Scan(&saldo_penarikan)
				akhir_saldo = saldo - saldo_penarikan
				fmt.Println("saldo anda sisa", akhir_saldo)
			} else if pilihan == 3 {
				fmt.Printf("Saldo anda sisa: %d\n", saldo)
				fmt.Println("Masukkan nominal saldo yang ingin anda setor? ")
				fmt.Scan(&saldo_penarikan)
				akhir_saldo = saldo + saldo_penarikan
				fmt.Println("saldo anda menjadi", akhir_saldo)
			} else {
				fmt.Println("input yang anda masukan tidak ada")
			}
		} else {
			fmt.Println("pin anda salah")
		}

		fmt.Println("Anda ingin Menginput lagi ? ya | tidak")
		fmt.Scan(&status_kata)

		if status_kata == "ya" {
			Status_ulang = true
		} else {
			Status_ulang = false
		}
	}

}
