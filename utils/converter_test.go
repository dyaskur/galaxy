package utils

import "testing"

func TestRomanToDecimal(t *testing.T) {
	type args struct {
		roman string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "I",
			args: args{roman: "I"},
			want: 1,
		},
		{
			name: "II",
			args: args{roman: "II"},
			want: 2,
		},
		{
			name: "III",
			args: args{roman: "III"},
			want: 3,
		},
		{
			name: "IV",
			args: args{roman: "IV"},
			want: 4,
		},
		{
			name: "V",
			args: args{roman: "V"},
			want: 5,
		},
		{
			name: "VI",
			args: args{roman: "VI"},
			want: 6,
		},
		{
			name: "VII",
			args: args{roman: "VII"},
			want: 7,
		},
		{
			name: "VIII",
			args: args{roman: "VIII"},
			want: 8,
		},
		{
			name: "IX",
			args: args{roman: "IX"},
			want: 9,
		},
		{
			name: "X",
			args: args{roman: "X"},
			want: 10,
		},
		{
			name: "XI",
			args: args{roman: "XI"},
			want: 11,
		},
		{
			name: "XII",
			args: args{roman: "XII"},
			want: 12,
		},
		{
			name: "XIII",
			args: args{roman: "XIII"},
			want: 13,
		},
		{
			name: "XIV",
			args: args{roman: "XIV"},
			want: 14,
		},
		{
			name: "XV",
			args: args{roman: "XV"},
			want: 15,
		},
		{
			name: "XVI",
			args: args{roman: "XVI"},
			want: 16,
		},
		{
			name: "XVII",
			args: args{roman: "XVII"},
			want: 17,
		},
		{
			name: "XVIII",
			args: args{roman: "XVIII"},
			want: 18,
		},
		{
			name: "XIX",
			args: args{roman: "XIX"},
			want: 19,
		},
		{
			name: "XX",
			args: args{roman: "XX"},
			want: 20,
		},
		{
			name: "XXI",
			args: args{roman: "XXI"},
			want: 21,
		},
		{
			name: "XXII",
			args: args{roman: "XXII"},
			want: 22,
		},
		{
			name: "XXIII",
			args: args{roman: "XXIII"},
			want: 23,
		},
		{
			name: "XXIV",
			args: args{roman: "XXIV"},
			want: 24,
		},
		{
			name: "XXV",
			args: args{roman: "XXV"},
			want: 25,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RomanToDecimal(tt.args.roman); got != float64(tt.want) {
				t.Errorf("RomanToDecimal() = %v, want %v", got, tt.want)
			}
		})
	}
}
