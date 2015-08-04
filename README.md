# 	bconv [![](http://gocover.io/_badge/github.com/everfore/bconv)](http://gocover.io/github.com/everfore/bconv)

Bson & Json Bytes Convert

----------------------------------

##	Get

		go get github.com/everfore/bconv

###		3rd party pkg 

		go get labix.org/v2/mgo/bson

##	Define

json represents for interface{}, for instance struct, map, slice.

bson represents for bson.M, actually which is a map[string]interface{}.

## 	Convert Conclusion

###		json <----> bson

json to bson, using bson.Marshal

bson to json, using json.Marshal

###		jsonbytes <--X--> bsonbytes

Well, jsonbytes do not equal bsonbytes. They are different []byte.

We can not use string(bsonbytes) as string.

##		Usage

###		Init

		var testBson = &bson.M{
			"Name": "Golang",
		}

		b := NewBson(testBson)


		type Lang struct {
			Name string
		}

		var testJson = &Lang{
			Name: "Golang ?",
		}

		j := NewJson(testJson)


###		Bson

		b := NewBson(testBson)

		bbbs := b.BBytes()
		t.Log(bbbs)

		bjbs := b.JBytes()
		t.Log(bjbs)

		bj := b.Json()
		t.Log(bj)

###		Json

		j := NewJson(testJson)

		jbbs := j.BBytes()
		t.Log(jbbs)

		jjbs := j.JBytes()
		t.Log(jjbs)

		jb := j.Bson()
		t.Log(jb)

###		Conv
	
	type Java struct {
		Name string
	}

	var java Java

1.	Json Conv

	err := j.Conv(&java)

	err := b.Json().Conv(&java)

2.	Bson Conv

	err := j.Bson().Conv(&java)
	
	err := b.Conv(&java)