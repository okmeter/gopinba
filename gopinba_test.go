package gopinba

import (
	"testing"
	"time"
)

func inStringSlice(haystack []string, needle string) bool {
	for _, s := range haystack {
		if s == needle {
			return true
		}
	}
	return false
}

func allInStringSlice(haystack []string, elements []string, t *testing.T) {
	for _, element := range elements {
		if !inStringSlice(haystack, element) {
			t.Errorf("%v should be in string slice %v, but it is not", element, haystack)
		}
	}
}

func TestUseAndUpdateDictionary(t *testing.T) {

	m1 := map[string]string{
		"marko": "kevac",
		"margo": "kevac",
		"kevac": "lazo",
	}

	m2 := map[string]string{
		"liza":   "kevac",
		"pontiy": "kevac",
		"kevac":  "kursor",
	}

	var dict []string

	_, dict = useAndUpdateDictionary(dict, m1)

	if len(dict) != 4 {
		t.Error("dict length should be 4")
	}

	allInStringSlice(dict, []string{"marko", "margo", "lazo", "kevac"}, t)

	_, dict = useAndUpdateDictionary(dict, m2)

	if len(dict) != 7 {
		t.Error("dict length should be 7")
	}

	allInStringSlice(dict, []string{"marko", "margo", "lazo", "kevac", "liza", "pontiy", "kursor"}, t)
}

func TestRequest(t *testing.T) {
	pc, err := NewClient("10.0.0.1:30002")
	if err != nil {
		t.Errorf("NewClient() returned error: %v", err)
	}

	req := Request{}

	for i := 0; i < 5; i++ {

		req.Hostname = "hostname"
		req.ServerName = "servername"
		req.ScriptName = "scriptname"
		req.RequestCount = 1
		req.RequestTime = 145987 * time.Microsecond
		req.DocumentSize = 1024

		err = pc.SendRequest(&req)
		if err != nil {
			t.Errorf("SendRequest() returned error: %v", err)
		}
	}
}
