// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/wiisportsresort/chat/server/ent/message"
	"github.com/wiisportsresort/chat/server/ent/privatemessage"
	"github.com/wiisportsresort/chat/server/ent/schema"
	"github.com/wiisportsresort/chat/server/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescContent is the schema descriptor for content field.
	messageDescContent := messageFields[1].Descriptor()
	// message.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	message.ContentValidator = messageDescContent.Validators[0].(func(string) error)
	// messageDescCreatedAt is the schema descriptor for created_at field.
	messageDescCreatedAt := messageFields[2].Descriptor()
	// message.DefaultCreatedAt holds the default value on creation for the created_at field.
	message.DefaultCreatedAt = messageDescCreatedAt.Default.(func() time.Time)
	privatemessageFields := schema.PrivateMessage{}.Fields()
	_ = privatemessageFields
	// privatemessageDescContent is the schema descriptor for content field.
	privatemessageDescContent := privatemessageFields[1].Descriptor()
	// privatemessage.ContentValidator is a validator for the "content" field. It is called by the builders before save.
	privatemessage.ContentValidator = privatemessageDescContent.Validators[0].(func(string) error)
	// privatemessageDescCreatedAt is the schema descriptor for created_at field.
	privatemessageDescCreatedAt := privatemessageFields[2].Descriptor()
	// privatemessage.DefaultCreatedAt holds the default value on creation for the created_at field.
	privatemessage.DefaultCreatedAt = privatemessageDescCreatedAt.Default.(func() time.Time)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[1].Descriptor()
	// user.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	user.UsernameValidator = userDescUsername.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.IDValidator is a validator for the "id" field. It is called by the builders before save.
	user.IDValidator = userDescID.Validators[0].(func(string) error)
}
