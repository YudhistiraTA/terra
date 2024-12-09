package mapper

import (
	"github.com/YudhistiraTA/terra/internal/application/command"
	"github.com/YudhistiraTA/terra/internal/interface/api/rest/dto/response"
)

func ToPostListResponse(postListCommandResult *command.PostListCommandResult) response.PostListReponse {
	var nextCursor *string
	if postListCommandResult.NextCursor != nil {
		nextCursorString := postListCommandResult.NextCursor.String()
		nextCursor = &nextCursorString
	}
	postListResponse := response.PostListReponse{
		Posts:      []response.PostListReponsePost{},
		NextCursor: nextCursor,
	}

	for _, post := range postListCommandResult.Posts {
		postListResponse.Posts = append(postListResponse.Posts, response.PostListReponsePost{
			ID:      post.ID.String(),
			Content: post.Content,
			Title:   post.Title,
		})
	}

	return postListResponse
}
