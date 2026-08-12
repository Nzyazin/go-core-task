package main

import (
	"testing"
)

func TestNewVariables(t *testing.T) {
	v := NewVariables()
	if v.NumDecimal != 42 {
		t.Errorf("NumDecimal = %d, want 42", v.NumDecimal)
	}
	if v.NumOctal != 052 {
		t.Errorf("NumOctal = %d, want 052 (42)", v.NumOctal)
	}
	if v.NumHexadecimal != 0x2A {
		t.Errorf("NumHexadecimal = %d, want 0x2A (42)", v.NumHexadecimal)
	}
	if v.Pi != 3.14 {
		t.Errorf("Pi = %v, want 3.14", v.Pi)
	}
	if v.Name != "Golang" {
		t.Errorf("Name = %q, want \"Golang\"", v.Name)
	}
	if !v.IsActive {
		t.Errorf("IsActive = false, want true")
	}
	if v.ComplexNum != 1+2i {
		t.Errorf("ComplexNum = %v, want (1+2i)", v.ComplexNum)
	}
}

func TestTypeOf(t *testing.T) {
	tests := []struct {
		value interface{}
		want  string
	}{
		{42, "int"},
		{3.14, "float64"},
		{"Golang", "string"},
		{true, "bool"},
		{complex64(1 + 2i), "complex64"},
	}
	for _, tt := range tests {
		got := TypeOf(tt.value)
		if got != tt.want {
			t.Errorf("TypeOf(%v) = %s, want %s", tt.value, got, tt.want)
		}
	}
}

func TestConvertToString(t *testing.T) {
	v := NewVariables()
	got := ConvertToString(v)
	want := "4242423.14Golangtrue(1+2i)"
	if got != want {
		t.Errorf("ConvertToString() = %q, want %q", got, want)
	}
}

func TestStringToRunes(t *testing.T) {
	s := "Go"
	runes := StringsToRunes(s)
	if len(runes) != 2 {
		t.Fatalf("len(runes) = %d, want 2", len(runes))
	}
	if runes[0] != 'G' || runes[1] != 'o' {
		t.Errorf("runes = %v, want [G o]", runes)
	}
}

func TestHashWithSalt(t *testing.T) {
	runes := []rune("hello")
	hash := HashWithSalt(runes, "go-2024")
	if hash == "" {
		t.Error("HashWithSalt returned empty string")
	}
	// SHA256 hex digest is always 64 chars
	if len(hash) != 64 {
		t.Errorf("len(hash) = %d, want 64", len(hash))
	}
}

func TestHashWithSaltDeterministic(t *testing.T) {
	runes := []rune("test")
	hash1 := HashWithSalt(runes, "go-2024")
	hash2 := HashWithSalt(runes, "go-2024")
	if hash1 != hash2 {
		t.Error("HashWithSalt is not deterministic for same input")
	}
}

func TestHashWithSaltDifferentFromNoSalt(t *testing.T) {
	runes := []rune("test")
	withSalt := HashWithSalt(runes, "go-2024")
	emptySalt := HashWithSalt(runes, "")
	if withSalt == emptySalt {
		t.Error("Hash with salt should differ from hash without salt")
	}
}
