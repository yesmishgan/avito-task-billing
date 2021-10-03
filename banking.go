package cashbox

type Transfer struct {
	Description string `json:"description"`
	Amount int `json:"amount"`
	Username string `json:"username"`
	Destination string `json:"destination"`
}

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
	Balance float32 `json:"balance" db:"balance"`
}
