package ierr

func BadRequest(code int, format string, args ...any) *Error {
	return New(400, code, format, args...)
}

func IsBadRequest(err error) bool {
	return Status(err) == 400
}

func Unauthorized(code int, format string, args ...any) *Error {
	return New(401, code, format, args...)
}

func IsUnauthorized(err error) bool {
	return Status(err) == 401
}

func Forbidden(code int, format string, args ...any) *Error {
	return New(403, code, format, args...)
}

func IsForbidden(err error) bool {
	return Status(err) == 403
}

func NotFound(code int, format string, args ...any) *Error {
	return New(404, code, format, args...)
}

func IsNotFound(err error) bool {
	return Status(err) == 404
}

func Conflict(code int, format string, args ...any) *Error {
	return New(409, code, format, args...)
}

func IsConflict(err error) bool {
	return Status(err) == 409
}

func InternalServer(code int, format string, args ...any) *Error {
	return New(500, code, format, args...)
}

func IsInternalServer(err error) bool {
	return Status(err) == 500
}
