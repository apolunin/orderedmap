package orderedmap

import (
	"testing"
)

func TestOrderedMap(t *testing.T) {
	const length = 8
	var (
		expectedKeys = [length]string{"d", "b", "c", "a", "aa", "bb", "cc", "dd"}
		expectedVals = [length]int{3, 8, 5, 9, 33, 88, 99, 55}
	)

	om := New[string, int]()

	if val, ok := om.Get("empty"); ok || val != 0 {
		t.Fatalf("it should not be possible to get anything from empty list")
	}

	if val, ok := om.Delete("empty"); ok || val != 0 {
		t.Fatalf("it should not be possible to delete anything from empty list")
	}

	for i := 0; i < length; i++ {
		om.Set(expectedKeys[i], expectedVals[i])
	}

	if om.Len() != length {
		t.Fatalf("wanted: %d, got: %d", length, om.Len())
	}

	if val, ok := om.Get("missing"); ok || val != 0 {
		t.Fatalf("element with key %q should not exist", "missing")
	}

	if val, ok := om.Delete("missing"); ok || val != 0 {
		t.Fatalf("element with key %q doesn't exist, it cannot be deleted", "missing")
	}

	actualKeys := [length]string{}
	actualVals := [length]int{}

	{
		i := 0
		next := om.Iterator()
		for k, v, ok := next(); ok; k, v, ok = next() {
			actualKeys[i] = k
			actualVals[i] = v
			i++
		}

		if actualKeys != expectedKeys {
			t.Fatalf("wanted: %q, got: %q", expectedKeys, actualKeys)
		}

		if actualVals != expectedVals {
			t.Fatalf("wanted: %q, got: %q", expectedVals, actualVals)
		}
	}

	if val, ok := om.Get(expectedKeys[0]); !ok || val != expectedVals[0] {
		t.Fatalf("get value, wanted: %q, got: %q", expectedVals[0], val)
	}

	if val, ok := om.Delete(expectedKeys[0]); !ok || val != expectedVals[0] {
		t.Fatalf("delete value, wanted: %q, got: %q", expectedVals[0], val)
	}

	if val, ok := om.Get(expectedKeys[0]); ok || val != 0 {
		t.Fatalf("value with key %q was not deleted as expected", expectedKeys[0])
	}
}
