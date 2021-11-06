package cache

import (
	"strconv"
	"sync"
	"testing"
)

func TestLruCache_Set(t *testing.T) {
	type fields struct {
		size     int
		elements ForwardList
		mp       map[string]*Node
	}
	type args struct {
		key   string
		value interface{}
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		wantOk bool
	}{
		{
			name: "test_set_get",
			fields: fields{
				1,
				NewList(),
				make(map[string]*Node),
			},
			args: args{
				key:   "firstValue",
				value: 100,
			},
			wantOk: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &LruCache{
				size:     tt.fields.size,
				elements: tt.fields.elements,
				mp:       tt.fields.mp,
			}
			if gotOk := c.Set(tt.args.key, tt.args.value); gotOk != tt.wantOk {
				t.Errorf("Set() = %v, want %v", gotOk, tt.wantOk)
			}
			if value, _ := c.Get(tt.args.key); value != tt.args.value {
				t.Errorf("Get() = %v, want %v", value, tt.args.value)
			}
			if gotOk := c.Set(tt.args.key, tt.args.value); gotOk == tt.wantOk {
				t.Errorf("Set() = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}

func TestLruCache_Set_One_Go(t *testing.T) {
	const maxCount = 10
	const maxCap = 10
	cache := NewCache(maxCap)

	for i := 0; i < maxCount; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i)
	}

	for i := maxCount - maxCap; i < maxCount; i++ {
		value := strconv.Itoa(i)
		expected := i
		if result, _ := cache.Get(value); expected != result.(int) {
			t.Fatalf("bad cache: got %v, expected %v", result, expected)
		}
	}
}

func TestLruCache_Set_One_Go_Many(t *testing.T) {
	const maxCount = 10000000
	const maxCap = 10
	cache := NewCache(maxCap)

	for i := 0; i < maxCount; i++ {
		key := strconv.Itoa(i)
		cache.Set(key, i)
	}

	for i := maxCount - maxCap; i < maxCount; i++ {
		value := strconv.Itoa(i)
		expected := i
		if result, _ := cache.Get(value); expected != result.(int) {
			t.Fatalf("bad cache: got %v, expected %v", result, expected)
		}
	}
}

func TestLruCache_Set_Some_Go(t *testing.T) {
	const maxCount = 10
	const maxCap = 10
	cache := NewCache(maxCap)

	wg := &sync.WaitGroup{}
	for i := 0; i < maxCount; i++ {
		key := strconv.Itoa(i)
		wg.Add(1)
		go func(key string, value int) {
			defer wg.Done()
			cache.Set(key, value)
		}(key, i)
	}
	wg.Wait()

	for i := maxCount - maxCap; i < maxCount; i++ {
		value := strconv.Itoa(i)
		expected := i
		if result, _ := cache.Get(value); expected != result.(int) {
			t.Fatalf("bad cache(go): got %v, expected %v", result, expected)
		}
	}
}

func TestLruCache_Set_Many_Go(t *testing.T) {
	const maxCount = 1000000
	const maxCap = 1000
	const maxGoro = 256
	cache := NewCache(maxCap)
	goMaxCount := make(chan struct{}, maxGoro)

	wg := &sync.WaitGroup{}
	for i := 0; i < maxCount; i++ {
		key := strconv.Itoa(i)
		wg.Add(1)
		goMaxCount <- struct{}{}
		go func(key string, value int) {
			defer wg.Done()
			cache.Set(key, value)
			<-goMaxCount
		}(key, i)
	}
	wg.Wait()

	for i := maxCount - maxCap; i < maxCount; i++ {
		wg.Add(1)
		goMaxCount <- struct{}{}
		go func(i int) {
			defer wg.Done()
			defer func() {
				<-goMaxCount
			}()

			value := strconv.Itoa(i)
			expected := i
			if result, _ := cache.Get(value); expected != result.(int) {
				t.Errorf("bad cache(go many): got %v, expected %v", result, expected)
				return
			}
		}(i)
	}

	wg.Wait()
}
