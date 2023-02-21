package dndml

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"

	v5 "github.com/theaufish-git/dndml/pkg/dndml/v5"
	yml "gopkg.in/yaml.v3"
)

var (
	// ErrParseScheme is returned when an invalid file source scheme is provided.
	ErrParseScheme = errors.New("invalid parse scheme")
)

type Object interface {
	GetKind() string
	GetName() string
}

type ObjMeta struct {
	Version string
	Kind    string
}

// Parser is a struct that parses dndml data.
type Parser struct {
	sources  []string
	registry *Registry
}

// NewParser returns a new Parser.
func NewParser() *Parser {
	return &Parser{
		registry: NewRegistry(),
	}
}

// Objects returns the objects parsed by the parser.
func (p *Parser) Objects() map[string]interface{} {
	return p.registry.objects
}

// Parse parses the resource specified and returns the representative objects.
func (p *Parser) Parse(source ...string) error {
	p.sources = append(p.sources, source...)
	for len(p.sources) > 0 {
		sources := p.sources
		p.sources = []string{}

		for _, source := range sources {
			log.Println("source:", source)
			u, err := url.Parse(source)
			if err != nil {
				return err
			}

			if !p.registry.needsParse(u.RequestURI()) {
				log.Println("skipping...")
				continue
			}

			objs, err := p.parseSource(u)
			if err != nil {
				return err
			}

			for _, obj := range objs {
				p.registry.register(u.RequestURI(), obj)

				withSource, ok := obj.(interface{ GetSource() string })
				if !ok {
					continue
				}

				source := withSource.GetSource()
				if source != "" {
					p.sources = append(p.sources, source)
				}
			}
		}
	}
	return nil
}

func (p *Parser) parseSource(u *url.URL) (objs []Object, err error) {
	if u.Scheme == "" {
		u.Scheme = "file"
	}

	fmt.Println(u)

	switch u.Scheme {
	case "file":
		objs, err = p.parseFile(u.Path)
	case "http", "https":
		objs, err = p.parseHTTP(u.String())
	default:
		err = fmt.Errorf("%w: scheme is invalid: - %s", ErrParseScheme, u.Scheme)
	}

	return
}

// parseFile parses a file and reutrns the representative objects.
func (p *Parser) parseFile(fname string) (objs []Object, err error) {
	f, err := os.Open(fname)
	if err != nil {
		return
	}
	defer f.Close()

	return p.parseIO(f)
}

// parseHTTP parses a file sourced from http/s and reutrns the representative objects.
func (p *Parser) parseHTTP(u string) (objs []Object, err error) {
	resp, err := http.Get(u)
	if err != nil {
		return
	}

	return p.parseIO(resp.Body)
}

func (p *Parser) parseIO(r io.Reader) (objs []Object, err error) {
	bits, err := io.ReadAll(r)
	if err != nil && !errors.Is(err, io.EOF) {
		return nil, err
	}

	metad := yml.NewDecoder(bytes.NewBuffer(bits))
	objd := yml.NewDecoder(bytes.NewBuffer(bits))
	for {
		om := &ObjMeta{}
		if err = metad.Decode(om); err == io.EOF {
			err = nil
			return
		} else if err != nil {
			return
		}

		var obj Object
		switch om.Version {
		case "v5":
			obj, err = p.parseV5(om, objd)
			if err != nil {
				return nil, err
			}
		}

		objs = append(objs, obj)
	}
}

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
	*/
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
