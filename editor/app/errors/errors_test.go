package errors_test

import (
	"editor/app/errors"
	lib_errors "errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var anyError = lib_errors.New("any error")
var nilErr = func() error {
	return nil
}
var withError = func() error {
	return anyError
}
var withPanic = func() error {
	panic("Any panic")
}

func TestCompose_Empty(t *testing.T) {
	assert := assert.New(t)
	err := errors.Compose()
	assert.Nil(err)
}

func TestCompose_Nil(t *testing.T) {
	assert := assert.New(t)
	err := errors.Compose(nilErr)
	assert.Nil(err)
}

func TestCompose_Error(t *testing.T) {
	assert := assert.New(t)
	err := errors.Compose(nilErr, withError)
	assert.ErrorIs(anyError, err)
}

func TestCompose_Panic(t *testing.T) {
	assert := assert.New(t)
	err := errors.Compose(nilErr, withError, withPanic)
	assert.ErrorIs(anyError, err)
}
