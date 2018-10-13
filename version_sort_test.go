package version_test

import (
	"testing"

	"github.com/bborbe/version"

	"sort"

	. "github.com/bborbe/assert"
)

func TestSortByNameNumber(t *testing.T) {
	versions := []version.Version{"1", "3", "2"}
	sort.Sort(version.VersionByName(versions))

	if err := AssertThat(versions[0], Is(version.Version("1"))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(versions[1], Is(version.Version("2"))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(versions[2], Is(version.Version("3"))); err != nil {
		t.Fatal(err)
	}
}

func TestSortByNameLetter(t *testing.T) {
	versions := []version.Version{"a", "c", "b"}
	sort.Sort(version.VersionByName(versions))

	if err := AssertThat(versions[0], Is(version.Version("a"))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(versions[1], Is(version.Version("b"))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(versions[2], Is(version.Version("c"))); err != nil {
		t.Fatal(err)
	}
}

func TestSortByNameNumberLength(t *testing.T) {
	versions := []version.Version{"3", "11", "2"}
	sort.Sort(version.VersionByName(versions))

	if err := AssertThat(versions[0], Is(version.Version("2"))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(versions[1], Is(version.Version("3"))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(versions[2], Is(version.Version("11"))); err != nil {
		t.Fatal(err)
	}
}

func TestSortByNameNumberWithDot(t *testing.T) {
	versions := []version.Version{"1.0.1", "1.0.3", "1.0.2"}
	sort.Sort(version.VersionByName(versions))

	if err := AssertThat(versions[0], Is(version.Version("1.0.1"))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(versions[1], Is(version.Version("1.0.2"))); err != nil {
		t.Fatal(err)
	}
	if err := AssertThat(versions[2], Is(version.Version("1.0.3"))); err != nil {
		t.Fatal(err)
	}
}

func TestLessEqualReturnTrueIfLess(t *testing.T) {
	if err := AssertThat(version.LessEqual(version.Version("1"), version.Version("2")), Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestLessEqualReturnTrueIfEqual(t *testing.T) {
	if err := AssertThat(version.LessEqual(version.Version("2"), version.Version("2")), Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestLessEqualReturnFalseIfGreater(t *testing.T) {
	if err := AssertThat(version.LessEqual(version.Version("2"), version.Version("1")), Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestLessThanReturnTrueIfLess(t *testing.T) {
	if err := AssertThat(version.LessThan(version.Version("1"), version.Version("2")), Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestLessThanReturnFalseIfEqual(t *testing.T) {
	if err := AssertThat(version.LessThan(version.Version("2"), version.Version("2")), Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestLessThanReturnFalseIfGreater(t *testing.T) {
	if err := AssertThat(version.LessThan(version.Version("2"), version.Version("1")), Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestEqualReturnFalseIfLess(t *testing.T) {
	if err := AssertThat(version.Equal(version.Version("1"), version.Version("2")), Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestEqualReturnTrueIfEqual(t *testing.T) {
	if err := AssertThat(version.Equal(version.Version("2"), version.Version("2")), Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestEqualReturnFalseIfGreater(t *testing.T) {
	if err := AssertThat(version.Equal(version.Version("2"), version.Version("1")), Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestGreaterEqualReturnFalseIfLess(t *testing.T) {
	if err := AssertThat(version.GreaterEqual(version.Version("1"), version.Version("2")), Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestGreaterEqualReturnTrueIfEqual(t *testing.T) {
	if err := AssertThat(version.GreaterEqual(version.Version("2"), version.Version("2")), Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestGreaterEqualReturnTrueIfGreater(t *testing.T) {
	if err := AssertThat(version.GreaterEqual(version.Version("2"), version.Version("1")), Is(true)); err != nil {
		t.Fatal(err)
	}
}

func TestGreaterThanReturnFalseIfLess(t *testing.T) {
	if err := AssertThat(version.GreaterThan(version.Version("1"), version.Version("2")), Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestGreaterThanReturnFalseIfEqual(t *testing.T) {
	if err := AssertThat(version.GreaterThan(version.Version("2"), version.Version("2")), Is(false)); err != nil {
		t.Fatal(err)
	}
}

func TestGreaterThanReturnTrueIfGreater(t *testing.T) {
	if err := AssertThat(version.GreaterThan(version.Version("2"), version.Version("1")), Is(true)); err != nil {
		t.Fatal(err)
	}
}
