// package main

// import "fmt"

// func main() {
//     var temp float64
//     var unit string

//     fmt.Scan(&temp, &unit)

//     if unit == "Celcius" {
//         fmt.Printf("Suhu dalam Fahrenheit adalah %.0f F\n", temp*9/5+32)
//     } else if unit == "Fahrenheit" {
//         fmt.Printf("Suhu dalam Celcius adalah %.0f C\n", (temp-32)*5/9)
//     } else {
//         fmt.Println("Unit suhu tidak valid.")
//     }
// }

// package main

// import "fmt"

// func main() {
//     var kartuid, suratizin bool

//     fmt.Scan(&kartuid, &suratizin)

//     if kartuid || suratizin {
//         fmt.Println("Diizinkan masuk")
//     } else {
//         fmt.Println("Tidak diizinkan masuk")
//     }
// }

// package main

// import "fmt"

// func main() {
//     var gram, x, karung int

//     fmt.Scan(&gram, &x)

//     karung = (gram + 999) / 1000 * x

//     fmt.Println(karung)
// }

//Furniture
// package main

// import "fmt"

// func main() {
// 	var uangbelanja float64
// 	fmt.Print("Masukkan uang belanja: ")
// 	fmt.Scanln(&uangbelanja)
// 	if uangbelanja >= 10000 {
// 		uangbelanja -= uangbelanja * 0.2
// 		fmt.Println(uangbelanja)
// 	} else if uangbelanja >= 500 && uangbelanja < 10000 {
// 		uangbelanja -= uangbelanja * 0.15
// 		fmt.Println(uangbelanja)
// 	} else if uangbelanja < 500 {
// 		uangbelanja -= uangbelanja * 0.05
// 		fmt.Println(uangbelanja)
// 	}
// }


//Nominal
// package main
// import "fmt"
// func main(){
// 	var nominal int
// 	fmt.Print("Masukkan nominal: ")
// 	fmt.Scanln(&nominal)
// 	if nominal < 10 {
// 		fmt.Println("Satuan")
// 	} else if nominal >= 10 && nominal < 100 {
// 		fmt.Println("Puluhan")
// 	} else if nominal >= 100 && nominal < 1000 {
// 		fmt.Println("Ratusan")
// 	} else if nominal >= 1000 && nominal < 10000 {
// 		fmt.Println("Ribuan")
// 	} else if nominal >= 10000 && nominal < 100000 {
// 		fmt.Println("Puluhan Ribu")
// 	} else if nominal >= 100000 && nominal < 1000000 {
// 		fmt.Println("Ratusan Ribu")
// 	}
// }

//BMI
// package main
// import "fmt"
// func main(){
// 	var berat, tinggi, BMI float64

// 	fmt.Print("Masukkan tinggi badan (cm): ")
// 	fmt.Scanln(&tinggi)
// 	fmt.Print("Masukkan berat badan (kg): ")
// 	fmt.Scanln(&berat)
	
// 	BMI = berat / (tinggi / 100 * tinggi / 100)

// 	if BMI < 18.5 {
// 		fmt.Println("Kekurangan berat badan")
// 	} else if BMI >= 18.5 && BMI <= 22.9 {
// 		fmt.Println("Berat badan normal")
// 	} else if BMI > 22.9 {
// 		fmt.Println("Kelebihan berat badan")
// 	}
// }





// package main
// import "fmt"
// func main() { 

// 	var bil1, bil2, hasil float64
// 	var operator rune

// 	fmt.Scanf("%f %c %f", &bil1, &operator, &bil2)
	
// 	switch operator {
// 	case '+':
// 		hasil = bil1 + bil2
// 		fmt.Printf("%.1f\n", hasil)
// 	case '-':
// 		hasil = bil1 - bil2
// 		fmt.Printf("%.1f\n", hasil)
// 	case '*':
// 		hasil = bil1 * bil2
// 		fmt.Println(hasil)
// 	case '/':
// 		hasil = bil1 / bil2
// 		fmt.Printf("%.3f\n", hasil)
// 	}
// }




// package main
// import "fmt"
// func main() {
// 	var level int
// 	fmt.Scan(&level)
// 	switch {
// 	case level >= 1 && level <= 10 :
// 		fmt.Println("Pemula")
// 	case level >= 11 && level <= 20 :
// 		fmt.Println("Menengah")
// 	case level >= 21 && level <= 30 :
// 		fmt.Println("Ahli")
// 	case level > 30:
// 		fmt.Println("Master")
// 	default:
// 		fmt.Println("Level tidak valid")
// 	}
// }



// package main
// import "fmt"
// func main() {
// 	var number int
// 	fmt.Scan(&number)
// 	switch {
// 	case number >= 1 && number <= 9:
// 		fmt.Println("Satuan")
// 	case number >= 10 && number <= 99:
// 		fmt.Println("Puluhan")
// 	case number >= 100 && number <= 999:
// 		fmt.Println("Ratusan")
// 	case number >= 1000 && number <= 9999:
// 		fmt.Println("Ribuan")
// 	case number >= 10000 && number <= 99999:
// 		fmt.Println("Puluhan ribu")
// 	case number >= 100000 && number <= 999999:
// 		fmt.Println("Ratusan ribu")
// 	}
// }