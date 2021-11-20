package utils

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

func ToExcel2(c *gin.Context, headers []string, values []map[string]interface{}) {
	xlsx := excelize.NewFile()
	for k, v := range headers {
		char := Div(k + 1)
		axis := fmt.Sprintf("%s1", char)
		xlsx.SetCellValue("Sheet1", axis, v)
	}
	for i := 0; i < len(values); i++ {
		row := i + 2 // 第一行为header行
		for j := 0; j < len(headers); j++ {
			col := j + 1
			char := Div(col)
			axis := fmt.Sprintf("%s%d", char, row)
			key := headers[j]
			currentCellValue, _ := values[i][key]
			xlsx.SetCellValue("Sheet1", axis, currentCellValue)
		}
	}

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"Workbook.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")

	//回写到web 流媒体 形成下载
	_ = xlsx.Write(c.Writer)
}

func Div(Num int) string {
	var (
		Str  string = ""
		k    int
		temp []int //保存转化后每一位数据的值，然后通过索引的方式匹配A-Z
	)
	//用来匹配的字符A-Z
	Slice := []string{"", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O",
		"P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	if Num > 26 { //数据大于26需要进行拆分
		for {
			k = Num % 26 //从个位开始拆分，如果求余为0，说明末尾为26，也就是Z，如果是转化为26进制数，则末尾是可以为0的，这里必须为A-Z中的一个
			if k == 0 {
				temp = append(temp, 26)
				k = 26
			} else {
				temp = append(temp, k)
			}
			Num = (Num - k) / 26 //减去Num最后一位数的值，因为已经记录在temp中
			if Num <= 26 {       //小于等于26直接进行匹配，不需要进行数据拆分
				temp = append(temp, Num)
				break
			}
		}
	} else {
		return Slice[Num]
	}

	for _, value := range temp {
		Str = Slice[value] + Str //因为数据切分后存储顺序是反的，所以Str要放在后面
	}
	return Str
}
