package blocks

import (
	"context"

	"github.com/hashicorp/hcl/v2/hclsyntax"
)

type Block struct {
	Type       string
	Labels     []string
	Attributes hclsyntax.Attributes
	Body       *hclsyntax.Body
}

type BlockHandler interface {
	Validate(b *Block) error
	Execute(ctx context.Context, b *Block) error
}
