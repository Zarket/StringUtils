package StringUtils

import "testing"

func TestIsBlank(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{""}, true},
		{"\\x09\\x0A\\x0B\\x0C\\x0D\\x20", args{"\x09\x0A\x0B\x0C\x0D\x20"}, true},
		{"\\x85\\xA0", args{string([]rune{'\x85', '\xA0'})}, true}, //WTF! "\x85" = 65533 but '\x85' is 133
		{"space\\t\\n\\v\\f\\r", args{" \t\n\v\f\r"}, true},
		{"\\u0009\\u000A\\u000B\\u000C\\u000D", args{"\u0009\u000a\u000b\u000c\u000d"}, true},
		{"\\u001C\\u001D\\u001E\\u001F", args{"\u001C\u001D\u001E\u001F"}, true},
		{"\\u0020\\u0085\\u00A0\\u2007\\u202F", args{"\u0020\u0085\u00A0\u2007\u202F"}, true},
		{"\\x00", args{"\x00"}, false},
		{"\\uFFFF", args{"\uFFFF"}, false},
		{"abc", args{"abc"}, false},
		{"  abc  ", args{"  abc  "}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBlank(tt.args.s); got != tt.want {
				t.Errorf("IsBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotBlank(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{""}, false},
		{"\\x09\\x0A\\x0B\\x0C\\x0D\\x20", args{"\x09\x0A\x0B\x0C\x0D\x20"}, false},
		{"\\x85\\xA0", args{string([]rune{'\x85', '\xA0'})}, false}, //WTF! "\x85" = 65533 but '\x85' is 133
		{"space\\t\\n\\v\\f\\r", args{" \t\n\v\f\r"}, false},
		{"\\u0009\\u000A\\u000B\\u000C\\u000D", args{"\u0009\u000A\u000B\u000C\u000D"}, false},
		{"\\u001C\\u001D\\u001E\\u001F", args{"\u001C\u001D\u001E\u001F"}, false},
		{"\\u0020\\u0085\\u00A0\\u2007\\u202F", args{"\u0020\u0085\u00A0\u2007\u202F"}, false},
		{"\\x00", args{"\x00"}, true},
		{"\\uFFFF", args{"\uFFFF"}, true},
		{"abc", args{"abc"}, true},
		{"  abc  ", args{"  abc  "}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotBlank(tt.args.s); got != tt.want {
				t.Errorf("IsNotBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllBlank(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"[]", args{[]string{}}, true, true},
		{"[empty]", args{[]string{""}}, true, false},
		{"[empty,\\x09,\\x0a,\\x0b,\\x0c,\\x0d,\\x20]",
			args{[]string{"", "\x09", "\x0a", "\x0b", "\x0c", "\x0d", "\x20"}}, true, false},
		{"[\\x85\\xA0]", args{[]string{string([]rune{'\x85', '\xA0'})}}, true, false}, //WTF! "\x85" = 65533 but '\x85' is 133
		{"[empty,space,\\t,\\n,\\v,\\f,\\r]",
			args{[]string{"", " ", "\t", "\n", "\v", "\f", "\r"}}, true, false},
		{"[empty,\\u0009,\\u000a,\\u000b,\\u000c,\\u000d]",
			args{[]string{"", "\u0009", "\u000a", "\u000b", "\u000c", "\u000d"}}, true, false},
		{"[empty,\\u001c,\\u001d,\\u001e,\\u001f]",
			args{[]string{"", "\u001c", "\u001d", "\u001e", "\u001f"}}, true, false},
		{"[empty,\\u0020,\\u0085,\\u00A0,\\u2007,\\u202F]",
			args{[]string{"", "\u0020", "\u0085", "\u00A0", "\u2007", "\u202F"}}, true, false},
		{"[empty,abc]", args{[]string{"", "abc"}}, false, false},
		{"[abc,empty]", args{[]string{"abc", ""}}, false, false},
		{"[space,abc]", args{[]string{" ", "abc"}}, false, false},
		{"[abc,cba]", args{[]string{"abc", "cba"}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsAllBlank(tt.args.ss...)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAllBlank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsAllBlank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllNotBlank(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name    string
		args    args
		wantB   bool
		wantErr bool
	}{
		{"[]", args{[]string{}}, false, true},
		{"[empty]", args{[]string{""}}, false, false},
		{"[empty,\\x09,\\x0A,\\x0B,\\x0C,\\x0D,\\x20]",
			args{[]string{"", "\x09", "\x0A", "\x0B", "\x0C", "\x0D", "\x20"}}, false, false},
		{"[\\x85\\xA0]", args{[]string{string([]rune{'\x85', '\xA0'})}}, false, false}, //WTF! "\x85" = 65533 but '\x85' is 133
		{"[empty,space,\\t,\\n,\\v,\\f,\\r]",
			args{[]string{"", " ", "\t", "\n", "\v", "\f", "\r"}}, false, false},
		{"[empty,\\u0009,\\u000A,\\u000B,\\u000C,\\u000D]",
			args{[]string{"", "\u0009", "\u000A", "\u000B", "\u000C", "\u000D"}}, false, false},
		{"[empty,\\u001C,\\u001D,\\u001E,\\u001F]",
			args{[]string{"", "\u001C", "\u001D", "\u001E", "\u001F"}}, false, false},
		{"[empty,\\u0020,\\u0085,\\u00A0,\\u2007,\\u202F]",
			args{[]string{"", "\u0020", "\u0085", "\u00A0", "\u2007", "\u202F"}}, false, false},
		{"[empty,abc]", args{[]string{"", "abc"}}, true, false},
		{"[abc,empty]", args{[]string{"abc", ""}}, true, false},
		{"[space,abc]", args{[]string{" ", "abc"}}, true, false},
		{"[abc,cba]", args{[]string{"abc", "cba"}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := IsAllNotBlank(tt.args.ss...)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAllNotBlank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotB != tt.wantB {
				t.Errorf("IsAllNotBlank() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestIsAnyBlank(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{"[]", args{[]string{}}, true, true},
		{"[empty]", args{[]string{""}}, true, false},
		{"[empty,\\x09,\\x0A,\\x0B,\\x0C,\\x0D,\\x20]",
			args{[]string{"", "\x09", "\x0A", "\x0B", "\x0C", "\x0D", "\x20"}}, true, false},
		{"[\\x85\\xA0]", args{[]string{string([]rune{'\x85', '\xA0'})}}, true, false}, //WTF! "\x85" = 65533 but '\x85' is 133
		{"[empty,space,\\t,\\n,\\v,\\f,\\r]",
			args{[]string{"", " ", "\t", "\n", "\v", "\f", "\r"}}, true, false},
		{"[empty,\\u0009,\\u000A,\\u000B,\\u000C,\\u000D]",
			args{[]string{"", "\u0009", "\u000A", "\u000B", "\u000C", "\u000D"}}, true, false},
		{"[empty,\\u001C,\\u001D,\\u001E,\\u001F]",
			args{[]string{"", "\u001C", "\u001D", "\u001E", "\u001F"}}, true, false},
		{"[empty,\\u0020,\\u0085,\\u00A0,\\u2007,\\u202F]",
			args{[]string{"", "\u0020", "\u0085", "\u00A0", "\u2007", "\u202F"}}, true, false},
		{"[empty,abc]", args{[]string{"", "abc"}}, true, false},
		{"[abc,empty]", args{[]string{"abc", ""}}, true, false},
		{"[space,abc]", args{[]string{" ", "abc"}}, true, false},
		{"[abc,cba]", args{[]string{"abc", "cba"}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsAnyBlank(tt.args.ss...)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAnyBlank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsAnyBlank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNoneBlank(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name    string
		args    args
		wantB   bool
		wantErr bool
	}{
		{"[]", args{[]string{}}, false, true},
		{"[empty]", args{[]string{""}}, false, false},
		{"[empty,\\x09,\\x0A,\\x0B,\\x0C,\\x0D,\\x20]",
			args{[]string{"", "\x09", "\x0A", "\x0B", "\x0C", "\x0D", "\x20"}}, false, false},
		{"[\\x85\\xA0]", args{[]string{string([]rune{'\x85', '\xA0'})}}, false, false}, //WTF! "\x85" = 65533 but '\x85' is 133
		{"[empty,space,\\t,\\n,\\v,\\f,\\r]",
			args{[]string{"", " ", "\t", "\n", "\v", "\f", "\r"}}, false, false},
		{"[empty,\\u0009,\\u000A,\\u000B,\\u000C,\\u000D]",
			args{[]string{"", "\u0009", "\u000A", "\u000B", "\u000C", "\u000D"}}, false, false},
		{"[empty,\\u001C,\\u001D,\\u001E,\\u001F]",
			args{[]string{"", "\u001C", "\u001D", "\u001E", "\u001F"}}, false, false},
		{"[empty,\\u0020,\\u0085,\\u00A0,\\u2007,\\u202F]",
			args{[]string{"", "\u0020", "\u0085", "\u00A0", "\u2007", "\u202F"}}, false, false},
		{"[empty,abc]", args{[]string{"", "abc"}}, false, false},
		{"[abc,empty]", args{[]string{"abc", ""}}, false, false},
		{"[space,abc]", args{[]string{" ", "abc"}}, false, false},
		{"[abc,cba]", args{[]string{"abc", "cba"}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := IsNoneBlank(tt.args.ss...)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsNoneBlank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotB != tt.wantB {
				t.Errorf("IsNoneBlank() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestDefaultIfBlank(t *testing.T) {
	type args struct {
		s string
		d string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"[empty,empty]", args{"", ""}, ""},
		{"[empty,abc]", args{"", "abc"}, "abc"},
		{"[\\x09\\x0A\\x0B\\x0C\\x0D\\x20,abc]", args{"\x09\x0A\x0B\x0C\x0D\x20", "abc"}, "abc"},
		{"[\\x85\\xA0,acb]", args{string([]rune{'\x85', '\xA0'}), "abc"}, "abc"}, //WTF! "\x85" = 65533 but '\x85' is 133
		{"[space\\t\\n\\v\\f\\r,abc]", args{" \t\n\v\f\r", "abc"}, "abc"},
		{"[\\u0009\\u000A\\u000B\\u000C\\u000D,abc]", args{"\u0009\u000A\u000B\u000C\u000D", "abc"}, "abc"},
		{"[\\u001C\\u001D\\u001E\\u001F,abc]", args{"\u001C\u001D\u001E\u001F", "abc"}, "abc"},
		{"[\\u0020\\u0085\\u00A0\\u2007\\u202F,abc]", args{"\u0020\u0085\u00A0\u2007\u202F", "abc"}, "abc"},
		{"[\\x00,abc]", args{"\x00", "abc"}, "\x00"},
		{"[abc,empty]", args{"abc", ""}, "abc"},
		{"[space,abc]", args{" ", "abc"}, "abc"},
		{"[abc,cba]", args{"abc", "cba"}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultIfBlank(tt.args.s, tt.args.d); got != tt.want {
				t.Errorf("DefaultIfBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstNonBlank(t *testing.T) {
	type args struct {
		ss []string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{"[]", args{[]string{}}, "", true},
		{"[empty]", args{[]string{""}}, "", true},
		{"[empty,\\x09,\\x0A,\\x0B,\\x0C,\\x0D,\\x20,\\x00]",
			args{[]string{"", "\x09", "\x0A", "\x0B", "\x0C", "\x0D", "\x20", "\x00"}}, "\x00", false},
		{"[\\x85\\xA0\\x00]", args{[]string{string([]rune{'\x85', '\xA0'}), "\x00"}}, "\x00", false}, //WTF! "\x85" = 65533 but '\x85' is 133
		{"[empty,space,\\t,\\n,\\v,\\f,\\r]",
			args{[]string{"", " ", "\t", "\n", "\v", "\f", "\r", "\x00"}}, "\x00", false},
		{"[empty,\\u0009,\\u000A,\\u000B,\\u000C,\\u000D,\\x00]",
			args{[]string{"", "\u0009", "\u000A", "\u000B", "\u000C", "\u000D", "\x00"}}, "\x00", false},
		{"[empty,\\u001C,\\u001D,\\u001E,\\u001F,\\x00]",
			args{[]string{"", "\u001C", "\u001D", "\u001E", "\u001F", "\x00"}}, "\x00", false},
		{"[empty,\\u0020,\\u0085,\\u00A0,\\u2007,\\u202F,\\x00]",
			args{[]string{"", "\u0020", "\u0085", "\u00A0", "\u2007", "\u202F", "\x00"}}, "\x00", false},
		{"[empty,abc]", args{[]string{"", "abc"}}, "abc", false},
		{"[abc,empty]", args{[]string{"abc", ""}}, "abc", false},
		{"[space,abc]", args{[]string{" ", "abc"}}, "abc", false},
		{"[abc,cba]", args{[]string{"abc", "cba"}}, "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstNonBlank(tt.args.ss...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstNonBlank() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FirstNonBlank() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIfBlank(t *testing.T) {
	type args struct {
		s string
		f func() string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"[empty, abc]", args{"", func() string { return "abc" }}, "abc"},
		{"[\\x09\\x0A\\x0B\\x0C\\x0D\\x20]",
			args{"\x09\x0A\x0B\x0C\x0D\x20", func() string { return "abc" }}, "abc"},
		{"\\x85\\xA0", args{string([]rune{'\x85', '\xA0'}), func() string { return "abc" }}, "abc"}, //WTF! "\x85" = 65533 but '\x85' is 133
		{"[space\\t\\n\\v\\f\\r]",
			args{" \t\n\v\f\r", func() string { return "abc" }}, "abc"},
		{"[\\u0009\\u000A\\u000B\\u000C\\u000D]",
			args{"\u0009\u000A\u000B\u000C\u000D", func() string { return "abc" }}, "abc"},
		{"[\\u001C\\u001D\\u001E\\u001F]",
			args{"\u001C\u001D\u001E\u001F", func() string { return "abc" }}, "abc"},
		{"[\\u0020\\u0085\\u00A0\\u2007\\u202F]",
			args{"\u0020\u0085\u00A0\u2007\u202F", func() string { return "abc" }}, "abc"},
		{"[\\x00, abc]", args{"\x00", func() string { return "abc" }}, "\x00"},
		{"[space, abc]", args{" ", func() string { return "abc" }}, "abc"},
		{"[abc, cba]", args{"abc", func() string { return "cba" }}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIfBlank(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("GetIfBlank() = %v, want %v", got, tt.want)
			}
		})
	}
}
