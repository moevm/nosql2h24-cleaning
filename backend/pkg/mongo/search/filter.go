package search

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type TimeInterval interface {
	GetBegin() *time.Time
	GetEnd() *time.Time
}

type NumberRange interface {
	GetMin() *int
	GetMax() *int
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

func (f Filter) AddNotEqual(key string, value string) {
	filter := bson.M{}
	if len(value) != 0 {
		filter["$ne"] = value
	}
	if len(filter) != 0 {
		f[key] = filter
	}
}

func (f Filter) AddEqualObjectID(key string, value string) error {
	if len(value) == 0 {
		return nil
	}
	objectID, err := bson.ObjectIDFromHex(value)
	if err != nil {
		return err
	}
	f[key] = objectID
	return nil
}

func (f Filter) AddTimeIterval(key string, interval TimeInterval) {
	filter := bson.M{}
	if interval.GetBegin() != nil {
		filter["$gte"] = interval.GetBegin()
	}
	if interval.GetEnd() != nil {
		filter["$lte"] = interval.GetEnd()
	}
	if len(filter) != 0 {
		f[key] = filter
	}
}

func (f Filter) AddNumberRange(key string, rangeValue NumberRange) {
	filter := bson.M{}
	if rangeValue.GetMin() != nil {
		filter["$gte"] = rangeValue.GetMin()
	}
	if rangeValue.GetMax() != nil {
		filter["$lte"] = rangeValue.GetMax()
	}
	if len(filter) != 0 {
		f[key] = filter
	}
}

func (f Filter) AddIn(key string, values []string) {
	if len(values) != 0 {
		f[key] = bson.M{"$in": values}
	}
}

func toObjectID(values []string) []bson.ObjectID {
	arr := make([]bson.ObjectID, len(values))
	for i, v := range values {
		arr[i], _ = bson.ObjectIDFromHex(v)
	}
	return arr
}

func (f Filter) AddInObjectIDs(key string, values []string) {
	if len(values) == 0 {
		return
	}
	f[key] = bson.M{"$in": toObjectID(values)}
}

func (f Filter) ContainsAll(key string, values []string) {
	if len(values) == 0 {
		return
	}
	f[key] = bson.M{"$all": toObjectID(values)}
}

func (f Filter) ToBson() bson.M {
	return bson.M(f)
}
