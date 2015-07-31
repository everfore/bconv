package bconv

import (
	"fmt"
	"labix.org/v2/mgo/bson"
)

type Json struct {
	Jv interface{}
}

func NewJson(v interface{}) *Json {
	return &Json{Jv: v}
}

type Bson struct {
	Bv *bson.M
}

func (b *Bson) String() string {
	return fmt.Sprintf("%#v", b.Bv)
}

func NewBson(v *bson.M) *Bson {
	return &Bson{Bv: v}
}

// return JsonBytes
func (j *Json) JBytes() []byte {
	b, err := Json2JsonBytes(j.Jv)
	if isError(err) {
		return nil
	}
	return b
}

func (j *Json) BBytes() []byte {
	Bv, err := Json2Bson(j.Jv)
	if isError(err) {
		return nil
	}
	b, err := Bson2BsonBytes(Bv)
	if isError(err) {
		return nil
	}
	return b
}

func (j *Json) Bson() *Bson {
	Bv, err := Json2Bson(j.Jv)
	if isError(err) {
		return nil
	}
	return NewBson(Bv)
}

func (b *Bson) BBytes() []byte {
	bs, err := Bson2BsonBytes(b.Bv)
	if isError(err) {
		return nil
	}
	return bs
}

func (b *Bson) JBytes() []byte {
	bs, err := Bson2JsonBytes(b.Bv)
	if isError(err) {
		return nil
	}
	return bs
}

func (b *Bson) Json() *Json {
	var i interface{}
	err := Bson2Json(b.Bv, &i)
	if isError(err) {
		return nil
	}
	return NewJson(i)
}
