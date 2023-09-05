package gocksum

import (
	"bufio"
	"os"
	"strings"
	"testing" //import go package for testing related functionality
)

func expectCksum(reader *bufio.Reader, expSz uint, expCk uint32, t *testing.T) {

	ck, sz, err := Cksum(reader)
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

func expectCksumFile(filename string, expSz uint, expCk uint32, t *testing.T) {
	f, _ := os.Open(filename)
	defer f.Close()
	if f == nil {
		t.Errorf("Unable to open fixture %s!", filename)
		return
	}
	reader := bufio.NewReader(f)

	expectCksum(reader, expSz, expCk, t)
}

func expectCksumString(str string, expSz uint, expCk uint32, t *testing.T) {
	f := strings.NewReader(str)
	reader := bufio.NewReader(f)

	expectCksum(reader, expSz, expCk, t)
}

func Test_Cksum(t *testing.T) {
	expectCksumFile("./test/fixtures/empty", 0, 4294967295, t)
	expectCksumFile("./test/fixtures/CoffeeScript.png", 3280, 3650423881, t)
	expectCksumString("", 0, 4294967295, t)
	expectCksumString("Hello world\n", 12, 3083891038, t)
}
