package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

type Variables struct {
	NumDecimal     int
	NumOctal       int
	NumHexadecimal int
	Pi             float64
	Name           string
	IsActive       bool
	ComplexNum     complex64
}

func NewVariables() Variables {
	return Variables{
		NumDecimal:     42,
		NumOctal:       052,
		NumHexadecimal: 0x2A,
		Pi:             3.14,
		Name:           "Golang",
		IsActive:       true,
		ComplexNum:     1 + 2i,
	}
}

func TypeOf(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func PrintTypes(v Variables) {
	fmt.Printf("numDecimal      type: %s\n", TypeOf(v.NumDecimal))
	fmt.Printf("numOctal        type: %s\n", TypeOf(v.NumOctal))
	fmt.Printf("numHexadecimal  type: %s\n", TypeOf(v.NumHexadecimal))
	fmt.Printf("pi              type: %s\n", TypeOf(v.Pi))
	fmt.Printf("name            type: %s\n", TypeOf(v.Name))
	fmt.Printf("isActive        type: %s\n", TypeOf(v.IsActive))
	fmt.Printf("complexNum      type: %s\n", TypeOf(v.ComplexNum))
}

func ConvertToString(v Variables) string {
	var sb strings.Builder
	sb.Grow(7)
	sb.WriteString(strconv.Itoa(v.NumDecimal))
	sb.WriteString(strconv.Itoa(v.NumOctal))
	sb.WriteString(strconv.Itoa(v.NumHexadecimal))
	sb.WriteString(strconv.FormatFloat(v.Pi, 'f', -1, 64))
	sb.WriteString(v.Name)
	sb.WriteString(strconv.FormatBool(v.IsActive))
	sb.WriteString(strconv.FormatComplex(complex128(v.ComplexNum), 'f', -1, 128))
	return sb.String()
}
func StringsToRunes(s string) []rune {
	return []rune(s)
}

func HashWithSalt(runes []rune, salt string) string {
	mid := len(runes) / 2
	saltRunes := []rune(salt)
	combined := make([]rune, 0, len(runes)+len(saltRunes))
	fmt.Println(len(combined), len(runes), len(saltRunes))
	combined = append(combined, runes[:mid]...)
	combined = append(combined, saltRunes...)
	combined = append(combined, runes[mid:]...)

	sum := sha256.Sum256([]byte(string(combined)))
	return fmt.Sprintf("%x", sum)
}

func main() {
	v := NewVariables()

	fmt.Println("=== Типы переменных ===")
	PrintTypes(v)

	fmt.Println("\n=== Объединённая строка ===")
	combined := ConvertToString(v)
	fmt.Println(combined)

	fmt.Println("\n=== Срез рун ===")
	runes := StringsToRunes(combined)
	fmt.Printf("%v\n", runes)

	fmt.Println("\n=== SHA256 с солью go-2024 ===")
	hash := HashWithSalt(runes, "go-2024")
	fmt.Println(hash)
}
