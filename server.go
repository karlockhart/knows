package knows

import (
	"fmt"
	"strings"
)

type knowRef struct {
	know    *Know
	deleted bool
	refs    []*refContainer
}

type refContainer struct {
	refs map[string]*knowRef
}

type tagIndex struct {
	idx map[string]*refContainer
}

func (t *tagIndex) index(uuid string, tags []string, ref *knowRef) {
	ref.refs = make([]*refContainer, 0)
	for _, tag := range tags {
		if _, ok := t.idx[tag]; !ok {
			rc := refContainer{}
			rc.refs = make(map[string]*knowRef)
			t.idx[tag] = &rc

		}
		t.idx[tag].refs[uuid] = ref
		rco := t.idx[tag]
		ref.refs = append(ref.refs, rco)
	}
}

func (t *tagIndex) remove(uuid string, ref *knowRef) {
	for _, r := range ref.refs {
		delete(r.refs, uuid)
	}
}

type Server struct {
	cache map[string]*Know
	index tagIndex
}

func NewServer() *Server {
	s := Server{}
	s.cache = make(map[string]*Know)
	s.index = tagIndex{}
	s.index.idx = map[string]*refContainer{}

	return &s
}

func (s *Server) Create(k Know) (string, error) {
	if k.UUID == "" {
		return "", fmt.Errorf("invalid UUID")
	}

	if _, ok := s.cache[k.UUID]; ok {
		return "", fmt.Errorf("%s already exists", k.UUID)
	}

	s.cache[k.UUID] = &k
	s.index.index(k.UUID, k.Tags, k.ref)

	return k.UUID, nil
}

func (s *Server) Read(uuid string) *Know {
	if v, ok := s.cache[uuid]; ok {
		return v
	}
	return nil
}

func (s *Server) Update(uuid string, k Know) error {
	if v, ok := s.cache[uuid]; ok {
		v.update(k)
		s.index.index(uuid, k.Tags, v.ref)

		return nil
	}

	return fmt.Errorf("not found")
}

func (s *Server) FindByTag(tag string) []Know {
	k := make([]Know, 0)

	if v, ok := s.index.idx[tag]; ok {
		fmt.Println(v)
		for _, ref := range v.refs {
			if !ref.deleted {
				k = append(k, *ref.know)
			}
		}
	}

	return k
}

func (s *Server) Delete(uuid string) error {
	if _, ok := s.cache[uuid]; !ok {
		return fmt.Errorf("not found")
	}

	k := s.cache[uuid]
	k.delete()
	s.index.remove(uuid, k.ref)
	delete(s.cache, uuid)

	return nil
}

func (s *Server) Dump() string {
	b := strings.Builder{}
	for _, item := range s.cache {
		b.WriteString(fmt.Sprintf("%s\n", item.String()))
	}

	return b.String()
}