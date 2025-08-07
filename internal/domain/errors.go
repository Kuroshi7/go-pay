package domain

import "errors"

var(
	ErrAccountNotFound = errors.New("account not found")
	ErrDuplicateAPIKey = errors.New("duplicate API key")
	ErrInvoiceNotFound = errors.New("invoice not found")
	ErrUnauthorizedAccess = errors.New("unauthorized access")
)