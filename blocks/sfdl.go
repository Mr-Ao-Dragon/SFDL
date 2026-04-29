package blocks

import (
	"context"
	"fmt"
)

func init() {
	RegisterBlockType("SFDL", &SFDLHandler{})
}

type SFDLHandler struct{}

func (h *SFDLHandler) Validate(b *Block) error {
	found := false
	for _, block := range b.Body.Blocks {
		if block.Type == "required_providers" {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("SFDL block must contain required_providers block")
	}
	return nil
}

func (h *SFDLHandler) Execute(_ context.Context, b *Block) error {
	return nil
}
