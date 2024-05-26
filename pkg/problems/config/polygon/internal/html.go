package internal

import (
	"embed"
	"encoding/base64"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/spf13/afero"
	"io"
	"mime"
	"path/filepath"
)

//go:embed problem-statement.css
var problemStatementCssEmbed embed.FS

const problemStatementCss = "problem-statement.css"

var (
	dataFormat  = `data:%s;charset=utf-8;base64,%s`
	styleFormat = `<style type="text/css">
%s
</style>`
)

// InlineHTML inlines images and link stylesheets from the incoming html document r and writes
// the result to w, which should be a standalone document equivalent to r.
// The referenced css and images files should be relative to the given fs.
func InlineHTML(fs afero.Fs, r io.Reader, w io.Writer) error {
	var (
		doc *goquery.Document
		err error
	)
	doc, err = goquery.NewDocumentFromReader(r)
	if err != nil {
		return err
	}
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		if err != nil {
			return
		}
		var src []byte
		srcAttr, exists := s.Attr("src")
		if !exists {
			return
		}
		src, err = afero.ReadFile(fs, srcAttr)
		if err != nil {
			return
		}
		s.SetAttr("src", fmt.Sprintf(dataFormat, mime.TypeByExtension(filepath.Ext(srcAttr)), base64.StdEncoding.EncodeToString(src)))
	})
	if err != nil {
		return err
	}
	doc.Find("link").Each(func(i int, s *goquery.Selection) {
		if err != nil {
			return
		}
		if s.AttrOr("rel", "") != "stylesheet" {
			return
		}
		var src []byte
		href, exists := s.Attr("href")
		if !exists {
			return
		}
		src, err = afero.ReadFile(fs, href)
		if err != nil {
			if href == problemStatementCss {
				src, err = problemStatementCssEmbed.ReadFile(problemStatementCss)
				if err != nil {
					return
				}
			} else {
				return
			}
		}
		style := fmt.Sprintf(styleFormat, string(src))
		s.ReplaceWithHtml(style)
	})
	if err != nil {
		return err
	}
	resHead, err := doc.Find("head").Html()
	if err != nil {
		return err
	}
	resBody, err := doc.Find("body").Html()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(resHead))
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(resBody))
	return err
}
