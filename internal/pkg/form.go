package pkg

import (
	"errors"
	"net/http"
)

func ParseForm(r *http.Request, t any) error {
	return errors.New("unable to parse form")
}
