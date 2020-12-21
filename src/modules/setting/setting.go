package setting

import "go.mongodb.org/mongo-driver/bson"

type Setting struct {
	Name  string      `json:"name" bson:"name"`
	Value interface{} `json:"value" bson:"value"`
}

func New(name string, value interface{}) *Setting {
	return &Setting{
		Name:  name,
		Value: value,
	}
}

func NewFilter(name string) bson.D {
	return bson.D{
		{
			Key:   "name",
			Value: name,
		},
	}
}

func (s *Setting) GetValue() interface{} {
	return s.Value
}
