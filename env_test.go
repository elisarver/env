package env

import (
	"os"
	"testing"
)

func init(){
	os.Setenv("FOUND", "found")
	os.Unsetenv("MISSING")
}

func TestFallback(t *testing.T) {
	type args struct {
		key      string
		fallback string
	}
	tests := []struct {
		name       string
		args       args
		wantString string
		wantBool   bool
	}{
		{
			name: "found",
			args: args{
				key:      "FOUND",
				fallback: "fallback",
			},
			wantString: "found",
			wantBool:   true,
		},
		{
			name: "fallback",
			args: args{
				key:      "MISSING",
				fallback: "fallback",
			},
			wantString: "fallback",
			wantBool:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotString, gotBool := FallbackLookup(tt.args.key, tt.args.fallback)
			if gotString != tt.wantString {
				t.Errorf("FallbackLookup() gotString = %v, wantString %v", gotString, tt.wantString)
			}
			if gotBool != tt.wantBool {
				t.Errorf("FallbackLookup() gotBool = %v, wantBool %v", gotBool, tt.wantBool)
			}
		})
	}
}

func TestGet(t *testing.T) {
	// core function, tested simply
	Get("FOUND")
}

func TestLookupError_Error(t *testing.T) {
	type fields struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "filled in",
			fields: fields{
				key: "KEY",
			},
			want: `env: MustLookup("KEY"): Key not found`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := LookupError{
				Key: tt.fields.key,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLookup(t *testing.T) {
	// core function, tested simply
	Lookup("FOUND")
}

func TestMustLookup(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name      string
		args      args
		want      string
		wantPanic bool
	}{
		{
			name: "found",
			args: args{
				key: "FOUND",
			},
			want: "found",
		},
		{
			name: "missing",
			args: args{
				key: "MISSING",
			},
			wantPanic: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func() {
				if err := recover(); err != nil && tt.wantPanic != true {
					t.Error(err)
					t.Fail()
				}
			}()
			if got := MustLookup(tt.args.key); got != tt.want {
				t.Errorf("MustLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
