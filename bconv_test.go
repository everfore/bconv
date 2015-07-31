package bconv

import (
	"labix.org/v2/mgo/bson"
	"testing"
)

type Lang struct {
	Name string
}

var (
	testBson = &bson.M{
		"Name": "Golang",
	}
	testJson = &Lang{
		Name: "Golang ?",
	}
)

func TestJson(t *testing.T) {
	j := NewJson(testJson)
	jbbs := j.BBytes()
	t.Log(jbbs)

	jjbs := j.JBytes()
	t.Log(jjbs)

	jb := j.Bson()
	t.Log(jb)
}

func TestBson(t *testing.T) {
	b := NewBson(testBson)
	bbbs := b.BBytes()
	t.Log(bbbs)

	bjbs := b.JBytes()
	t.Log(bjbs)

	bj := b.Json()
	t.Log(bj)
}

func TestBsonJson(t *testing.T) {
	b := NewBson(testBson)
	bbbs := b.BBytes()
	t.Log(bbbs)

	bjbs := b.JBytes()
	t.Log(bjbs)

	bj := b.Json()
	t.Log(bj)

	jbbs := bj.BBytes()
	t.Log(jbbs)

	jjbs := bj.JBytes()
	t.Log(jjbs)

	jb := bj.Bson()
	t.Log(jb)
}

func TestJsonBson(t *testing.T) {
	j := NewJson(testJson)
	jbbs := j.BBytes()
	t.Log(jbbs)

	jjbs := j.JBytes()
	t.Log(jjbs)

	jb := j.Bson()
	t.Log(jb)

	bbbs := jb.BBytes()
	t.Log(bbbs)

	bjbs := jb.JBytes()
	t.Log(bjbs)

	bj := jb.Json()
	t.Log(bj)
}
