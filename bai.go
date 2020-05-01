package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type nhieuBai []string

func taoNhieuBai() nhieuBai {
	kq := nhieuBai{}
	nhieunuoc := []string{"cơ", "rô", "chuồn", "bích"}
	nhieunut := []string{"1", "2", "3"}
	for _, nuoc := range nhieunuoc {
		for _, nut := range nhieunut {
			bai := nut + " " + nuoc
			kq = append(kq, bai)
		}
	}
	return kq
}
func (n nhieuBai) in() {
	fmt.Println(n)
}
func chiaBai(n nhieuBai, sl int) (nhieuBai, nhieuBai) {
	return n[:sl], n[sl:]
}
func (n nhieuBai) chuyenThanhString() string {
	return strings.Join(n, ",")
}

func (n nhieuBai) luuFile(tenFile string) error {
	data := []byte(n.chuyenThanhString())
	return ioutil.WriteFile(tenFile, data, 0666)
}

func taoBaiTuFile(tenFile string) nhieuBai {
	data, err := ioutil.ReadFile(tenFile)
	if err != nil {
		log.Printf("doc file khong duoc %v", err)
		os.Exit(1)
	}
	chuoiBai := string(data)
	mang := strings.Split(chuoiBai, ",")
	nhieuLaBai := nhieuBai(mang)
	return nhieuLaBai
}
