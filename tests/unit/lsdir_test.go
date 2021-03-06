package test

import (
	"fleetsearch/dirman"
	"testing"
)

func TestLsAllGood(t *testing.T) {
	ls := dirman.ListDir{}

	objs, err := ls.LsAll("./")
	if err != nil {
		t.Errorf("Expected Error: %v, got: %v", nil, err.Error())
	}
	if len(objs) < 1 {
		t.Errorf("Expected Length: +1, got: %v", len(objs))
	}

	exists := false
	for _, obj := range objs {
		if obj.Name() == "lsdir_test.go" {
			exists = true
		}
	}
	if !exists {
		t.Error("File lsdir_test.go does not exists in the returned slice.")
	}
}

func TestLsAllBad(t *testing.T) {
	ls := dirman.ListDir{}

	_, err := ls.LsAll("./fake_substring")
	if err == nil {
		t.Errorf("Expected Error got %v", err.Error())
	}
}

func TestLsContainsGood(t *testing.T) {
	ls := dirman.ListDir{}

	objs, err := ls.LsContains("./", "lsdir")
	if err != nil {
		t.Errorf("Expected Error: %v, got: %v", nil, err.Error())
	}
	if len(objs) != 1 {
		t.Errorf("Expected Length: 1, got: %v", len(objs))
	}

	if objs[0].Name() != "lsdir_test.go" {
		t.Error("File lsdir_test.go does not exists in the returned slice.")
	}
}

func TestLsContainsBad(t *testing.T) {
	ls := dirman.ListDir{}

	objs, err := ls.LsContains("./", "fake_substring")
	if err != nil {
		t.Errorf("Expected Error: %v, got: %v", nil, err.Error())
	}
	if len(objs) != 0 {
		t.Errorf("Expected Length: 0, got: %v", len(objs))
	}

	_, err = ls.LsContains("./fake_substring", "lsdir")
	if err == nil {
		t.Errorf("Expected Error got %v", err.Error())
	}
}

func TestLsStartsWithGood(t *testing.T) {
	ls := dirman.ListDir{}

	objs, err := ls.LsStartsWith("./", "lsdir")
	if err != nil {
		t.Errorf("Expected Error: %v, got: %v", nil, err.Error())
	}
	if len(objs) != 1 {
		t.Errorf("Expected Length: 1, got: %v", len(objs))
	}

	if objs[0].Name() != "lsdir_test.go" {
		t.Error("File lsdir_test.go does not exists in the returned slice.")
	}
}

func TestLsStartsWithBad(t *testing.T) {
	ls := dirman.ListDir{}

	objs, err := ls.LsStartsWith("./", ".go")
	if err != nil {
		t.Errorf("Expected Error: %v, got: %v", nil, err.Error())
	}
	if len(objs) != 0 {
		t.Errorf("Expected Length: 0, got: %v", len(objs))
	}

	_, err = ls.LsStartsWith("./fake_substring", "lsdir")
	if err == nil {
		t.Errorf("Expected Error got %v", err.Error())
	}
}

func TestLsEndsWithGood(t *testing.T) {
	ls := dirman.ListDir{}

	objs, err := ls.LsEndsWith("./", ".go")
	if err != nil {
		t.Errorf("Expected Error: %v, got: %v", nil, err.Error())
	}
	if len(objs) != 1 {
		t.Errorf("Expected Length: 1, got: %v", len(objs))
	}

	if objs[0].Name() != "lsdir_test.go" {
		t.Errorf("Expected: lsdir_test.go, got: %v", objs[0].Name())
	}
}

func TestLsEndsWithBad(t *testing.T) {
	ls := dirman.ListDir{}

	objs, err := ls.LsEndsWith("./", "lsdir")
	if err != nil {
		t.Errorf("Expected Error: %v, got: %v", nil, err.Error())
	}
	if len(objs) != 0 {
		t.Errorf("Expected Length: 0, got: %v", len(objs))
	}

	_, err = ls.LsStartsWith("./fake_substring", ".go")
	if err == nil {
		t.Errorf("Expected Error got %v", err.Error())
	}
}
