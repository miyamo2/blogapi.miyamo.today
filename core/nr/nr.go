package nr

import (
	"fmt"
	"reflect"
)

type Segment interface {
	AddAttribute(key string, val any)
}

type Attribute struct {
	Key   string
	Value string
}

type Attributer interface {
	SegmentAttributes() []Attribute
}

func Add(segment Segment, keyPrefix string, value any) {
	if attributer, ok := value.(Attributer); ok {
		for _, attr := range attributer.SegmentAttributes() {
			segment.AddAttribute(fmt.Sprintf("%s.%s", keyPrefix, attr.Key), attr.Value)
		}
		return
	}
	rv := reflect.ValueOf(value)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		segment.AddAttribute(keyPrefix, fmt.Sprintf("%v", value))
		return
	}
	rt := rv.Type()
	for i := 0; i < rt.NumField(); i++ {
		field := rv.Field(i)
		if !field.CanInterface() {
			continue
		}
		fieldValue := field.Interface()
		segment.AddAttribute(fmt.Sprintf("%s.%s", keyPrefix, rt.Field(i).Name), fmt.Sprintf("%v", fieldValue))
	}
}
