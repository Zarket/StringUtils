package stringutils

import (
	"errors"
	"unicode"
)

// ErrNoArguments means the function get 0 arguments or array zero length
var ErrNoArguments = errors.New("haven't arguments")

// ErrArrIsEmpty means the all strings, transferred to function, is empty
var ErrArrIsEmpty = errors.New("all strings is empty")

// ErrArrIsBlank means the all strings, transferred to function, is blank
var ErrArrIsBlank = errors.New("all strings is blank")

// IsEmpty Checks if a string is empty ("").
//  stringutils.isEmpty("")        = true
//  stringutils.isEmpty(" ")       = false
//  stringutils.isEmpty("abc")     = false
//  stringutils.isEmpty("  abc  ") = false
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsNotEmpty Checks if a string is not empty ("").
//  stringutils.isNotEmpty("")        = false
//  stringutils.isNotEmpty(" ")       = true
//  stringutils.isNotEmpty("abc")     = true
//  stringutils.isNotEmpty("  abc  ") = true
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsAllEmpty Checks if all the strings are empty ("").
//  stringutils.isAllEmpty()             = true, error
//  stringutils.isAllEmpty("")           = true
//  stringutils.isAllEmpty("", "abc")    = false
//  stringutils.isAllEmpty("abc", "")    = false
//  stringutils.isAllEmpty(" ", "abc")   = false
//  stringutils.isAllEmpty("abc", "cba") = false
func IsAllEmpty(ss ...string) (bool, error) {
	if len(ss) == 0 {
		return true, ErrNoArguments
	}
	for _, s := range ss {
		if IsNotEmpty(s) {
			return false, nil
		}
	}
	return true, nil
}

// IsNotAllEmpty Checks if not all the strings are empty ("").
//  stringutils.IsNotAllEmpty()             = false, error
//  stringutils.IsNotAllEmpty("")           = false
//  stringutils.IsNotAllEmpty("", "abc")    = true
//  stringutils.IsNotAllEmpty("abc", "")    = true
//  stringutils.IsNotAllEmpty(" ", "abc")   = true
//  stringutils.IsNotAllEmpty("abc", "cba") = true
func IsNotAllEmpty(ss ...string) (b bool, e error) {
	b, e = IsAllEmpty(ss...)
	return !b, e
}

// IsAnyNotEmpty Checks if any the strings are not empty ("").
//  stringutils.IsAnyNotEmpty()             = false, error
//  stringutils.IsAnyNotEmpty("")           = false
//  stringutils.IsAnyNotEmpty("", "abc")    = true
//  stringutils.IsAnyNotEmpty("abc", "")    = true
//  stringutils.IsAnyNotEmpty(" ", "abc")   = true
//  stringutils.IsAnyNotEmpty("abc", "cba") = true
func IsAnyNotEmpty(ss ...string) (b bool, e error) {
	return IsNotAllEmpty(ss...)
}

// IsAnyEmpty Checks if any the strings are empty ("").
//  stringutils.IsAnyEmpty()             = true, error
//  stringutils.IsAnyEmpty("")           = true
//  stringutils.IsAnyEmpty("", "abc")    = true
//  stringutils.IsAnyEmpty("abc", "")    = true
//  stringutils.IsAnyEmpty(" ", "abc")   = false
//  stringutils.IsAnyEmpty("abc", "cba") = false
func IsAnyEmpty(ss ...string) (bool, error) {
	if len(ss) == 0 {
		return true, ErrNoArguments
	}
	for _, s := range ss {
		if IsEmpty(s) {
			return true, nil
		}
	}
	return false, nil
}

// IsNoneEmpty Checks if none of the strings are empty ("").
//  stringutils.isNoneEmpty()             = false, error
//  stringutils.isNoneEmpty("")           = false
//  stringutils.isNoneEmpty("", "abc")    = false
//  stringutils.isNoneEmpty("abc", "")    = false
//  stringutils.isNoneEmpty(" ", "abc")   = true
//  stringutils.isNoneEmpty("abc", "cba") = true
func IsNoneEmpty(ss ...string) (b bool, e error) {
	b, e = IsAnyEmpty(ss...)
	return !b, e
}

// DefaultIfEmpty Returns either the passed in s, or if the s is empty, the value of d.
//  stringutils.defaultIfEmpty("", "abc")    = "abc"
//  stringutils.defaultIfEmpty(" ", "abc")   = " "
//  stringutils.defaultIfEmpty("abc", "cba") = "abc"
func DefaultIfEmpty(s string, d string) string {
	if IsEmpty(s) {
		return d
	}
	return s
}

// FirstNonEmpty Returns the first value which is not empty.
//  stringutils.FirstNonEmpty()             = "", error
//  stringutils.FirstNonEmpty("")           = "", error
//  stringutils.FirstNonEmpty("", "abc")    = "abc"
//  stringutils.FirstNonEmpty("abc", "")    = "abc"
//  stringutils.FirstNonEmpty(" ", "abc")   = " "
//  stringutils.FirstNonEmpty("abc", "cba") = "abc"
func FirstNonEmpty(ss ...string) (string, error) {
	if len(ss) == 0 {
		return "", ErrNoArguments
	}
	for _, s := range ss {
		if IsNotEmpty(s) {
			return s, nil
		}
	}
	return "", ErrArrIsEmpty
}

// GetIfEmpty Returns either the passed in s, or if the s is empty, the value supplied by f.
//  stringutils.GetIfEmpty("", func() string { return "abc" }) 	  = "abc"
//  stringutils.GetIfEmpty(" ", func() string { return "abc" })   = " "
//  stringutils.GetIfEmpty("abc", func() string { return "cba" }) = "abc"
func GetIfEmpty(s string, f func() string) string {
	if IsEmpty(s) {
		return f()
	}
	return s
}

// IsInformationSeparator reports whether the rune is information separator character
// as defined by Unicode's this is
// U+001C (FS), U+001D (GS), U+001E (RS), U+001F (ES)
func IsInformationSeparator(r rune) bool {
	if uint32(r) > unicode.MaxLatin1 {
		return false
	}
	switch r {
	case 0x1C, 0x1D, 0x1E, 0x1F:
		return true
	}
	return false
}

// IsBlank Checks if a string is empty or whitespace only.
//  stringutils.isBlank("")        = true
//  stringutils.isBlank(" ")       = true
//  stringutils.isBlank("abc")     = false
//  stringutils.isBlank("  abc  ") = false
func IsBlank(s string) bool {
	if IsEmpty(s) {
		return true
	}
	for _, r := range s {
		if !unicode.IsSpace(r) && !IsInformationSeparator(r) {
			return false
		}
	}
	return true
}

// IsNotBlank Checks if a string is not empty and not whitespace only.
//  stringutils.isNotBlank("")        = false
//  stringutils.isNotBlank(" ")       = false
//  stringutils.isNotBlank("abc")     = true
//  stringutils.isNotBlank("  abc  ") = true
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsAllBlank Checks if all the strings are empty or whitespace only.
//  stringutils.isAllBlank()             = true, error
//  stringutils.isAllBlank("", " ")      = true
//  stringutils.isAllBlank("", "abc")    = false
//  stringutils.isAllBlank("abc", "")    = false
//  stringutils.isAllBlank(" ", "abc")   = false
//  stringutils.isAllBlank("abc", "cba") = false
func IsAllBlank(ss ...string) (bool, error) {
	if len(ss) == 0 {
		return true, ErrNoArguments
	}
	for _, s := range ss {
		if IsNotBlank(s) {
			return false, nil
		}
	}
	return true, nil
}

// IsNotAllBlank Checks if not all the strings are empty or whitespace only.
//  stringutils.IsNotAllBlank()             = false, error
//  stringutils.IsNotAllBlank("", " ")      = false
//  stringutils.IsNotAllBlank("", "abc")    = true
//  stringutils.IsNotAllBlank("abc", "")    = true
//  stringutils.IsNotAllBlank(" ", "abc")   = true
//  stringutils.IsNotAllBlank("abc", "cba") = true
func IsNotAllBlank(ss ...string) (b bool, e error) {
	b, e = IsAllBlank(ss...)
	return !b, e
}

// IsAnyNotBlank Checks if any the strings are not empty or whitespace only.
//  stringutils.IsAnyNotBlank()             = false, error
//  stringutils.IsAnyNotBlank("", " ")      = false
//  stringutils.IsAnyNotBlank("", "abc")    = true
//  stringutils.IsAnyNotBlank("abc", "")    = true
//  stringutils.IsAnyNotBlank(" ", "abc")   = true
//  stringutils.IsAnyNotBlank("abc", "cba") = true
func IsAnyNotBlank(ss ...string) (b bool, e error) {
	return IsNotAllBlank(ss...)
}

// IsAnyBlank Checks if any the strings are empty or whitespace only.
//  stringutils.IsAnyBlank()             = true, error
//  stringutils.IsAnyBlank("", " ")      = true
//  stringutils.IsAnyBlank(" ", "abc")   = true
//  stringutils.IsAnyBlank("abc", "")    = true
//  stringutils.IsAnyBlank(" ", "abc")   = true
//  stringutils.IsAnyBlank("abc", "cba") = false
func IsAnyBlank(ss ...string) (bool, error) {
	if len(ss) == 0 {
		return true, ErrNoArguments
	}
	for _, s := range ss {
		if IsBlank(s) {
			return true, nil
		}
	}
	return false, nil
}

// IsNoneBlank Checks if none of the strings are empty or whitespace only.
//  stringutils.isNoneBlank()             = false, error
//  stringutils.isNoneBlank("", " ")      = false
//  stringutils.isNoneBlank("", "abc")    = false
//  stringutils.isNoneBlank("abc", "")    = false
//  stringutils.isNoneBlank(" ", "abc")   = false
//  stringutils.isNoneBlank("abc", "cba") = true
func IsNoneBlank(ss ...string) (b bool, e error) {
	b, e = IsAnyBlank(ss...)
	return !b, e
}

// DefaultIfBlank Returns either the passed in s, or if the s is empty or whitespace only, the value of d.
//  stringutils.defaultIfBlank("", "abc")    = "abc"
//  stringutils.defaultIfBlank(" ", "abc")   = "abc"
//  stringutils.defaultIfBlank("abc", "cba") = "abc"
func DefaultIfBlank(s string, d string) string {
	if IsBlank(s) {
		return d
	}
	return s
}

// FirstNonBlank Returns the first value which is not empty or whitespace only.
//  stringutils.FirstNonBlank()             = "", error
//  stringutils.FirstNonBlank("")           = "", error
//  stringutils.FirstNonBlank("", "abc")    = "abc"
//  stringutils.FirstNonBlank("abc", "")    = "abc"
//  stringutils.FirstNonBlank(" ", "abc")   = "abc"
//  stringutils.FirstNonBlank("abc", "cba") = "abc"
func FirstNonBlank(ss ...string) (string, error) {
	if len(ss) == 0 {
		return "", ErrNoArguments
	}
	for _, s := range ss {
		if IsNotBlank(s) {
			return s, nil
		}
	}
	return "", ErrArrIsBlank
}

// GetIfBlank Returns either the passed in s, or if the s is empty or whitespace only, the value supplied by f.
//  stringutils.GetIfBlank("", func() string { return "abc" }) 	  = "abc"
//  stringutils.GetIfBlank(" ", func() string { return "abc" })   = "abc"
//  stringutils.GetIfBlank("abc", func() string { return "cba" }) = "abc"
func GetIfBlank(s string, f func() string) string {
	if IsBlank(s) {
		return f()
	}
	return s
}
