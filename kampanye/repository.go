package kampanye

import "gorm.io/gorm"

type Repository interface{
	FindAll() ([]Kampanye, error)
	FindByUserID(userID int) ([]Kampanye, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

func (r *repository) FindAll() ([]Kampanye, error){
	var kampanyes []Kampanye

	err := r.db.Preload("KampanyeImages", "kampanye_images.is_primary = 1").Find(&kampanyes).Error
	if err != nil {
		return kampanyes, err
	}
	return kampanyes, nil
}

func (r *repository) FindByUserID(userID int) ([]Kampanye, error){
	var kampanyes []Kampanye

	err := r.db.Where("user_id = ?", userID).Preload("KampanyeImages", "kampanye_images.is_primary = 1").Find(&kampanyes).Error
	if err != nil {
		return kampanyes, err
	}
	return kampanyes, nil
}