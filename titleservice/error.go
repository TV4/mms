package titleservice

import "errors"

var (
	// ErrUnexpectedContentType is returned if response isn't JSON
	ErrUnexpectedContentType = errors.New("unexpected Content-Type")

	// ErrMissingParameter is returned if missing a parameter required by the MMS TitleService API
	ErrMissingParameter = errors.New("missing parameter")

	// ErrInvalidParameter is returned if a parameter is invalid for some reason
	ErrInvalidParameter = errors.New("invalid parameter")
)

// ErrorWithMessage annotates err with a new message.
// If err is nil, ErrorWithMessage returns nil.
func ErrorWithMessage(err error, msg string) error {
	if err == nil {
		return nil
	}

	return &errorWithMessage{
		cause: err,
		msg:   msg,
	}
}

type errorWithMessage struct {
	cause error
	msg   string
}

func (e *errorWithMessage) Error() string {
	return e.msg + ": " + e.cause.Error()
}

func (e *errorWithMessage) Cause() error {
	return e.cause
}

// ErrorCause returns the underlying cause of the error, if possible.
// An error value has a cause if it implements the following
// interface:
//
//     type causer interface {
//            Cause() error
//     }
//
// If the error does not implement Cause, the original error will
// be returned. If the error is nil, nil will be returned without further
// investigation.
func ErrorCause(err error) error {
	type causer interface {
		Cause() error
	}

	for err != nil {
		cause, ok := err.(causer)
		if !ok {
			break
		}
		err = cause.Cause()
	}
	return err
}
