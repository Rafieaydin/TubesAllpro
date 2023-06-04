package main

import (
	"fmt"
)

type prabrik struct {
	nama      string
	alamat    string
	noTelp    string
	listMobil [100]mobil
	nMobil    int
}

type mobil struct {
	merk            string
	tipe            string
	warna           string
	harga           int
	total_penjualan int
}

type listPrabrik struct {
	prabrik [100]prabrik
	nPabrik int
}

func findIdx(list listPrabrik, key string, category string) int {
	if category == "pabrik" {
		for i := 0; i < list.nPabrik; i++ {
			if key == list.prabrik[i].nama {
				return i
			}
		}
	} else if category == "mobil" {
		for i := 0; i < list.nPabrik; i++ {
			for j := 0; j < list.prabrik[i].nMobil; j++ {
				if key == list.prabrik[i].listMobil[j].merk {
					return j
				}
			}
		}
	}
	return -1
}

// buat function untuk menambahkan prabrik
func addPabrik(list *listPrabrik) {
	fmt.Print("Masukkan nama prabrik: ")
	fmt.Scanln(&list.prabrik[list.nPabrik].nama)
	fmt.Print("Masukkan alamat prabrik: ")
	fmt.Scanln(&list.prabrik[list.nPabrik].alamat)
	fmt.Print("Masukkan no telp prabrik: ")
	fmt.Scanln(&list.prabrik[list.nPabrik].noTelp)
	list.nPabrik++
	fmt.Println()
	fmt.Println("Data berhasil di tammbahkan")
	fmt.Println()
	menuPabrik(list)
}

func editPabrik(list *listPrabrik) {
	var nama string

	fmt.Println("Data Silahkan masukan nama yang ingin di ubah")
	fmt.Scanln(&nama)
	id := findIdx(*list, nama, "pabrik")
	if id != -1 {
		fmt.Println("Masukan data yang ingin di ubah")
		fmt.Println("Masukkan nama prabrik: ")
		fmt.Scanln(&list.prabrik[id].nama)
		fmt.Println("Masukkan alamat prabrik: ")
		fmt.Scanln(&list.prabrik[id].alamat)
		fmt.Print("Masukkan no telp prabrik: ")
		fmt.Scanln(&list.prabrik[id].noTelp)
		fmt.Println()
		fmt.Println("Data berhasil di update")
		fmt.Println()
	} else {
		fmt.Println("Nama tidak ditemukan")
		fmt.Println("Silahkan input data lagi : ")
		fmt.Scanln(&nama)
		editPabrik(list)
	}
}

func delPabrik(list *listPrabrik) {
	var nama string
	fmt.Println("Silahkan masukkan nama pabrik yang ingin dihapus:")
	fmt.Scanln(&nama)
	id := findIdx(*list, nama, "pabrik")
	if id != -1 {
		for i := id; i < list.nPabrik-1; i++ {
			list.prabrik[i] = list.prabrik[i+1]
		}
		list.nPabrik--
		fmt.Println("Data berhasil dihapus")
		fmt.Println()
		menuPabrik(list)
	} else {
		fmt.Println("Nama tidak ditemukan")
		fmt.Println("Silahkan input data lagi:")
		fmt.Scanln(&nama)
		delPabrik(list)
	}
}

func listPabrik(list *listPrabrik) {
	if list.nPabrik != 0 {
		fmt.Println()
		fmt.Println("Berikut list pabrik yang tersedia")
		fmt.Println()
		for i := 0; i < list.nPabrik; i++ {
			fmt.Println("Nama :", list.prabrik[i].nama)
			fmt.Println("alamat :", list.prabrik[i].alamat)
			fmt.Println("alamat :", list.prabrik[i].noTelp)
		}
	} else {
		fmt.Println("pabrik masih kosong")
	}

}
func menu(list *listPrabrik) {
	var pick int
	fmt.Println("===== Selamat Datang di DEMOB =========================")
	fmt.Println("===== Dealer Mobil Terpercaya dan amanah ==============")
	fmt.Println("===== Selamat Datang User =============================")
	fmt.Println("===== Berikut Feature yang di sediakan  ===============")
	fmt.Println("1. Kelola Daftar pabrik")
	fmt.Println("2. Kelola Mobil dealer")
	fmt.Println("3. daftar Mobil dengan penjualan tertinggi")
	fmt.Println("===== Silahkan Pilih feature ===============")
	fmt.Scanln(&pick)
	switch pick {
	case 1:
		menuPabrik(list)
	case 2:

	}
}
func varepeat() int {
	var input int
	// fmt.Println()
	fmt.Println("Apakah ingin kembali ? ")
	fmt.Println("1. Ya")
	fmt.Println("2. Tidak")
	fmt.Scanln(&input)
	switch input {
	case 1:
		return 1
	case 2:
		return 0
	}
	return 0
}

func menuPabrik(list *listPrabrik) {
	var pick int
	fmt.Println()
	fmt.Println("==== Selamat datang di menu pabrik ===========")
	fmt.Println("==== Berikut feature yang di sediakan  =======")
	fmt.Println("1. List Daftar pabrik")
	fmt.Println("2. Tambah daftar pabrik")
	fmt.Println("3. edit data pabrik")
	fmt.Println("4. hapus data pabrik")
	fmt.Println("5. kembali")
	fmt.Println("===== Silahkan Pilih feature ===============")
	fmt.Scanln(&pick)
	switch pick {
	case 1:
		if list.nPabrik == 0 {
			fmt.Println()
			fmt.Println("Pabrik masih kosong")
			fmt.Println("Silahkan kembali")
			fmt.Println()
			if varepeat() != 0 {
				menuPabrik(list)
			}
			fmt.Println()
		} else {
			listPabrik(list)
			if varepeat() != 0 {
				menuPabrik(list)
			}
			fmt.Println()
		}
	case 2:
		addPabrik(list)
		fmt.Println()
		menuPabrik(list)
	case 3:
		if list.nPabrik != 0 {
			fmt.Println()
			fmt.Println("berikut list nama pabrik yang tersedia")
			listPabrik(list)
			fmt.Println()
			editPabrik(list)
			menuPabrik(list)
		} else {
			fmt.Println("data pabrik masih kosong")
			if varepeat() != 0 {
				menuPabrik(list)
			}
		}
	case 4:
		if list.nPabrik != 0 {
			fmt.Println()
			fmt.Println("Berikut list nama pabrik yang tersedia:")
			listPabrik(list)
			fmt.Println()
			delPabrik(list)
			menuPabrik(list)
		} else {
			fmt.Println()
			fmt.Println("Data pabrik masih kosong")
			fmt.Println()
			if varepeat() != 0 {
				menuPabrik(list)
			}
		}
	case 5:
		menu(list)
	}
}
func main() {
	var list listPrabrik
	menu(&list)
}
