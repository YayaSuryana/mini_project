package kampanye

import "yayasuryana/user"

type GetKampanyeDetailInput struct{
	ID 			int 	`uri:"id" binding:"required"`
}

type CreateKampanyeInput struct{
	Name			 string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description 	 string `json:"description" binding:"required"`
	GoalAmount		 int	`json:"goal_amount" binding:"required"`
	Perks			 string `json:"perks" binding:"required"`
	User 			 user.User
}

type CreateKampanyeImage struct {
	KampanyeID			int `form:"kampanye_id" binding:"required"`
	IsPrimary			bool `form:"is_primary" binding:"required"`
	User 				user.User
}