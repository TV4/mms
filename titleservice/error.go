package titleservice

import "errors"

var (
	// ErrUnexpectedContentType is returned if response isn't JSON
	ErrUnexpectedContentType = errors.New("unexpected Content-Type")

	// ErrMissingParameter is returned if missing a parameter required by the MMS TitleService API
	ErrMissingParameter = errors.New("missing parameter")

	// ErrInvalidParameter is returned if a parameter is invalid for some reason
	ErrInvalidParameter = errors.New("invalid parameter")

	// ErrNoUsername is returned if username is empty
	ErrNoUsername = errors.New("no username")

	// ErrNoPassword is returned if password is empty
	ErrNoPassword = errors.New("no password")

	// ErrInvalidInputData is returned on status 400 from the MMS TitleService API
	ErrInvalidInputData = errors.New("invalid input data (bad request)")

	// ErrAuthenticationFailure is returned on status 403 from the MMS TitleService API
	ErrAuthenticationFailure = errors.New("authentication failure (forbidden)")

	// ErrAlreadyRegistered is returned on status 409 from the MMS TitleService API
	ErrAlreadyRegistered = errors.New("already registered (conflict)")

	// ErrInternalServerError is returned on status 500 from the MMS TitleService API
	ErrInternalServerError = errors.New("internal server error")
)

// newErrorWithMessage annotates err with a new message.
// If err is nil, newErrorWithMessage returns nil.
func newErrorWithMessage(err error, msg string) error {
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
