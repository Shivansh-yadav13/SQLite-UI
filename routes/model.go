package routes

// Table struct helps to bind the post requests to the API.
type Table struct {
	Name string `json:"table_name" form:"table_name"`
}
