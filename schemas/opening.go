package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Opening struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Role      string             `bson:"role"`
	Company   string             `bson:"company"`
	Location  string             `bson:"location"`
	Remote    bool               `bson:"remote"`
	Link      string             `bson:"link"`
	Salary    int64              `bson:"salary"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
	DeletedAt *time.Time         `bson:"deletedAt,omitempty"`
}
