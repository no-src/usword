package _http

import (
	"errors"
	"github.com/no-src/usword/res/lang"
)

var (
	MethodNotSupported = errors.New(lang.HTTP_Error_MethodNotSupported)
)
