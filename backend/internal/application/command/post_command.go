package command

import "github.com/google/uuid"

type PostListCommand struct {
	Search *string
	Cursor *uuid.UUID
	UserId uuid.UUID
}

type PostListCommandResultPost struct {
	ID      uuid.UUID
	Content string
	Title   string
}

type PostListCommandResult struct {
	Posts          []PostListCommandResultPost
	NextCursor     *uuid.UUID
	PreviousCursor *uuid.UUID
}

type CreatePostCommand struct {
	Title   string
	Content string
	UserId  uuid.UUID
}
