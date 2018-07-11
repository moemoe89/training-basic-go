package kelas

import(
	// jika saling import akan errror
	//"training-basic-go/session1/kelas_beda"
)

var Hello string

func init(){
	Hello = "mengisi hello kelas lain"
}

func KelasLain() string {
	// jika saling import akan error
	//return kelas_beda.KelasBeda()
	return ""
}
