package kampanye

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