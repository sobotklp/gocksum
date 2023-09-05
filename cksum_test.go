package gocksum

import (
	"os"
	"testing" //import go package for testing related functionality
)

func expectCksum(filename string, expSz uint, expCk uint32, t *testing.T) {
	f, _ := os.Open(filename)
	defer f.Close()
	if f == nil {
		t.Errorf("Unable to open fixture %s!", filename)
		return
	}
	ck, sz, err := Cksum(f)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if sz != expSz {
		t.Errorf("Expected %d to equal %d", sz, expSz)
	}
	if ck != expCk {
		t.Errorf("Expected %d to equal %d", ck, expCk)
	}
}

func Test_Cksum(t *testing.T) {
	expectCksum("./test/fixtures/empty", 0, 4294967295, t)
	expectCksum("./test/fixtures/CoffeeScript.png", 3280, 3650423881, t)
}
