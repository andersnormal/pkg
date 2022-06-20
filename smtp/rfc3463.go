package smtp

import (
	"strconv"
	"strings"
)

// EnhancedStatusCodeClass represents a class of enhanced status codes.
type EnhancedStatusCodeClass int

// EnhancedStatusCodeSubject represents a classification of the status codes.
type EnhancedStatusCodeSubject int

// EnhancedStatusCodeDetail represents the detailed status.
type EnhancedStatusCodeDetail int

// EnhancedStatusCodeDescription represents the message of this error code.
type EnhancedStatusCodeDescription string

// EnhancedStatusCode is a data structure to contain enhanced
// mail system status codes from RFC 3463 (https://datatracker.ietf.org/doc/html/rfc3463).
type EnhancedStatusCode struct {
	class   EnhancedStatusCodeClass
	subject EnhancedStatusCodeSubject
	detail  EnhancedStatusCodeDetail
	desc    EnhancedStatusCodeDescription
}

const (
	// Signals a positive delivery action.
	EnhancedStatusCodeSuccess EnhancedStatusCodeClass = 2
	// Signals that there is a temporary failure in positively delivery action.
	EnhancedStatusCodePersistentTransientFailure EnhancedStatusCodeClass = 4
	// Signals that there is a permanent failure in the delivery action.
	EnhancedStatusCodePermanentFailure EnhancedStatusCodeClass = 5
)

// String returns formated string of the enhanced status code.
func (e *EnhancedStatusCode) String() string {
	return strings.Join([]string{strconv.Itoa(int(e.class)), strconv.Itoa(int(e.subject)), strconv.Itoa(int(e.detail))}, ".")
}

// NewStatusCode returns a new enhanced mail system status code.
func NewStatusCode(class EnhancedStatusCodeClass, subject EnhancedStatusCodeSubject, detail EnhancedStatusCodeDetail, desc EnhancedStatusCodeDescription) *EnhancedStatusCode {
	return &EnhancedStatusCode{class, subject, detail, desc}
}

// NewSuccess returns a new enhanced mail system status code of a successful delivery action.
func NewSuccess(subject EnhancedStatusCodeSubject, detail EnhancedStatusCodeDetail, desc EnhancedStatusCodeDescription) *EnhancedStatusCode {
	return &EnhancedStatusCode{class: EnhancedStatusCodeSuccess, subject: subject, detail: detail, desc: desc}
}

// NewTransientFailure returns a new enhanced mail system status code of a transient failure delivery action.
func NewTransientFailure(subject EnhancedStatusCodeSubject, detail EnhancedStatusCodeDetail, desc EnhancedStatusCodeDescription) *EnhancedStatusCode {
	return &EnhancedStatusCode{class: EnhancedStatusCodePersistentTransientFailure, subject: subject, detail: detail, desc: desc}
}

// NewPermanentFailure returns a new enhanced mail system status code of a permanently failing delivery action.
func NewPermanentFailure(subject EnhancedStatusCodeSubject, detail EnhancedStatusCodeDetail, desc EnhancedStatusCodeDescription) *EnhancedStatusCode {
	return &EnhancedStatusCode{class: EnhancedStatusCodePermanentFailure, subject: subject, detail: detail, desc: desc}
}

// SetDetail is setting the details on the status code.
func (e *EnhancedStatusCode) SetDetail(detail EnhancedStatusCodeDetail) {
	e.detail = detail
}

// Detail is returning the current detail of the status code.
func (e *EnhancedStatusCode) Detail() EnhancedStatusCodeDetail {
	return e.detail
}

// SetSubject is setting the subject on the status code.
func (e *EnhancedStatusCode) SetSubject(subject EnhancedStatusCodeSubject) {
	e.subject = subject
}

// Subject is returning the current subject of the status code.
func (e *EnhancedStatusCode) Subject() EnhancedStatusCodeSubject {
	return e.subject
}

// SetClass is setting the class on the status code.
func (e *EnhancedStatusCode) SetClass(class EnhancedStatusCodeClass) {
	e.class = class
}

// Class is returning the current class of the status code.
func (e *EnhancedStatusCode) Class() EnhancedStatusCodeClass {
	return e.class
}

var (
	PermanentBadDestinationMailboxAddress = NewPermanentFailure(1, 1, "The mailbox specified in the address does not exist.")
)
