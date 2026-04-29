package sfdl

import (
	"context"
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclsyntax"
	sfdlblocks "github.com/lightDproject/SFDL/blocks"
)

type Config struct {
	Filename string
	Content  []byte
}

type File struct {
	*hcl.File
}

func (f *File) ProcessBlocks(ctx context.Context) error {
	body, err := f.SyntaxBody()
	if err != nil {
		return err
	}
	for _, hclBlock := range body.Blocks {
		b := &sfdlblocks.Block{
			Type:       hclBlock.Type,
			Labels:     hclBlock.Labels,
			Attributes: hclBlock.Body.Attributes,
			Body:       hclBlock.Body,
		}
		if err := sfdlblocks.ProcessBlock(b); err != nil {
			return err
		}
	}
	return nil
}

func GetRegisteredBlockTypes() []string {
	return sfdlblocks.GetRegisteredBlockTypes()
}

type Parser struct {
	filename string
	content  []byte
	diags    hcl.Diagnostics
}

func Parse(config Config) (*File, error) {
	p := NewParser(config.Filename, config.Content)
	file := p.Parse()
	if p.Errs().HasErrors() {
		return nil, p.Errs()
	}
	return file, nil
}

func NewParser(filename string, content []byte) *Parser {
	return &Parser{
		filename: filename,
		content:  content,
	}
}

func (p *Parser) Parse() *File {
	file, diags := hclsyntax.ParseConfig(p.content, p.filename, hcl.InitialPos)
	p.diags = diags
	if diags.HasErrors() {
		return nil
	}
	return &File{File: file}
}

func (p *Parser) Errs() hcl.Diagnostics {
	return p.diags
}

func (f *File) SyntaxBody() (*hclsyntax.Body, error) {
	body, ok := f.Body.(*hclsyntax.Body)
	if !ok {
		return nil, fmt.Errorf("body is not hclsyntax.Body")
	}
	return body, nil
}

func (f *File) MustSyntaxBody() *hclsyntax.Body {
	body, err := f.SyntaxBody()
	if err != nil {
		panic(err)
	}
	return body
}
