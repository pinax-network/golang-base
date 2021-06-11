package helper

import (
	"fmt"
	"github.com/friendsofgo/errors"
)

func WrapErrorWithEmail(err error, email string) error {
	return errors.WithMessage(err, fmt.Sprintf("email '%s'", email))
}

func WrapErrorWithName(err error, name string) error {
	return errors.WithMessage(err, fmt.Sprintf("name '%s'", name))
}

func WrapErrorWithEosnId(err error, eosnId string) error {
	return errors.WithMessage(err, fmt.Sprintf("eosn id '%s'", eosnId))
}

func WrapErrorWithId(err error, id int) error {
	return errors.WithMessage(err, fmt.Sprintf("id '%d'", id))
}
