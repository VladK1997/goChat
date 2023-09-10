package chat_domain

type ChatFilter struct {
	Search string
	Limit  *int
	Offset *int
	Sort   string
}
