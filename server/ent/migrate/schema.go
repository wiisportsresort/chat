// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "content", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "message_children", Type: field.TypeString, Nullable: true},
		{Name: "user_messages", Type: field.TypeString, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_messages_children",
				Columns:    []*schema.Column{MessagesColumns[3]},
				RefColumns: []*schema.Column{MessagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "messages_users_messages",
				Columns:    []*schema.Column{MessagesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PrivateMessagesColumns holds the columns for the "private_messages" table.
	PrivateMessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString, Unique: true},
		{Name: "content", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "private_message_children", Type: field.TypeString, Nullable: true},
		{Name: "user_private_messages", Type: field.TypeString, Nullable: true},
		{Name: "user_private_messages_received", Type: field.TypeString, Nullable: true},
	}
	// PrivateMessagesTable holds the schema information for the "private_messages" table.
	PrivateMessagesTable = &schema.Table{
		Name:       "private_messages",
		Columns:    PrivateMessagesColumns,
		PrimaryKey: []*schema.Column{PrivateMessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "private_messages_private_messages_children",
				Columns:    []*schema.Column{PrivateMessagesColumns[3]},
				RefColumns: []*schema.Column{PrivateMessagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "private_messages_users_private_messages",
				Columns:    []*schema.Column{PrivateMessagesColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "private_messages_users_private_messages_received",
				Columns:    []*schema.Column{PrivateMessagesColumns[5]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "username", Type: field.TypeString},
		{Name: "avatar_url", Type: field.TypeString, Nullable: true},
		{Name: "secret_hash", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		MessagesTable,
		PrivateMessagesTable,
		UsersTable,
	}
)

func init() {
	MessagesTable.ForeignKeys[0].RefTable = MessagesTable
	MessagesTable.ForeignKeys[1].RefTable = UsersTable
	PrivateMessagesTable.ForeignKeys[0].RefTable = PrivateMessagesTable
	PrivateMessagesTable.ForeignKeys[1].RefTable = UsersTable
	PrivateMessagesTable.ForeignKeys[2].RefTable = UsersTable
}