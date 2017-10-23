package Dao

import (
	"ReportingC2C/App/Models"

	"github.com/jinzhu/gorm"
)

func GetDataAlfa(start string, end string, db *gorm.DB) (int, []Models.C2CAlfa) {
	var arr []Models.C2CAlfa
	i := 0
	rows, _ := db.Raw(AlfaSql(), start, end).Rows() // (*sql.Rows, error)
	for rows.Next() {
		var result Models.C2CAlfa
		db.ScanRows(rows, &result)
		arr = append(arr, result)
		i += 1
	}
	return i, arr
}
func GetDataNonAlfa(start string, end string, db *gorm.DB) (int, []Models.C2CNonAlfa) {
	var arr []Models.C2CNonAlfa
	i := 0
	rows, _ := db.Raw(NonAlfaSql(), start, end).Rows() // (*sql.Rows, error)

	for rows.Next() {
		var result Models.C2CNonAlfa
		db.ScanRows(rows, &result)
		arr = append(arr, result)
		i += 1
	}
	return i, arr
}
func AlfaSql() string {
	sql := `SELECT
	*
FROM
	"TMN2_C2C2_Alfamart"
WHERE
	"Trx Date" :: DATE BETWEEN ? AND ?`
	return sql
}
func NonAlfaSql() string {
	sql := `SELECT
	*
FROM
	"TMN2_C2C2_NonAlfamart"
WHERE
	"Trx Date" :: DATE BETWEEN ? AND ?`
	return sql
}
