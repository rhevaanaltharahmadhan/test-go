package handlers

import (
	"coding-test/database"
	"coding-test/entities"
	"coding-test/helpers"
	"errors"
	"math/rand"
	"time"
)

type Transaction entities.Transaction

var letterRunes = []rune("0123456789")

func RandStringRunes() string {

	rand.Seed(time.Now().UTC().UnixNano())

	b := make([]rune, 12)
	l := len(letterRunes)
	for i := range b {
		b[i] = letterRunes[rand.Intn(l)]
	}

	return string(b)
}

func (transaction Transaction) H_CreateTransaction(customerId int) (*Transaction, error) {

	var customer entities.Customer

	tx := database.GetDB().Begin()

	if tx.Error != nil {
		helpers.Logger("fail", "In Server: "+tx.Error.Error())
		return nil, tx.Error
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Where("id = ?", customerId).First(&customer).Error; err != nil {
		helpers.Logger("fail", "In Server: "+err.Error())
		return nil, err
	}

	customerTransaction := Transaction{}
	customerTransaction.NomorKontrak = RandStringRunes()
	customerTransaction.Otr = transaction.Otr
	customerTransaction.AdminFee = transaction.AdminFee
	customerTransaction.JumlahCicilan = transaction.JumlahCicilan
	customerTransaction.JumlahBunga = transaction.JumlahBunga
	customerTransaction.NamaAset = transaction.NamaAset
	customerTransaction.CustomerId = customer.Nik

	var limitError = errors.New("Limit tidak mencukupi")

	var limit entities.Limit

	if err := tx.Where("customer_id = ?", customer.Nik).First(&limit).Error; err != nil {
		helpers.Logger("fail", "In Server: "+err.Error())
		return nil, err
	}

	switch {
	case customerTransaction.JumlahCicilan <= limit.Limit1 && customerTransaction.JumlahCicilan < customer.Gaji:
		if err := tx.Create(&customerTransaction).Error; err != nil {
			tx.Rollback()
			helpers.Logger("fail", "In Server: "+err.Error())
			return nil, err
		}

		if err := tx.Model(entities.Limit{}).Where("id = ?", customerId).Update("Limit1", limit.Limit1-customerTransaction.JumlahCicilan).Error; err != nil {
			tx.Rollback()
			helpers.Logger("fail", "In Server: "+err.Error())
			return nil, err
		}
	case customerTransaction.JumlahCicilan <= limit.Limit2 && customerTransaction.JumlahCicilan < customer.Gaji:
		if err := tx.Create(&customerTransaction).Error; err != nil {
			tx.Rollback()
			helpers.Logger("fail", "In Server: "+err.Error())
			return nil, err
		}

		if err := tx.Model(entities.Limit{}).Where("id = ?", customerId).Update("Limit2", limit.Limit2-customerTransaction.JumlahCicilan).Error; err != nil {
			tx.Rollback()
			helpers.Logger("fail", "In Server: "+err.Error())
			return nil, err
		}
	case customerTransaction.JumlahCicilan <= limit.Limit3 && customerTransaction.JumlahCicilan < customer.Gaji:
		if err := tx.Create(&customerTransaction).Error; err != nil {
			tx.Rollback()
			helpers.Logger("fail", "In Server: "+err.Error())
			return nil, err
		}

		if err := tx.Model(entities.Limit{}).Where("id = ?", customerId).Update("Limit3", limit.Limit3-customerTransaction.JumlahCicilan).Error; err != nil {
			tx.Rollback()
			helpers.Logger("fail", "In Server: "+err.Error())
			return nil, err
		}
	case customerTransaction.JumlahCicilan <= limit.Limit4 && customerTransaction.JumlahCicilan < customer.Gaji:
		if err := tx.Create(&customerTransaction).Error; err != nil {
			tx.Rollback()
			helpers.Logger("fail", "In Server: "+err.Error())
			return nil, err
		}

		if err := tx.Model(entities.Limit{}).Where("id = ?", customerId).Update("Limit4", limit.Limit4-customerTransaction.JumlahCicilan).Error; err != nil {
			tx.Rollback()
			helpers.Logger("fail", "In Server: "+err.Error())
			return nil, err
		}

	case limit.Limit1 == 0 && limit.Limit2 == 0 && limit.Limit3 == 0 && limit.Limit4 == 0:
		tx.Rollback()
		helpers.Logger("fail", "In Server: "+limitError.Error())
		return nil, limitError
	default:
		tx.Rollback()
		helpers.Logger("fail", "In Server: "+limitError.Error())
		return nil, limitError
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		helpers.Logger("fail", "In Server: "+err.Error())
		return nil, err
	}

	return &customerTransaction, nil
}
