package db

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Message struct {
	ID     string `json:"id"`
	Sender string `json:"sender"`
	Text   string `json:"text"`
	Media  string `json:"media"`
	Date   int64  `json:"date"` // func (t time.Time) UnixMilli() int64
}

type Chat struct {
	ID       string   `json:"id"`
	Users    []string `json:"users"`
	Messages []string `json:"messages"`
	Active   int64    `json:"active"`
}
