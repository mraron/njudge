package internal

import (
	"bytes"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestInlineHTML(t *testing.T) {
	fs := afero.NewMemMapFs()
	_ = afero.WriteFile(fs, "test.jpg", []byte("lul"), 0777)
	res := bytes.NewBuffer(nil)
	assert.NoError(t, InlineHTML(fs, strings.NewReader(`<img src="test.jpg">`), res))
	assert.Equal(t, `<img src="data:image/jpeg;charset=utf-8;base64,bHVs"/>`, res.String())

	_ = afero.WriteFile(fs, "test.css", []byte("cool"), 0777)
	res.Reset()
	assert.NoError(t, InlineHTML(fs, strings.NewReader(`<link rel="stylesheet" href="test.css">`), res))
	assert.Equal(t, `<style type="text/css">
cool
</style>`, res.String())

	res.Reset()
	assert.NoError(t, InlineHTML(fs, strings.NewReader(`<link rel="stylesheet" href="test.css"><img src="test.jpg">`), res))
	assert.Equal(t, `<style type="text/css">
cool
</style><img src="data:image/jpeg;charset=utf-8;base64,bHVs"/>`, res.String())

	res.Reset()
	assert.NoError(t, InlineHTML(fs, strings.NewReader(`<img src="test.jpg"><link rel="stylesheet" href="test.css">`), res))
	assert.Equal(t, `<img src="data:image/jpeg;charset=utf-8;base64,bHVs"/><style type="text/css">
cool
</style>`, res.String())

	_ = afero.WriteFile(fs, "problem-statement.css", []byte(`#ads {
	basd: streter;
}`), 0777)
	res.Reset()
	assert.NoError(t, InlineHTML(fs, strings.NewReader(`<link href="problem-statement.css" rel="stylesheet" type="text/css"/><div></div>`), res))
	assert.Equal(t, `<style type="text/css">
#ads {
	basd: streter;
}
</style><div></div>`, res.String())
}
