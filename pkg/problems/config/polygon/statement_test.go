package polygon_test

import (
	"github.com/mraron/njudge/pkg/problems/config/polygon"
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestParseJSONStatement(t *testing.T) {
	r := io.NopCloser(strings.NewReader(`{"legend": null}`))
	res, err := polygon.NewJSONStatement(r, "hungarian")
	assert.NoError(t, err)
	assert.Nil(t, res.Legend)
}

func TestJSONStatement_Html(t *testing.T) {

	tests := []struct {
		name          string
		jsonStatement polygon.JSONStatement
		want          string
		wantErr       bool
	}{
		{name: "empty", jsonStatement: polygon.JSONStatement{
			Locale:      "hungarian",
			Name:        "teszt",
			TimeLimit:   0,
			MemoryLimit: 0, InputFile: "stdin", OutputFile: "stdout"},
			want: `<link href="problem-statement.css" rel="stylesheet" type="text/css"><div class="problem-statement">
<div class="header">
	<div class="title">teszt</div>
	<div class="time-limit"><div class="property-title">tesztenkénti időlimit</div> 0 ms</DIV>
	<div class="memory-limit"><div class="property-title">tesztenkénti memórialimit</div> 0 MiB</div>
	<div class="input-file"><div class="property-title">inputfájl</div>  stdin </div>
	<div class="output-file"><div class="property-title">outputfájl</div>  stdout </div>
</div><p></p><p></p>


</div>`, wantErr: false},
	}
	remSpaces := func(s string) string {
		return strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\n", "")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.jsonStatement.Html()
			assert.NoError(t, err)
			assert.Equal(t, remSpaces(tt.want), remSpaces(string(got)))
		})
	}
}
