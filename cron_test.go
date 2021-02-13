package thenovadiary

import (
	"testing"
	"time"
)

func TestHappyTimeDeltaDays(t *testing.T) {
	t2020, err := time.Parse("2006-01-02", "2020-01-02")
	if err != nil {
		t.Errorf("error testing time: %v", err)
	}
	t2021, err := time.Parse("2006-01-02", "2021-01-02")
	if err != nil {
		t.Errorf("error testing time: %v", err)
	}
	n1 := TimeDeltaDays(t2020, t2021)
	n2 := TimeDeltaDays(t2021, t2020)
	if n1 != n2 {
		t.Errorf("unequal delta days %d != %d", n1, n2)
		t.FailNow()
	}
	if n1 != 366 {
		t.Errorf("n1 != 366")
	}
	if n2 != 366 {
		t.Errorf("n2 != 366")
	}
}
