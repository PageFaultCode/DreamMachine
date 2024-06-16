// Package commonwealth for common data
package commonwealth

import (
	"strings"

	"github.com/google/uuid"
)

const (
	CookieTableName = "cookie"
)

// StatusData is for each application to report in
type StatusData struct {
	ApplicationName    string `json:"app_name"`
	ApplicationHealth  int    `json:"app_health"`
	ApplicationMessage string `json:"app_msg"`
}

// DatabaseQuery is for query the database
type DatabaseQuery struct {
	Query string `json:"query"`
	Error error  `json:"error"`
}

// Generic response data
type DatabaseResponse struct {
	Entries map[string]any
}

// CreateSubject create a subject from a number of pieces
func CreateSubject(pieces []string) string {
	return strings.Join(pieces, "/")
}

// StripBrackets removes the starting and ending brackets if any
func StripBrackets(value string) string {
	modified := strings.ReplaceAll(value, "[", "")

	modified = strings.ReplaceAll(modified, "]", "")

	return modified
}

// CookieEntry is the cookie structures that are used
type CookieEntry struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Value  string    `json:"value"`
	Active bool      `json:"active"`
}

// Cookie is a collection of cookie entries
// these may move to the web interface
type Cookie struct {
	CookieEntry CookieEntry `db:"cookie"`
}
