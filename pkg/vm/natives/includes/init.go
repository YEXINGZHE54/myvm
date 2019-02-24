package includes

import (
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/class"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/object"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/system"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/float"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/double"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/string"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/sun/misc/vm"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/sun/misc/unsafe"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/reflect"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/throwable"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/security"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/java/io"
	_ "github.com/YEXINGZHE54/myvm/pkg/vm/natives/java/lang"
)