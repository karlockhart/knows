package knows

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Know struct {
	UUID   string    `json:"uuid"`
	Title  string    `json:"time"`
	Author string    `json:"author"`
	Tags   []string  `json:"tags"`
	Body   []byte    `json:"body"`
	Create time.Time `json:"created_dts"`
	Update time.Time `json:"updated_dts"`
	ref    *knowRef
}

func EmptyKnow() *Know {
	k := &Know{
		ref: &knowRef{},
	}

	k.updateRef()

	return k
}

func NewKnow(title string, tags []string, body []byte) *Know {
	k := EmptyKnow()
	k.UUID = uuid.New().String()
	k.Title = title
	k.Tags = tags
	k.Body = body
	k.Create = time.Now()
	k.Update = time.Now()

	return k
}

func KnowFromData(data []byte) (*Know, error) {
	k := EmptyKnow()

	err := json.Unmarshal(data, &k)
	if err != nil {
		return nil, err
	}

	return k, nil
}

func (k *Know) update(u Know) {
	k.Title = u.Title
	k.Body = u.Body
	k.Update = time.Now()
}

func (k *Know) updateRef() {
	k.ref.know = k
	k.ref.deleted = false
}

func (k *Know) delete() {
	k.ref.know = nil
	k.ref.deleted = true
}

func (k *Know) String() string {
	return fmt.Sprintf("%s\n%v\n%s", k.Title, k.Tags, string(k.Body))
}
