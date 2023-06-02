// Command gotext-update-templates merge translations and generates a catalog.
//
// Unlike gotext update, it also extracts messages for translation from HTML
// templates. For that purpose it accepts an additional flag "trfunc", which
// defaults to "Tr". It extracts strings from pipelines ".Tr" and "$.Tr" and recurses
// recursive nodes
//
// Templates are read from the first non flag arg. If you use
// go generate, note that "the generator is run in the package's source
// directory".
//
// Original source: https://github.com/dys2p/eco/tree/main/language/gotext-update-templates
package main

import (
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"text/template/parse"

	"golang.org/x/exp/slices"
	"golang.org/x/text/language"
	"golang.org/x/text/message/pipeline"
)

type Config struct {
	Dir               string
	Lang              string
	Out               string
	TemplateDir       string
	Packages          []string
	SrcLang           string
	TranslateFuncName string
}

func main() {
	// own FlagSet because the global one is polluted
	fs := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	dir := fs.String("dir", "locales", "default subdirectory to store translation files")
	lang := fs.String("lang", "en-US", "comma-separated list of languages to process")
	out := fs.String("out", "catalog.go", "output file to write to")
	srcLang := fs.String("srclang", "en-US", "the source-code language")
	trFunc := fs.String("trfunc", "Tr", "name of translate method which is used in templates")
	fs.Parse(os.Args[1:])

	config := Config{
		Dir:               *dir,
		Lang:              *lang,
		Out:               *out,
		TemplateDir:       fs.Args()[0],
		Packages:          fs.Args()[1:],
		SrcLang:           *srcLang,
		TranslateFuncName: *trFunc,
	}
	if err := config.Run(); err != nil {
		log.Fatalln(err)
	}
}

func (config Config) Run() error {
	var templateMessages = []pipeline.Message{}
	err := filepath.Walk(config.TemplateDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if ext := filepath.Ext(info.Name()); ext == ".gohtml" {
			file, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			// similar to parse.Parse but wirh SkipFuncCheck
			trees := make(map[string]*parse.Tree)
			t := parse.New("name")
			t.Mode |= parse.SkipFuncCheck
			if _, err := t.Parse(string(file), "", "", trees); err != nil {
				return err
			}

			var dfs func(parse.Node)
			dfs = func(node parse.Node) {
				if node == nil || reflect.ValueOf(node).IsNil() {
					return
				}

				if node.Type() == parse.NodeAction {
					if actionNode, ok := node.(*parse.ActionNode); ok {
						for _, cmd := range actionNode.Pipe.Cmds {
							if !containsIdentifier(cmd, config.TranslateFuncName) {
								continue
							}
							for _, arg := range cmd.Args {
								if arg.Type() == parse.NodeString {
									if stringNode, ok := arg.(*parse.StringNode); ok {
										text := stringNode.Text
										message := pipeline.Message{
											ID:  pipeline.IDList{text},
											Key: text,
											Message: pipeline.Text{
												Msg: text,
											},
										}
										templateMessages = append(templateMessages, message)
									}
								}
							}
						}
					}
				} else if node.Type() == parse.NodeList {
					if nodeList, ok := node.(*parse.ListNode); ok && nodeList != nil {
						for _, elem := range nodeList.Nodes {
							dfs(elem)
						}
					}
				} else if node.Type() == parse.NodeIf {
					dfs(node.(*parse.IfNode).List)
					dfs(node.(*parse.IfNode).ElseList)
				} else if node.Type() == parse.NodeRange {
					dfs(node.(*parse.RangeNode).List)
					dfs(node.(*parse.RangeNode).ElseList)
				} else if node.Type() == parse.NodeWith {
					dfs(node.(*parse.WithNode).List)
					dfs(node.(*parse.WithNode).ElseList)
				}
			}
			for _, tree := range trees {
				for _, node := range tree.Root.Nodes {
					dfs(node)
				}
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	supported := []language.Tag{}
	for _, l := range strings.FieldsFunc(config.Lang, func(r rune) bool { return r == ',' }) {
		supported = append(supported, language.Make(l))
	}

	pconf := &pipeline.Config{
		Supported:      supported,
		SourceLanguage: language.Make(config.SrcLang),
		Packages:       config.Packages,
		Dir:            config.Dir,
		GenFile:        config.Out,
	}

	// see https://cs.opensource.google/go/x/text/+/master:cmd/gotext/update.go
	state, err := pipeline.Extract(pconf)
	if err != nil {
		return err
	}
	state.Extracted.Messages = append(state.Extracted.Messages, templateMessages...)
	if err := state.Import(); err != nil {
		return err
	}
	if err := state.Merge(); err != nil {
		return err
	}
	if err := state.Export(); err != nil {
		return err
	}
	if err := state.Generate(); err != nil {
		return err
	}
	return nil
}

func containsIdentifier(cmd *parse.CommandNode, identifier string) bool {
	if len(cmd.Args) == 0 {
		return false
	}
	arg := cmd.Args[0]
	var identifiers []string
	switch arg.Type() {
	case parse.NodeField:
		identifiers = arg.(*parse.FieldNode).Ident
	case parse.NodeVariable:
		identifiers = arg.(*parse.VariableNode).Ident
	case parse.NodeIdentifier:
		identifiers = []string{arg.(*parse.IdentifierNode).Ident}
	}
	return slices.Contains(identifiers, identifier)
}
