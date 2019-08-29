package utils

type Apierror struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
