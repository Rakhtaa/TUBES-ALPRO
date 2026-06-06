package main

import "fmt"

const NMAX int = 100

type menu struct {
	Nama, Kategori, Status string
	Harga, JumlahKomposisi int
	Komposisi              [10]string
}

type tabMenu [NMAX]menu

func main() {
	// Program katalog cafe dan program akan berhenti ketika user input 7

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
		} else if pilihan == 7 {
			fmt.Println("\nProgram Selesai\n")
		} else {
			fmt.Println("\nPILIHAN TIDAK TERSEDIA\n")
		}
	}
}

func gabungKomposisi(d menu) string {
	// Function untuk menggabungkan array komposisi menjadi satu string yang dipisahkan koma

	var hasil string
	var i int
	hasil = ""
	for i = 0; i < d.JumlahKomposisi; i++ {
		if i == 0 {
			hasil = d.Komposisi[i]
		} else {
			hasil = hasil + ", " + d.Komposisi[i]
		}
	}
	return hasil
}

// ============================ PENCARIAN (SEARCHING) ============================

func SequentialSearch(d tabMenu, n int, target string) bool {
	// Function searching berdasarkan kategori dengan Metode Sequential Search

	var found bool
	var i, nomor int
	var textKomposisi string

	found = false
	nomor = 1
	for i = 0; i < n; i++ {
		if d[i].Kategori == target {
			textKomposisi = gabungKomposisi(d[i])
			fmt.Printf("%-3d | %-20s | %-10d | %-15s | %-15s | %-30s\n", nomor, d[i].Nama, d[i].Harga, d[i].Kategori, d[i].Status, textKomposisi)
			nomor++
			found = true
		}
	}
	return found
}

func BinarySearch(d tabMenu, n int, target string) bool {
	// Function searching berdasaran kategori dengan Metode Binary Search

	var left, right, mid, foundIdx, i, nomor int
	var textKomposisi string

	SortingBerdasarkanKategori(&d, n)

	left = 0
	right = n - 1
	foundIdx = -1

	for left <= right && foundIdx == -1 {
		mid = (left + right) / 2
		if d[mid].Kategori < target {
			left = mid + 1
		} else if d[mid].Kategori > target {
			right = mid - 1
		} else {
			foundIdx = mid
		}
	}

	if foundIdx != -1 {
		nomor = 1
		for i = 0; i < n; i++ {
			if d[i].Kategori == target {
				textKomposisi = gabungKomposisi(d[i])
				fmt.Printf("%-3d | %-20s | %-10d | %-15s | %-15s | %-30s\n", nomor, d[i].Nama, d[i].Harga, d[i].Kategori, d[i].Status, textKomposisi)
				nomor++
			}
		}
		return true
	}

	return false
}

// ============================ PENGURUTAN (SORTING) ============================

func SortingBerdasarkanKategori(d *tabMenu, n int) {
	// Procedure sorting berdasarkan kategori menggunakan Metode Insertion Sort secara Ascending

	var pass, i int
	var temp menu

	pass = 1
	for pass < n {
		i = pass
		temp = d[pass]
		for i > 0 && d[i-1].Kategori > temp.Kategori {
			d[i] = d[i-1]
			i--
		}
		d[i] = temp
		pass++
	}
}

func SelectionSort(d *tabMenu, n int) {
	// Procedure sorting berdasarkan harga menggunakan Metode Selection Sort secara Ascending

	var pass, i, idx int
	var temp menu

	pass = 1
	for pass <= n-1 {
		idx = pass - 1
		i = pass
		for i < n {
			if d[idx].Harga > d[i].Harga {
				idx = i
			}
			i++
		}
		temp = d[pass-1]
		d[pass-1] = d[idx]
		d[idx] = temp
		pass++
	}
}

func InsertionSort(d *tabMenu, n int) {
	// Procedure sorting berdasarkan harga menggunakan Metode Insertion Sort secara Ascending

	var pass, i int
	var temp menu

	pass = 1
	for pass < n {
		i = pass
		temp = d[pass]
		for i > 0 && d[i-1].Harga > temp.Harga {
			d[i] = d[i-1]
			i--
		}
		d[i] = temp
		pass++
	}
}

// ============================ DAFTAR MENU ============================

func keluarkanMenu(d tabMenu, n int) {
	// Procedure untuk mengeluarkan daftar menu dengan metode Selection Sort atau Insertion Sort

	var i, pilihan, nomor int
	var textKomposisi string

	if n == 0 {
		fmt.Println("\nKatalog Masih Kosong :)\n")
	} else {
		fmt.Println("\nDaftar Menu Cafe")
		fmt.Println("1. Selection Sort")
		fmt.Println("2. Insertion Sort")
		fmt.Print("Silahkan pilih (1/2) : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 || pilihan == 2 {
			if pilihan == 1 {
				SelectionSort(&d, n)
			} else if pilihan == 2 {
				InsertionSort(&d, n)
			}

			fmt.Println("\n==============================================================================================================================")
			fmt.Printf("%-3s | %-20s | %-10s | %-15s | %-15s | %-30s\n", "No", "Nama Menu", "Harga", "Kategori", "Status", "Komposisi")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")

			nomor = 1
			for i = 0; i < n; i++ {
				textKomposisi = gabungKomposisi(d[i])
				fmt.Printf("%-3d | %-20s | %-10d | %-15s | %-15s | %-30s\n", nomor, d[i].Nama, d[i].Harga, d[i].Kategori, d[i].Status, textKomposisi)
				nomor++
			}
			fmt.Println("==============================================================================================================================")
		} else {
			fmt.Println("\nPILIHAN TIDAK TERSEDIA\n")
		}
	}
}

// ============================ TAMBAH MENU ============================

func tambahMenu(d *tabMenu, n *int) {
	// Procedure untuk menambahkan menu ke dalam array dan akan berakhir ketika user input "No"

	var keputusan, inputBahan string
	var i int

	keputusan = "Yes"
	for keputusan != "No" && *n < NMAX {
		fmt.Println("\n========================================")
		fmt.Println("          TAMBAH MENU CAFE              ")
		fmt.Println("========================================")
		fmt.Print("Nama Menu : ")
		fmt.Scan(&d[*n].Nama)
		fmt.Print("Harga Menu : ")
		fmt.Scan(&d[*n].Harga)
		fmt.Print("Kategori Menu (Coffee/Non-Coffee) : ")
		fmt.Scan(&d[*n].Kategori)
		fmt.Println("Masukkan Komposisi (Ketik 'Selesai' jika sudah, Maksimal 10 bahan):")
		d[*n].JumlahKomposisi = 0
		inputBahan = ""
		i = 0
		for i < 10 && inputBahan != "Selesai" {
			fmt.Printf("Bahan ke-%d : ", i+1)
			fmt.Scan(&inputBahan)
			if inputBahan != "Selesai" {
				d[*n].Komposisi[i] = inputBahan
				d[*n].JumlahKomposisi++
				i++
			}
		}
		fmt.Print("Status Menu (Tersedia/Tidak-Tersedia) : ")
		fmt.Scan(&d[*n].Status)
		*n++
		fmt.Print("Tambahkan Menu Lagi? (Yes/No) : ")
		fmt.Scan(&keputusan)
	}
}

// ============================ HAPUS MENU ============================

func hapusMenu(d *tabMenu, n *int) {
	// Procedure untuk menghapus menu tertentu dalam array

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
	// Procedure mengubah data menu tertentu dalam array

	var target, inputBahan string
	var found bool
	var idxFound, i, j int

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
			fmt.Print("Kategori Menu (Coffee/Non-Coffee): ")
			fmt.Scan(&d[idxFound].Kategori)
			fmt.Println("Masukkan Komposisi Baru (Ketik 'Selesai' jika sudah, Maksimal 10):")
			d[idxFound].JumlahKomposisi = 0
			inputBahan = ""
			j = 0
			for j < 10 && inputBahan != "Selesai" {
				fmt.Printf("Bahan ke-%d : ", j+1)
				fmt.Scan(&inputBahan)
				if inputBahan != "Selesai" {
					d[idxFound].Komposisi[j] = inputBahan
					d[idxFound].JumlahKomposisi++
					j++
				}
			}
			fmt.Print("Status Ketersediaan (Tersedia/Tidak-Tersedia): ")
			fmt.Scan(&d[idxFound].Status)
			fmt.Println("\nData berhasil diubah.")
		} else {
			fmt.Println("\nMenu tidak ditemukan.\n")
		}
	}
}

// ============================ KATEGORI MENU ============================

func kategoriMenu(d tabMenu, n int) {
	// Procedure searching untuk mencari menu berdasarkan kategori menggunakan metode Sequential Search atau Binary Search

	var target string
	var pilihan int
	var found bool

	if n == 0 {
		fmt.Println("\nKatalog kosong.\n")
	} else {
		fmt.Println("\nMetode Pencarian Kategori : ")
		fmt.Println("1. Sequential Search")
		fmt.Println("2. Binary Search")
		fmt.Print("Silahkan pilih (1/2) : ")
		fmt.Scan(&pilihan)

		if pilihan == 1 || pilihan == 2 {
			fmt.Print("\nMasukkan kategori (Coffee/Non-Coffee) : ")
			fmt.Scan(&target)
			fmt.Println("\n==============================================================================================================================")
			fmt.Printf("%-3s | %-20s | %-10s | %-15s | %-15s | %-30s\n", "No", "Nama Menu", "Harga", "Kategori", "Status", "Komposisi")
			fmt.Println("------------------------------------------------------------------------------------------------------------------------------")

			if pilihan == 1 {
				found = SequentialSearch(d, n, target)
			} else if pilihan == 2 {
				found = BinarySearch(d, n, target)
			}
			if !found {
				fmt.Println("\nKategori Tidak Ditemukan.\n")
			}
			fmt.Println("==============================================================================================================================")
		} else {
			fmt.Println("\nPILIHAN TIDAK TERSEDIA\n")
		}
	}
}

// ============================ STATISTIK ============================

func statistikCafe(d tabMenu, n int) {
	// Procedure untuk menampilkan data statistik menu

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
