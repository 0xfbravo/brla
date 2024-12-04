package model

type Pix struct {
	ID     string `json:"id" validate:"required" example:"1e6db88f-578a-4c0c-8175-7ce6711b22f4"`
	BrCode string `json:"brCode" validate:"required" example:"00020126420014br.gov.bcb.pix0107213213402091423412345204000053039865409213412.345802BR59063412346004412362070503***6304FA6D"`
	Due    string `json:"due" validate:"required" example:"2021-09-01T12:00:00Z"`
}
