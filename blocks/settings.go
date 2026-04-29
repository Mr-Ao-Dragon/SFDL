package blocks

import (
	"context"
	"fmt"
)

func init() {
	RegisterBlockType("settings", &SettingsHandler{})
}

type SettingsHandler struct{}

func (h *SettingsHandler) Validate(b *Block) error {
	if len(b.Attributes) == 0 {
		return fmt.Errorf("settings block must contain attributes")
	}
	return nil
}

func (h *SettingsHandler) Execute(_ context.Context, b *Block) error {
	return nil
}
