package StringUtils

import (
	"errors"
)

// IsEmpty Checks if a string is empty ("").
//  StringUtils.isEmpty("")        = true
//  StringUtils.isEmpty(" ")       = false
//  StringUtils.isEmpty("abc")     = false
//  StringUtils.isEmpty("  abc  ") = false
func IsEmpty(s string) bool {
	return len(s) == 0
}

// IsNotEmpty Checks if a string is not empty ("").
//  StringUtils.isNotEmpty("")        = false
//  StringUtils.isNotEmpty(" ")       = true
//  StringUtils.isNotEmpty("abc")     = true
//  StringUtils.isNotEmpty("  abc  ") = true
func IsNotEmpty(s string) bool {
	return !IsEmpty(s)
}

// IsAllEmpty Checks if all the strings are empty ("").
//  StringUtils.isAllEmpty()             = true, error
//  StringUtils.isAllEmpty("")           = true
//  StringUtils.isAllEmpty("", "abc")    = false
//  StringUtils.isAllEmpty("abc", "")    = false
//  StringUtils.isAllEmpty(" ", "abc")   = false
//  StringUtils.isAllEmpty("abc", "cba") = false
func IsAllEmpty(ss ...string) (bool, error) {
	if len(ss) == 0 {
		return true, errors.New("haven't arguments to check")
	}
	for _, s := range ss {
		if IsNotEmpty(s) {
			return false, nil
		}
	}
	return true, nil
}

// IsAllNotEmpty Checks if all the strings are not empty ("").
//  StringUtils.IsAllNotEmpty()             = false, error
//  StringUtils.IsAllNotEmpty("")           = false
//  StringUtils.IsAllNotEmpty("", "abc")    = true
//  StringUtils.IsAllNotEmpty("abc", "")    = true
//  StringUtils.IsAllNotEmpty(" ", "abc")   = true
//  StringUtils.IsAllNotEmpty("abc", "cba") = true
func IsAllNotEmpty(ss ...string) (b bool, e error) {
	b, e = IsAllEmpty(ss...)
	return !b, e
}

// IsAnyEmpty Checks if any the strings are empty ("").
//  StringUtils.IsAnyEmpty()             = true, error
//  StringUtils.IsAnyEmpty("")           = true
//  StringUtils.IsAnyEmpty("", "abc")    = true
//  StringUtils.IsAnyEmpty("abc", "")    = true
//  StringUtils.IsAnyEmpty(" ", "abc")   = false
//  StringUtils.IsAnyEmpty("abc", "cba") = false
func IsAnyEmpty(ss ...string) (bool, error) {
	if len(ss) == 0 {
		return true, errors.New("haven't arguments to check")
	}
	for _, s := range ss {
		if IsEmpty(s) {
			return true, nil
		}
	}
	return false, nil
}

// IsNoneEmpty Checks if none of the strings are empty ("").
//  StringUtils.isNoneEmpty()             = false, error
//  StringUtils.isNoneEmpty("")           = false
//  StringUtils.isNoneEmpty("", "abc")    = false
//  StringUtils.isNoneEmpty("abc", "")    = false
//  StringUtils.isNoneEmpty(" ", "abc")   = true
//  StringUtils.isNoneEmpty("abc", "cba") = true
func IsNoneEmpty(ss ...string) (b bool, e error) {
	b, e = IsAnyEmpty(ss...)
	return !b, e
}

// DefaultIfEmpty Returns either the passed in s, or if the s is empty, the value of d.
//  StringUtils.defaultIfEmpty("", "abc")    = "abc"
//  StringUtils.defaultIfEmpty(" ", "abc")   = " "
//  StringUtils.defaultIfEmpty("abc", "cba") = "abc"
func DefaultIfEmpty(s string, d string) string {
	if IsEmpty(s) {
		return d
	}
	return s
}

// FirstNonEmpty Returns the first value which is not empty.
//  StringUtils.FirstNonEmpty()             = "", error
//  StringUtils.FirstNonEmpty("")           = "", error
//  StringUtils.FirstNonEmpty("", "abc")    = "abc"
//  StringUtils.FirstNonEmpty("abc", "")    = "abc"
//  StringUtils.FirstNonEmpty(" ", "abc")   = " "
//  StringUtils.FirstNonEmpty("abc", "cba") = "abc"
func FirstNonEmpty(ss ...string) (string, error) {
	if len(ss) == 0 {
		return "", errors.New("haven't arguments to check")
	}
	for _, s := range ss {
		if IsNotEmpty(s) {
			return s, nil
		}
	}
	return "", errors.New("haven't arguments which is not empty")
}

// GetIfEmpty Returns either the passed in s, or if the s is empty, the value supplied by f.
//  StringUtils.GetIfEmpty("", func() string { return "abc" }) 	  = "abc"
//  StringUtils.GetIfEmpty(" ", func() string { return "abc" })   = " "
//  StringUtils.GetIfEmpty("abc", func() string { return "cba" }) = "abc"
func GetIfEmpty(s string, f func() string) string {
	if IsEmpty(s) {
		return f()
	}
	return s
}
