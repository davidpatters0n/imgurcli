package types

type JsonResponse struct {
	Data    bool `json:"data"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}
