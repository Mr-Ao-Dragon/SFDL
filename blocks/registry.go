package blocks

import (
	"fmt"
	"maps"
	"slices"
)

var blockRegistry = make(map[string]BlockHandler)

func RegisterBlockType(blockType string, handler BlockHandler) {
	blockRegistry[blockType] = handler
}

func GetRegisteredBlockTypes() []string {
	keys := slices.Collect(maps.Keys(blockRegistry))
	slices.Sort(keys)
	return keys
}

func ProcessBlock(b *Block) error {
	handler, ok := blockRegistry[b.Type]
	if !ok {
		return fmt.Errorf("unsupported block type: %s", b.Type)
	}
	if err := handler.Validate(b); err != nil {
		return fmt.Errorf("block %s validate failed: %w", b.Type, err)
	}
	return nil
}
