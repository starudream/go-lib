package ierr

func BadRequest(reason, format string, args ...any) *Error {
	return New(400, reason, format, args...)
}

func IsBadRequest(err error) bool {
	return Code(err) == 400
}

func Unauthorized(reason, format string, args ...any) *Error {
	return New(401, reason, format, args...)
}

func IsUnauthorized(err error) bool {
	return Code(err) == 401
}

func Forbidden(reason, format string, args ...any) *Error {
	return New(403, reason, format, args...)
}

func IsForbidden(err error) bool {
	return Code(err) == 403
}

func NotFound(reason, format string, args ...any) *Error {
	return New(404, reason, format, args...)
}

func IsNotFound(err error) bool {
	return Code(err) == 404
}

func Conflict(reason, format string, args ...any) *Error {
	return New(409, reason, format, args...)
}

func IsConflict(err error) bool {
	return Code(err) == 409
}

func InternalServer(reason, format string, args ...any) *Error {
	return New(500, reason, format, args...)
}

func IsInternalServer(err error) bool {
	return Code(err) == 500
}
