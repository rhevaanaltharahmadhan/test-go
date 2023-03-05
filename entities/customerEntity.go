package entities

type Customer struct {
	Id           uint          `json:"id" gorm:"primary_key"`
	Nik          int           `json:"nik" gorm:"null;unique" binding:"required"`
	Fullname     string        `json:"fullname" gorm:"size:255;null" binding:"required"`
	Legalname    string        `json:"legalname" gorm:"size:255;null" binding:"required"`
	TempatLahir  string        `json:"tempat_lahir" gorm:"size:255;null" binding:"required"`
	TanggalLahir string        `json:"tanggal_lahir" gorm:"size:255;null" binding:"required"`
	Gaji         int           `json:"gaji" gorm:"null" binding:"required"`
	FotoKtp      string        `json:"foto_ktp" gorm:"size:255;null" binding:"required"`
	FotoSelfie   string        `json:"foto_selfie" gorm:"size:255;null" binding:"required"`
	Limit        Limit         `json:"limit" gorm:"foreignKey:CustomerId;references:Nik"`
	Transaction  []Transaction `json:"transactions" gorm:"foreignKey:CustomerId;references:Nik"`
	CreatedBase
}
