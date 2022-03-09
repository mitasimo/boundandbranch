package matrix

import (
	"testing"
)

func TestMatrix_checkIndex(t *testing.T) {
	var (
		x = 5
		y = 5
	)
	m := New(x, y)
	if err := m.checkIndex(2, 2); err != nil {
		t.Errorf("checkIndex(2, 2) = %v, want nil", err)
	}
	if err := m.checkIndex(5, 2); err != errorColumnIndexOutOfBound {
		t.Errorf("checkIndex(2, 2) = %v, want %v", err, errorColumnIndexOutOfBound)
	}
	if err := m.checkIndex(2, 5); err != errorRowIndexOutOfBound {
		t.Errorf("checkIndex(2, 2) = %v, want %v", err, errorRowIndexOutOfBound)
	}
}

func TestNew(t *testing.T) {
	var (
		x = 5
		y = 5
	)
	m := New(x, y)
	if m.x != x {
		t.Errorf("m.x = %d, want %d", m.x, x)
	}
	if m.y != y {
		t.Errorf("m.y = %v, want %d", m.y, y)
	}
	if len(m.data) != x*y {
		t.Errorf("len(m.data) = %d, want %d", len(m.data), x*y)
	}
}

func TestMatrix_At(t *testing.T) {
	want := 10
	m := New(3, 3)
	m.Set(0, 0, want)

	if v, _ := m.At(0, 0); v != want {
		t.Errorf("m[0,0] = %d, want %d", v, want)
	}
}

func TestMatrix_Set(t *testing.T) {
	want := 10
	m := New(3, 3)
	m.Set(0, 0, want)
	if m.data[0] != want {
		t.Errorf("m[0,0] = %v, want %v", m.data[0], want)
	}
}
