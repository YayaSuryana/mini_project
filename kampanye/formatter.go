package kampanye

import "strings"

type KampanyeFormatter struct{
	ID					int 		`json:"id"`
	UserID				int			`json:"user_id"`
	Name				string		`json:"name"`
	ShortDescription 	string		`json:"short_description"`
	ImageURL			string		`json:"img_url"`
	GoalAmount			int			`json:"goal_amount"`
	CurrentAmount 		int			`json:"current_amount"`
	Slug 				string		`json:"slug"`
}

func FormatKampanye(kampanye Kampanye) KampanyeFormatter{
	kampanyeFormatter := KampanyeFormatter{}
	kampanyeFormatter.ID = kampanye.ID
	kampanyeFormatter.UserID = kampanye.UserID
	kampanyeFormatter.Name = kampanye.Name
	kampanyeFormatter.ShortDescription = kampanye.ShortDescription
	kampanyeFormatter.GoalAmount = kampanye.GoalAmount
	kampanyeFormatter.GoalAmount = kampanye.GoalAmount
	kampanyeFormatter.CurrentAmount = kampanye.CurrentAmount
	kampanyeFormatter.Slug = kampanye.Slug
	kampanyeFormatter.ImageURL = ""

	// periksa jika ada kampanye yang tidak memiliki gambar
	if len(kampanye.KampanyeImages) > 0 {
		kampanyeFormatter.ImageURL = kampanye.KampanyeImages[0].FileName
	}

	return kampanyeFormatter
}

// fungsi untuk merubah single kampanye ke slice of kampanye
func FormatKampanyes(kampanyes []Kampanye) []KampanyeFormatter{
	// pengecekan data kampanye dengan user id yang tidak sesuai
	// if len(kampanyes) == 0{
	// 	return []KampanyeFormatter{}
	// }
	// refaktor
	kampanyesFormatter := []KampanyeFormatter{}

	for _, kampanye := range kampanyes {
		kampanyeFormatter := FormatKampanye(kampanye)
		kampanyesFormatter = append(kampanyesFormatter, kampanyeFormatter)
	}

	return kampanyesFormatter
}

type KampnyeDetailFormater struct{
	ID 						int			`json:"id"`
	Name 					string			`json:"name"`
	ShortDescription 		string			`json:"short_description"`
	Description 			string			`json:"description"`
	ImageURL				string			`json:"image_url"`
	GoalAmount				int				`json:"goal_amount"`
	CurrentAmount			int				`json:"current_amount"`
	UserID					int				`json:"user_id"`
	Slug					string			`json:"slug"`
	Perks					[]string		`json:"perks"`
	User 					KampanyeUserFormatter `json:"user"`
	Images					[]KampanyeImageFormatter `json:"images"`
}

type KampanyeUserFormatter struct{
	Name 			string `json:"name"`
	ImageURL		string	`json:"image_url"`
}

type KampanyeImageFormatter struct{
	ImageURL 		string	`json:"image_url"`
	IsPrimary		bool 	`json:"is_primary"`
}

func FormatKampanyeDetail(kampanye Kampanye) KampnyeDetailFormater{
	kampanyeDetailFormatter := KampnyeDetailFormater{}
	kampanyeDetailFormatter.ID = kampanye.ID
	kampanyeDetailFormatter.Name = kampanye.Name
	kampanyeDetailFormatter.ShortDescription = kampanye.ShortDescription
	kampanyeDetailFormatter.Description = kampanye.Description
	kampanyeDetailFormatter.GoalAmount = kampanye.GoalAmount
	kampanyeDetailFormatter.GoalAmount = kampanye.GoalAmount
	kampanyeDetailFormatter.CurrentAmount = kampanye.CurrentAmount
	kampanyeDetailFormatter.Slug = kampanye.Slug
	kampanyeDetailFormatter.ImageURL = ""

	// periksa jika ada kampanye yang tidak memiliki gambar
	if len(kampanye.KampanyeImages) > 0 {
		kampanyeDetailFormatter.ImageURL = kampanye.KampanyeImages[0].FileName
	}

	var perks []string

	for _, perk := range strings.Split(kampanye.Perks, ","){
		perks = append(perks, strings.TrimSpace(perk))
	}
	kampanyeDetailFormatter.Perks = perks

	// formater preload user
	user := kampanye.User
	kampanyeUserFormatter := KampanyeUserFormatter{}
	kampanyeUserFormatter.Name = user.Name
	kampanyeUserFormatter.ImageURL = user.AvatarFileName
	kampanyeDetailFormatter.User = kampanyeUserFormatter


	images := []KampanyeImageFormatter{}
	for _, image := range kampanye.KampanyeImages{
		var kampanyeImageFormatter KampanyeImageFormatter
		kampanyeImageFormatter.ImageURL = image.FileName

		isPrimary := false
		if image.IsPrimary == 1{
			isPrimary = true
		}
		kampanyeImageFormatter.IsPrimary = isPrimary

		images = append(images, kampanyeImageFormatter)
	}

	kampanyeDetailFormatter.Images = images
	return kampanyeDetailFormatter
}