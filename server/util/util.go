package util

type Userlist struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authentication struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Question struct {
	Question string `json:"question"`
}

type Response struct {
	Answer string `json:"answer"`
}
