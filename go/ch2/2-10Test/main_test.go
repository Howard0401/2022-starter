package main

import "testing"

func TestShow(t *testing.T) {
	result := Show()
	wanted := "Show Func stringsss"
	if result == wanted {
		t.Logf("Show() = %v, want = %v", result, wanted)
	} else {
		t.Errorf("t.Errorf")
	}
}

func TestShowWithTable(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			"str1",
			"Show Func string", // "str2", failed
		},
		// {
		// 	"str1",
		// 	"str2", // "str2", failed
		// },
	}
	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			if got := Show(); got != v.want {
				t.Errorf("inside T error, Show=%v, want=%v", got, v.want)
			}
		})
	}
}
