package dndml

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"

	"github.com/jinzhu/copier"
)

const (
	sourceKeyFmt   = "%s/%s"
	urlKeyFmt      = "%s/%s#%s"
	registryKeyFmt = "%s/%s#%s:%s"
)

var (
	// ErrUnknownKind is returned when a kind is unknown.
	ErrUnknownKind = errors.New("unknown kind")
	// ErrNotAnObject is returned when a value cannot be cast as an Object.
	ErrNotAnObject = errors.New("not an object")
)

func sourceKey(u *url.URL) string {
	return ""
}

func registryKey(u *url.URL) string {
	return ""
}

// Registry is a store of all previously seen object types.
type Registry struct {
	sources  map[string]struct{}
	resolved map[string]Object
	objects  map[string]Object
	internal map[string]struct{}
}

// NewRegistry returns a new Registry.
func NewRegistry() *Registry {
	return &Registry{
		sources:  map[string]struct{}{},
		resolved: map[string]Object{},
		objects:  map[string]Object{},
		internal: map[string]struct{}{},
	}
}

func (r *Registry) Object(key string) (obj Object, err error) {
	fmt.Println("OBJ KEY", key)

	u, err := url.Parse(key)
	if err != nil {
		return nil, err
	}
	fmt.Println("URL ", u.Host, u.Path, u.Fragment, u.RawFragment)

	key = fmt.Sprintf(urlKeyFmt, u.Host, u.Path, u.Fragment)

	for k, v := range r.objects {
		fmt.Println(k, "->", v)
	}

	var ok bool
	if obj, ok = r.resolved[key]; ok {
		return
	}

	obj, err = r.resolve(key)
	return
}

func (r *Registry) resolve(key string) (obj Object, err error) {
	base, ok := r.objects[key]
	if !ok {
		err = fmt.Errorf("%w: %s", ErrUnknownKind, key)
		return
	}

	ptr := reflect.New(reflect.TypeOf(base))
	obj, ok = ptr.Elem().Interface().(Object)
	if !ok {
		err = fmt.Errorf("%w: %T", ErrNotAnObject, obj)
		return
	}

	var parent Object
	src := base.GetSource()
	if src != "" {
		parent, err = r.Object(src)
		if err != nil {
			return
		}

		err = copier.Copy(obj, parent)
		if err != nil {
			return
		}
	}

	err = copier.Copy(obj, base)
	if err != nil {
		return
	}

	r.resolved[key] = obj
	return
}

// register parses the given file and adds its objects to the Registry.
func (r *Registry) register(u *url.URL, obj Object, internal bool) {
	objkey := fmt.Sprintf(registryKeyFmt, u.Host, u.Path, obj.GetKind(), obj.GetName())
	sourceKey := fmt.Sprintf(sourceKeyFmt, u.Host, u.Path)
	r.sources[sourceKey] = struct{}{}
	r.objects[objkey] = obj
	if internal {
		r.internal[objkey] = struct{}{}
	}
}

func (r *Registry) needsParse(u *url.URL) (ok bool) {
	source := fmt.Sprintf(sourceKeyFmt, u.Host, u.Path)
	_, ok = r.sources[source]
	return !ok
}
