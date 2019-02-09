package keva

import (
	"github.com/aws/aws-sdk-go/aws"
	"testing"
)

var kv = NewWithConfig("keva", &aws.Config{
	Endpoint: aws.String("http://localhost:8000"),
})

func TestString(t *testing.T) {
	key, val := "test:string", "abc@#€$"
	kv.Set(key, val)
	defer kv.Delete(key)
	res := kv.Get(key)
	if res != val {
		t.Errorf("result differs: %v != %v", res, val)
	}
}

func TestBool(t *testing.T) {
	key, val := "test:bool", true
	kv.Set(key, val)
	res := kv.Get(key) //.(bool)
	defer kv.Delete(key)
	if res != val {
		t.Errorf("result not %+v, is %+v", val, res)
	}
}

func TestFloat(t *testing.T) {
	key, val := "test:float", 1234.4321
	kv.Set(key, val)
	defer kv.Delete(key)
	res := kv.Get(key) //.(float64)
	if res != val {
		t.Errorf("result differs: %v != %v", res, val)
	}
}

func TestStringSlice(t *testing.T) {
	key, val := "test:stringslice", []string{"one", "two", "many"}
	kv.Set(key, val)
	defer kv.Delete(key)
	res := kv.GetSlice(key)
	for i := range res {
		if res[i] != val[i] {
			t.Errorf("result differs at %d : %v != %v", i, res[i], val[i])
		}
	}
}

func TestFloatSlice(t *testing.T) {
	key, val := "test:floatslice", []float64{1, 2, 3, 4}
	kv.Set(key, val)
	defer kv.Delete(key)
	res := kv.GetSlice(key)
	for i := range res {
		if res[i] != val[i] {
			t.Errorf("result differs at %d : %v != %v", i, res[i], val[i])
		}
	}
}

func TestStringMap(t *testing.T) {
	key, val := "test:stringmap", map[string]string{
		"a": "A",
		"w": "W",
		"s": "s",
	}
	kv.Set(key, val)
	defer kv.Delete(key)
	res := kv.GetStringMap(key)
	if res["a"] != val["a"] || res["w"] != val["w"] || res["s"] != val["s"] {
		t.Errorf("result differs: %v != %v", res, val)
	}
}

func TestFloatMap(t *testing.T) {
	key, val := "test:floatmap", map[string]float64{
		"a": 1.2,
		"w": 2.3,
		"s": 4.3,
	}
	kv.Set(key, val)
	defer kv.Delete(key)
	res := kv.GetFloatMap(key)
	if res["a"] != val["a"] || res["w"] != val["w"] || res["s"] != val["s"] {
		t.Errorf("result differs: %v != %v", res, val)
	}
}

func TestInterfaceMap(t *testing.T) {
	// same as GetFloatMap
	key, val := "test:floatmapinterface", map[string]float64{
		"a": 3.4,
		"w": 4.5,
		"s": 5.6,
	}
	kv.Set(key, val)
	res := kv.Get(key).(map[string]interface{})
	defer kv.Delete(key)
	if res["a"] != val["a"] || res["w"] != val["w"] || res["s"] != val["s"] {
		t.Errorf("result differs: %v != %v", res, val)
	}
}
