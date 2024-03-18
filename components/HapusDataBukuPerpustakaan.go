package components

import (
	"bufio"
	"fmt"
	"os"

	cls "github.com/MasterDimmy/go-cls"
)

func HapusDataBukuPerpustakaan() {
	cls.CLS()
	fmt.Println("==============================")
	fmt.Println("Hapus Buku dari Perpustakaan")
	fmt.Println("==============================")

	if len(listBook) == 0 {
		fmt.Println(listBook)
		fmt.Println("Tidak ada buku yang tersedia untuk dihapus.")
		fmt.Println("Tekan 'Enter' untuk kembali...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	TampilkanListBuku()

	fmt.Print("Masukkan nomor urutan buku yang ingin dihapus: ")
	var nomorUrutan int
	_, err := fmt.Scanln(&nomorUrutan)
	if err != nil || nomorUrutan < 1 || nomorUrutan > len(listBook) {
		fmt.Println("Input tidak valid.")
		fmt.Println("Tekan 'Enter' untuk kembali...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	bukuYangDihapus := listBook[nomorUrutan-1]
	fmt.Printf("Apakah Anda yakin ingin menghapus buku '%s' (Kode Buku: %s)? (y/n): ", bukuYangDihapus.JudulBuku, bukuYangDihapus.KodeBuku)
	var konfirmasi string
	_, err = fmt.Scanln(&konfirmasi)
	if err != nil || (konfirmasi != "y" && konfirmasi != "Y") {
		fmt.Println("Penghapusan dibatalkan.")
		fmt.Println("Tekan 'Enter' untuk kembali...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	listBook = append(listBook[:nomorUrutan-1], listBook[nomorUrutan:]...)

	err = os.Remove(fmt.Sprintf("books/%s.json", bukuYangDihapus.KodeBuku))
	if err != nil {
		fmt.Println("Terjadi error saat menghapus file:", err)
		fmt.Println("Tekan 'Enter' untuk kembali...")
		bufio.NewReader(os.Stdin).ReadBytes('\n')
		return
	}

	fmt.Println("Buku berhasil dihapus dari perpustakaan.")
	fmt.Println("Tekan 'Enter' untuk kembali...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}