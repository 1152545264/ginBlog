package export

import "github.com/EDDYCJY/go-gin-example/pkg/setting"

const EXT = ".xlsx"

// GetExcelFullUrl get the full access path of the Excel file
func GetExcelFullUrl(name string) string {
	var (
		excelPath, res string
	)
	excelPath = GetExcelPath()         //< fixme: 此处是空值
	res = setting.AppSetting.PrefixUrl ///< fixme: 此处是空值
	res += "/" + excelPath
	res += name
	return res
}

// GetExcelPath get the relative save path of the Excel file
func GetExcelPath() string {
	return setting.AppSetting.ExportSavePath
}

// GetExcelFullPath Get the full save path of the Excel file
func GetExcelFullPath() string {
	return setting.AppSetting.RuntimeRootPath + GetExcelPath()
}
