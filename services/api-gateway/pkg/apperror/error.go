package apperror

type Error struct {
	Description    string `json:"description"`
	HTTPStatusCode int    `json:"status_code"`
	GRPCStatusCode int    `json:"-"`
}

func (e Error) Error() string {
	return e.Description
}
