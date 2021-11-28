package StringUtils

import (
	"errors"
	"unicode"
)

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
//  StringUtils.isBlank("")        = true
//  StringUtils.isBlank(" ")       = true
//  StringUtils.isBlank("abc")     = false
//  StringUtils.isBlank("  abc  ") = false
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
//  StringUtils.isNotBlank("")        = false
//  StringUtils.isNotBlank(" ")       = false
//  StringUtils.isNotBlank("abc")     = true
//  StringUtils.isNotBlank("  abc  ") = true
func IsNotBlank(s string) bool {
	return !IsBlank(s)
}

// IsAllBlank Checks if all the strings are empty or whitespace only.
//  StringUtils.isAllBlank()             = true, error
//  StringUtils.isAllBlank("", " ")      = true
//  StringUtils.isAllBlank("", "abc")    = false
//  StringUtils.isAllBlank("abc", "")    = false
//  StringUtils.isAllBlank(" ", "abc")   = false
//  StringUtils.isAllBlank("abc", "cba") = false
func IsAllBlank(ss ...string) (bool, error) {
	if len(ss) == 0 {
		return true, errors.New("haven't arguments to check")
	}
	for _, s := range ss {
		if IsNotBlank(s) {
			return false, nil
		}
	}
	return true, nil
}

// IsAllNotBlank Checks if all the strings are empty or whitespace only.
//  StringUtils.IsAllNotBlank()             = false, error
//  StringUtils.IsAllNotBlank("", " ")      = false
//  StringUtils.IsAllNotBlank("", "abc")    = true
//  StringUtils.IsAllNotBlank("abc", "")    = true
//  StringUtils.IsAllNotBlank(" ", "abc")   = true
//  StringUtils.IsAllNotBlank("abc", "cba") = true
func IsAllNotBlank(ss ...string) (b bool, e error) {
	b, e = IsAllBlank(ss...)
	return !b, e
}

// IsAnyBlank Checks if any the strings are empty or whitespace only.
//  StringUtils.IsAnyBlank()             = true, error
//  StringUtils.IsAnyBlank("", " ")      = true
//  StringUtils.IsAnyBlank(" ", "abc")   = true
//  StringUtils.IsAnyBlank("abc", "")    = true
//  StringUtils.IsAnyBlank(" ", "abc")   = true
//  StringUtils.IsAnyBlank("abc", "cba") = false
func IsAnyBlank(ss ...string) (bool, error) {
	if len(ss) == 0 {
		return true, errors.New("haven't arguments to check")
	}
	for _, s := range ss {
		if IsBlank(s) {
			return true, nil
		}
	}
	return false, nil
}

// IsNoneBlank Checks if none of the strings are empty or whitespace only.
//  StringUtils.isNoneBlank()             = false, error
//  StringUtils.isNoneBlank("", " ")      = false
//  StringUtils.isNoneBlank("", "abc")    = false
//  StringUtils.isNoneBlank("abc", "")    = false
//  StringUtils.isNoneBlank(" ", "abc")   = false
//  StringUtils.isNoneBlank("abc", "cba") = true
func IsNoneBlank(ss ...string) (b bool, e error) {
	b, e = IsAnyBlank(ss...)
	return !b, e
}

// DefaultIfBlank Returns either the passed in s, or if the s is empty or whitespace only, the value of d.
//  StringUtils.defaultIfBlank("", "abc")    = "abc"
//  StringUtils.defaultIfBlank(" ", "abc")   = "abc"
//  StringUtils.defaultIfBlank("abc", "cba") = "abc"
func DefaultIfBlank(s string, d string) string {
	if IsBlank(s) {
		return d
	}
	return s
}

// FirstNonBlank Returns the first value which is not empty or whitespace only.
//  StringUtils.FirstNonBlank()             = "", error
//  StringUtils.FirstNonBlank("")           = "", error
//  StringUtils.FirstNonBlank("", "abc")    = "abc"
//  StringUtils.FirstNonBlank("abc", "")    = "abc"
//  StringUtils.FirstNonBlank(" ", "abc")   = "abc"
//  StringUtils.FirstNonBlank("abc", "cba") = "abc"
func FirstNonBlank(ss ...string) (string, error) {
	if len(ss) == 0 {
		return "", errors.New("haven't arguments to check")
	}
	for _, s := range ss {
		if IsNotBlank(s) {
			return s, nil
		}
	}
	return "", errors.New("haven't arguments which is not Blank")
}

// GetIfBlank Returns either the passed in s, or if the s is empty or whitespace only, the value supplied by f.
//  StringUtils.GetIfBlank("", func() string { return "abc" }) 	  = "abc"
//  StringUtils.GetIfBlank(" ", func() string { return "abc" })   = "abc"
//  StringUtils.GetIfBlank("abc", func() string { return "cba" }) = "abc"
func GetIfBlank(s string, f func() string) string {
	if IsBlank(s) {
		return f()
	}
	return s
}
