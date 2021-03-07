package flags

import "testing"

func TestCountTheWays_String(t *testing.T) {
	var c = CountTheWays([]int{1, 2})
	tests := []struct {
		name string
		c    *CountTheWays
		want string
	}{
		{"base-case", &c, "1 ... 2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.String(); got != tt.want {
				t.Errorf("CountTheWays.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountTheWays_Set(t *testing.T) {
	var c = CountTheWays([]int{1, 2})
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		c       *CountTheWays
		args    args
		wantErr bool
	}{
		{"base-case", &c, args{"1,2"}, false},
		{"bad args", &c, args{"1,a"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.Set(tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("CountTheWays.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
