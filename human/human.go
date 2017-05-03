package human

// type hints allow for JSON parsing of mismatched names
type Human struct {
	Name string `json:"personName"`
	Age  int    `json:"personAge"`
}
