package main

import "fmt"

const NMAX int = 1000

type menu struct {
	Nama, Kategori, Komposisi, Tersedia string
	Harga                               int
}

type tabMenu [NMAX]menu

func main() {
	var pilihan, jumlahMenu int
	var daftarMenu tabMenu

	for pilihan != 7 {

		fmt.Println("\n========================================")
		fmt.Println("          PROGRAM KATALOG CAFE          ")
		fmt.Println("========================================")
		fmt.Println("1. Daftar Menu")
		fmt.Println("2. Tambahkan Menu")
		fmt.Println("3. Hapus Menu")
		fmt.Println("4. Ubah Menu")
		fmt.Println("5. Kategori Menu")
		fmt.Println("6. Statistik Cafe")
		fmt.Println("7. EXIT")
		fmt.Print("Silahkan Pilih Nomor : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 {
			keluarkanMenu(daftarMenu, jumlahMenu)
		} else if pilihan == 2 {
			tambahMenu(&daftarMenu, &jumlahMenu)
		} else if pilihan == 3 {
			hapusMenu(&daftarMenu, &jumlahMenu)
		} else if pilihan == 4 {
			ubahMenu(&daftarMenu, jumlahMenu)
		} else if pilihan == 5 {
			kategoriMenu(daftarMenu, jumlahMenu)
		} else if pilihan == 6 {
			statistikCafe(daftarMenu, jumlahMenu)
		} else {
			fmt.Println("\nProgram Selesai\n")
		}
	}
}

// ============================ DAFTAR MENU ============================

func keluarkanMenu(d tabMenu, n int) {
	var pass, pilihan, i, j, idx int
	var temp menu

	if n == 0 {
		fmt.Println("\nKatalog Masih Kosong :)\n")
	} else {
		fmt.Println("\nDaftar Menu Cafe")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Print("Silahkan pilih (1/2) : ")
		fmt.Scan(&pilihan)
		// ================= SELECTION SORT =================
		if pilihan == 1 {
			pass = 1
			for pass <= n-1 {
				idx = pass - 1
				i = pass
				for i < n {
					if d[idx].Harga < d[i].Harga {
						idx = i
					}
					i++
				}
				temp = d[pass-1]
				d[pass-1] = d[idx]
				d[idx] = temp
				pass++
			}
			// ================= INSERTION SORT =================
		} else if pilihan == 2 {
			pass = 1
			for pass < n {
				i = pass
				temp = d[pass]
				for i > 0 && d[i-1].Harga < temp.Harga {
					d[i] = d[i-1]
					i--
				}
				d[i] = temp
				pass++
			}
		} else {
			fmt.Println("\nPILIHAN TIDAK TERSEDIA\n")
		}
		// ================= OUTPUT =================
		fmt.Println("\n==============================================================================================================")
		fmt.Printf("%-3s | %-20s | %-10s | %-15s | %-20s | %-10s\n",
			"No", "Nama Menu", "Harga", "Kategori", "Komposisi", "Status")
		fmt.Println("--------------------------------------------------------------------------------------------------------------")
		for j = 0; j < n; j++ {
			fmt.Printf("%-3d | %-20s | %-10d | %-15s | %-20s | %-10s\n", j+1, d[j].Nama, d[j].Harga, d[j].Kategori, d[j].Komposisi, d[j].Tersedia)
		}
		fmt.Println("==============================================================================================================")
	}
}

// ============================ TAMBAH MENU ============================

func tambahMenu(d *tabMenu, n *int) {
	var keputusan string
	keputusan = "Yes"
	for keputusan != "No" && *n < NMAX {
		fmt.Println("\n========================================")
		fmt.Println("          TAMBAH MENU CAFE              ")
		fmt.Println("========================================")
		fmt.Print("Nama Menu : ")
		fmt.Scan(&d[*n].Nama)
		fmt.Print("Harga Menu : ")
		fmt.Scan(&d[*n].Harga)
		fmt.Print("Komposisi Menu : ")
		fmt.Scan(&d[*n].Komposisi)
		fmt.Print("Kategori Menu (Coffee/Non-Coffee) : ")
		fmt.Scan(&d[*n].Kategori)
		fmt.Print("Status Ketersediaan (Yes/No) : ")
		fmt.Scan(&d[*n].Tersedia)
		*n++
		fmt.Print("Tambahkan Menu Lagi? (Yes/No) : ")
		fmt.Scan(&keputusan)
	}
}

// ============================ HAPUS MENU ============================

func hapusMenu(d *tabMenu, n *int) {
	var target string
	var found bool
	var idxFound, i int
	if *n == 0 {
		fmt.Println("\nKatalog kosong.\n")
	} else {
		fmt.Print("\nMasukkan Nama Menu yang ingin dihapus : ")
		fmt.Scan(&target)
		found = false
		i = 0
		for i < *n && !found {
			if d[i].Nama == target {
				found = true
				idxFound = i
			}
			i++
		}
		if found {
			for i = idxFound; i < *n-1; i++ {
				d[i] = d[i+1]
			}
			(*n)--
			fmt.Println("\nMenu berhasil dihapus.\n")
		} else {
			fmt.Println("\nMenu tidak ditemukan.\n")
		}
	}
}

// ============================ UBAH MENU ============================

func ubahMenu(d *tabMenu, n int) {
	var target string
	var found bool
	var idxFound, i int

	if n == 0 {
		fmt.Println("\nKatalog kosong.\n")
	} else {
		fmt.Print("\nMasukkan Nama Menu yang ingin diubah : ")
		fmt.Scan(&target)
		found = false
		i = 0
		for i < n && !found {
			if d[i].Nama == target {
				found = true
				idxFound = i
			}
			i++
		}
		if found {
			fmt.Println("\nMasukkan Data Baru")
			fmt.Print("Nama Menu : ")
			fmt.Scan(&d[idxFound].Nama)
			fmt.Print("Harga Menu : ")
			fmt.Scan(&d[idxFound].Harga)
			fmt.Print("Komposisi Menu : ")
			fmt.Scan(&d[idxFound].Komposisi)
			fmt.Print("Kategori Menu : ")
			fmt.Scan(&d[idxFound].Kategori)
			fmt.Print("Status Ketersediaan : ")
			fmt.Scan(&d[idxFound].Tersedia)
			fmt.Println("\nData berhasil diubah.")
		} else {
			fmt.Println("\nMenu tidak ditemukan.\n")
		}
	}
}

// ============================ KATEGORI MENU ============================

func kategoriMenu(d tabMenu, n int) {
	var target string
	var found bool
	var i int
	if n == 0 {
		fmt.Println("\nKatalog kosong.\n")
	} else {
		fmt.Print("\nMasukkan kategori (Coffee/Non-Coffee) : ")
		fmt.Scan(&target)

		found = false
		fmt.Println("\n==============================================================================================================")
		fmt.Printf("%-3s | %-20s | %-10s | %-15s | %-20s | %-10s\n", "No", "Nama Menu", "Harga", "Kategori", "Komposisi", "Status")
		fmt.Println("--------------------------------------------------------------------------------------------------------------")
		for i = 0; i < n; i++ {
			if d[i].Kategori == target {
				fmt.Printf("%-3d | %-20s | %-10d | %-15s | %-20s | %-10s\n", i+1, d[i].Nama, d[i].Harga, d[i].Kategori, d[i].Komposisi, d[i].Tersedia)
				found = true
			}
		}
		if !found {
			fmt.Println("\nKategori tidak ditemukan.\n")
		}
		fmt.Println("==============================================================================================================")
	}
}

// ============================ STATISTIK ============================

func statistikCafe(d tabMenu, n int) {
	var totalHarga, rataRata int
	var avgCoffee, avgNonCoffee int
	var totalHargaCoffee, totalHargaNonCoffee int
	var jmlCoffee, jmlNonCoffee int
	var i int
	if n == 0 {
		fmt.Println("\nBelum ada data statistik.\n")
	} else {
		for i = 0; i < n; i++ {
			totalHarga += d[i].Harga
			if d[i].Kategori == "Coffee" {
				jmlCoffee++
				totalHargaCoffee += d[i].Harga
			} else if d[i].Kategori == "Non-Coffee" {
				jmlNonCoffee++
				totalHargaNonCoffee += d[i].Harga
			}
		}
		rataRata = totalHarga / n
		if jmlCoffee > 0 {
			avgCoffee = totalHargaCoffee / jmlCoffee
		}
		if jmlNonCoffee > 0 {
			avgNonCoffee = totalHargaNonCoffee / jmlNonCoffee
		}
		fmt.Println("\n========================================")
		fmt.Println("          STATISTIK CAFE                ")
		fmt.Println("========================================")
		fmt.Printf("Total Menu           : %d\n", n)
		fmt.Printf("Jumlah Coffee        : %d\n", jmlCoffee)
		fmt.Printf("Jumlah Non-Coffee    : %d\n", jmlNonCoffee)
		fmt.Println("----------------------------------------")
		fmt.Printf("Rata-rata Coffee     : Rp %d\n", avgCoffee)
		fmt.Printf("Rata-rata NonCoffee  : Rp %d\n", avgNonCoffee)
		fmt.Printf("Rata-rata Semua Menu : Rp %d\n", rataRata)
		fmt.Println("----------------------------------------")
		fmt.Printf("Total Nilai Menu     : Rp %d\n", totalHarga)
		fmt.Println("========================================")
	}
}
