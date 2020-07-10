package assert

import (
	"reflect"
	"testing"
	"time"
)

func TestEqual(t *testing.T) {
	type innerStruct struct {
		Value string
	}

	type testStruct struct {
		ID          int
		Value       string
		CreatedAt   time.Time
		InnerStruct innerStruct
	}

	type args struct {
		actual        interface{}
		expected      interface{}
		skippedFields []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				actual: &testStruct{
					ID:          1,
					Value:       "test",
					CreatedAt:   time.Now(),
					InnerStruct: innerStruct{Value: "123"},
				},
				expected: &testStruct{
					ID:          1,
					Value:       "test",
					CreatedAt:   time.Now().Add(time.Hour),
					InnerStruct: innerStruct{Value: "123"},
				},
				skippedFields: []string{"CreatedAt"},
			},

			want: true,
		},
		{
			name: "fail",
			args: args{
				actual: &testStruct{
					ID:          1,
					Value:       "test",
					CreatedAt:   time.Now(),
					InnerStruct: innerStruct{Value: "123"},
				},
				expected: &testStruct{
					ID:          1,
					Value:       "test",
					CreatedAt:   time.Now().Add(time.Hour),
					InnerStruct: innerStruct{Value: "123"},
				},
			},

			want: false,
		},
		{
			name: "success with primitive",
			args: args{
				actual:   "hello world",
				expected: "hello world",
			},
			want: true,
		},
		{
			name: "fail with primitive",
			args: args{
				actual:   "hello world",
				expected: "hello world - 123",
			},
			want: false,
		},
		{
			name: "success with primitive ptr",
			args: args{
				actual: func() *string {
					s := "hello world"
					return &s
				}(),
				expected: func() *string {
					s := "hello world"
					return &s
				}(),
			},
			want: true,
		},
		{
			name: "fail with primitive ptr",
			args: args{
				actual: func() *string {
					s := "hello world"
					return &s
				}(),
				expected: func() *string {
					s := "hello world - 123"
					return &s
				}(),
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equals(tt.args.actual, tt.args.expected, tt.args.skippedFields...); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualsWithDiffFunc(t *testing.T) {
	type innerStruct struct {
		Value string
	}

	type testStruct struct {
		ID          int
		Value       string
		CreatedAt   time.Time
		InnerStruct innerStruct
	}

	type args struct {
		actual        interface{}
		expected      interface{}
		skippedFields []string
		diffFn        func(actual, expected interface{}) bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				actual: &testStruct{
					ID:          1,
					Value:       "test",
					CreatedAt:   time.Now(),
					InnerStruct: innerStruct{Value: "123"},
				},
				expected: &testStruct{
					ID:          1,
					Value:       "test",
					CreatedAt:   time.Now().Add(time.Hour),
					InnerStruct: innerStruct{Value: "123"},
				},
				skippedFields: []string{"CreatedAt"},
				diffFn: func(actual, expected interface{}) bool {
					return reflect.DeepEqual(actual, expected)
				},
			},
			want: true,
		},
		{
			name: "fail",
			args: args{
				actual: &testStruct{
					ID:          1,
					Value:       "test",
					CreatedAt:   time.Now(),
					InnerStruct: innerStruct{Value: "123"},
				},
				expected: &testStruct{
					ID:          1,
					Value:       "test",
					CreatedAt:   time.Now().Add(time.Hour),
					InnerStruct: innerStruct{Value: "123"},
				},
				diffFn: func(actual, expected interface{}) bool {
					return reflect.DeepEqual(actual, expected)
				},
			},
			want: false,
		},
		{
			name: "success with primitive",
			args: args{
				actual:   "hello world",
				expected: "hello world",
				diffFn: func(actual, expected interface{}) bool {
					return reflect.DeepEqual(actual, expected)
				},
			},
			want: true,
		},
		{
			name: "fail with primitive",
			args: args{
				actual:   "hello world",
				expected: "hello world - 123",
				diffFn: func(actual, expected interface{}) bool {
					return reflect.DeepEqual(actual, expected)
				},
			},
			want: false,
		},
		{
			name: "success with primitive ptr",
			args: args{
				actual: func() *string {
					s := "hello world"
					return &s
				}(),
				expected: func() *string {
					s := "hello world"
					return &s
				}(),
				diffFn: func(actual, expected interface{}) bool {
					return reflect.DeepEqual(actual, expected)
				},
			},
			want: true,
		},
		{
			name: "fail with primitive ptr",
			args: args{
				actual: func() *string {
					s := "hello world"
					return &s
				}(),
				expected: func() *string {
					s := "hello world - 123"
					return &s
				}(),
				diffFn: func(actual, expected interface{}) bool {
					return reflect.DeepEqual(actual, expected)
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualsWithDiffFunc(tt.args.actual, tt.args.expected, tt.args.skippedFields, tt.args.diffFn); got != tt.want {
				t.Errorf("EqualsWithDiffFunc() = %v, want %v", got, tt.want)
			}
		})
	}
}
