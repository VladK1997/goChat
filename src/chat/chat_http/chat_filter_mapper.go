package chat_http

import "goChat/src/chat/chat_domain"

func ChatFilterFromQuery(query *TodoQuery) *chat_domain.ChatFilter {
	if query == nil {
		return nil
	}

	return &chat_domain.ChatFilter{
		Search: query.Search,
		Limit:  query.Limit,
		Offset: query.Offset,
		Sort:   query.Sort,
	}
}
