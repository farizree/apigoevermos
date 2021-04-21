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
	Merchant struct {
		ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		MerchantName string             `json:"merchantname,omitempty" bson:"merchantname,omitempty"`
		Address      string             `json:"address,omitempty" bson:"address,omitempty"`
		Phone        string             `json:"phone,omitempty" bson:"phone,omitempty"`
		Owner        string             `json:"owner,omitempty" bson:"owner,omitempty"`
		IsActive     int32              `json:"isactive,omitempty" bson:"isactive,omitempty"`
		Dtmcrt       time.Time          `bson:"dtmcrt"`
		Dtmupd       time.Time          `bson:"dtmupd"`
	}

	FindMerchantJSON struct {
		ID string `json:"id"`
	}

	Product struct {
		ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
		MerchantID   primitive.ObjectID `json:"merchantid,omitempty" bson:"merchantid,omitempty"`
		MerchantName string             `json:"merchantname,omitempty" bson:"merchantname,omitempty"`
		ProductName  string             `json:"productname,omitempty" bson:"productname,omitempty"`
		Price        string             `json:"price,omitempty" bson:"price,omitempty"`
		Category     string             `json:"category,omitempty" bson:"category,omitempty"`
		Stock        int32              `json:"stock,omitempty" bson:"stock,omitempty"`
		IsActive     int32              `json:"isactive,omitempty" bson:"isactive,omitempty"`
		Dtmcrt       time.Time          `bson:"dtmcrt"`
		Dtmupd       time.Time          `bson:"dtmupd"`
	}

	FindProductJSON struct {
		ID string `json:"id"`
	}

	InsertMerchantJSON struct {
		MerchantName string `json:"merchantname"`
		Address      string `json:"address"`
		Phone        string `json:"phone"`
		Owner        string `json:"owner"`
		IsActive     int32  `json:"isactive"`
	}

	InsertProductJSON struct {
		MerchantID   primitive.ObjectID `json:"merchantid"`
		MerchantName string             `json:"merchantname"`
		ProductName  string             `json:"productname"`
		Price        string             `json:"price"`
		Category     string             `json:"category"`
		Stock        int32              `json:"stock"`
		IsActive     int32              `json:"isactive"`
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
