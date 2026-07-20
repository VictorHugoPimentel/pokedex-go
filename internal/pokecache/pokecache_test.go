package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Errorf("Expected cache to be initialized, got nil")
	}
}

func TestAddCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{
			inputKey: "val1",
			inputVal: []byte("val1"),
		},
		{
			inputKey: "val2",
			inputVal: []byte("val2"),
		},
		{
			inputKey: "",
			inputVal: []byte("val3"),
		},
	}
	for _, cas := range cases {
		cache.Add(cas.inputKey, cas.inputVal)
		actual, ok := cache.Get(cas.inputKey)
		if !ok {
			t.Error(cas.inputKey + " not found in cache")
			continue
		}
		if string(actual) != string(cas.inputVal) {
			t.Errorf("Expected value '%s', got '%s'", string(cas.inputVal), string(actual))
			continue
		}
	}
}

func TestReap(t *testing.T) {
	internal := time.Millisecond * 10
	c := NewCache(internal)
	keyone := "keyone"
	c.Add(keyone, []byte("val1"))
	time.Sleep(internal + time.Millisecond)
	_, ok := c.Get(keyone)
	if ok {
		t.Errorf("Expected %s to be reaped", keyone)
	}
}

func TestReapFail(t *testing.T) {
	internal := time.Millisecond * 10
	c := NewCache(internal)
	keyone := "keyone"
	c.Add(keyone, []byte("val1"))
	time.Sleep(internal/2)
	_, ok := c.Get(keyone)
	if !ok {
		t.Errorf("Expected %s to still be active", keyone)
	}
}