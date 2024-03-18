package components

import (
	"danuhputra/miniproject2/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func UpdateDataBukuPerpustakaan() {
	fmt.Println("=================================")
	fmt.Println("Edit Buku")
	fmt.Println("=================================")
	TampilkanListBuku()
	var kode int
	var buku models.DataBukuPerpustakaan
	var bukuDitemukan bool
	var err error
	for {
		fmt.Print("Masukan Urutan Buku : ")
		_, err = fmt.Scanln(&kode)
		if err != nil {
			fmt.Println("Terjadi error:", err)
			return
		}
		if kode-1 >= 0 && kode-1 < len(listBook) {
			buku = listBook[kode-1]
			bukuDitemukan = true
			break
		} else {
			break
		}
	}
	if !bukuDitemukan {
		fmt.Println("Buku tidak ditemukan")
		return
	}
	var kodeBaru, judulBaru, pengarangBaru, penerbitBaru string
	var jumlahHalamanBaru, tahunterbitBaru int

	for {
		fmt.Print("Kode Buku Baru: ")
		kodeBaru, err = inputUser.ReadString('\n')
		if err != nil {
			fmt.Println("Error: ", err)
			return
		}
		kodeBaru = strings.TrimSpace(kodeBaru)
		if kodeBukuExists(kodeBaru) {
			fmt.Println("Kode buku sudah digunakan. Masukkan kode buku yang berbeda.")
		} else {
			break
		}
	}

	fmt.Print("Judul Baru: ")
	judulBaru, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	judulBaru = strings.TrimSpace(judulBaru)

	fmt.Print("Pengarang Baru: ")
	pengarangBaru, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	pengarangBaru = strings.TrimSpace(pengarangBaru)

	fmt.Print("Penerbit Baru: ")
	penerbitBaru, err = inputUser.ReadString('\n')
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	penerbitBaru = strings.TrimSpace(penerbitBaru)

	fmt.Print("Jumlah Halaman Baru :")
	_, err = fmt.Fscanf(inputUser, "%d\n", &jumlahHalamanBaru)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Print("Tahun Terbit Buru :")
	_, err = fmt.Fscanf(inputUser, "%d\n", &tahunterbitBaru)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	filePath := fmt.Sprintf("books/%d.json", kode)
	file, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Terjadi error saat membaca file JSON:", err)
		return
	}

	err = json.Unmarshal(file, &buku)
	if err != nil {
		fmt.Println("Terjadi error saat mengurai data JSON:", err)
		return
	}

	buku.KodeBuku = kodeBaru
	buku.JudulBuku = judulBaru
	buku.Pengarang = pengarangBaru
	buku.Penerbit = penerbitBaru
	buku.JumlahHalaman = jumlahHalamanBaru
	buku.TahunTerbit = tahunterbitBaru

	updatedData, err := json.Marshal(buku)
	if err != nil {
		fmt.Println("Terjadi error saat mengubah data buku menjadi JSON:", err)
		return
	}

	newPath := fmt.Sprintf("books/%s.json", kodeBaru)
	err = os.WriteFile(newPath, updatedData, 0644)
	if err != nil {
		fmt.Println("Terjadi error saat menulis ke file JSON:", err)
		return
	}
	err = os.Remove(filePath)
	if err != nil {
		fmt.Println("Terjadi error saat menghapus file lama:", err)
		return
	}
	listBook[kode] = buku

	fmt.Println("Data buku berhasil diperbarui.")

}