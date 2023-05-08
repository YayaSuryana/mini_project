package kampanye

import "gorm.io/gorm"

type Repository interface{
	FindAll() ([]Kampanye, error)
	FindByUserID(userID int) ([]Kampanye, error)
	FindByID(ID int) (Kampanye, error)
	Save(kampanye Kampanye) (Kampanye, error)
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

func(r *repository) FindByID(ID int) (Kampanye, error){
	var kampanye Kampanye

	err := r.db.Preload("User").Preload("KampanyeImages").Where("id = ?", ID).Find(&kampanye).Error

	if err != nil {
		return kampanye, err
	}
	return kampanye, nil
}

func (r *repository) Save(kampanye Kampanye) (Kampanye, error){
	err := r.db.Create(&kampanye).Error
	if err != nil{
		return kampanye, err
	}

	return kampanye, nil
}