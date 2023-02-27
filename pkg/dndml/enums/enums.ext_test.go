package enums_test

import (
	"errors"
	"testing"

	"github.com/theaufish-git/dndml/pkg/dndml/enums"
)

func TestClass_MarshalText(t *testing.T) {
	testcases := []struct {
		desc   string
		input  enums.Class
		expect []byte
		err    error
	}{
		{
			desc:   "valid enum",
			input:  enums.Class_barbarian,
			expect: []byte("barbarian"),
		},

		{
			desc:  "invalid enum",
			input: enums.Class(-1),
			err:   enums.ErrIvalidEnumInt,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			actual, err := tc.input.MarshalText()
			if err != nil {
				if !errors.Is(err, tc.err) {
					t.Error("an error of the wrong type was returned:", err, tc.err)
					return
				}
				return
			}

			if tc.err != nil {
				t.Error("an error was not returned when one was expected")
				return
			}

			if string(actual) != string(tc.expect) {
				t.Error("unmarshaled enum not as expected:", string(actual), string(tc.expect))
			}
		})
	}
}

func TestClass_UnmarshalText(t *testing.T) {
	testcases := []struct {
		desc   string
		input  []byte
		expect enums.Class
		err    error
	}{
		{
			desc:   "valid enum",
			input:  []byte("barbarian"),
			expect: enums.Class_barbarian,
		},

		{
			desc:  "invalid enum",
			input: []byte("bardbarbarian"),
			err:   enums.ErrIvalidEnumString,
		},
		{
			desc:   "empty string is undefined",
			input:  []byte(""),
			expect: enums.Class_undefined_class,
		},
		{
			desc:   "nil string is undefined",
			input:  nil,
			expect: enums.Class_undefined_class,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			var actual enums.Class

			if err := actual.UnmarshalText(tc.input); err != nil {
				if !errors.Is(err, tc.err) {
					t.Error("an error of the wrong type was returned:", err, tc.err)
					return
				}
				return
			}

			if tc.err != nil {
				t.Error("an error was not returned when one was expected")
				return
			}

			if actual != tc.expect {
				t.Error("unmarshaled enum not as expected:", actual, tc.expect)
			}
		})
	}
}

func TestPeriod_MarshalText(t *testing.T) {
	testcases := []struct {
		desc   string
		input  enums.Period
		expect []byte
		err    error
	}{
		{
			desc:   "valid enum",
			input:  enums.Period_day,
			expect: []byte("day"),
		},

		{
			desc:  "invalid enum",
			input: enums.Period(-1),
			err:   enums.ErrIvalidEnumInt,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			actual, err := tc.input.MarshalText()
			if err != nil {
				if !errors.Is(err, tc.err) {
					t.Error("an error of the wrong type was returned:", err, tc.err)
					return
				}
				return
			}

			if tc.err != nil {
				t.Error("an error was not returned when one was expected")
				return
			}

			if string(actual) != string(tc.expect) {
				t.Error("unmarshaled enum not as expected:", string(actual), string(tc.expect))
			}
		})
	}
}

func TestPeriod_UnmarshalText(t *testing.T) {
	testcases := []struct {
		desc   string
		input  []byte
		expect enums.Period
		err    error
	}{
		{
			desc:   "valid enum",
			input:  []byte("day"),
			expect: enums.Period_day,
		},

		{
			desc:  "invalid enum",
			input: []byte("not-a-day"),
			err:   enums.ErrIvalidEnumString,
		},
		{
			desc:   "empty string is undefined",
			input:  []byte(""),
			expect: enums.Period_undefined_period,
		},
		{
			desc:   "nil string is undefined",
			input:  nil,
			expect: enums.Period_undefined_period,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			var actual enums.Period

			if err := actual.UnmarshalText(tc.input); err != nil {
				if !errors.Is(err, tc.err) {
					t.Error("an error of the wrong type was returned:", err, tc.err)
					return
				}
				return
			}

			if tc.err != nil {
				t.Error("an error was not returned when one was expected")
				return
			}

			if actual != tc.expect {
				t.Error("unmarshaled enum not as expected:", actual, tc.expect)
			}
		})
	}
}
