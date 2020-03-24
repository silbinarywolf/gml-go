package assetidgen

import (
	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
)

func lazyLoad() {
	if node == nil {
		var err error
		node, err = snowflake.NewNode(0)
		if err != nil {
			panic("error creating NewNode: " + err.Error())
		}
	}
}

// NewID will generate a new unique distributed ID for an asset.
// This should only be used by debug code.
func NewID() uint64 {
	lazyLoad()
	return uint64(node.Generate())
}
