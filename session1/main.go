// paket utama
package main

// lib yang digunakan
import (
	"fmt"
	"strconv"

	"training-basic-go/session1/kelas"
	"training-basic-go/session1/kelas_beda"
)

// global variabel
var hello string

// fungsi yang dijalankan sebelum main
func init(){
	hello = "hello world"
}

// fungsi utama
func main(){

	// deklarasi variabel dg tipe data string
	var exampleString1 string
	exampleString1 = "test"
	fmt.Println(exampleString1)

	// deklarasi variabel langsung dengan nilainya
	exampleString2 := "test"
	fmt.Println(exampleString2)

	// akan error jika kita deklarasi nama variabel yang sama
	// exampleString1 := "error"
	// fmt.Println(exampleString1)

	// contoh deklarasi tipe data lain
	var exampleInteger int
	exampleInteger = 0
	fmt.Println(exampleInteger)

	var exampleFloat float64
	fmt.Println(exampleFloat)

	var exampleBool bool
	fmt.Println(exampleBool)

	// variabel hello akan mengoverride variabel global diatas
	// hello := 0

	// contoh deklarasi dengan pointer (*) untuk support null
	var pointerInt *int

	// pengecekan apakah pointerInt mempunyai nilai
	// karena akan error apabila pointerInt nilainya nil dan kt paksa mendapatkan valuenya (*pointerInt)
	if pointerInt != nil {
		fmt.Println(*pointerInt)
	}

	// contoh assign value ke pointer int, dg menggunakan tanda &
	nonPointerInt := 5
	pointerInt = &nonPointerInt

	fmt.Println(pointerInt)
	fmt.Println(*pointerInt)

	// converting string to int
	// di go akan selalu banyak return dengan tipe error
	// jika kita ingin meng-ignore balikan tersebut, bisa kita gunakan underscore _ ( convStringToInt, _ := strconv.Atoi("5") )
	// sebaiknya tipe error kita cek dulu seperti pada pointer
	convStringToInt, err := strconv.Atoi("5")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(convStringToInt)

	// contoh multiple return suatu function
	iniInt, iniString := MultipleReturn()
	fmt.Println(iniInt)
	fmt.Println(iniString)


	// contoh tipe data interface{} , bisa diassign value apa saja, tipe data, struct, map
	var iniInterface interface{}

	iniInterface = "asdasdas"
	fmt.Println(iniInterface)

	iniInterface = 100
	fmt.Println(iniInterface)

	// contoh array
	iniArrayString := []interface{}{
		4,
		"def",
	}

	// menambahkan value ke array
	iniArrayString = append(iniArrayString, 5)
	iniArrayString = append(iniArrayString, "asd")

	// contoh looping array
	for k, v := range iniArrayString {
		fmt.Println(k)
		fmt.Println(v)
	}

	// contoh deklarasi variabel dengan tipe map int string
	// int sebagai key, string sebagai value
	iniMap := make(map[int]string)
	iniMap[0] = "ini nol"
	iniMap[1] = "ini satu" + " tambah"

	// contoh looping map
	for k, v := range iniMap {
		fmt.Println(k)
		fmt.Println(v)
	}

	// map bisa langsung kita tembak key nya
	fmt.Println(iniMap[1])

	// deklarasi struct
	type IniStruct struct {
		IniString    string
		IniInt       int
		IniInterface interface{}
	}

	// deklarasi variabel dengan tipe data struct, serta assign valuenya
	var iniVarStruct IniStruct
	iniVarStruct.IniInt       = 5
	iniVarStruct.IniString    = "lima"
	iniVarStruct.IniInterface = true
	fmt.Println(iniVarStruct)

	// contoh import variabel dari package lain
	fmt.Println(kelas.Hello)
	fmt.Println(kelas_beda.KelasBeda())
}

func MultipleInput(kata string, input ...int) (int, string) {
	return len(input), kata
}

func MultipleReturn() (int, string) {
	return 0, "nol"
}