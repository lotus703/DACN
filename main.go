package main

import "fmt"
type nguoi struct{
	ho string
	ten string
	thongTin thongTin
}
 type thongTin struct{
	 diaChi string
 }
func(n nguoi) in(){
	fmt.Printf("%+v\n",n)
}
func(n *nguoi) doiTen(tenMoi string){
	(*n).ten = tenMoi
}
func main() {
	mrC := nguoi{
		ho: "mrC",
		ten: "deV",
		thongTin: thongTin{
			diaChi: "ở đâu còn lâu mới nói",
		},
	}
	mrCPointer := &mrC
	mrCPointer.doiTen("Tín")
	mrCPointer.in()
}

