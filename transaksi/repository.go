package transaksi

import "gorm.io/gorm"

type Reqpository interface{
	GetByKampanyeID(kampanyeID int) ([]Transaksi, error)
	GetByUserID(UserID int) ([]Transaksi, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

// get list transaksi berdasarkan kampanye id (list detail kampanye transaksi)
func (r *repository) GetByKampanyeID(kampanyeID int) ([]Transaksi, error){
	var transaksi []Transaksi

	err := r.db.Preload("User").Where("kampanye_id = ?", kampanyeID).Order("id desc").Find(&transaksi).Error

	if err != nil {
		return transaksi, err
	}

	return transaksi, nil
}

// get list transaksi yang sudah pernah di bayar kan oleh user
func (r *repository) GetByUserID(UserID int) ([]Transaksi, error){
	var transaksi []Transaksi

	err := r.db.Preload("Kampanye.KampanyeImages", "kampanye_images.is_primary = 1").Where("user_id = ?", UserID).Order("id desc").Find(&transaksi).Error
	if err != nil {
		return transaksi, err
	}
	return transaksi, nil
}