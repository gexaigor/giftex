package apiserver

import (
	"errors"

	"github.com/lib/pq"
)

var (
	errorNotFound                 = errors.New("not found")
	errorBadRequest               = errors.New("bad request")
	errorIncorrectLoginOrPassword = errors.New("incorrect login or password")
	errorConvertation             = errors.New("convertation failed")
	errorForbidden                = errors.New("dont have permision")
	errorTokenIsNotValid          = errors.New("token is not valid")
	errorMalfordedToken           = errors.New("malformed authentication token")
	errorInvalidToken             = errors.New("invalid/malformed auth token")
	errorMissingToken             = errors.New("missing auth token")
	errorUniqueViolation          = errors.New("not unique value")
)

const (
	pgUniqueViolation = pq.ErrorCode("23505")
)
