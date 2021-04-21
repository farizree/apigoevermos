package Mevermos

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ISODate struct {
	time.Time
}

func (t *ISODate) String() string {
	return t.Format("2006-01-02")
}

type (
	evermos struct {
		ID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		Title  string             `json:"title,omitempty" bson:"title,omitempty"`
		Author string             `json:"author,omitempty" bson:"author,omitempty"`
		Tags   []string           `json:"tags,omitempty" bson:"tags,omitempty"`
		Dtmcrt time.Time          `bson:"dtmcrt"`
		Dtmupd time.Time          `bson:"dtmupd"`
	}
	// evermosJSON struct {
	// 	ID     primitive.ObjectID `json:"_id"`
	// 	Title  string             `json:"title"`
	// 	Author string             `json:"author"`
	// 	Tags   []string           `json:"tags"`
	// 	Dtmcrt string             `json:"dtmcrt"`
	// 	Dtmupd string             `json:"dtmupd"`
	// }

	FindevermosJSON struct {
		ID string `json:"id"`
	}
	InsertEpisodeevermosJSON struct {
		Title       string   `json:"title"`
		Author      string   `json:"author"`
		Tags        []string `json:"tags"`
		Episode     string   `json:"episode"`
		Description string   `json:"description"`
		Duration    int32    `json:"duration"`
	}

	Episode struct {
		ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		evermos     primitive.ObjectID `json:"evermos,omitempty" bson:"evermos,omitempty"`
		Episode     string             `json:"episode,omitempty" bson:"episode,omitempty"`
		Description string             `json:"description,omitempty" bson:"description,omitempty"`
		Duration    int32              `json:"duration,omitempty" bson:"duration,omitempty"`
	}

	DeleteevermosJSON struct {
		ID string `json:"id"`
	}

	UpdateevermosJSON struct {
		ID     string   `json:"id" bson:"id"`
		Title  string   `json:"title"`
		Author string   `json:"author"`
		Tags   []string `json:"tags"`
	}
)
