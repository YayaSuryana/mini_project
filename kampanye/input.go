package kampanye

type GetKampanyeDetailInput struct{
	ID 			int 	`uri:"id" binding:"required"`
}