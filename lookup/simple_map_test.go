package lookup

import (
	"testing"
)

func TestMapLookup(t *testing.T) {
	mapLookup := &MapLookup{
		Vars: map[string]string{"VAR": "VALUE"},
	}

	envs := mapLookup.Lookup("VAR", nil)
	if len(envs) != 1 {
		t.Fatalf("Expected envs to contains one element, but was %v", envs)
	}

	envs = mapLookup.Lookup("var", nil)
	if len(envs) != 0 {
		t.Fatalf("Expected envs to be empty, but was %v", envs)
	}

	envs = mapLookup.Lookup("DOES_NOT_EXIST", nil)
	if len(envs) != 0 {
		t.Fatalf("Expected envs to be empty, but was %v", envs)
	}
}
