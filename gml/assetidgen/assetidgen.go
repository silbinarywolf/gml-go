package assetidgen

import (
	"github.com/bwmarrin/snowflake"
)

var (
	node *snowflake.Node
)

func lazyLoad() {
	var err error
	node, err = snowflake.NewNode(0)
	if err != nil {
		panic("error creating NewNode: " + err.Error())
	}
}

func NewID() uint64 {
	lazyLoad()
	return uint64(node.Generate())
}
