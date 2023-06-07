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

type daftarPabrikan struct {
	daftarPabrikan [100]Pabrikan
	totalPabrikan  int
}

func main() {
	var daf daftarPabrikan
	menuUtama(&daf)
}

func cariIndexPabrikan(daf daftarPabrikan, namaPabrikan string) int {
	// var i = 0
	for i := 0; i < daf.totalPabrikan; i++ {
		if daf.daftarPabrikan[i].Nama == namaPabrikan {
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

func getPabrikanByMobil(daf daftarPabrikan, mobil Mobil) Pabrikan {
	for i := 0; i < daf.totalPabrikan; i++ {
		for _, m := range daf.daftarPabrikan[i].Mobil {
			if m == mobil {
				return daf.daftarPabrikan[i]
			}
		}
	}
	return Pabrikan{}
}

func menuUtama(daf *daftarPabrikan) {
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
			if daf.totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama(daf)
				}
			} else {
				fmt.Println()
				lihatPabrikan(*daf)
			}
		case 69:
			if daf.totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama(daf)
				}
			} else {
				fmt.Println()
				lihatMobilPabrikan(*daf)
			}
		case 1:
			tambahPabrikan(daf)
		case 2:
			if daf.totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama(daf)
				}
			} else {
				fmt.Println()
				EditPabrikan(daf)
			}
		case 3:
			if daf.totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama(daf)
				}
			} else {
				fmt.Println()
				HapusPabrikan(daf)
			}
		case 4:
			if daf.totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama(daf)
				}
			} else {
				fmt.Println()
				tambahMobil(daf)
			}
		case 5:
			if daf.totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama(daf)
				}
			} else {
				fmt.Println()
				editMobil(daf)
			}
		case 6:
			if daf.totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama(daf)
				}
			} else {
				fmt.Println()
				hapusMobil(daf)
			}
		case 7:
			if daf.totalPabrikan == 0 {
				fmt.Println("Data masih kosong")
				if varepeat() != 0 {
					menuUtama(daf)
				}
			} else {
				fmt.Println()
				tampilkanMobilByPabrikan(*daf)
			}
		case 8:
			tampilkanMobilByKriteria(daf)
		case 9:
			tampilkanTop3Penjualan(*daf)
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

func tambahPabrikan(daf *daftarPabrikan) {
	fmt.Println("===== Tambah Data Pabrikan =====")
	fmt.Print("Masukkan nama pabrikan: ")

	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	idx := daf.totalPabrikan
	daf.daftarPabrikan[idx].Nama = namaPabrikan
	// daftarPabrikan[idx].Mobil = []Mobil{}
	daf.daftarPabrikan[idx].nMobil = 0
	daf.totalPabrikan++
	fmt.Println("Data pabrikan berhasil ditambahkan.")
	fmt.Println()
	if varepeat() != 0 {
		menuUtama(daf)
	} else {
		tambahPabrikan(daf)
	}
}
func EditPabrikan(daf *daftarPabrikan) {
	fmt.Println("===== Berikut pabrik mobil yg tersediav =====")
	lihatPabrikan(*daf)
	fmt.Println("===== Silahkan pilih pabrik mobil =====")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(*daf, namaPabrikan)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama(daf)
		} else {
			EditPabrikan(daf)
		}
	}
	var nama string
	fmt.Println("Masukan Edit pabrik")
	fmt.Scanln(&nama)
	daf.daftarPabrikan[index].Nama = nama
	if varepeat() != 0 {
		menuUtama(daf)
	} else {
		EditPabrikan(daf)
	}
}

func HapusPabrikan(daf *daftarPabrikan) {
	fmt.Println("===== Berikut pabrik mobil yg tersediav =====")
	lihatPabrikan(*daf)
	fmt.Println("===== Masukan pabrik yang pengen di hapus =====")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(*daf, namaPabrikan)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama(daf)
		} else {
			HapusPabrikan(daf)
		}
	}
	fmt.Println(index)
	for i := index; i < daf.totalPabrikan-1; i++ {
		daf.daftarPabrikan[index] = daf.daftarPabrikan[index+1]
	}
	daf.totalPabrikan--
	fmt.Println("Data berhasil di hapus", " ")
	if varepeat() != 0 {
		menuUtama(daf)
	} else {
		HapusPabrikan(daf)
	}
}

func lihatPabrikan(daf daftarPabrikan) {
	fmt.Print()
	fmt.Println("===== Lihat Data Pabrikan =====")
	fmt.Println("===== List Data Pabrikan =====")
	fmt.Print()
	for i := 0; i < daf.totalPabrikan; i++ {
		fmt.Println("Nama pabrik :", daf.daftarPabrikan[i].Nama)
	}
}

func lihatMobilPabrikan(daf daftarPabrikan) {
	fmt.Print()
	fmt.Println("===== List Data Mobil =====")
	fmt.Print()
	for i := 0; i < daf.totalPabrikan; i++ {
		for j := 0; j < daf.daftarPabrikan[i].nMobil; j++ {
			fmt.Println()
			fmt.Println("Nama pabrik :", daf.daftarPabrikan[i].Nama)
			fmt.Println("Model mobil :", daf.daftarPabrikan[i].Mobil[j].Model)
			fmt.Println("Harga :", daf.daftarPabrikan[i].Mobil[j].Harga)
			fmt.Println("Penjualan :", daf.daftarPabrikan[i].Mobil[j].Penjualan)
			fmt.Println("Tahun keluar :", daf.daftarPabrikan[i].Mobil[j].TahunKeluar)
			fmt.Println()
		}
	}
}

func tambahMobil(daf *daftarPabrikan) {
	fmt.Println("===== Berikut pabrik mobil yg tersediav =====")
	lihatPabrikan(*daf)
	fmt.Println("===== Silahkan pilih pabrik mobil =====")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(*daf, namaPabrikan)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama(daf)
		} else {
			tambahMobil(daf)
		}
	}
	fmt.Print("Masukkan model mobil: ")
	fmt.Scanln(&daf.daftarPabrikan[index].Mobil[daf.daftarPabrikan[index].nMobil].Model)
	fmt.Print("Masukkan tahun keluaran mobil: ")
	fmt.Scanln(&daf.daftarPabrikan[index].Mobil[daf.daftarPabrikan[index].nMobil].TahunKeluar)
	fmt.Print("Masukkan harga mobil: ")
	fmt.Scanln(&daf.daftarPabrikan[index].Mobil[daf.daftarPabrikan[index].nMobil].Harga)
	fmt.Print("Masukkan jumlah penjualan mobil: ")
	fmt.Scanln(&daf.daftarPabrikan[index].Mobil[daf.daftarPabrikan[index].nMobil].Penjualan)
	daf.daftarPabrikan[index].nMobil++
	fmt.Println("Data mobil berhasil ditambahkan.")
	fmt.Println()
	if varepeat() != 0 {
		menuUtama(daf)
	} else {
		tambahMobil(daf)
	}
}

func editMobil(daf *daftarPabrikan) {
	fmt.Println("===== Berikut pabrik mobil yg tersediav =====")
	lihatMobilPabrikan(*daf)
	fmt.Println("===== Silahkan pilih pabrik mobil =====")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(*daf, namaPabrikan)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama(daf)
		} else {
			editMobil(daf)
		}

	}

	pabrikan := &daf.daftarPabrikan[index]
	fmt.Print("Masukkan model mobil yang ingin diubah: ")

	var modelMobil string
	fmt.Scanln(&modelMobil)

	mobilIndex := cariIndexMobil(pabrikan, modelMobil)
	if mobilIndex == -1 {
		fmt.Println("Mobil tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama(daf)
		} else {
			editMobil(daf)
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
		menuUtama(daf)
	} else {
		editMobil(daf)
	}
}

func hapusMobil(daf *daftarPabrikan) {
	fmt.Println("===== Berikut pabrik mobil yg tersediav =====")
	lihatMobilPabrikan(*daf)
	fmt.Println("===== Silahkan pilih pabrik mobil =====")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(*daf, namaPabrikan)
	fmt.Println(index)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama(daf)
		} else {
			hapusMobil(daf)
		}
	}

	pabrikan := &daf.daftarPabrikan[index]
	fmt.Print("Masukkan model mobil yang ingin hapus: ")
	var modelMobil string
	fmt.Scanln(&modelMobil)
	mobilIndex := cariIndexMobil(pabrikan, modelMobil)
	if mobilIndex == -1 {
		fmt.Println("Mobil tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama(daf)
		} else {
			hapusMobil(daf)
		}
	}

	for i := mobilIndex; i < daf.daftarPabrikan[index].nMobil-1; i++ {
		daf.daftarPabrikan[index].Mobil[i] = daf.daftarPabrikan[index].Mobil[i+1]
	}

	daf.daftarPabrikan[index].nMobil--

	fmt.Println("Data mobil berhasil dihapus.")
	fmt.Println()
	if varepeat() != 0 {
		menuUtama(daf)
	} else {
		hapusMobil(daf)
	}
}

func tampilkanMobilByPabrikan(daf daftarPabrikan) {

	fmt.Println("===== Tampilkan Daftar Mobil Berdasarkan Pabrikan =====")
	fmt.Println("===== Berikut List pabrikan yg tersedia")
	lihatPabrikan(daf)
	fmt.Println()

	fmt.Print("Masukkan nama pabrikan: ")
	var namaPabrikan string
	fmt.Scanln(&namaPabrikan)
	index := cariIndexPabrikan(daf, namaPabrikan)
	if index == -1 {
		fmt.Println("Pabrikan tidak ditemukan.")
		if varepeat() != 0 {
			menuUtama(&daf)
		} else {
			tampilkanMobilByPabrikan(daf)
		}

	}

	fmt.Println("Daftar mobil dari pabrikan", daf.daftarPabrikan[index].Nama, ":")

	for i := 0; i < daf.daftarPabrikan[index].nMobil; i++ {
		fmt.Printf("Model: %s, Tahun Keluar: %d, Harga: Rp%d, Penjualan: %d\n", daf.daftarPabrikan[index].Mobil[i].Model, daf.daftarPabrikan[index].Mobil[i].TahunKeluar, daf.daftarPabrikan[index].Mobil[i].Harga, daf.daftarPabrikan[index].Mobil[i].Penjualan)
	}
	// if varepeat() != 0 {
	// 	menuUtama()
	// } else {
	// 	tampilkanMobilByPabrikan()
	// }
}

func tampilkanMobilByKriteria(daf *daftarPabrikan) {
	fmt.Println("===== Tampilkan Daftar Mobil Berdasarkan Kriteria =====")
	fmt.Println("1. Tampilkan mobil berdasarkan tahun keluaran")
	fmt.Println("2. Tampilkan mobil berdasarkan harga")
	fmt.Println("3. Tampilkan mobil berdasarkan jumlah penjualan")
	fmt.Print("Pilih kriteria: ")

	var pilihan int
	fmt.Scanln(&pilihan)

	switch pilihan {
	case 1:
		tampilkanMobilByTahunKeluar(*daf)
	case 2:
		tampilkanMobilByHarga(*daf)
		if varepeat() != 0 {
			menuUtama(daf)
		} else {
			tampilkanMobilByTahunKeluar(*daf)
		}
	case 3:
		tampilkanMobilByPenjualan(*daf)
		// if varepeat() != 0 {
		// 	menuUtama(daf)
		// } else {
		// 	tampilkanMobilByTahunKeluar(*daf)
		// }
	default:
		fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
	}
	if varepeat() != 0 {
		menuUtama(daf)
	} else {
		tampilkanMobilByKriteria(daf)
	}
}

func tampilkanMobilByTahunKeluar(daf daftarPabrikan) {
	fmt.Print("Masukkan tahun keluaran mobil: ")

	var tahun int
	fmt.Scanln(&tahun)

	var mobil []Mobil

	for i := 0; i < daf.totalPabrikan; i++ {
		for j := 0; j < daf.daftarPabrikan[i].nMobil; j++ {
			if daf.daftarPabrikan[i].Mobil[j].TahunKeluar == tahun {
				mobil = append(mobil, daf.daftarPabrikan[i].Mobil[j])
			}
		}
	}

	if len(mobil) == 0 {
		fmt.Println("Tidak ditemukan mobil dengan tahun keluaran tersebut.")
	}

	fmt.Println("Daftar mobil dengan tahun keluaran", tahun, ":")

	for i := 0; i < len(mobil); i++ {
		pabrikan := getPabrikanByMobil(daf, mobil[i])
		fmt.Printf("Model: %s, Pabrikan: %s, Harga: Rp%d, Penjualan: %d\n", mobil[i].Model, pabrikan.Nama, mobil[i].Harga, mobil[i].Penjualan)
	}

}

func tampilkanMobilByHarga(daf daftarPabrikan) {
	fmt.Print("Masukkan rentang harga mobil (contoh: 10000000-20000000): ")

	var hargaMin, hargaMax int
	fmt.Scanf("%d-%d", &hargaMin, &hargaMax)

	var mobil []Mobil

	for i := 0; i < daf.totalPabrikan; i++ {
		for j := 0; j < daf.daftarPabrikan[i].nMobil; j++ {
			if daf.daftarPabrikan[i].Mobil[j].Harga >= hargaMin && daf.daftarPabrikan[i].Mobil[j].Harga <= hargaMax {
				mobil = append(mobil, daf.daftarPabrikan[i].Mobil[j])
			}
		}
	}

	if len(mobil) == 0 {
		fmt.Println("Tidak ditemukan mobil dengan harga dalam rentang tersebut.")
		if varepeat() != 0 {
			menuUtama(&daf)
		} else {
			tampilkanMobilByHarga(daf)
		}
	}

	fmt.Println("Daftar mobil dengan harga antara Rp", hargaMin, "dan Rp", hargaMax, ":")

	for i := 0; i < len(mobil); i++ {
		pabrikan := getPabrikanByMobil(daf, mobil[i])
		fmt.Printf("Model: %s, Pabrikan: %s, Harga: Rp%d, Penjualan: %d Tahun: %d\n", mobil[i].Model, pabrikan.Nama, mobil[i].Harga, mobil[i].Penjualan, mobil[i].TahunKeluar)
	}
	fmt.Println("ini halaman menu")
	if varepeat() >	 0 {
		menuUtama(&daf)
	} else {
		tampilkanMobilByHarga(daf)
	}
}

func tampilkanMobilByPenjualan(daf daftarPabrikan) {
	fmt.Print("Masukkan jumlah penjualan mobil: ")

	var penjualan int
	fmt.Scanln(&penjualan)

	var mobil []Mobil

	for i := 0; i < daf.totalPabrikan; i++ {
		for j := 0; j < daf.daftarPabrikan[i].nMobil; j++ {
			if daf.daftarPabrikan[i].Mobil[j].Penjualan == penjualan {
				mobil = append(mobil, daf.daftarPabrikan[i].Mobil[j])
			}
		}
	}

	if len(mobil) == 0 {
		fmt.Println("Tidak ditemukan mobil dengan jumlah penjualan tersebut.")
	}

	fmt.Println("Daftar mobil dengan jumlah penjualan", penjualan, ":")

	for i := 0; i < len(mobil); i++ {
		pabrikan := getPabrikanByMobil(daf, mobil[i])
		fmt.Printf("Model: %s, Pabrikan: %s, Harga: Rp%d, Penjualan: %d Tahun: %d\n", mobil[i].Model, pabrikan.Nama, mobil[i].Harga, mobil[i].Penjualan, mobil[i].TahunKeluar)
	}
	if varepeat() != 0 {
		menuUtama(&daf)
	} else {
		tampilkanMobilByPenjualan(daf)
	}
}

func tampilkanTop3Penjualan(daf daftarPabrikan) {
	fmt.Println("===== Tampilkan 3 Daftar Mobil dan Pabrikan dengan Penjualan Tertinggi =====")

	var mobil []Mobil

	for i := 0; i < daf.totalPabrikan; i++ {
		for j := 0; j < daf.daftarPabrikan[i].nMobil; j++ {
			mobil = append(mobil, daf.daftarPabrikan[i].Mobil[j]) // insert arr to mobil
		}
	}

	if len(mobil) == 0 {
		fmt.Println("Belum ada data mobil.")
		if varepeat() != 0 {
			menuUtama(&daf)
		} else {
			tampilkanTop3Penjualan(daf)
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
		pabrikan := getPabrikanByMobil(daf, m)
		fmt.Printf("Model: %s, Pabrikan: %s, Penjualan: %d\n", m.Model, pabrikan.Nama, m.Penjualan)
	}

	if varepeat() != 0 {
		menuUtama(&daf)
	} else {
		tampilkanTop3Penjualan(daf)
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
