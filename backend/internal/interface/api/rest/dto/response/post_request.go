package response

type PostListReponsePost struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	Title   string `json:"title"`
}

type PostListReponse struct {
	Posts      []PostListReponsePost `json:"posts"`
	NextCursor *string               `json:"next_cursor"`
}
