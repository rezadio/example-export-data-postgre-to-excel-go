package main

import (
	"ReportingC2C/App/Dao"
	"ReportingC2C/App/Datasources/DbDevel"
	"ReportingC2C/App/Functions"
	"fmt"

	"github.com/gin-gonic/gin"
)

var db = DbDevel.GetConnectionDevel()

func main() {

	r := gin.Default()
	v1 := r.Group("report/api/")
	{
		// v1.GET("alfa/:start/:end/:name", GetAlfa)
		v1.GET("nonalfa/:start/:end", GetNonAlfa)
		v1.POST("/alfa", GetAlfa)
		v1.POST("/nonalfa", GetNonAlfa)

	}

	r.Run(":8090")
}
func GetAlfa(c *gin.Context) {
	start := c.PostForm("datestart")
	end := c.PostForm("dateend")
	path := "Attachment/"
	fileName := "Alfa_Remittance_join_" + start + "_" + end + ".xlsx"
	fullPath := path + fileName

	fmt.Println(fileName)
	i, rows := Dao.GetDataAlfa(start, end, db)
	Functions.SaveToExcelAlfa(i, rows, fullPath)
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(fullPath)
}
func GetNonAlfa(c *gin.Context) {
	start := c.PostForm("datestart")
	end := c.PostForm("dateend")
	path := "Attachment/"
	fileName := "TMN_Transfer_E-Money_join_" + start + "_" + end + ".xlsx"
	fullPath := path + fileName
	i, rows := Dao.GetDataNonAlfa(start, end, db)
	Functions.SaveToExcelNonAlfa(i, rows, fullPath)

	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(fullPath)

}
