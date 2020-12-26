package apiserver

// AccountCreateRequest ...
type AccountCreateRequest struct {
	Login     string `json:"login"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsCompany bool   `json:"isCompany"`
}
