package StringUtils

import "testing"

func TestIsEmpty(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{""}, true},
		{"\\x00", args{"\x00"}, false},
		{"\\x00\\x01", args{"\x00\x01"}, false},
		{"\\uFFFF", args{"\uFFFF"}, false},
		{"space", args{" "}, false},
		{"abc", args{"abc"}, false},
		{"  abc  ", args{"  abc  "}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.s); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNotEmpty(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"empty", args{""}, false},
		{"\\x00", args{"\x00"}, true},
		{"\\x00\\x01", args{"\x00\x01"}, true},
		{"\\uFFFF", args{"\uFFFF"}, true},
		{"space", args{" "}, true},
		{"abc", args{"abc"}, true},
		{"  abc  ", args{"  abc  "}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsNotEmpty(tt.args.s); got != tt.want {
				t.Errorf("IsNotEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllEmpty(t *testing.T) {
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
		{"[empty,\\x00]", args{[]string{"", "\x00"}}, false, false},
		{"[empty,abc]", args{[]string{"", "abc"}}, false, false},
		{"[abc,empty]", args{[]string{"abc", ""}}, false, false},
		{"[space,abc]", args{[]string{" ", "abc"}}, false, false},
		{"[abc,cba]", args{[]string{"abc", "cba"}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsAllEmpty(tt.args.ss...)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAllEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsAllEmpty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsAllNotEmpty(t *testing.T) {
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
		{"[empty,\\x00]", args{[]string{"", "\x00"}}, true, false},
		{"[empty,abc]", args{[]string{"", "abc"}}, true, false},
		{"[abc,empty]", args{[]string{"abc", ""}}, true, false},
		{"[space,abc]", args{[]string{" ", "abc"}}, true, false},
		{"[abc,cba]", args{[]string{"abc", "cba"}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := IsAllNotEmpty(tt.args.ss...)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAllNotEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotB != tt.wantB {
				t.Errorf("IsAllNotEmpty() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestIsAnyEmpty(t *testing.T) {
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
		{"[empty,\\x00]", args{[]string{"", "\x00"}}, true, false},
		{"[empty,abc]", args{[]string{"", "abc"}}, true, false},
		{"[abc,empty]", args{[]string{"abc", ""}}, true, false},
		{"[space,abc]", args{[]string{" ", "abc"}}, false, false},
		{"[abc,cba]", args{[]string{"abc", "cba"}}, false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsAnyEmpty(tt.args.ss...)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsAnyEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsAnyEmpty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsNoneEmpty(t *testing.T) {
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
		{"[empty,\\x00]", args{[]string{"", "\x00"}}, false, false},
		{"[empty,abc]", args{[]string{"", "abc"}}, false, false},
		{"[abc,empty]", args{[]string{"abc", ""}}, false, false},
		{"[space,abc]", args{[]string{" ", "abc"}}, true, false},
		{"[abc,cba]", args{[]string{"abc", "cba"}}, true, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotB, err := IsNoneEmpty(tt.args.ss...)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsNoneEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotB != tt.wantB {
				t.Errorf("IsNoneEmpty() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func TestDefaultIfEmpty(t *testing.T) {
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
		{"[\\x00,abc]", args{"\x00", "abc"}, "\x00"},
		{"[abc,empty]", args{"abc", ""}, "abc"},
		{"[space,abc]", args{" ", "abc"}, " "},
		{"[abc,cba]", args{"abc", "cba"}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DefaultIfEmpty(tt.args.s, tt.args.d); got != tt.want {
				t.Errorf("DefaultIfEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstNonEmpty(t *testing.T) {
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
		{"[empty,\\x00]", args{[]string{"", "\x00"}}, "\x00", false},
		{"[empty,abc]", args{[]string{"", "abc"}}, "abc", false},
		{"[abc,empty]", args{[]string{"abc", ""}}, "abc", false},
		{"[space,abc]", args{[]string{" ", "abc"}}, " ", false},
		{"[abc,cba]", args{[]string{"abc", "cba"}}, "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstNonEmpty(tt.args.ss...)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstNonEmpty() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("FirstNonEmpty() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetIfEmpty(t *testing.T) {
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
		{"[\\x00, abc]", args{"\x00", func() string { return "abc" }}, "\x00"},
		{"[space, abc]", args{" ", func() string { return "abc" }}, " "},
		{"[abc, cba]", args{"abc", func() string { return "cba" }}, "abc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetIfEmpty(tt.args.s, tt.args.f); got != tt.want {
				t.Errorf("GetIfEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
