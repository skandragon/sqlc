// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package querytest

import (
	"database/sql/driver"
	"fmt"
)

type FooDigit string

const (
	FooDigit0       FooDigit = "0"
	FooDigit1       FooDigit = "1"
	FooDigit2       FooDigit = "2"
	FooDigit3       FooDigit = "3"
	FooDigit4       FooDigit = "4"
	FooDigit5       FooDigit = "5"
	FooDigit6       FooDigit = "6"
	FooDigit7       FooDigit = "7"
	FooDigit8       FooDigit = "8"
	FooDigit9       FooDigit = "9"
	FooDigitValue10 FooDigit = "#"
	FooDigitValue11 FooDigit = "*"
)

func (e *FooDigit) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FooDigit(s)
	case string:
		*e = FooDigit(s)
	default:
		return fmt.Errorf("unsupported scan type for FooDigit: %T", src)
	}
	return nil
}

type NullFooDigit struct {
	FooDigit FooDigit
	Valid    bool // Valid is true if FooDigit is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFooDigit) Scan(value interface{}) error {
	if value == nil {
		ns.FooDigit, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FooDigit.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFooDigit) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FooDigit), nil
}

type FooFoobar string

const (
	FooFoobarFooA FooFoobar = "foo-a"
	FooFoobarFooB FooFoobar = "foo_b"
	FooFoobarFooC FooFoobar = "foo:c"
	FooFoobarFooD FooFoobar = "foo/d"
	FooFoobarFooe FooFoobar = "foo@e"
	FooFoobarFoof FooFoobar = "foo+f"
	FooFoobarFoog FooFoobar = "foo!g"
)

func (e *FooFoobar) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = FooFoobar(s)
	case string:
		*e = FooFoobar(s)
	default:
		return fmt.Errorf("unsupported scan type for FooFoobar: %T", src)
	}
	return nil
}

type NullFooFoobar struct {
	FooFoobar FooFoobar
	Valid     bool // Valid is true if FooFoobar is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullFooFoobar) Scan(value interface{}) error {
	if value == nil {
		ns.FooFoobar, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.FooFoobar.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullFooFoobar) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.FooFoobar), nil
}

type Foo struct {
	Foobar FooFoobar
	Digit  FooDigit
}
