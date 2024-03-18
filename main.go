package main

import (
	"danuhputra/miniproject2/components"
	"fmt"
	"os"

	"github.com/MasterDimmy/go-cls"
)

func main(){
	cls.CLS()
	var PilihanAksi int

	fmt.Println("===========================================")
	fmt.Println("Aplikasi Manajemen Daftar Buku Perpustakaan")
	fmt.Println("===========================================")
	fmt.Println("silahkan pilih menu : ")
	fmt.Println("1. Menambahkan Buku Baru Perpustakaan")
	fmt.Println("2. Menampilkan Buku Perpustakaan")
	fmt.Println("3. Hapus Buku Perpustakaan")
	fmt.Println("4. Edit Buku Perpustakaan")
	fmt.Println("5. Print Semua Buku Perpustakaan")
	fmt.Println("6. Print Buku")
	fmt.Println("7. Keluar dari Program")
	fmt.Println("===========================================")
	fmt.Print("masukan pilihan : ")
	_, err := fmt.Scanln(&PilihanAksi)
	if err != nil {
		fmt.Println("Ups, Terjadi error pada aksi yang kamu pilih!", err)
	}

	switch PilihanAksi{
		case 1 :
			components.TambahBukuBaru()
		case 2 :
			components.TampilkanListBuku()
		case 3 :
			components.HapusDataBukuPerpustakaan()
		case 4 :
			components.UpdateDataBukuPerpustakaan()
		case 5 : 
			components.GeneratePdf()
		case 6 :
			components.PrintSelectedBook()
		case 7 : 
			os.Exit(0)
	}

	main()
}

