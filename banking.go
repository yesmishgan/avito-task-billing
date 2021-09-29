package cashbox

type Bill struct {
	Username    string `json:"username"`
	Amount      int    `json:"amount"`
	Flag        bool   `json:"flag"`
	Description string `json:"description"`
}

type User struct {
	Username string `json:"username" binding:"required"`
}

type Account struct {
	Balance int `json:"balance" db:"balance"`
}
