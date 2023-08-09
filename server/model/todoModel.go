package model

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"` // omitempty means that if the field is empty, it will not be included in the JSON response
	Completed   bool   `json:"completed"`
}
