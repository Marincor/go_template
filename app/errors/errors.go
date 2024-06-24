package constants

import "errors"

var (
	ErrInvalidUserAgent     = errors.New("invalid user agent")
	ErrMissingUserAgent     = errors.New("user agent is missing")
	ErrDatabaseNotConnected = errors.New("database is not connected")
	ErrAssertDBResponse     = errors.New("error while asserting database response")
)
