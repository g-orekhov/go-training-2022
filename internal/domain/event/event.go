package event

import "github.com/upper/db/v4"

type Event struct {
	Id               int64    `db:"id,omitempty" json:"id"`
	Name             string   `db:"name" json:"name"`
	Description      string   `db:"descr,omitempty" json:"description"`
	Images           []string `db:"images,omitempty" json:"images"`
	ShortDescription string   `db:"short_descr,omitempty" json:"short_descr"`
	PreviewImage     string   `db:"preview_img,omitempty" json:"preview_img"`
	Coords           Coords   `db:"coords,omitempty" json:"coords"`
}

func (*Event) Store(sess db.Session) db.Store {
	return sess.Collection("events")
}

type EventPreview struct {
	Id               int64  `db:"id" json:"id"`
	Name             string `db:"name" json:"name"`
	ShortDescription string `db:"short_descr" json:"description"`
	PreviewImage     string `db:"preview_img" json:"image"`
}
