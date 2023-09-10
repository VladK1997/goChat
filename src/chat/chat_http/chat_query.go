package chat_http

type TodoQuery struct {
	Search string `query:"search"`
	Limit  *int   `query:"limit"`
	Offset *int   `query:"offset"`
	Sort   string `query:"sort"`
}
