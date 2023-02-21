package dndml

import "fmt"

const (
	registryKeyFmt = "%s#%s:%s"
)

// Registry is a store of all previously seen object types.
type Registry struct {
	sources map[string]struct{}
	objects map[string]interface{}
}

// NewRegistry returns a new Registry.
func NewRegistry() *Registry {
	return &Registry{
		sources: map[string]struct{}{},
		objects: map[string]interface{}{},
	}
}

// register parses the given file and adds its objects to the Registry.
func (r *Registry) register(source string, obj Object) {
	key := fmt.Sprintf(registryKeyFmt, source, obj.GetKind(), obj.GetName())
	r.sources[source] = struct{}{}
	r.objects[key] = obj
}

func (r *Registry) needsParse(source string) bool {
	_, ok := r.sources[source]
	return !ok
}
