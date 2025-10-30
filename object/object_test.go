package object

import "testing"

func TestStringHashkey(t *testing.T) {
	hello1 := &String{Value: "Hello World"}
	hello2 := &String{Value: "Hello World"}
	diff1 := &String{Value: "My name is eren"}
	diff2 := &String{Value: "My name is eren"}

	if hello1.hashKey() != hello2.hashKey() {
		t.Errorf("strings with same content but different hash")
	}

	if diff1.hashKey() != diff2.hashKey() {
		t.Errorf("strings with same content have different hash keys")
	}

	if hello1.hashKey() != diff1.hashKey() {
		t.Errorf("strings with different content have same hash keys")
	}
}