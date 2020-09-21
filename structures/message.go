package structures

type Message struct {
	Error bool 	`json:"error"`
	Msg   string 	`json:"msg"`
	Id    int       `json:"id"`
}
