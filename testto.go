// package main

// import "fmt"

// func main() {
// 	var jumlahsemester, skoreprt int
// 	var cumlaude bool

// 	fmt.Print("masukan jumlah semester dan skor eprt: ")
// 	fmt.Scan(&jumlahsemester, &skoreprt)

// 	cumlaude = jumlahsemester <= 8 && skoreprt >= 500
// 	fmt.Print(cumlaude)
// }

//------------------------------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	type waktu struct {
// 		jam, menit, detik int
// 	}

// 	var waktu1, waktu2 waktu

// 	fmt.Scan(&waktu1.jam, &waktu1.menit, &waktu1.detik)
// 	fmt.Scan(&waktu2.jam, &waktu2.menit, &waktu2.detik)

// 	fmt.Print("waktu 1: ")
// 	fmt.Print(waktu1)
// 	fmt.Print("waktu 2: ")
// 	fmt.Print(waktu2)
// }

//----------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	type waktu struct {
// 		jam, menit, detik int
// 	}

// 	var waktu1, waktu2 waktu
// 	var detikwaktu1, detikwaktu2, selisihtotal int

// 	fmt.Scan(&waktu1.jam, &waktu1.menit, &waktu1.detik)
// 	fmt.Scan(&waktu2.jam, &waktu2.menit, &waktu2.detik)

// 	detikwaktu1 = waktu1.jam*3600 + waktu1.menit*60 + waktu1.detik
// 	detikwaktu2 = waktu2.jam*3600 + waktu2.menit*60 + waktu2.detik
// 	selisihtotal = detikwaktu2 - detikwaktu1

// 	fmt.Print("selisih total detik: ")
// 	fmt.Print(selisihtotal)

// }

////--------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	type waktu struct {
// 		jam, menit, detik int
// 	}

// 	var waktu1, waktu2 waktu
// 	var detikwaktu1, detikwaktu2 int
// 	var jam, menit, detik int

// 	fmt.Scan(&waktu1.jam, &waktu1.menit, &waktu1.detik)
// 	fmt.Scan(&waktu2.jam, &waktu2.menit, &waktu2.detik)

// 	detikwaktu1 = waktu1.jam*3600 + waktu1.menit*60 + waktu1.detik
// 	detikwaktu2 = waktu2.jam*3600 + waktu2.menit*60 + waktu2.detik
// 	jam = (detikwaktu2 - detikwaktu1) / 3600
// 	menit = (detikwaktu2 - detikwaktu1) % 3600 / 60
// 	detik = (detikwaktu2 - detikwaktu1) % 3600 % 60

// 	fmt.Print("selisih total detik: ")
// 	fmt.Print(jam, ":", menit, ":", detik)
// }

//-------------------------------------------------------------------------------------

// package main

// import ("fmt")
// func main() {
// 	type transaksi struct {
// 		nama string
// 		jumlah int
// 		hargasatuan float64
// 		totalharga float64
// }
// 	var data transaksi

// 	fmt.Scan(&data.nama, &data.jumlah, &data.hargasatuan)
// 	data.totalharga = float64(data.jumlah) * data.hargasatuan

// 	fmt.Println("informasi transaksi: ")
// 	fmt.Printf("nama barang: %s\n", data.nama)
// 	fmt.Printf("jumlah: %v\n", data.jumlah)
// 	fmt.Printf("harga satuan: Rp %v\n", data.hargasatuan)
// 	fmt.Printf("total harga: Rp  %v\n", data.totalharga)

// }

//-----------------------------------------------------------------------------------

// package main
// import ("fmt")

// func main(){
// 	type persegipanjang struct {
// 		panjang, lebar, luas, keliling float64
// 	}
// 	var data persegipanjang

// 	fmt.Scan(&data.panjang, &data.lebar)
// 	data.keliling = 2 * data.panjang + 2*data.lebar
// 	data.luas = data.panjang * data.lebar

// 	fmt.Println("informasi persegi panjang: ")
// 	fmt.Printf("panjang: %.1f\n", data.panjang)
// 	fmt.Printf("lebar: %.1f\n", data.lebar)
// 	fmt.Printf("luas: %.1f\n", data.luas)
// 	fmt.Printf("keliling: %.1f\n", data.keliling)
// }

// ---------------------------------------------------------------------------------------------------------
// package main

// import (
// 	"fmt"
// )

// func main() {
// 	type bmi struct {
// 		bb, tb, bmi float64
// 		nama string
// 	}
// 	var data bmi

// 	fmt.Scan(&data.nama, &data.bb, &data.tb)
// 	data.bmi = data.bb / (data.tb * data.tb)

// 	fmt.Println("informasi BMI: ")
// 	fmt.Printf("nama: %s\n", data.nama)
// 	fmt.Printf("bb: %.1f kg\n", data.bb)
// 	fmt.Printf("tb: %.2f m\n", data.tb)
// 	fmt.Printf("bmi:%.1f\n", data.bmi)
// }

//---------------------------------------------------------------------------------------------------------------

// package main

// import ("fmt")
// func main(){
// 	type karyawan struct {
// 		nama string
// 		gajipokok, tunjangan, potongan, totalgaji float64
// 	}
// 	var data karyawan

// 	fmt.Scan(&data.nama, &data.gajipokok, &data.tunjangan, &data.potongan)
// 	data.totalgaji = data.gajipokok + data.tunjangan - data.potongan

// 	fmt.Println("informasi karyawan: ")
// 	fmt.Printf("nama: %s\n", data.nama)
// 	fmt.Printf("gaji pokok: Rp %.2f\n", data.gajipokok)
// 	fmt.Printf("tunjangan: Rp %.2f\n", data.tunjangan)
// 	fmt.Printf("potongan: Rp %.2f\n", data.potongan)
// 	fmt.Printf("total gaji: Rp %.2f\n", data.totalgaji)
// }