package models

import (
	"time"
)

type DataBukuPerpustakaan struct {
	KodeBuku 		string 
	JudulBuku 		string
	Pengarang 		string
	Penerbit 		string
	JumlahHalaman 	int	
	TahunTerbit 	int
	Tanggal        	time.Time
}