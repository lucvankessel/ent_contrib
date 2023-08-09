// Code generated by ent, DO NOT EDIT.

package task

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the task type in the database.
	Label = "task"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldComplete holds the string denoting the complete field in the database.
	FieldComplete = "complete"
	// FieldSignature holds the string denoting the signature field in the database.
	FieldSignature = "signature"
	// Table holds the table name of the task in the database.
	Table = "tasks"
)

// Columns holds all SQL columns for task fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldDescription,
	FieldComplete,
	FieldSignature,
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

// OrderOption defines the ordering options for the Task queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByDescription orders the results by the description field.
func ByDescription(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDescription, opts...).ToFunc()
}

// ByComplete orders the results by the complete field.
func ByComplete(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldComplete, opts...).ToFunc()
}

// BySignature orders the results by the signature field.
func BySignature(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSignature, opts...).ToFunc()
}
