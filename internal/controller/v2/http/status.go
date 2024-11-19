package http

type StatusBadRequest struct {
	Error string `json:"error" example:"Incorrect request body"`
}

type StatusInternalServerError struct {
	Error string `json:"error" example:"Internal error"`
}

type StatusUnauthorized struct {
	Error string `json:"error" example:"Cant login user"`
}
