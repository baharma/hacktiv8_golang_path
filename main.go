package main

import "fmt"

func main() {
	// name := "bahar"
	// age := 99

	// name0, name1, name2 := "daa", "fsa", "dasf"

	pertama, kedua, ketiga := "1", 2, "3ads"

	_, _, _ = pertama, kedua, ketiga

	// fmt.Println("hello : ", name)
	// fmt.Println("age : ", age)

	// fmt.Println(name0, name1, name2)

	// nama, age, address := "Aditya", 12, "bali"

	// fmt.Printf("halo Nama ku %s , umurku adalah %d, dan aku tinggal di %s", nama, age, address)

	// fmt.Printf("tipe ini adalah %T \n", nama)
	// tessfmt.Printf("tipe ini adalah %T \n", age)3ss

	// const age = 2024
	// if chekc := age - 2001; chekc > 17 {
	// 	fmt.Println("sim terbuat")
	// } else {
	// 	fmt.Println("Sim gagal")
	// }

	// age := 10

	// if check := 30; check < age {
	// 	fmt.Println("Tua Ke anjeng")
	// } else {
	// 	fmt.Println("haiya")
	// }

	// score := 4

	// if score > 7 {
	// 	switch {
	// 	case score == 8:
	// 		fmt.Println("nice")
	// 	case (score < 8) && (score > 5):
	// 		fmt.Println("good")
	// 	default:
	// 		fmt.Println("tolol")
	// 	}
	// } else {
	// 	if score == 5 {
	// 		fmt.Println("nyerah aja")
	// 	} else if score == 4 {
	// 		fmt.Println("kasian")
	// 	} else {
	// 		fmt.Println("kamu pasti gak bisa")
	// 		if score == 0 {
	// 			fmt.Println("pengangguran")
	// 		}
	// 	}
	// }
	// i := 0
	// for {
	// 	fmt.Println("haha Tolol :", i)
	// 	i++

	// 	if i == 6 {
	// 		break
	// 	}
	// }

	// for i := 1; i <= 10; i++ {
	// 	if i%2 == 1 {
	// 		continue
	// 	}
	// 	if i > 8 {
	// 		break
	// 	}

	// 	fmt.Println("angka", i)
	// }

	// for i := 0; i < 5; i++ {
	// 	for j := i; j < 5; j++ {
	// 		fmt.Print(j, " ")
	// 	}
	// 	fmt.Println()
	// }

	// outerLoop:
	// 	for i := 0; i < 3; i++ {
	// 		fmt.Println("perulangan ke -", i+1)
	// 		for j := 0; j < 3; j++ {
	// 			if i == 2 {
	// 				break outerLoop
	// 			}
	// 			fmt.Print(j, " ")
	// 		}
	// 		fmt.Print("\n")

	// 	}

	// var number [4]int

	// number = [4]int{1, 2, 3, 4}

	// var name = [3]string{"yayan", "yuyun", "yeyen"}
	// fmt.Printf("%#v\n", number)
	// fmt.Printf("%#v\n", name)

	// biji := [4]string{"apel", "titit", "tutut", "haha"}
	// biji[0] = "haiya"

	// fmt.Printf("%#v\n", biji)

	// var buah = [3]string{"buah", "dada", "data"}

	// for i, v := range buah {
	// 	fmt.Printf("Index : %d , value :%s\n", i, v)
	// }
	// fmt.Println(strings.Repeat("#", 19))

	// for i := 0; i < len(buah); i++ {
	// 	fmt.Printf("index : %d , Value : %s\n", i, buah[i])
	// }

	// balances := [2][3]int{{4, 3, 4}, {5, 5, 3}}

	// _ = balances
	// for _, arr := range balances {
	// 	for _, value := range arr {
	// 		fmt.Printf("%d", value)
	// 	}
	// 	fmt.Println()
	// }

	var buah = make([]string, 3)

	_ = buah

	buah[0] = "Haiya"
	buah[1] = "dada"
	buah[2] = "haha"
	fmt.Printf("%#v", buah)

}
