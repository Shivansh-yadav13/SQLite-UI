package routes

// Table struct helps to bind the post requests to the API.
type TableStruct struct {
	Name string `json:"table_name" form:"table_name"`
}
// Table struct helps to bind the post requests to the API.
type ColumnStruct struct {
	Name string `json:"column_name" form:"column_name"`
	Type string `json:"column_type" form:"column_type"`
}