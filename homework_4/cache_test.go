package homework_4

import (
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
