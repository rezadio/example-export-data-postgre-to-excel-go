package Models

import "time"

func (u *C2CNonAlfa) TableName() string {
	return "TMN2_C2C2_NonAlfamart"
}

type C2CNonAlfa struct {
	Trx_reqres_id            int64     `gorm:"column:trx_reqres_id"`
	Trx_reqres_id_receive    string    `gorm:"column:trx_reqres_id_receive"`
	IdTransaksiSender        string    `gorm:"column:id transaksi sender"`
	IdTransaksiReceiver      string    `gorm:"column:id transaksi receiver"`
	TrxDate                  time.Time `gorm:"column:Trx Date"`
	CashOutDate              time.Time `gorm:"column:Cash Out Date"`
	SenderName               string    `gorm:"column:Sender Name"`
	Trx_receiver_name        string    `gorm:"column:trx_receiver_name"`
	ID_MEMBER_SENDER         string    `gorm:"column:ID_MEMBER_SENDER"`
	ID_MEMBER_RECEIVER       string    `gorm:"column:ID_MEMBER_RECEIVER"`
	ID_AGENT_SENDER          string    `gorm:"column:ID_AGENT_SENDER"`
	ID_AGENT_RECEIVER        string    `gorm:"column:ID_AGENT_RECEIVER"`
	ID_AGENTACCOUNT_SENDER   string    `gorm:"column:ID_AGENTACCOUNT_SENDER"`
	ID_AGENTACCOUNT_RECEIVER string    `gorm:"column:ID_AGENTACCOUNT_RECEIVER"`
	TrxSenderNohp            string    `gorm:"column:trx_sender_nohp"`
	TrxReceiverNohp          string    `gorm:"column:trx_receiver_nohp"`
	CHANNEL                  string    `gorm:"column:CHANNEL"`
	KeteranganSend           string    `gorm:"column:Keterangan Send"`
	KeteranganReceive        string    `gorm:"column:Keterangan Receive"`
	TrxRemittanceStatus      string    `gorm:"column:trx_remittance_status"`
	Nominal                  int64     `gorm:"column:Nominal"`
	Biaya                    int64     `gorm:"column:Biaya"`
	GrossRevenue             int64     `gorm:"column:Gross Revenue"`
	KomisiAgentSender        int64     `gorm:"column:Komisi Agent Sender"`
	KomisiAgentReceiver      int64     `gorm:"column:Komisi Agent Receiver"`
}
