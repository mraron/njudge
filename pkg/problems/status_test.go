package problems

import (
	"encoding/json"
	"github.com/mraron/njudge/pkg/language/memory"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestFeedbackFromShortString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want FeedbackType
	}{
		{"ioi", args{"ioi"}, FeedbackIOI},
		{"cf", args{"cf"}, FeedbackCF},
		{"acm", args{"acm"}, FeedbackACM},
		{"default", args{"sdfsdf"}, FeedbackCF},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FeedbackTypeFromShortString(tt.args.str); got != tt.want {
				t.Errorf("FeedbackFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_FirstNonAC(t *testing.T) {
	type fields struct {
		Name         string
		Scoring      ScoringType
		Testcases    []Testcase
		Dependencies []string
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{"AC", fields{Testcases: []Testcase{{VerdictName: VerdictAC}, {VerdictName: VerdictAC}}}, -1},
		{"WA", fields{Testcases: []Testcase{{VerdictName: VerdictAC}, {VerdictName: VerdictWA}}}, 2},
		{"DR", fields{Testcases: []Testcase{{VerdictName: VerdictAC}, {VerdictName: VerdictAC}, {VerdictName: VerdictDR}}}, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Group{
				Name:         tt.fields.Name,
				Scoring:      tt.fields.Scoring,
				Testcases:    tt.fields.Testcases,
				Dependencies: tt.fields.Dependencies,
			}
			if got := g.FirstNonAC(); got != tt.want {
				t.Errorf("FirstNonAC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_IsAC(t *testing.T) {
	type fields struct {
		Name         string
		Scoring      ScoringType
		Testcases    []Testcase
		Dependencies []string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"AC", fields{Testcases: []Testcase{{VerdictName: VerdictAC}, {VerdictName: VerdictAC}}}, true},
		{"WA", fields{Testcases: []Testcase{{VerdictName: VerdictAC}, {VerdictName: VerdictWA}}}, false},
		{"DR", fields{Testcases: []Testcase{{VerdictName: VerdictAC}, {VerdictName: VerdictAC}, {VerdictName: VerdictDR}}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Group{
				Name:         tt.fields.Name,
				Scoring:      tt.fields.Scoring,
				Testcases:    tt.fields.Testcases,
				Dependencies: tt.fields.Dependencies,
			}
			if got := g.IsAC(); got != tt.want {
				t.Errorf("IsAC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_MaxMemoryUsage(t *testing.T) {
	type fields struct {
		Name         string
		Scoring      ScoringType
		Testcases    []Testcase
		Dependencies []string
	}
	tests := []struct {
		name   string
		fields fields
		want   memory.Amount
	}{
		{"last", fields{Testcases: []Testcase{{MemoryUsed: 0}, {MemoryUsed: 2}}}, 2},
		{"first", fields{Testcases: []Testcase{{MemoryUsed: 4}, {MemoryUsed: 2}}}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Group{
				Name:         tt.fields.Name,
				Scoring:      tt.fields.Scoring,
				Testcases:    tt.fields.Testcases,
				Dependencies: tt.fields.Dependencies,
			}
			if got := g.MaxMemoryUsage(); got != tt.want {
				t.Errorf("MaxMemoryUsage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_MaxScore(t *testing.T) {
	type fields struct {
		Name         string
		Scoring      ScoringType
		Testcases    []Testcase
		Dependencies []string
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"last", fields{Scoring: ScoringSum, Testcases: []Testcase{{MaxScore: 0.5}, {MaxScore: 0.5}}}, 1},
		{"first", fields{Scoring: ScoringSum, Testcases: []Testcase{{MaxScore: 1}, {MaxScore: 2}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Group{
				Name:         tt.fields.Name,
				Scoring:      tt.fields.Scoring,
				Testcases:    tt.fields.Testcases,
				Dependencies: tt.fields.Dependencies,
			}
			if got := g.MaxScore(); got != tt.want {
				t.Errorf("MaxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_MaxTimeSpent(t *testing.T) {
	type fields struct {
		Name         string
		Scoring      ScoringType
		Testcases    []Testcase
		Dependencies []string
	}
	tests := []struct {
		name   string
		fields fields
		want   time.Duration
	}{
		{"last", fields{Testcases: []Testcase{{TimeSpent: 1 * time.Second}, {TimeSpent: 2 * time.Second}}}, 2 * time.Second},
		{"first", fields{Testcases: []Testcase{{TimeSpent: 500 * time.Second}, {TimeSpent: 400 * time.Second}}}, 500 * time.Second},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Group{
				Name:         tt.fields.Name,
				Scoring:      tt.fields.Scoring,
				Testcases:    tt.fields.Testcases,
				Dependencies: tt.fields.Dependencies,
			}
			if got := g.MaxTimeSpent(); got != tt.want {
				t.Errorf("MaxTimeSpent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_Score(t *testing.T) {
	type fields struct {
		Name         string
		Scoring      ScoringType
		Testcases    []Testcase
		Dependencies []string
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{"scoring sum ac", fields{Scoring: ScoringSum, Testcases: []Testcase{{Score: 0.1, MaxScore: 0.5}, {Score: 0.15, MaxScore: 0.5}}}, 0.25},
		{"scoring sum wa", fields{Scoring: ScoringSum, Testcases: []Testcase{{VerdictName: VerdictWA, Score: 0.0, MaxScore: 0.5}, {Score: 0.15, MaxScore: 0.5}}}, 0.15},
		{"scoring group ac", fields{Scoring: ScoringGroup, Testcases: []Testcase{{VerdictName: VerdictPC, Score: 0.1, MaxScore: 0.5}, {VerdictName: VerdictAC, Score: 0.5, MaxScore: 0.5}}}, 0.6},
		{"scoring group pc", fields{Scoring: ScoringGroup, Testcases: []Testcase{{VerdictName: VerdictPC, Score: 0.1, MaxScore: 0.5}, {VerdictName: VerdictAC, Score: 0.5, MaxScore: 0.5}}}, 0.6},
		{"scoring group wa", fields{Scoring: ScoringGroup, Testcases: []Testcase{{VerdictName: VerdictWA, Score: 0.0, MaxScore: 0.5}, {VerdictName: VerdictAC, Score: 0.5, MaxScore: 0.5}}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := Group{
				Name:         tt.fields.Name,
				Scoring:      tt.fields.Scoring,
				Testcases:    tt.fields.Testcases,
				Dependencies: tt.fields.Dependencies,
			}
			if got := g.Score(); got != tt.want {
				t.Errorf("Score() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScoringFromString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want ScoringType
	}{
		{"sum", args{"sum"}, ScoringSum},
		{"group", args{"group"}, ScoringGroup},
		{"default", args{"asd"}, ScoringSum},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ScoringFromString(tt.args.str); got != tt.want {
				t.Errorf("ScoringFromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTestset_Testcases(t *testing.T) {
	type fields struct {
		Name   string
		Groups []Group
	}
	tests := []struct {
		name   string
		fields fields
		group  int
		test   int
	}{
		{"test single", fields{Groups: []Group{
			{Testcases: []Testcase{{Index: 1}, {Index: 2}}},
		}}, 0, 1},
		{"test 2 groups", fields{Groups: []Group{
			{Testcases: []Testcase{{Index: 1}, {Index: 2}}},
			{Testcases: []Testcase{{Index: 3}, {Index: 4}}},
		}}, 1, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := Testset{
				Name:   tt.fields.Name,
				Groups: tt.fields.Groups,
			}

			testcases := ts.Testcases()
			testcases[len(testcases)-1].Index = -1
			if idx := ts.Groups[tt.group].Testcases[tt.test].Index; idx != -1 {
				t.Errorf("got %d, want -1", idx)
			}
		})
	}
}

func TestTestset_FirstNonAC(t *testing.T) {
	type fields struct {
		Name   string
		Groups []Group
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "all AC",
			fields: fields{
				Name: "testset",
				Groups: []Group{
					{
						Name:    "subtask1",
						Scoring: ScoringGroup,
						Testcases: []Testcase{
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
						},
						Dependencies: nil,
					},
					{
						Name:    "subtask2",
						Scoring: ScoringGroup,
						Testcases: []Testcase{
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
						},
						Dependencies: nil,
					},
				},
			},
			want: -1,
		},
		{
			name: "WA in first",
			fields: fields{
				Name: "testset",
				Groups: []Group{
					{
						Name:    "subtask1",
						Scoring: ScoringGroup,
						Testcases: []Testcase{
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictWA},
							{VerdictName: VerdictAC},
						},
						Dependencies: nil,
					},
					{
						Name:    "subtask2",
						Scoring: ScoringGroup,
						Testcases: []Testcase{
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
						},
						Dependencies: nil,
					},
				},
			},
			want: 3,
		},
		{
			name: "WA in second",
			fields: fields{
				Name: "testset",
				Groups: []Group{
					{
						Name:    "subtask1",
						Scoring: ScoringGroup,
						Testcases: []Testcase{
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
						},
						Dependencies: nil,
					},
					{
						Name:    "subtask2",
						Scoring: ScoringGroup,
						Testcases: []Testcase{
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictWA},
							{VerdictName: VerdictAC},
						},
						Dependencies: nil,
					},
				},
			},
			want: 7,
		},
		{
			name: "WA in both",
			fields: fields{
				Name: "testset",
				Groups: []Group{
					{
						Name:    "subtask1",
						Scoring: ScoringGroup,
						Testcases: []Testcase{
							{VerdictName: VerdictAC},
							{VerdictName: VerdictWA},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
						},
						Dependencies: nil,
					},
					{
						Name:    "subtask2",
						Scoring: ScoringGroup,
						Testcases: []Testcase{
							{VerdictName: VerdictAC},
							{VerdictName: VerdictAC},
							{VerdictName: VerdictWA},
							{VerdictName: VerdictAC},
						},
						Dependencies: nil,
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ts := Testset{
				Name:   tt.fields.Name,
				Groups: tt.fields.Groups,
			}
			if got := ts.FirstNonAC(); got != tt.want {
				t.Errorf("FirstNonAC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerdictName_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		jsonString string
		want       VerdictName
		wantErr    assert.ErrorAssertionFunc
	}{
		{"compatibility_ac", "0", VerdictAC, assert.NoError},
		{"compatibility_wa", "1", VerdictWA, assert.NoError},
		{"compatibility_re", "2", VerdictRE, assert.NoError},
		{"compatibility_tl", "3", VerdictTL, assert.NoError},
		{"compatibility_ml", "4", VerdictML, assert.NoError},
		{"compatibility_xx", "5", VerdictXX, assert.NoError},
		{"compatibility_dr", "6", VerdictDR, assert.NoError},
		{"compatibility_pc", "7", VerdictPC, assert.NoError},
		{"compatibility_pe", "8", VerdictPE, assert.NoError},
		{"compatibility_null", "null", VerdictUnknown, assert.NoError},
		{"ac", "\"AC\"", VerdictAC, assert.NoError},
		{"wa", "\"WA\"", VerdictWA, assert.NoError},
		{"re", "\"RE\"", VerdictRE, assert.NoError},
		{"tl", "\"TL\"", VerdictTL, assert.NoError},
		{"ml", "\"ML\"", VerdictML, assert.NoError},
		{"xx", "\"XX\"", VerdictXX, assert.NoError},
		{"dr", "\"DR\"", VerdictDR, assert.NoError},
		{"pc", "\"PC\"", VerdictPC, assert.NoError},
		{"pe", "\"PE\"", VerdictPE, assert.NoError},
		{"other", "\"shalal\"", VerdictUnknown, assert.Error},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res VerdictName
			tt.wantErr(t, json.Unmarshal([]byte(tt.jsonString), &res))
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestFeedbackType_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		jsonString string
		want       FeedbackType
		wantErr    assert.ErrorAssertionFunc
	}{
		{"compatibility_cf", "0", FeedbackCF, assert.NoError},
		{"compatibility_ioi", "1", FeedbackIOI, assert.NoError},
		{"compatibility_acm", "2", FeedbackACM, assert.NoError},
		{"compatibility_lazyioi", "3", FeedbackLazyIOI, assert.NoError},
		{"compatibility_null", "null", FeedbackUnknown, assert.NoError},
		{"cf", "\"FeedbackCF\"", FeedbackCF, assert.NoError},
		{"ioi", "\"FeedbackIOI\"", FeedbackIOI, assert.NoError},
		{"acm", "\"FeedbackACM\"", FeedbackACM, assert.NoError},
		{"lazyioi", "\"FeedbackLazyIOI\"", FeedbackLazyIOI, assert.NoError},
		{"other", "\"lol\"", FeedbackUnknown, assert.Error},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res FeedbackType
			tt.wantErr(t, json.Unmarshal([]byte(tt.jsonString), &res))
			assert.Equal(t, tt.want, res)
		})
	}
}

func TestScoringType_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name       string
		jsonString string
		want       ScoringType
		wantErr    assert.ErrorAssertionFunc
	}{
		{"compatibility_group", "0", ScoringGroup, assert.NoError},
		{"compatibility_sum", "1", ScoringSum, assert.NoError},
		{"compatibility_min", "2", ScoringMin, assert.NoError},
		{"compatibility_null", "null", ScoringUnknown, assert.NoError},
		{"group", "\"ScoringGroup\"", ScoringGroup, assert.NoError},
		{"sum", "\"ScoringSum\"", ScoringSum, assert.NoError},
		{"min", "\"ScoringMin\"", ScoringMin, assert.NoError},
		{"other", "\"lol\"", ScoringUnknown, assert.Error},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var res ScoringType
			tt.wantErr(t, json.Unmarshal([]byte(tt.jsonString), &res))
			assert.Equal(t, tt.want, res)
		})
	}
}
