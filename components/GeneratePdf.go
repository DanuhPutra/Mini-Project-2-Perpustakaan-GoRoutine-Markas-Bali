package components

import (
	"danuhputra/miniproject2/models"
	"fmt"
	"time"

	"github.com/go-pdf/fpdf"
)

func GeneratePdf() {
	TampilkanListBuku()
	fmt.Println("=================================")
	fmt.Println("Membuat Daftar Buku ...")
	fmt.Println("=================================")
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)


	for i, buku := range listBook {
		bukuText := fmt.Sprintf(
			"====================================\nbuku  #%d:\nKodeBuku : %s\nJudulBuku : %s\nPengarang : %s\nPenerbit : %s\nJumlahHalaman : %d\nTahunTerbit :  %d\nTanggal : %s\n====================================\n",
			i+1, buku.KodeBuku, buku.JudulBuku,
			buku.Pengarang, buku.Penerbit, buku.JumlahHalaman, buku.TahunTerbit,
			buku.Tanggal.Format("2006-01-02 15:04:05"))

		pdf.MultiCell(0, 10, bukuText, "0", "L", false)
		pdf.Ln(5)
	}

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("daftar_buku _%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
}

func GenerateSelected(selectedBook models.DataBukuPerpustakaan) {
	fmt.Println("=================================")
	fmt.Println("Membuat Daftar Buku ...")
	fmt.Println("=================================")

	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	bukuText := fmt.Sprintf(
		"KodeBuku : %s\nJudulBuku : %s\nPengarang : %s\nPenerbit : %s\nJumlahHalaman : %d\nTahunTerbit :  %d\nTanggal : %s\n",
		selectedBook.KodeBuku, selectedBook.JudulBuku,
		selectedBook.Pengarang, selectedBook.Penerbit, selectedBook.JumlahHalaman, selectedBook.TahunTerbit,
		selectedBook.Tanggal.Format("2006-01-02 15:04:05"))

	pdf.MultiCell(0, 10, bukuText, "0", "L", false)

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("daftar_buku_%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}

}
func PrintSelectedBook() {
	TampilkanListBuku()
	fmt.Print("Masukkan nomor urut buku yang ingin dicetak: ")
	var selectedNumber int
	_, err := fmt.Scanln(&selectedNumber)
	if err != nil {
		fmt.Println("Terjadi error:", err)
		return
	}
	if selectedNumber < 1 || selectedNumber > len(listBook) {
		fmt.Println("Nomor urut buku tidak valid.")
		return
	}
	selectedBook := listBook[selectedNumber-1]
	Selected(selectedBook)
}

func Selected(selectedBook models.DataBukuPerpustakaan) {
	pdf := fpdf.New("P", "mm", "A4", "")
	pdf.AddPage()

	pdf.SetFont("Arial", "", 12)
	pdf.SetLeftMargin(10)
	pdf.SetRightMargin(10)

	bukuText := fmt.Sprintf(
		"====================================\nKodeBuku : %s\nJudulBuku : %s\nPengarang : %s\nPenerbit : %s\nJumlahHalaman : %d\nTahunTerbit :  %d\nTanggal : %s\n====================================\n",
		selectedBook.KodeBuku, selectedBook.JudulBuku,
		selectedBook.Pengarang, selectedBook.Penerbit, selectedBook.JumlahHalaman, selectedBook.TahunTerbit,
		selectedBook.Tanggal.Format("2006-01-02 15:04:05"))

	pdf.MultiCell(0, 10, bukuText, "0", "L", false)

	err := pdf.OutputFileAndClose(
		fmt.Sprintf("daftar_buku_%s.pdf",
			time.Now().Format("2006-01-02-15-04-05")))

	if err != nil {
		fmt.Println("Terjadi error:", err)
	}
	fmt.Println("Buku berhasil dicetak dalam file PDF.")
}
