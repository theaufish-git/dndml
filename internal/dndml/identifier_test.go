package dndml_test

import (
	"errors"
	"testing"

	"github.com/go-test/deep"
	"github.com/theaufish-git/dndml/internal/dndml"
	"github.com/theaufish-git/dndml/pkg/dndml/enums"
)

func TestParseIdentifier(t *testing.T) {
	testcases := []struct {
		desc        string
		input       string
		expect      dndml.Identifier
		expectError error
	}{
		{
			desc:  "from url",
			input: "https://www.host.com/path/to/file.yaml#creature:aspect of tiamat",
			expect: dndml.Identifier{
				Host:   "www.host.com",
				Source: "/path/to/file.yaml",
				Kind:   enums.Object_creature,
				Name:   "aspect of tiamat",
			},
		},
		{
			desc:  "from file",
			input: "./path/to/file.yaml#creature:aspect of tiamat",
			expect: dndml.Identifier{
				Host:   "",
				Source: "/path/to/file.yaml",
				Kind:   enums.Object_creature,
				Name:   "aspect of tiamat",
			},
		},
		{
			desc:        "error - no fragment",
			input:       "/path/to/file.yaml",
			expect:      dndml.Identifier{},
			expectError: dndml.ErrIdentifierParse,
		},
		{
			desc:        "error - invalid fragment - too few tokens",
			input:       "/path/to/file.yaml#invalid",
			expect:      dndml.Identifier{},
			expectError: dndml.ErrIdentifierParse,
		},
		{
			desc:        "error - invalid fragment - too many tokens",
			input:       "/path/to/file.yaml#too:many:tokens",
			expect:      dndml.Identifier{},
			expectError: dndml.ErrIdentifierParse,
		},
		{
			desc:        "error - invalid fragment - invalid kind",
			input:       "/path/to/file.yaml#table:aspect of tiamat",
			expect:      dndml.Identifier{},
			expectError: dndml.ErrIdentifierParse,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			actual, err := dndml.ParseIdentifier(tc.input)

			if !errors.Is(err, tc.expectError) {
				t.Error("returned errors is not as expected:", err, tc.expectError)
			}

			if diff := deep.Equal(actual, tc.expect); diff != nil {
				t.Error("returned identifier not as expected")
				for _, d := range diff {
					t.Log(d)
				}
			}
		})
	}
}
