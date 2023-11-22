package snowflake

import "github.com/bwmarrin/snowflake"

type Snowflake struct {
	node *snowflake.Node
}

func NewSnowflake() (*Snowflake, error) {
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}
	return &Snowflake{node: node}, nil
}

func (s *Snowflake) GenId() snowflake.ID {
	return s.node.Generate()
}
