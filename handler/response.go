package handler

type Expense struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
type HTTPError struct {
	Code    int32  `json:"code"`
	Message string `json:"message"`
}
