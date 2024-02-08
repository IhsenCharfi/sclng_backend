package main

type Owner struct {
	Login string `json:"login"`
}
type Repository struct {
	ID            int         `json:"id"`
	Name          string      `json:"name"`
	FullName      string      `json:"full_name"`
	Owner         Owner       `json:"owner"`
	Languages_URL string      `json:"languages_url"`
	Languages     interface{} `json:"languages"`
}
