package bconv

import (
	"gopkg.in/mgo.v2/bson"
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

func TestJConv(t *testing.T) {
	j := NewJson(testJson)
	t.Logf("%#v", j.Jv)

	type Java struct {
		Name string
	}

	var java0 Java
	err0 := j.Conv(&java0)
	log_err(err0, t)
	t.Logf("%#v", java0)

	var java Java
	err := j.Bson().Conv(&java)
	log_err(err, t)
	t.Logf("%#v", java)
}

func TestBConv(t *testing.T) {
	b := NewBson(testBson)
	t.Logf("%#v", b.Bv)

	type Java struct {
		Name string
	}

	var java0 Java
	err0 := b.Conv(&java0)
	log_err(err0, t)
	t.Logf("%#v", java0)

	var java Java
	err := b.Json().Conv(&java)
	log_err(err, t)
	t.Logf("%#v", java)
}
