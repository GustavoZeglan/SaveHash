package password

type Password struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Value  string `json:"value"`
	UserID uint64 `json:"user_id"`
}
