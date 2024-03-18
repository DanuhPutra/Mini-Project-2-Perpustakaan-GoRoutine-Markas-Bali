package components

import (
	"bufio"
	"danuhputra/miniproject2/models"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"

	cls "github.com/MasterDimmy/go-cls"
)

var listBook []models.DataBukuPerpustakaan
var inputUser = bufio.NewReader(os.Stdin)

func TambahBukuBaru(){
	cls.CLS()

	// deklarasi variabel
	var KodeBukuTambah string
	JudulBukuBaru := bufio.NewReader(os.Stdin)
	PengarangBukuBaru := bufio.NewReader(os.Stdin)
	PenerbitBukuBaru := bufio.NewReader(os.Stdin)
	JumlahHalamanBukuBaru := 0
	TahunTerbitBukuBaru := 0 

	fmt.Println("==============================")
	fmt.Println("Menambahkan Buku Baru")
	fmt.Printf("==============================\n")
	simpanBuku := []models.DataBukuPerpustakaan{}

	for {
		// kode buku
		for{
			fmt.Print("Kode Buku Baru : ")
			_, err := fmt.Scanln(&KodeBukuTambah)
			if err != nil {
				fmt.Println("Ups, Terjadi error pada Kode Buku!", err)
				return
			}
			KodeBukuTambah = strings.TrimSpace(KodeBukuTambah)

			if kodeBukuExists(KodeBukuTambah) {
				fmt.Println("Kode buku sudah digunakan. Masukkan kode buku yang berbeda.")
			} else {
				break
			}

		}

		// judul buku
		fmt.Print("Judul Buku Baru : ")
		JudulBukuTambah, err := JudulBukuBaru.ReadString('\n')
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Judul Buku!", err)
			return
		}
		JudulBukuTambah = strings.TrimSpace(JudulBukuTambah)

		// pengarang buku
		fmt.Print("Pengarang Buku Baru : ")
		PengarangBukuTambah, err := PengarangBukuBaru.ReadString('\n')
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Judul Buku!", err)
			return
		}
		PengarangBukuTambah = strings.TrimSpace(PengarangBukuTambah)

		// penerbit buku
		fmt.Print("Penerbit Buku Baru : ")
		PenerbitBukuTambah, err := PenerbitBukuBaru.ReadString('\n')
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Judul Buku!",err)
			return
		}
		PenerbitBukuTambah = strings.TrimSpace(PenerbitBukuTambah)

		// jumlah halaman buku
		fmt.Print("silahkan masukan Jumlah Halaman pada Buku Baru :")
		_, err = fmt.Scanln(&JumlahHalamanBukuBaru)
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Jumlah Halaman Buku! :", err)
			return
		}

		// tahun terbit buku
		fmt.Print("silahkan masukan Tahun Terbit Buku :")
		_, err = fmt.Scanln(&TahunTerbitBukuBaru)
		if err != nil {
			fmt.Println("Ups, Terjadi error pada Tahun Terbit Buku! :", err)
			return
		}

		simpanBuku = append(simpanBuku, models.DataBukuPerpustakaan{
			KodeBuku : KodeBukuTambah,
			JudulBuku : JudulBukuTambah,
			Pengarang : PengarangBukuTambah,
			Penerbit : PenerbitBukuTambah,
			JumlahHalaman : JumlahHalamanBukuBaru,
			TahunTerbit : TahunTerbitBukuBaru,
			Tanggal : time.Now(),
		})

		var pilihanMenuBuku = 0
		fmt.Println("Ketik 1 untuk menambah buku lagi, Ketik 0 untuk kembali")
		_, err = fmt.Scanln(&pilihanMenuBuku)
		if err != nil {
			fmt.Println("Terjadi Error : ", err)
			return
		}

		if pilihanMenuBuku == 0 {
			break
		}

	}

	fmt.Println("Menambahkan Buku Kedalam Perpustakaan...")
	_ = os.Mkdir("books", 0755)
	ch := make(chan models.DataBukuPerpustakaan)
	wg := sync.WaitGroup{}
	jumlahStafBuku := 5

	for i := 0; i < jumlahStafBuku; i++ {
		wg.Add(1)
		go simpanBukuTambahan(ch, &wg, i)
	}

	for _, kodeBuku := range simpanBuku {
		ch <- kodeBuku
	}

	close(ch)
	wg.Wait()

	fmt.Println("berhasil menambahkan buku baru kedalam perpustakaan!")
	fmt.Println("\n======================================")
	fmt.Println("Tekan 'Enter' untuk melanjutkan...")		
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func simpanBukuTambahan(ch <-chan models.DataBukuPerpustakaan, wg *sync.WaitGroup, noStaff int){
	for buku := range ch {
		dataJson, err := json.Marshal(buku)
		if err != nil {
			fmt.Println("terjadi error!", err)
		}

		err = os.WriteFile(fmt.Sprintf("books/%s.json", buku.KodeBuku), dataJson, 0644)
		if err != nil {
			fmt.Println("Terjadi error:", err)
		}

		fmt.Printf("staff No %d Memproses buku baru dengan KodeBuku : %s!\n", noStaff, buku.KodeBuku)
	}
	wg.Done()
}