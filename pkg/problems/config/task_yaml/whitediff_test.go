package task_yaml

import (
	"io"
	"strings"
	"testing"
)

func TestWhitediff(t *testing.T) {
	type args struct {
		a io.Reader
		b io.Reader
	}

	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		//tests from cms
		{"test_no_diff_one_token1", args{strings.NewReader(""), strings.NewReader("")}, 1.0, false},
		{"test_no_diff_one_token2", args{strings.NewReader("1"), strings.NewReader("1")}, 1.0, false},
		{"test_no_diff_one_token3", args{strings.NewReader("a"), strings.NewReader("a")}, 1.0, false},
		{"test_no_diff_one_token4", args{strings.NewReader("你好"), strings.NewReader("你好")}, 1.0, false},

		{"test_no_diff_one_token_and_whites1", args{strings.NewReader("1   "), strings.NewReader("1")}, 1.0, false},
		{"test_no_diff_one_token_and_whites2", args{strings.NewReader("   1"), strings.NewReader("1")}, 1.0, false},
		{"test_no_diff_one_token_and_whites3", args{strings.NewReader("1" + string(whitespaces)), strings.NewReader("1")}, 1.0, false},

		{"test_no_diff_one_token_and_trailing_blank_lines1", args{strings.NewReader("1\n"), strings.NewReader("1")}, 1.0, false},
		{"test_no_diff_one_token_and_trailing_blank_lines2", args{strings.NewReader("1\n\n\n\n"), strings.NewReader("1")}, 1.0, false},
		{"test_no_diff_one_token_and_trailing_blank_lines3", args{strings.NewReader("1\n\n\n\n"), strings.NewReader("1\n")}, 1.0, false},
		{"test_no_diff_one_token_and_trailing_blank_lines4", args{strings.NewReader("1\n\n\r \n\n \n \n"), strings.NewReader("1   \n\r   ")}, 1.0, false},

		{"test_no_diff_multiple_tokens1", args{strings.NewReader("1 asd\n\n\n"), strings.NewReader("   1\tasd  \n")}, 1.0, false},
		{"test_no_diff_multiple_tokens2", args{strings.NewReader("1 2\n\n\n"), strings.NewReader("1 2\n")}, 1.0, false},
		{"test_no_diff_multiple_tokens3", args{strings.NewReader("1\t\r2"), strings.NewReader("1 2")}, 1.0, false},

		{"test_diff_wrong_tokens1", args{strings.NewReader("1 2"), strings.NewReader("12")}, 0.0, false},
		{"test_diff_wrong_tokens2", args{strings.NewReader("1 23"), strings.NewReader("12 3")}, 0.0, false},
		{"test_diff_wrong_tokens3", args{strings.NewReader("1"), strings.NewReader("01")}, 0.0, false},
		{"test_diff_wrong_tokens4", args{strings.NewReader("1.0"), strings.NewReader("1")}, 0.0, false},

		{"test_diff_wrong_line1", args{strings.NewReader("\n1"), strings.NewReader("1")}, 0.0, false},
		{"test_diff_wrong_line2", args{strings.NewReader("1 2"), strings.NewReader("1\n2")}, 0.0, false},
		{"test_diff_wrong_line3", args{strings.NewReader("1\n\n2"), strings.NewReader("1\n2")}, 0.0, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Whitediff(tt.args.a, tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("Whitediff() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Whitediff() = %v, want %v", got, tt.want)
			}
		})
	}
}
