package export

import (
	"ginBlog/pkg/setting"
	"os"
)

const EXT = ".xlsx"

// GetExcelFullUrl get the full access path of the Excel file
func GetExcelFullUrl(name string) string {
	var (
		excelPath, res string
	)
	excelPath = GetExcelPath()
	res += excelPath
	res += name
	return res
}

// GetExcelPath get the relative save path of the Excel file
func GetExcelPath() string {
	res := setting.AppSetting.ExportSavePath
	return res
}

// GetExcelFullPath Get the full save path of the Excel file
func GetExcelFullPath() string {
	curPath, _ := os.Getwd()
	return curPath + "/" + GetExcelPath()
}
