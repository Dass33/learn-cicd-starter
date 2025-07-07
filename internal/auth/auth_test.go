package auth

import (
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := make(map[string][]string, 0)

	h["akldjf"] = []string{"kdsaf", "akdslf"}
	_, err := GetAPIKey(h)
	if err == nil {
		t.Fatal("Nonsense header should error")
	}

	h["Authorization"] = []string{"ApiKey topsecreatekeya1893841", " ths is some garbaseg"}
	key, err := GetAPIKey(h)
	if err != nil {
		t.Fatal("Should find the key")
	}
	if !reflect.DeepEqual("topsecreatekeya1893841", key) {
		t.Fatal("Key does not match")
	}
}
