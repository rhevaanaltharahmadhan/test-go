package entities

type Transaction struct {
	Id            uint   `json:"id" gorm:"primary_key"`
	NomorKontrak  string `json:"nomor_kontrak" gorm:"size:255;null"`
	Otr           int    `json:"otr" gorm:"null" binding:"required"`
	AdminFee      int    `json:"admin_fee" gorm:"null" binding:"required"`
	JumlahCicilan int    `json:"jumlah_cicilan" gorm:"null" binding:"required"`
	JumlahBunga   int    `json:"jumlah_bunga" gorm:"null" binding:"required"`
	NamaAset      string `json:"nama_aset" gorm:"size:255;null" binding:"required"`
	CustomerId    int    `json:"customer_id" gorm:"null"`
	CreatedBase
}
