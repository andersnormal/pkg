package smtp

// EnhancedStatusCodeClass represents a class of enhanced status codes.
type EnhancedStatusCodeClass int

// EnhancedStatusCodeSubject represents a classification of the status codes.
type EnhancedStatusCodeSubject int

// EnhancedStatusCodeDetail represents the detailed status.
type EnhancedStatusCodeDetail int

// EnhancedStatusCode is a data structure to contain enhanced
// mail system status codes from RFC 3463 (https://datatracker.ietf.org/doc/html/rfc3463).
type EnhancedStatusCode [3]int

const (
	// Signals a positive delivery action.
	EnhancedStatusCodeSuccess EnhancedStatusCodeClass = 2
	// Signals that there is a temporary failure in positively delivery action.
	EnhancedStatusCodePersistentTransientFailure EnhancedStatusCodeClass = 4
	// Signals that there is a permanent failure in the delivery action.
	EnhancedStatusCodePermanentFailure EnhancedStatusCodeClass = 5
)
