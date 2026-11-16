package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	cases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for _, c := range cases {
		cache.Add(c.key, c.val)

		actual, ok := cache.Get(c.key)
		if !ok {
			t.Errorf("expected to find key %s", c.key)
			continue
		}
		if !bytes.Equal(actual, c.val) {
			t.Errorf("expected value %s, got %s", string(c.val), string(actual))
		}
	}
}

func TestMissingKey(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	val, ok := cache.Get("nonexistent_key")
	if ok {
		t.Errorf("expected false for missing key, got true")
	}
	if val != nil {
		t.Errorf("expected nil for missing key, got %v", val)
	}
}

func TestReplaceValue(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	key := "test_key"
	val1 := []byte("first_value")
	val2 := []byte("second_value")

	cache.Add(key, val1)
	actual1, ok1 := cache.Get(key)
	if !ok1 || !bytes.Equal(actual1, val1) {
		t.Fatalf("expected initial value %s, got %s", string(val1), string(actual1))
	}

	// Replace the value with val2
	cache.Add(key, val2)
	actual2, ok2 := cache.Get(key)
	if !ok2 {
		t.Fatalf("expected key to exist after replacement")
	}
	if !bytes.Equal(actual2, val2) {
		t.Errorf("expected replaced value %s, got %s", string(val2), string(actual2))
	}
}

func TestMultipleKeysDoNotOverwrite(t *testing.T) {
	const interval = 5 * time.Second
	cache := NewCache(interval)

	k1, v1 := "key1", []byte("val1")
	k2, v2 := "key2", []byte("val2")
	k3, v3 := "key3", []byte("val3")

	cache.Add(k1, v1)
	cache.Add(k2, v2)
	cache.Add(k3, v3)

	got1, ok1 := cache.Get(k1)
	if !ok1 || !bytes.Equal(got1, v1) {
		t.Errorf("expected key1 to have value %s, got %s", string(v1), string(got1))
	}

	got2, ok2 := cache.Get(k2)
	if !ok2 || !bytes.Equal(got2, v2) {
		t.Errorf("expected key2 to have value %s, got %s", string(v2), string(got2))
	}

	got3, ok3 := cache.Get(k3)
	if !ok3 || !bytes.Equal(got3, v3) {
		t.Errorf("expected key3 to have value %s, got %s", string(v3), string(got3))
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond
	cache := NewCache(baseTime)

	cache.Add("https://example.com", []byte("testdata"))

	// Fresh entry should stay available immediately
	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	// After waitTime (which is longer than interval), old entry should be reaped
	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key after reap interval")
		return
	}
}
