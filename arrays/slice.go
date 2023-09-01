package main

import "fmt"

func main() {
	var months = []string{
		/** 0 **/ "january",
		/** 1 **/ "february",
		/** 2 **/ "march",
		/** 3 **/ "april",
		/** 4 **/ "may",
		/** 5 **/ "june",
		/** 6 **/ "july",
		/** 7 **/ "august",
		/** 8 **/ "september",
		/** 9 **/ "october",
		/** 10 **/ "november",
		/** 11 **/ "december",
	}
	println("================ print data slice ================")
	var slice = months[5:7]
	fmt.Println(slice)

	/** Untuk mendapatkan panjang  **/
	fmt.Println(len(slice))

	/** Untuk mendapat kapasitas **/
	fmt.Println(cap(slice))

	println("================ update data slice ================")
	//bisa dirubah juga datanya
	months[5] = "update"
	fmt.Println(slice)

	//index 0 di line 20
	slice[0] = "jadi gini lee"
	fmt.Println(months)

	println("================ menambahkan data slice baru ================")
	/** untuk menambahkan element baru pada slice **/
	targetSlice := months[10:]
	appendSlice := append(targetSlice, "indoapril")
	fmt.Println(appendSlice)

	println("================ membuat slice baru ================")

	makeSlice := make([]string, 2, 6)

	makeSlice[0] = "saleh"
	makeSlice[1] = "rashid"

	fmt.Println(makeSlice)
	fmt.Println(len(makeSlice))
	fmt.Println(cap(makeSlice))

	println("================ copy slice baru ================")

	copySlice := make([]string, len(makeSlice), cap(makeSlice))

	//copySlice "destination" <- makeSlice "source"
	copy(copySlice, makeSlice)
	fmt.Println(copySlice)

	println("================ perbedaan slice dan array saat deklarasi ================")

	var isArray = [...]int{1, 2, 3, 4, 5, 6}
	var isSlice = []int{1, 2, 3, 4, 5, 6}

	fmt.Println("ini adalah array", isArray, "\n", "ini adalah slice", isSlice)
}
