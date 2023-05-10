package transaksi

import "gorm.io/gorm"

type Reqpository interface{
	GetByKampanyeID(kampanyeID int) ([]Transaksi, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) GetByKampanyeID(kampanyeID int) ([]Transaksi, error){
	var transaction []Transaksi

	err := r.db.Preload("User").Where("kampanye_id = ?", kampanyeID).Order("id desc").Find(&transaction).Error

	if err != nil {
		return transaction, err
	}

	return transaction, nil
}
