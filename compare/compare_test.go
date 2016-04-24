package compare

import (
	"testing"

	"bytes"

	. "github.com/bborbe/assert"
	"github.com/bborbe/version"
)

func TestNoArg(t *testing.T) {
	writer := bytes.NewBufferString("")
	returnCode, err := do(writer, version.LessThan, []string{})
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(returnCode, Is(1))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("can only compare to versions\n"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestOneArg(t *testing.T) {
	writer := bytes.NewBufferString("")
	returnCode, err := do(writer, version.LessThan, []string{"1.2.3"})
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(returnCode, Is(1))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is("can only compare to versions\n"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestTwoArgLess(t *testing.T) {
	writer := bytes.NewBufferString("")
	returnCode, err := do(writer, version.LessThan, []string{"1.2.2", "1.2.3"})
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(returnCode, Is(0))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is(""))
	if err != nil {
		t.Fatal(err)
	}
}

func TestTwoArgNotLess(t *testing.T) {
	writer := bytes.NewBufferString("")
	returnCode, err := do(writer, version.LessThan, []string{"1.2.4", "1.2.3"})
	err = AssertThat(err, NilValue())
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(returnCode, Is(1))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(writer.String(), Is(""))
	if err != nil {
		t.Fatal(err)
	}
}
