package main

import (
	"fmt"
)

type Mobil struct {
	Model       string
	TahunKeluar int
	Harga       int
	Penjualan   int
}

type Pabrikan struct {
	Nama   string
	Mobil  [100]Mobil
	nMobil int
}

var daftarPabrikan [100]Pabrikan
var totalPabrikan int

func main() {
	menuUtama()
}

func cariIndexPabrikan(namaPabrikan string) int {
	// var i = 0

	for i := 0; i < totalPabrikan; i++ {
		if daftarPabrikan[i].Nama == namaPabrikan {
			return i
		}
	}
	return -1
}

func cariIndexMobil(pabrikan *Pabrikan, modelMobil string) int {
	for i := 0; i < pabrikan.nMobil; i++ {
		if pabrikan.Mobil[i].Model == modelMobil {
			return i
		}
	}
	return -1
}

func getPabrikanByMobil(mobil Mobil) Pabrikan {
	for i := 0; i < totalPabrikan; i++ {
		for _, m := range daftarPabrikan[i].Mobil {
			if m == mobil {
				return daftarPabrikan[i]
			}
		}
	}
	return Pabrikan{}
}

func menuUtama() {
	for {
		fmt.Println("===== Aplikasi Penjualan Mobil =====")
		fmt.Println("1. Tambah Data Pabrikan")
		fmt.Println("2. Edit Data Pabrikan")
		fmt.Println("3. Hapus Data Pabrikan")
		fmt.Println("4. Tambah Data Mobil")
		fmt.Println("5. Edit Data Mobil")
		fmt.Println("6. Hapus Data Mobil")
		fmt.Println("7. Tampilkan Daftar Mobil Berdasarkan Pabrikan")
		fmt.Println("8. Tampilkan Daftar Mobil Berdasarkan Kriteria")
		fmt.Println("9. Tampilkan 3 Daftar Mobil dan Pabrikan dengan Penjualan Tertinggi")
		fmt.Println("0. Keluar")
		fmt.Print("Pilih menu: ")

		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 99:
			if totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama()
				}
			} else {
				fmt.Println()
				lihatPabrikan()
			}
		case 69:
			if totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama()
				}
			} else {
				fmt.Println()
				lihatMobilPabrikan()
			}
		case 1:
			tambahPabrikan()
		case 2:
			if totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama()
				}
			} else {
				fmt.Println()
				EditPabrikan()
			}
		case 3:
			if totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama()
				}
			} else {
				fmt.Println()
				HapusPabrikan()
			}
		case 4:
			if totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama()
				}
			} else {
				fmt.Println()
				tambahMobil()
			}
		case 5:
			if totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama()
				}
			} else {
				fmt.Println()
				editMobil()
			}
		case 6:
			if totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama()
				}
			} else {
				fmt.Println()
				hapusMobil()
			}
		case 7:
			if totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama()
				}
			} else {
				fmt.Println()
				tampilkanMobilByPabrikan()
			}
		case 8:
			tampilkanMobilByKriteria()
		case 9:
			tampilkanTop3Penjualan()
		case 0:
			fmt.Println("Terima kasih telah menggunakan aplikasi ini. Sampai jumpa!")
			return
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
		}

		fmt.Println()
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

func tambahPabrikan() {
	fmt.Println("===== Tambah Data Pabrikan =====")
	fmt.Print("Masukkan nama pabrikan: ")

	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	idx := totalPabrikan
	daftarPabrikan[idx].Nama = namaPabrikan
	// daftarPabrikan[idx].Mobil = []Mobil{}
	daftarPabrikan[idx].nMobil = 0
	totalPabrikan++
	fmt.Println("Data pabrikan berhasil ditambahkan.")
	fmt.Println()
	if varepeat() != 0 {
		menuUtama()
	} else {
		tambahPabrikan()
	}
}
func EditPabrikan() {
	fmt.Println("===== Berikut pabrik mobil yg tersediav =====")
	lihatPabrikan()
	fmt.Println("===== Silahkan pilih pabrik mobil =====")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(namaPabrikan)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama()
		} else {
			EditPabrikan()
		}
	}
	var nama string
	fmt.Println("Masukan Edit pabrik")
	fmt.Scanln(&nama)
	daftarPabrikan[index].Nama = nama
	if varepeat() != 0 {
		menuUtama()
	} else {
		EditPabrikan()
	}
}

func HapusPabrikan() {
	fmt.Println("===== Berikut pabrik mobil yg tersediav =====")
	lihatPabrikan()
	fmt.Println("===== Masukan pabrik yang pengen di hapus =====")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(namaPabrikan)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama()
		} else {
			HapusPabrikan()
		}
	}
	fmt.Println(index)
	for i := index; i < totalPabrikan-1; i++ {
		daftarPabrikan[index] = daftarPabrikan[index+1]
	}
	totalPabrikan--
	fmt.Println("Data berhasil di hapus", " ")
	if varepeat() != 0 {
		menuUtama()
	} else {
		HapusPabrikan()
	}
}

func lihatPabrikan() {
	fmt.Print()
	fmt.Println("===== Lihat Data Pabrikan =====")
	fmt.Println("===== List Data Pabrikan =====")
	fmt.Print()
	for i := 0; i < totalPabrikan; i++ {
		fmt.Println("Nama pabrik :", daftarPabrikan[i].Nama)
	}
}

func lihatMobilPabrikan() {
	fmt.Print()
	fmt.Println("===== List Data Mobil =====")
	fmt.Print()
	for i := 0; i < totalPabrikan; i++ {
		for j := 0; j < daftarPabrikan[i].nMobil; j++ {
			fmt.Println()
			fmt.Println("Nama pabrik :", daftarPabrikan[i].Nama)
			fmt.Println("Model mobil :", daftarPabrikan[i].Mobil[j].Model)
			fmt.Println("Harga :", daftarPabrikan[i].Mobil[j].Harga)
			fmt.Println("Penjualan :", daftarPabrikan[i].Mobil[j].Penjualan)
			fmt.Println("Tahun keluar :", daftarPabrikan[i].Mobil[j].TahunKeluar)
			fmt.Println()
		}
	}
}

func tambahMobil() {
	fmt.Println("===== Berikut pabrik mobil yg tersediav =====")
	lihatPabrikan()
	fmt.Println("===== Silahkan pilih pabrik mobil =====")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(namaPabrikan)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama()
		} else {
			tambahMobil()
		}
	}
	fmt.Print("Masukkan model mobil: ")
	fmt.Scanln(&daftarPabrikan[index].Mobil[daftarPabrikan[index].nMobil].Model)
	fmt.Print("Masukkan tahun keluaran mobil: ")
	fmt.Scanln(&daftarPabrikan[index].Mobil[daftarPabrikan[index].nMobil].TahunKeluar)
	fmt.Print("Masukkan harga mobil: ")
	fmt.Scanln(&daftarPabrikan[index].Mobil[daftarPabrikan[index].nMobil].Harga)
	fmt.Print("Masukkan jumlah penjualan mobil: ")
	fmt.Scanln(&daftarPabrikan[index].Mobil[daftarPabrikan[index].nMobil].Penjualan)
	daftarPabrikan[index].nMobil++
	fmt.Println("Data mobil berhasil ditambahkan.")
	fmt.Println()
	if varepeat() != 0 {
		menuUtama()
	} else {
		tambahMobil()
	}
}

func editMobil() {
	fmt.Println("===== Berikut pabrik mobil yg tersediav =====")
	lihatMobilPabrikan()
	fmt.Println("===== Silahkan pilih pabrik mobil =====")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(namaPabrikan)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama()
		} else {
			editMobil()
		}

	}

	pabrikan := &daftarPabrikan[index]
	fmt.Print("Masukkan model mobil yang ingin diubah: ")

	var modelMobil string
	fmt.Scanln(&modelMobil)

	mobilIndex := cariIndexMobil(pabrikan, modelMobil)
	if mobilIndex == -1 {
		fmt.Println("Mobil tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama()
		} else {
			editMobil()
		}
	}

	var mobil Mobil
	fmt.Print("Masukkan model mobil baru: ")
	fmt.Scanln(&mobil.Model)
	fmt.Print("Masukkan tahun keluaran mobil baru: ")
	fmt.Scanln(&mobil.TahunKeluar)
	fmt.Print("Masukkan harga mobil baru: ")
	fmt.Scanln(&mobil.Harga)
	fmt.Print("Masukkan jumlah penjualan mobil baru: ")
	fmt.Scanln(&mobil.Penjualan)

	pabrikan.Mobil[mobilIndex] = mobil

	fmt.Println("Data mobil berhasil diubah.")
	if varepeat() != 0 {
		menuUtama()
	} else {
		editMobil()
	}
}

func hapusMobil() {
	fmt.Println("===== Berikut pabrik mobil yg tersediav =====")
	lihatMobilPabrikan()
	fmt.Println("===== Silahkan pilih pabrik mobil =====")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(namaPabrikan)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama()
		} else {
			hapusMobil()
		}
	}

	pabrikan := &daftarPabrikan[index]
	fmt.Print("Masukkan model mobil yang ingin hapus: ")
	var modelMobil string
	fmt.Scanln(&modelMobil)
	mobilIndex := cariIndexMobil(pabrikan, modelMobil)
	if mobilIndex == -1 {
		fmt.Println("Mobil tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama()
		} else {
			hapusMobil()
		}
	}

	for i := mobilIndex; i < daftarPabrikan[index].nMobil-1; i++ {
		daftarPabrikan[index].Mobil[i] = daftarPabrikan[index].Mobil[i+1]
	}

	daftarPabrikan[index].nMobil--

	fmt.Println("Data mobil berhasil dihapus.")
	fmt.Println()
	if varepeat() != 0 {
		menuUtama()
	} else {
		hapusMobil()
	}
}

func tampilkanMobilByPabrikan() {

	fmt.Println("===== Tampilkan Daftar Mobil Berdasarkan Pabrikan =====")
	fmt.Println("===== Berikut List pabrikan yg tersedia")
	lihatPabrikan()
	fmt.Println()

	fmt.Print("Masukkan nama pabrikan: ")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(namaPabrikan)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama()
		} else {
			tampilkanMobilByPabrikan()
		}

	}

	fmt.Println("Daftar mobil dari pabrikan", daftarPabrikan[index].Nama, ":")

	for i := 0; i < daftarPabrikan[index].nMobil; i++ {
		fmt.Printf("Model: %s, Tahun Keluar: %d, Harga: Rp%d, Penjualan: %d\n", daftarPabrikan[index].Mobil[i].Model, daftarPabrikan[index].Mobil[i].TahunKeluar, daftarPabrikan[index].Mobil[i].Harga, daftarPabrikan[index].Mobil[i].Penjualan)
	}
	// if varepeat() != 0 {
	// 	menuUtama()
	// } else {
	// 	tampilkanMobilByPabrikan()
	// }
}

func tampilkanMobilByKriteria() {
	fmt.Println("===== Tampilkan Daftar Mobil Berdasarkan Kriteria =====")
	fmt.Println("1. Tampilkan mobil berdasarkan tahun keluaran")
	fmt.Println("2. Tampilkan mobil berdasarkan harga")
	fmt.Println("3. Tampilkan mobil berdasarkan jumlah penjualan")
	fmt.Print("Pilih kriteria: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tampilkanMobilByTahunKeluar()
	case 2:
		tampilkanMobilByHarga()
	case 3:
		tampilkanMobilByPenjualan()
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
	if varepeat() != 0 {
		menuUtama()
	} else {
		tampilkanMobilByKriteria()
	}
}

func tampilkanMobilByTahunKeluar() {
	fmt.Print("Masukkan tahun keluaran mobil: ")

	var tahun int
	fmt.Scanln(&tahun)

	var mobil []Mobil

	for i := 0; i < totalPabrikan; i++ {
		for j := 0; j < daftarPabrikan[i].nMobil; j++ {
			if daftarPabrikan[i].Mobil[j].TahunKeluar == tahun {
				mobil = append(mobil, daftarPabrikan[i].Mobil[j])
			}
		}
	}

	if len(mobil) == 0 {
		fmt.Println("Tidak ditemukan mobil dengan tahun keluaran tersebut.")
	}

	fmt.Println("Daftar mobil dengan tahun keluaran", tahun, ":")

	for i := 0; i < len(mobil); i++ {
		pabrikan := getPabrikanByMobil(mobil[i])
		fmt.Printf("Model: %s, Pabrikan: %s, Harga: Rp%d, Penjualan: %d\n", mobil[i].Model, pabrikan.Nama, mobil[i].Harga, mobil[i].Penjualan)
	}
	if varepeat() != 0 {
		menuUtama()
	} else {
		tampilkanMobilByTahunKeluar()
	}
}

func tampilkanMobilByHarga() {
	fmt.Print("Masukkan rentang harga mobil (contoh: 10000000-20000000): ")

	var hargaMin, hargaMax int
	fmt.Scanf("%d-%d", &hargaMin, &hargaMax)

	var mobil []Mobil

	for i := 0; i < totalPabrikan; i++ {
		for j := 0; j < daftarPabrikan[i].nMobil; j++ {
			if daftarPabrikan[i].Mobil[j].Harga >= hargaMin && daftarPabrikan[i].Mobil[j].Harga <= hargaMax {
				mobil = append(mobil, daftarPabrikan[i].Mobil[j])
			}
		}
	}

	if len(mobil) == 0 {
		fmt.Println("Tidak ditemukan mobil dengan harga dalam rentang tersebut.")
		if varepeat() != 0 {
			menuUtama()
		} else {
			tampilkanMobilByHarga()
		}
	}

	fmt.Println("Daftar mobil dengan harga antara Rp", hargaMin, "dan Rp", hargaMax, ":")

	for i := 0; i < len(mobil); i++ {
		pabrikan := getPabrikanByMobil(mobil[i])
		fmt.Printf("Model: %s, Pabrikan: %s, Harga: Rp%d, Penjualan: %d Tahun: %d\n", mobil[i].Model, pabrikan.Nama, mobil[i].Harga, mobil[i].Penjualan, mobil[i].TahunKeluar)
	}
	if varepeat() != 0 {
		menuUtama()
	} else {
		tampilkanMobilByHarga()
	}
}

func tampilkanMobilByPenjualan() {
	fmt.Print("Masukkan jumlah penjualan mobil: ")

	var penjualan int
	fmt.Scanln(&penjualan)

	var mobil []Mobil

	for i := 0; i < totalPabrikan; i++ {
		for j := 0; j < daftarPabrikan[i].nMobil; j++ {
			if daftarPabrikan[i].Mobil[j].Penjualan == penjualan {
				mobil = append(mobil, daftarPabrikan[i].Mobil[j])
			}
		}
	}

	if len(mobil) == 0 {
		fmt.Println("Tidak ditemukan mobil dengan jumlah penjualan tersebut.")
	}

	fmt.Println("Daftar mobil dengan jumlah penjualan", penjualan, ":")

	for i := 0; i < len(mobil); i++ {
		pabrikan := getPabrikanByMobil(mobil[i])
		fmt.Printf("Model: %s, Pabrikan: %s, Harga: Rp%d, Penjualan: %d Tahun: %d\n", mobil[i].Model, pabrikan.Nama, mobil[i].Harga, mobil[i].Penjualan, mobil[i].TahunKeluar)
	}
	if varepeat() != 0 {
		menuUtama()
	} else {
		tampilkanMobilByPenjualan()
	}
}

func tampilkanTop3Penjualan() {
	fmt.Println("===== Tampilkan 3 Daftar Mobil dan Pabrikan dengan Penjualan Tertinggi =====")

	var mobil []Mobil

	for i := 0; i < totalPabrikan; i++ {
		for j := 0; j < daftarPabrikan[i].nMobil; j++ {
			mobil = append(mobil, daftarPabrikan[i].Mobil[j]) // insert arr to mobil
		}
	}

	if len(mobil) == 0 {
		fmt.Println("Belum ada data mobil.")
		if varepeat() != 0 {
			menuUtama()
		} else {
			tampilkanTop3Penjualan()
		}
	}

	// Selection sort
	pass := 0
	for pass <= len(mobil)-1 {
		maxIndex := pass // idx =0
		for i := pass; i < len(mobil); i++ {
			if mobil[i].Penjualan > mobil[maxIndex].Penjualan {
				maxIndex = i
			}
		}
		//mobil[i], mobil[maxIndex] = mobil[maxIndex], mobil[i]
		var temp = mobil[pass] // i = 0
		mobil[pass] = mobil[maxIndex]
		mobil[maxIndex] = temp
		pass++

	}

	fmt.Println("Daftar mobil dan pabrikan dengan penjualan tertinggi:")

	for i := 0; i < 3 && i < len(mobil); i++ {
		m := mobil[i]
		pabrikan := getPabrikanByMobil(m)
		fmt.Printf("Model: %s, Pabrikan: %s, Penjualan: %d\n", m.Model, pabrikan.Nama, m.Penjualan)
	}

	if varepeat() != 0 {
		menuUtama()
	} else {
		tampilkanTop3Penjualan()
	}
}

// func tampilkanTop3Penjualan() {
// 	fmt.Println("===== Tampilkan 3 Daftar Mobil dan Pabrikan dengan Penjualan Tertinggi =====")

// 	var mobil []Mobil

// 	for i := 0; i < totalPabrikan; i++ {
// 		for j := 0; j < daftarPabrikan[i].nMobil; j++ {
// 			mobil = append(mobil, daftarPabrikan[i].Mobil[j]) // insert arr to mobil
// 		}
// 	}

// 	if len(mobil) == 0 {
// 		fmt.Println("Belum ada data mobil.")
// 		if varepeat() != 0 {
// 			menuUtama()
// 		} else {
// 			tampilkanTop3Penjualan()
// 		}
// 		return
// 	}

// 	// Selection sort
// 	for pass := 0; pass < len(mobil)-1; pass++ {
// 		maxIndex := pass
// 		for i := pass + 1; i < len(mobil); i++ {
// 			if mobil[i].Penjualan > mobil[maxIndex].Penjualan {
// 				maxIndex = i
// 			}
// 		}
// 		mobil[pass], mobil[maxIndex] = mobil[maxIndex], mobil[pass]
// 	}

// 	fmt.Println("Daftar mobil dan pabrikan dengan penjualan tertinggi:")

// 	for i := 0; i < 3 && i < len(mobil); i++ {
// 		m := mobil[i]
// 		pabrikan := getPabrikanByMobil(m)
// 		fmt.Printf("Model: %s, Pabrikan: %s, Penjualan: %d\n", m.Model, pabrikan.Nama, m.Penjualan)
// 	}

// 	if varepeat() != 0 {
// 		menuUtama()
// 	} else {
// 		tampilkanTop3Penjualan()
// 	}
// }
