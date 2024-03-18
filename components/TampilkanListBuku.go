package components

import (
	"bufio"
	"danuhputra/miniproject2/models"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"sync"

	cls "github.com/MasterDimmy/go-cls"
)

func LihatBuku(ch <-chan string, chPesanan chan models.DataBukuPerpustakaan, wg *sync.WaitGroup) {
	var listBuku models.DataBukuPerpustakaan
	for idBuku := range ch {
		dataJSON, err := os.ReadFile(fmt.Sprintf("Books/%s", idBuku))
		if err != nil {
			fmt.Println("Terjadi error saat membaca file!", err)
		}

		err = json.Unmarshal(dataJSON, &listBuku)
		if err != nil {
			fmt.Println("Terjadi error saat menerima data json!", err)
		}

		chPesanan <- listBuku
	}
	wg.Done()
}

func TampilkanListBuku(){
	cls.CLS()
	fmt.Println("======================================")
	fmt.Println("List Buku yang ada di Perpustakaan ini")
	fmt.Printf("======================================\n")

	listBook = []models.DataBukuPerpustakaan{}

	listJsonBuku, err := os.ReadDir("books")
	if err != nil {
		fmt.Println("Terjadi error: ", err)
	}

	wg := sync.WaitGroup{}

	ch := make(chan string)
	chPesanan := make(chan models.DataBukuPerpustakaan, len(listJsonBuku))

	jumlahPelayan := 5

	for i := 0; i < jumlahPelayan; i++ {
		wg.Add(1)
		go LihatBuku(ch, chPesanan, &wg)
	}

	for _, filePesanan := range listJsonBuku {
		ch <- filePesanan.Name()
	}

	close(ch)

	wg.Wait()

	close(chPesanan)

	for dataPesanan := range chPesanan {
		listBook = append(listBook, dataPesanan)
	}

	sort.Slice(listBook, func(i, j int) bool {
		return listBook[i].Tanggal.Before(listBook[j].Tanggal)
	})


	for banyakBuku, DataBukuPerpustakaan := range listBook{
		fmt.Printf(" | %d. | Kode Buku : %s | Judul Buku : %s | Pengarang Buku : %s | Penerbit Buku :%s | Jumlah Halaman Buku : %d | Tahun Terbit Buku : %d |\n",
			banyakBuku + 1,
			DataBukuPerpustakaan.KodeBuku,
			DataBukuPerpustakaan.JudulBuku,
			DataBukuPerpustakaan.Pengarang,
			DataBukuPerpustakaan.Penerbit,
			DataBukuPerpustakaan.JumlahHalaman,
			DataBukuPerpustakaan.TahunTerbit,
		)
	}

	fmt.Println("\n======================================")
	fmt.Println("Tekan 'Enter' untuk melanjutkan...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

}

