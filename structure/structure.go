package structure

type Students struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}
type Read struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
}
