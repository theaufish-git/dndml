package dndml

import (
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/theaufish-git/dndml/pkg/dndml/enums"
)

// Identifier is the structure used to represent an identifier for a specific
// object within dndml.
type Identifier struct {
	Host   string
	Source string
	Kind   enums.Object
	Name   string
}

var (
	// ErrIdentifier is the base error that all identifier errors report.
	ErrIdentifier = errors.New("identifier error")
	// ErrIdentifierParse is an error returned when an identifier cannot be parsed.
	ErrIdentifierParse = errors.New("identifier parse error")
)

// ParseIdentifier returns a Key from the given string.
func ParseIdentifier(key string) (Identifier, error) {
	u, err := url.Parse(key)
	if err != nil {
		return Identifier{}, err
	}

	if u.Fragment == "" {
		return Identifier{}, fmt.Errorf("%w: no fragment", ErrIdentifierParse)
	}

	tokens := strings.Split(u.Fragment, ":")
	if len(tokens) != 2 {
		return Identifier{}, fmt.Errorf("%w: invalid fragment", ErrIdentifierParse)
	}

	kind, ok := enums.Object_value[tokens[0]]
	if !ok {
		return Identifier{}, fmt.Errorf("%w: invalid fragment kind - %s", ErrIdentifierParse, tokens[0])
	}

	return Identifier{
		Host:   u.Host,
		Source: u.Path,
		Kind:   enums.Object(kind),
		Name:   tokens[1],
	}, nil
}
