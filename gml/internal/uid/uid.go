package uid

import "github.com/bwmarrin/snowflake"

var (
	node *snowflake.Node
)

const (
	// NOTE(Jake): 2019-09-08
	// Chose the last number from the end so
	// custom asset libraries can generate on a different ID
	// if they also use Snowflake for generating int64's.
	internalSnowflakeNodeId = 1023
)

func lazyInit() {
	if node == nil {
		var err error
		node, err = snowflake.NewNode(internalSnowflakeNodeId)
		if err != nil {
			panic(err)
		}
	}
}

func GenerateUniqueID() int64 {
	lazyInit()
	return node.Generate().Int64()
}
