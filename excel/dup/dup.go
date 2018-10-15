package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/OuSatoru/fuckgo/gbk"
)

func main() {
	if len(os.Args) <= 1 {
		panic("没有xlsx文件")
	}
	xlpath := os.Args[1]
	if path.Ext(xlpath) != ".xlsx" {
		panic("xlsx格式")
	}
	txtpath := strings.TrimRight(xlpath, ".xlsx") + ".txt"
	xlsx, err := excelize.OpenFile(xlpath)
	if err != nil {
		log.Fatal("无法打开文件", err)
	}
	txt, err := os.OpenFile(txtpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("创建txt文件失败", err)
	}
	defer txt.Close()
	shm := xlsx.GetSheetMap()
	for i := 1; ; i++ {
		nums := xlsx.GetCellValue(shm[1], fmt.Sprintf("C%d", i))
		num, err := strconv.Atoi(nums)
		if err != nil {
			break
		}
		if nums == "" { // useless
			break
		}
		towrite := []byte(xlsx.GetCellValue(shm[1], fmt.Sprintf("A%d", i)) + "\t" + xlsx.GetCellValue(shm[1], fmt.Sprintf("B%d", i)) + "\r\n")
		for j := 0; j < num; j++ {
			txt.Write(gbk.UTF8toGBKbyte(towrite))
		}
	}
}
