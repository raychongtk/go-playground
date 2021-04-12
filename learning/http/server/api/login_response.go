package api

type LoginResponse struct {
	Success      bool    `json:"success"`
	ErrorMessage *string `json:"errorMessage"`
}
