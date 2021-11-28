package StringUtils

import "errors"

// ErrNoArguments means the function get 0 arguments or array zero length
var ErrNoArguments = errors.New("haven't arguments")

// ErrArrIsEmpty means the all strings, transferred to function, is empty
var ErrArrIsEmpty = errors.New("all strings is empty")

// ErrArrIsBlank means the all strings, transferred to function, is blank
var ErrArrIsBlank = errors.New("all strings is blank")
