package smtp

// EnhancedMailSystemStatusCode ...
// https://datatracker.ietf.org/doc/html/rfc3463
type EnhancedMailSystemStatusCode [3]int

// StatusCode ...
type StatusCode struct {
	replyCode    int
	enhancedCode EnhancedStatusCode
	message      string
}

// SMTPError ...
type SMTPError struct {
	statusCode StatusCode
}

// Error ...
func (e *SMTPError) Error() string {
	return e.statusCode.message
}

// NewStatusCode ...
func (s *StatusCode) NewStatusCode(replyCode int, enhancedCode EnhancedStatusCode, message string) *StatusCode {
	return &StatusCode{
		replyCode:    replyCode,
		enhancedCode: enhancedCode,
		message:      message,
	}
}

// Error ...
func Error(replyCode int, enhancedCode EnhancedStatusCode, message string) error {
	return &SMTPError{
		statusCode: StatusCode{
			replyCode:    replyCode,
			enhancedCode: enhancedCode,
			message:      message,
		},
	}
}
