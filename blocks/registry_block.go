package blocks

import (
	"context"
	"fmt"
)

func init() {
	RegisterBlockType("registry", &RegistryHandler{})
}

type RegistryHandler struct{}

func (h *RegistryHandler) Validate(b *Block) error {
	if len(b.Labels) == 0 {
		return fmt.Errorf("registry block must have a label")
	}
	return nil
}

func (h *RegistryHandler) Execute(_ context.Context, b *Block) error {
	return nil
}
