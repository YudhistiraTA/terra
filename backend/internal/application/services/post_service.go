package services

import (
	"context"

	"github.com/YudhistiraTA/terra/internal/application/command"
	"github.com/YudhistiraTA/terra/internal/infrastructure/db/sqlc"
	"github.com/google/uuid"
)

type PostService struct {
	ctx context.Context
	db  *sqlc.Queries
}

func NewPostService(ctx context.Context, db *sqlc.Queries) *PostService {
	return &PostService{ctx, db}
}

func (ps *PostService) CreatePost(cmd command.CreatePostCommand) error {
	err := ps.db.CreatePost(ps.ctx, sqlc.CreatePostParams{
		Title:   cmd.Title,
		Content: cmd.Content,
		UserID:  cmd.UserId,
	})
	if err != nil {
		return err
	}
	return nil

}

func (ps *PostService) GetPostList(cmd command.PostListCommand) (*command.PostListCommandResult, error) {
	var cursor *string
	if cmd.Cursor != nil {
		cursorString := cmd.Cursor.String()
		cursor = &cursorString
	}
	var previousCursor *uuid.UUID
	if cursor != nil {
		currentCursor := uuid.MustParse(*cursor)
		pc, err := ps.db.GetPreviousCursor(ps.ctx, sqlc.GetPreviousCursorParams{
			Cursor:     currentCursor,
			UserID:     cmd.UserId,
			SearchTerm: cmd.Search,
		})
		if err != nil {
			if err.Error() != "no rows in result set" {
				return nil, err
			}
		}
		if pc != uuid.Nil {
			previousCursor = &pc
		}
	}
	dbPost, err := ps.db.FuzzySearchPosts(ps.ctx, sqlc.FuzzySearchPostsParams{
		SearchTerm: cmd.Search,
		Cursor:     cursor,
		UserID:     cmd.UserId,
	})
	if err != nil {
		return nil, err
	}
	var finalPost sqlc.Post
	var nextCursor *uuid.UUID

	if len(dbPost) > 5 {
		finalPost = dbPost[len(dbPost)-1]
		dbPost = dbPost[:len(dbPost)-1]
		if finalPost.ID != uuid.Nil {
			nextCursor = &finalPost.ID
		}
	}

	var posts []command.PostListCommandResultPost
	for _, post := range dbPost {
		posts = append(posts, command.PostListCommandResultPost{
			ID:      post.ID,
			Content: post.Content,
			Title:   post.Title,
		})
	}
	return &command.PostListCommandResult{
		Posts:          posts,
		NextCursor:     nextCursor,
		PreviousCursor: previousCursor,
	}, nil
}
