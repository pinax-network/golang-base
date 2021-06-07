package helper

import (
	"fmt"
	"github.com/friendsofgo/errors"
)

func WrapErrorWithName(err error, name string) error {
	return errors.WithMessage(err, fmt.Sprintf("with name %q", name))
}

func WrapErrorWithId(err error, id int) error {
	return errors.WithMessage(err, fmt.Sprintf("with id %q", id))
}
