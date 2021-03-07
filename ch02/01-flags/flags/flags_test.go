package flags

import "testing"

func TestConfig_Setup(t *testing.T) {
	type fields struct {
		subject      string
		isAwesome    bool
		howAwesome   int
		countTheWays CountTheWays
	}
	tests := []struct {
		name   string
		fields fields
	}{
		{"base-case", fields{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				subject:      tt.fields.subject,
				isAwesome:    tt.fields.isAwesome,
				howAwesome:   tt.fields.howAwesome,
				countTheWays: tt.fields.countTheWays,
			}
			c.Setup()
		})
	}
}

func TestConfig_GetMessage(t *testing.T) {
	type fields struct {
		subject      string
		isAwesome    bool
		howAwesome   int
		countTheWays CountTheWays
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"base-case", fields{subject: "something"}, "something is NOT awesome with a certainty of 0 out of 10. Let me count the ways "},
		{"base-case", fields{subject: "something", isAwesome: true}, "something is awesome with a certainty of 0 out of 10. Let me count the ways "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Config{
				subject:      tt.fields.subject,
				isAwesome:    tt.fields.isAwesome,
				howAwesome:   tt.fields.howAwesome,
				countTheWays: tt.fields.countTheWays,
			}
			if got := c.GetMessage(); got != tt.want {
				t.Errorf("Config.GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
