package cashbox

type Bill struct {
	Username string `json:"username"`
	Money    int    `json:"money"`
	Flag     bool   `json:"flag"`
}
