package kampanye

import "gorm.io/gorm"

type Repository interface{
	FindAll() ([]Kampanye, error)
	FindByUserID(userID int) ([]Kampanye, error)
	FindByID(ID int) (Kampanye, error)
	Save(kampanye Kampanye) (Kampanye, error)
	Update(kampanye Kampanye) (Kampanye, error)
	CreateImage(kampanyeImage KampanyeImage) (KampanyeImage, error)
	MarkNonIsPrimary(kampanyeID int) (bool, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository{
	return &repository{db}
}

// menampilkan semua kampanye
func (r *repository) FindAll() ([]Kampanye, error){
	var kampanyes []Kampanye

	err := r.db.Preload("KampanyeImages", "kampanye_images.is_primary = 1").Find(&kampanyes).Error
	if err != nil {
		return kampanyes, err
	}
	return kampanyes, nil
}

// menampilkan 
func (r *repository) FindByUserID(userID int) ([]Kampanye, error){
	var kampanyes []Kampanye

	err := r.db.Where("user_id = ?", userID).Preload("KampanyeImages", "kampanye_images.is_primary = 1").Find(&kampanyes).Error
	if err != nil {
		return kampanyes, err
	}
	return kampanyes, nil
}

// menampilkan detail kampanyer
func(r *repository) FindByID(ID int) (Kampanye, error){
	var kampanye Kampanye

	err := r.db.Preload("User").Preload("KampanyeImages").Where("id = ?", ID).Find(&kampanye).Error

	if err != nil {
		return kampanye, err
	}
	return kampanye, nil
}

// save kampanye
func (r *repository) Save(kampanye Kampanye) (Kampanye, error){
	err := r.db.Create(&kampanye).Error
	if err != nil{
		return kampanye, err
	}

	return kampanye, nil
}

// update kampanye
func (r *repository) Update(kampanye Kampanye) (Kampanye, error){
	err := r.db.Save(&kampanye).Error
	if err != nil {
		return kampanye, err
	}

	return kampanye, nil
}

// save ke table
func (r *repository) CreateImage(kampanyeImage KampanyeImage) (KampanyeImage, error){
	err := r.db.Save(&kampanyeImage).Error
	if err != nil {
		return kampanyeImage, err
	}

	return kampanyeImage, nil
}

// merubah is_primary jika ingin mengupdate is_primary true dan menimpanya menjadi false diganti dengan yang lain
func (r *repository) MarkNonIsPrimary(kampanyeID int) (bool, error){
	// update kampanye_images set is_primary = false where kampanye_id = (dari parameter)

	err := r.db.Model(&KampanyeImage{}).Where("kampanye_id = ?", kampanyeID).Update("is_primary", false).Error
	if err != nil {
		return false, err
	}

	return true, nil
}