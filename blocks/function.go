package blocks

import (
	"context"
	"fmt"
)

func init() {
	RegisterBlockType("function", &FunctionHandler{})
}

type FunctionHandler struct{}

func (h *FunctionHandler) Validate(b *Block) error {
	if len(b.Labels) == 0 {
		return fmt.Errorf("function block must have a label")
	}
	return nil
}

func (h *FunctionHandler) Execute(_ context.Context, b *Block) error {
	return nil
}
