package convert

import (
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions"
	"github.com/YEXINGZHE54/myvm/pkg/vm/engine/instructions/constants"
)

const (
	i2c_op = 0x92
)

func init()  {
	instructions.Register(i2c_op, &constants.NoopInst{})
}