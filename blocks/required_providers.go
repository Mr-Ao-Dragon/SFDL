package blocks

import (
	"context"
	"fmt"
)

func init() {
	RegisterBlockType("required_providers", &RequiredProvidersHandler{})
}

type RequiredProvidersHandler struct{}

func (h *RequiredProvidersHandler) Validate(b *Block) error {
	if len(b.Attributes) == 0 {
		return fmt.Errorf("required_providers block must contain at least one provider")
	}
	for name := range b.Attributes {
		if name == "" {
			return fmt.Errorf("provider name cannot be empty")
		}
	}
	return nil
}

func (h *RequiredProvidersHandler) Execute(_ context.Context, b *Block) error {
	return nil
}
