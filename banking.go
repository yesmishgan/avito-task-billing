package cashbox

type Bill struct {
	Username    string `json:"username" binding:"required"`
	Amount      int    `json:"amount" binding:"required"`
	Flag        bool   `json:"flag" binding:"required"`
	Description string `json:"description"`
}

type User struct {
	Username string `json:"username" binding:"required"`
}

type Account struct {
	Balance int `json:"balance" db:"balance"`
}
