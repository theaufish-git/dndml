package dndml

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"

	"github.com/theaufish-git/dndml/internal/ref"
	"github.com/theaufish-git/dndml/pkg/dndml"
	"github.com/theaufish-git/dndml/pkg/dndml/enums"
	yml "gopkg.in/yaml.v3"
)

var (
	// ErrParseScheme is returned when an invalid file source scheme is provided.
	ErrParseScheme = errors.New("invalid parse scheme")
)

// ObjUnmarshaller is a function that returns a reference to an object that values can be
// unmarshalled into.
type ObjUnmarshaller func() Object

// ObjUnmarshallerKey represents the metadata components that the unmarshaller uses to
// generate a version specific interpretation of an object.
type ObjUnmarshallerKey struct {
	Kind    enums.Object
	Version string
}

// Unmarshaller is a struct that parses dndml data.
type Unmarshaller struct {
	sources          []string
	objUnmarshallers map[ObjUnmarshallerKey]ObjUnmarshaller
	registry         *Registry
}

// NewUnmarshaller returns a new Parser.
func NewUnmarshaller() *Unmarshaller {
	return &Unmarshaller{
		objUnmarshallers: map[ObjUnmarshallerKey]ObjUnmarshaller{
			{Kind: enums.Object_trait}: func() Object { return &dndml.Trait{} },
		},
		registry: NewRegistry(),
	}
}

// Register registers a parser.
func (u *Unmarshaller) Register(kind enums.Object, version string, fn ObjUnmarshaller) {
	u.objUnmarshallers[ObjUnmarshallerKey{Kind: kind, Version: version}] = fn
}

// Objects returns the objects parsed by the parser.
func (u *Unmarshaller) Objects() (map[string]Object, error) {
	objs := map[string]Object{}
	for key := range u.registry.objects {
		fmt.Println("obj key:", key)
		if _, ok := u.registry.internal[key]; ok {
			continue
		}

		obj, err := u.registry.Object(key)
		if err != nil {
			return nil, err
		}
		objs[key] = obj
	}

	return objs, nil
}

// Unmarshal parses the resource specified and returns the representative objects.
func (u *Unmarshaller) Unmarshal(source ...string) error {
	initialSources := ref.SetFromSlice(source)

	u.sources = append(u.sources, source...)
	for len(u.sources) > 0 {
		sources := u.sources
		u.sources = []string{}

		for _, source := range sources {
			url, err := url.Parse(source)
			if err != nil {
				return err
			}

			if !u.registry.needsParse(url) {
				continue
			}

			objs, err := u.parseSource(url)
			if err != nil {
				return err
			}

			for _, obj := range objs {
				_, isInitial := initialSources[source]
				u.registry.register(url, obj, !isInitial)

				withSource, ok := obj.(interface{ GetSource() string })
				if !ok {
					continue
				}

				source := withSource.GetSource()
				if source != "" {
					u.sources = append(u.sources, source)
				}
			}
		}
	}
	return nil
}

func (u *Unmarshaller) parseSource(url *url.URL) (objs []Object, err error) {
	if url.Scheme == "" {
		url.Scheme = "file"
	}

	switch url.Scheme {
	case "file":
		objs, err = u.parseFile(url.Path)
	case "http", "https":
		objs, err = u.parseHTTP(url)
	default:
		err = fmt.Errorf("%w: scheme is invalid: - %s", ErrParseScheme, url.Scheme)
	}

	return
}

// parseFile parses a file and reutrns the representative objects.
func (u *Unmarshaller) parseFile(fname string) (objs []Object, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()

	return u.parseIO(f)
}

// parseHTTP parses a file sourced from http/s and reutrns the representative objects.
func (u *Unmarshaller) parseHTTP(url *url.URL) (objs []Object, err error) {
	resp, err := http.Get(url.String())
	if err != nil {
		return
	}

	return u.parseIO(resp.Body)
}

func (u *Unmarshaller) parseIO(r io.Reader) (objs []Object, err error) {
	bits, err := io.ReadAll(r)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	metad := yml.NewDecoder(bytes.NewBuffer(bits))
	objd := yml.NewDecoder(bytes.NewBuffer(bits))
	for {
		om := &dndml.Object{}
		if err = metad.Decode(om); err == io.EOF {
			err = nil
			return
		} else if err != nil {
			return
		}

		objkey := ObjUnmarshallerKey{
			Kind:    om.GetKind(),
			Version: om.GetVersion(),
		}

		objparser, ok := u.unmarshalFn(objkey)
		if !ok {
			continue
		}

		obj := objparser()
		if err = objd.Decode(obj); err != nil {
			return
		}

		objs = append(objs, obj)
	}
}

func (u *Unmarshaller) unmarshalFn(key ObjUnmarshallerKey) (fn ObjUnmarshaller, ok bool) {
	fn, ok = u.objUnmarshallers[key]
	if ok {
		return
	}

	key.Version = ""
	fn, ok = u.objUnmarshallers[key]
	return
}

/*
func (p *Parser) parseV5(om *ObjMeta, ymld *yml.Decoder) (obj Object, err error) {
	switch om.Kind {
	case "creature":
	/*
		obj := &v5.Creature{}
		ymld := yml.NewDecoder(bytes.NewBuffer(bits))

		if err := ymld.Decode(obj); err == io.EOF {
			return obj, nil
		} else if err != nil {
			return nil, err
		}
		return obj, nil
	/
	case "trait":
		obj = &v5.Trait{}
	default:
		return nil, fmt.Errorf("invalid object kind for v5: %s", om.Kind)
	}

	err = ymld.Decode(obj)
	if err == io.EOF {
		err = nil
	}

	return
}
*/
