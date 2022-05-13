package log

type Log struct {
	ID        string `json:"id"`
	At        string `json:"at"`
	Host      string `json:"host"`
	Method    string `json:"method"`
	Path      string `json:"path"`
	IP        string `json:"ip"`
	Status    int    `json:"status"`
	UserAgent string `json:"user_agent"`
}
