package knows

import (
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

func NewKnow(title string, tags []string, body []byte) *Know {
	k := &Know{
		UUID:   uuid.New().String(),
		Title:  title,
		Tags:   tags,
		Body:   body,
		Create: time.Now(),
		Update: time.Now(),
		ref:    &knowRef{},
	}

	k.updateRef()

	return k
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
