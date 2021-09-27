package cashbox

type Bill struct {
	Username string `json:"username"`
	Money    int    `json:"money"`
	Flag     bool   `json:"flag"`
}

type User struct {
	Username string `json:"username" binding:"required"`
}

type Account struct {
	Balance int `json:"balance" db:"balance"`
}