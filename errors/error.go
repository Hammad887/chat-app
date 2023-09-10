package errors

import "errors"

var (
	// ErrNoContent custom error for no record found
	ErrNoContent = errors.New("no model exist in database")
	// ErrConflict custom error for record already exist
	ErrConflict = errors.New("user model is already exist in database")
	// ErrChatRoomConflict custom error for record already exist
	ErrChatRoomConflict = errors.New("chatroom model is already exist in database")
	// ErrMessageConflict custom error for record already exist
	ErrMessageConflict = errors.New("message model is already exist in database")
)
