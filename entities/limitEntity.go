package entities

type Limit struct {
	Id         uint `json:"id" gorm:"primary_key"`
	Limit1     int  `json:"limit_1" gorm:"null" binding:"required"`
	Limit2     int  `json:"limit_2" gorm:"null" binding:"required"`
	Limit3     int  `json:"limit_3" gorm:"null" binding:"required"`
	Limit4     int  `json:"limit_4" gorm:"null" binding:"required"`
	CustomerId int  `json:"customer_id" gorm:"null"`
	CreatedBase
}
