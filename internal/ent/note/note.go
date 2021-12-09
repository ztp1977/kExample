// Code generated by entc, DO NOT EDIT.

package note

import (
	"time"
)

const (
	// Label holds the string label denoting the note type in the database.
	Label = "note"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNoteID holds the string denoting the note_id field in the database.
	FieldNoteID = "note_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// Table holds the table name of the note in the database.
	Table = "notes"
)

// Columns holds all SQL columns for note fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldNoteID,
	FieldCreatedAt,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
