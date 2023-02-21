package enums

import (
	"errors"
	"fmt"
)

var (
	// ErrIvalidEnumInt is returned when an int value given is not valid for an enum.
	ErrIvalidEnumInt = errors.New("invalid enum int")
	// ErrIvalidEnumString is returned when a string value given is not a valid enum.
	ErrIvalidEnumString = errors.New("invalid enum string")
)

// MarshalText attempts to marshal the given Class into a string representation.
func (x *Class) MarshalText() (text []byte, err error) {
	class, ok := Class_name[int32(*x)]
	if !ok {
		return nil, fmt.Errorf("%w: %d is not a Class", ErrIvalidEnumInt, *x)
	}

	return []byte(class), nil
}

// UnmarshalText attempts to unmarshal the text given into the Class
func (x *Class) UnmarshalText(text []byte) error {
	var class Class
	if len(text) == 0 {
		class = Class_undefined_class
	} else if c, ok := Class_value[string(text)]; ok {
		class = Class(c)
	} else {
		return fmt.Errorf("%w: %s is not a Class", ErrIvalidEnumString, text)
	}

	*x = class
	return nil
}

// MarshalText attempts to marshal the given Period into a string representation.
func (x *Period) MarshalText() (text []byte, err error) {
	Period, ok := Period_name[int32(*x)]
	if !ok {
		return nil, fmt.Errorf("%w: %d is not a Period", ErrIvalidEnumInt, *x)
	}

	return []byte(Period), nil
}

// UnmarshalText attempts to unmarshal the text given into the Period
func (x *Period) UnmarshalText(text []byte) error {
	var period Period
	if len(text) == 0 {
		period = Period_undefined_period
	} else if p, ok := Period_value[string(text)]; ok {
		period = Period(p)
	} else {
		return fmt.Errorf("%w: %s is not a Period", ErrIvalidEnumString, text)
	}

	*x = period
	return nil
}
