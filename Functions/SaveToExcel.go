package Functions

import (
	"ReportingC2C/App/Models"
	"fmt"
	"os"
	"strconv"

	"github.com/Luxurioust/excelize"
)

func SaveToExcelAlfa(arrSize int, db []Models.C2CAlfa, fileName string) int {

	xlsx := excelize.CreateFile()
	hasil := 0
	//Header
	xlsx.SetCellValue("Sheet1", "A1", "TRX REQRES ID")
	xlsx.SetCellValue("Sheet1", "B1", "TRX REQRES ID RECEIVE")
	xlsx.SetCellValue("Sheet1", "C1", "ID TRANSAKSI SENDER")
	xlsx.SetCellValue("Sheet1", "D1", "ID TRANSAKSI RECEIVER")
	xlsx.SetCellValue("Sheet1", "E1", "TRX DATE")
	xlsx.SetCellValue("Sheet1", "F1", "CASH OUT DATE")
	xlsx.SetCellValue("Sheet1", "G1", "TRX RECEIVER NAME")
	xlsx.SetCellValue("Sheet1", "H1", "ID MEMBER_SENDER")
	xlsx.SetCellValue("Sheet1", "I1", "ID MEMBER_RECEIVER")
	xlsx.SetCellValue("Sheet1", "J1", "ID AGENT_SENDER")
	xlsx.SetCellValue("Sheet1", "K1", "ID AGENT_RECEIVER")
	xlsx.SetCellValue("Sheet1", "L1", "ID AGENTACCOUNT_SENDER")
	xlsx.SetCellValue("Sheet1", "M1", "ID AGENTACCOUNT_RECEIVER")
	xlsx.SetCellValue("Sheet1", "N1", "TRX SENDER NOHP")
	xlsx.SetCellValue("Sheet1", "O1", "TRX RECEIVER NOHP")
	xlsx.SetCellValue("Sheet1", "P1", "CHANNEL")
	xlsx.SetCellValue("Sheet1", "Q1", "KETERANGAN SEND")
	xlsx.SetCellValue("Sheet1", "R1", "KETERANGAN RECEIVE")
	xlsx.SetCellValue("Sheet1", "S1", "TRX REMITTANCE STATUS")
	xlsx.SetCellValue("Sheet1", "T1", "NOMINAL")
	xlsx.SetCellValue("Sheet1", "U1", "BIAYA")
	xlsx.SetCellValue("Sheet1", "V1", "GROSS REVENUE")
	xlsx.SetCellValue("Sheet1", "W1", "KOMISI AGENT SENDER")
	xlsx.SetCellValue("Sheet1", "X1", "KOMISI AGENT RECEIVER")

	fmt.Println("Data Size : ", arrSize)
	for i := 0; i < arrSize; i++ {
		rows1 := "A" + strconv.Itoa(i+2)
		rows2 := "B" + strconv.Itoa(i+2)
		rows3 := "C" + strconv.Itoa(i+2)
		rows4 := "D" + strconv.Itoa(i+2)
		rows5 := "E" + strconv.Itoa(i+2)
		rows6 := "F" + strconv.Itoa(i+2)
		rows7 := "G" + strconv.Itoa(i+2)
		rows8 := "H" + strconv.Itoa(i+2)
		rows9 := "I" + strconv.Itoa(i+2)
		rows10 := "J" + strconv.Itoa(i+2)
		rows11 := "K" + strconv.Itoa(i+2)
		rows12 := "L" + strconv.Itoa(i+2)
		rows13 := "M" + strconv.Itoa(i+2)
		rows14 := "N" + strconv.Itoa(i+2)
		rows15 := "O" + strconv.Itoa(i+2)
		rows16 := "P" + strconv.Itoa(i+2)
		rows17 := "Q" + strconv.Itoa(i+2)
		rows18 := "R" + strconv.Itoa(i+2)
		rows19 := "S" + strconv.Itoa(i+2)
		rows20 := "T" + strconv.Itoa(i+2)
		rows21 := "U" + strconv.Itoa(i+2)
		rows22 := "V" + strconv.Itoa(i+2)
		rows23 := "W" + strconv.Itoa(i+2)
		rows24 := "X" + strconv.Itoa(i+2)
		xlsx.SetCellValue("Sheet1", rows1, db[i].Trx_reqres_id)
		xlsx.SetCellValue("Sheet1", rows2, db[i].Trx_reqres_id_receive)
		xlsx.SetCellValue("Sheet1", rows3, db[i].IdTransaksiSender)
		xlsx.SetCellValue("Sheet1", rows4, db[i].IdTransaksiReceiver)
		xlsx.SetCellValue("Sheet1", rows5, db[i].TrxDate.Format("2006-01-02"))
		xlsx.SetCellValue("Sheet1", rows6, db[i].CashOutDate.Format("2006-01-02"))
		xlsx.SetCellValue("Sheet1", rows7, db[i].Trx_receiver_name)
		xlsx.SetCellValue("Sheet1", rows8, db[i].ID_MEMBER_SENDER)
		xlsx.SetCellValue("Sheet1", rows9, db[i].ID_MEMBER_RECEIVER)
		xlsx.SetCellValue("Sheet1", rows10, db[i].ID_AGENT_SENDER)
		xlsx.SetCellValue("Sheet1", rows11, db[i].ID_AGENT_RECEIVER)
		xlsx.SetCellValue("Sheet1", rows12, db[i].ID_AGENTACCOUNT_SENDER)
		xlsx.SetCellValue("Sheet1", rows13, db[i].ID_AGENTACCOUNT_RECEIVER)
		xlsx.SetCellValue("Sheet1", rows14, db[i].TrxSenderNohp)
		xlsx.SetCellValue("Sheet1", rows15, db[i].TrxReceiverNohp)
		xlsx.SetCellValue("Sheet1", rows16, db[i].CHANNEL)
		xlsx.SetCellValue("Sheet1", rows17, db[i].KeteranganSend)
		xlsx.SetCellValue("Sheet1", rows18, db[i].KeteranganReceive)
		xlsx.SetCellValue("Sheet1", rows19, db[i].TrxRemittanceStatus)
		xlsx.SetCellValue("Sheet1", rows20, db[i].Nominal)
		xlsx.SetCellValue("Sheet1", rows21, db[i].Biaya)
		xlsx.SetCellValue("Sheet1", rows22, db[i].GrossRevenue)
		xlsx.SetCellValue("Sheet1", rows23, db[i].KomisiAgentSender)
		xlsx.SetCellValue("Sheet1", rows24, db[i].KomisiAgentReceiver)

	}

	xlsx.SetActiveSheet(1)

	err := xlsx.WriteTo(fileName)
	if err != nil {
		fmt.Println(err)
		hasil = 1
		os.Exit(1)

	}
	return hasil
}
func SaveToExcelNonAlfa(arrSize int, db []Models.C2CNonAlfa, fileName string) int {

	xlsx := excelize.CreateFile()
	hasil := 0
	//Header
	xlsx.SetCellValue("Sheet1", "A1", "TRX REQRES ID")
	xlsx.SetCellValue("Sheet1", "B1", "TRX REQRES ID RECEIVE")
	xlsx.SetCellValue("Sheet1", "C1", "ID TRANSAKSI SENDER")
	xlsx.SetCellValue("Sheet1", "D1", "ID TRANSAKSI RECEIVER")
	xlsx.SetCellValue("Sheet1", "E1", "TRX DATE")
	xlsx.SetCellValue("Sheet1", "F1", "CASH OUT DATE")
	xlsx.SetCellValue("Sheet1", "G1", "TRX RECEIVER NAME")
	xlsx.SetCellValue("Sheet1", "H1", "ID MEMBER SENDER")
	xlsx.SetCellValue("Sheet1", "I1", "ID MEMBER RECEIVER")
	xlsx.SetCellValue("Sheet1", "J1", "ID AGENT SENDER")
	xlsx.SetCellValue("Sheet1", "K1", "ID AGENT RECEIVER")
	xlsx.SetCellValue("Sheet1", "L1", "ID AGENTACCOUNT SENDER")
	xlsx.SetCellValue("Sheet1", "M1", "ID AGENTACCOUNT RECEIVER")
	xlsx.SetCellValue("Sheet1", "N1", "TRX SENDER NOHP")
	xlsx.SetCellValue("Sheet1", "O1", "TRX RECEIVER NOHP")
	xlsx.SetCellValue("Sheet1", "P1", "CHANNEL")
	xlsx.SetCellValue("Sheet1", "Q1", "KETERANGAN SEND")
	xlsx.SetCellValue("Sheet1", "R1", "KETERANGAN RECEIVE")
	xlsx.SetCellValue("Sheet1", "S1", "TRX REMITTANCE STATUS")
	xlsx.SetCellValue("Sheet1", "T1", "NOMINAL")
	xlsx.SetCellValue("Sheet1", "U1", "BIAYA")
	xlsx.SetCellValue("Sheet1", "V1", "GROSS REVENUE")
	xlsx.SetCellValue("Sheet1", "W1", "KOMISI AGENT SENDER")
	xlsx.SetCellValue("Sheet1", "X1", "KOMISI AGENT RECEIVER")
	xlsx.SetCellValue("Sheet1", "y1", "TIPETRANSAKSI")

	fmt.Println("Data Size : ", arrSize-1)
	for i := 0; i < arrSize-1; i++ {
		rows1 := "A" + strconv.Itoa(i+2)
		rows2 := "B" + strconv.Itoa(i+2)
		rows3 := "C" + strconv.Itoa(i+2)
		rows4 := "D" + strconv.Itoa(i+2)
		rows5 := "E" + strconv.Itoa(i+2)
		rows6 := "F" + strconv.Itoa(i+2)
		rows7 := "G" + strconv.Itoa(i+2)
		rows8 := "H" + strconv.Itoa(i+2)
		rows9 := "I" + strconv.Itoa(i+2)
		rows10 := "J" + strconv.Itoa(i+2)
		rows11 := "K" + strconv.Itoa(i+2)
		rows12 := "L" + strconv.Itoa(i+2)
		rows13 := "M" + strconv.Itoa(i+2)
		rows14 := "N" + strconv.Itoa(i+2)
		rows15 := "O" + strconv.Itoa(i+2)
		rows16 := "P" + strconv.Itoa(i+2)
		rows17 := "Q" + strconv.Itoa(i+2)
		rows18 := "R" + strconv.Itoa(i+2)
		rows19 := "S" + strconv.Itoa(i+2)
		rows20 := "T" + strconv.Itoa(i+2)
		rows21 := "U" + strconv.Itoa(i+2)
		rows22 := "V" + strconv.Itoa(i+2)
		rows23 := "W" + strconv.Itoa(i+2)
		rows24 := "X" + strconv.Itoa(i+2)
		rows25 := "Y" + strconv.Itoa(i+2)
		xlsx.SetCellValue("Sheet1", rows1, db[i].Trx_reqres_id)
		xlsx.SetCellValue("Sheet1", rows2, db[i].Trx_reqres_id_receive)
		xlsx.SetCellValue("Sheet1", rows3, db[i].IdTransaksiSender)
		xlsx.SetCellValue("Sheet1", rows4, db[i].IdTransaksiReceiver)
		xlsx.SetCellValue("Sheet1", rows5, db[i].TrxDate.Format("2006-01-02"))
		xlsx.SetCellValue("Sheet1", rows6, db[i].CashOutDate.Format("2006-01-02"))
		xlsx.SetCellValue("Sheet1", rows7, db[i].Trx_receiver_name)
		xlsx.SetCellValue("Sheet1", rows8, db[i].ID_MEMBER_SENDER)
		xlsx.SetCellValue("Sheet1", rows9, db[i].ID_MEMBER_RECEIVER)
		xlsx.SetCellValue("Sheet1", rows10, db[i].ID_AGENT_SENDER)
		xlsx.SetCellValue("Sheet1", rows11, db[i].ID_AGENT_RECEIVER)
		xlsx.SetCellValue("Sheet1", rows12, db[i].ID_AGENTACCOUNT_SENDER)
		xlsx.SetCellValue("Sheet1", rows13, db[i].ID_AGENTACCOUNT_RECEIVER)
		xlsx.SetCellValue("Sheet1", rows14, db[i].TrxSenderNohp)
		xlsx.SetCellValue("Sheet1", rows15, db[i].TrxReceiverNohp)
		xlsx.SetCellValue("Sheet1", rows16, db[i].CHANNEL)
		xlsx.SetCellValue("Sheet1", rows17, db[i].KeteranganSend)
		xlsx.SetCellValue("Sheet1", rows18, db[i].KeteranganReceive)
		xlsx.SetCellValue("Sheet1", rows19, db[i].TrxRemittanceStatus)
		xlsx.SetCellValue("Sheet1", rows20, db[i].Nominal)
		xlsx.SetCellValue("Sheet1", rows21, db[i].Biaya)
		xlsx.SetCellValue("Sheet1", rows22, db[i].GrossRevenue)
		xlsx.SetCellValue("Sheet1", rows23, db[i].KomisiAgentSender)
		xlsx.SetCellValue("Sheet1", rows24, db[i].KomisiAgentReceiver)
		xlsx.SetCellValue("Sheet1", rows25, "E-MONEY TRANSFER")

	}

	xlsx.SetActiveSheet(1)

	err := xlsx.WriteTo(fileName)
	if err != nil {
		fmt.Println(err)
		hasil = 1
		os.Exit(1)

	}
	return hasil
}
