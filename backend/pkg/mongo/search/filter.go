package search

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type TimeInterval struct {
	Begin *time.Time
	End   *time.Time
}

type Filter bson.M

func (f Filter) AddRegex(key string, value string) {
	if len(value) != 0 {
		f[key] = bson.M{"$regex": value, "$options": "i"}
	}
}

func (f Filter) AddEqual(key string, value string) {
	if len(value) != 0 {
		f[key] = value
	}
}

func (f Filter) AddTimeIterval(key string, interval TimeInterval) {
	filter := bson.M{}
	if interval.Begin != nil {
		filter["$gte"] = interval.Begin
	}
	if interval.End != nil {
		filter["$lte"] = interval.End
	}
	if len(filter) != 0 {
		f[key] = filter
	}
}

func (f Filter) ToBson() bson.M {
	return bson.M(f)
}
