package dict

import "testing"

func TestPopulate(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.PopulateByUNIXCommand()
	if !dictionary.InDictionary("waterman") {
		t.Error("Word supposed to be in dictionary.")
	}
}
