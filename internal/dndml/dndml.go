package dndml

import (
	"github.com/theaufish-git/dndml/pkg/dndml/enums"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

// Renderable is an object that can be rendered.
//counterfeiter:generate . Renderable
type Renderable interface {
	Render(map[string]interface{}) ([]byte, error)
}

// Object provides the base metadata about an object that the parser needs to generate objects.
//counterfeiter:generate . Object
type Object interface {
	Renderable

	GetKind() enums.Object
	GetName() string
	GetVersion() string
	GetSource() string
}
