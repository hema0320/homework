package excel

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func InitExcel(sheet string, titles map[string]string, beActive bool) *excelize.File { //创建一个excel
	file := excelize.NewFile()
	index := file.NewSheet("Sheet1")
	// 设置工作簿的默认工作表
	if beActive {
		file.SetActiveSheet(index)
	}
	for k, v := range titles {
		file.SetCellValue(sheet, k, v)
	}
	return file
}

func ExportExcel(file *excelize.File, sheet string, contents map[string]string) {
	for k, v := range contents {
		file.SetCellValue(sheet, k, v)
	}
}
func SaveExcel(file *excelize.File) {
	//generate file
	err := file.SaveAs("./log.xlsx")
	if err != nil {
		fmt.Errorf("生成excel出错")
	}
}
