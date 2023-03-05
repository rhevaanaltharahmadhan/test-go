package handlers

import (
	"coding-test/database"
	"coding-test/entities"
	"coding-test/helpers"
)

type Customer entities.Customer

func (customer *Customer) H_CreateCustomer() (*Customer, error) {
	tx := database.GetDB().Begin()

	data := Customer{}
	data.Nik = customer.Nik
	data.Fullname = customer.Fullname
	data.Legalname = customer.Legalname
	data.TempatLahir = customer.TempatLahir
	data.TanggalLahir = customer.TanggalLahir
	data.Gaji = customer.Gaji
	data.FotoKtp = customer.FotoKtp
	data.FotoSelfie = customer.FotoSelfie
	data.Limit.Limit1 = customer.Limit.Limit1
	data.Limit.Limit2 = customer.Limit.Limit2
	data.Limit.Limit3 = customer.Limit.Limit3
	data.Limit.Limit4 = customer.Limit.Limit4
	data.Limit.CustomerId = customer.Nik

	if err := tx.Create(&customer).Error; err != nil {
		tx.Rollback()
		helpers.Logger("fail", "In Server: "+err.Error())
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		helpers.Logger("fail", "In Server: "+err.Error())
		return nil, err
	}

	return customer, nil
}

func H_GetAllCustomer() (*[]Customer, error) {
	var data []Customer

	if err := database.GetDB().Preload("Limit").Preload("Transaction").Find(&data).Error; err != nil {
		helpers.Logger("fail", "In Server: "+err.Error())
		return nil, err
	}

	return &data, nil
}
