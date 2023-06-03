package problems

import (
	"testing"
	"time"
)

func TestFeedbackFromString(t *testing.T) {
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
			if got := FeedbackFromString(tt.args.str); got != tt.want {
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
		want   int
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
		{"scoring group ac", fields{Scoring: ScoringGroup, Testcases: []Testcase{{Score: 0.1, MaxScore: 0.5}, {Score: 0.5, MaxScore: 0.5}}}, 0.6},
		{"scoring group pc", fields{Scoring: ScoringGroup, Testcases: []Testcase{{VerdictName: VerdictPC, Score: 0.1, MaxScore: 0.5}, {Score: 0.5, MaxScore: 0.5}}}, 0.6},
		{"scoring group wa", fields{Scoring: ScoringGroup, Testcases: []Testcase{{VerdictName: VerdictWA, Score: 0.0, MaxScore: 0.5}, {Score: 0.5, MaxScore: 0.5}}}, 0},
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
