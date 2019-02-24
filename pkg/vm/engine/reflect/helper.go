package reflect

import "github.com/YEXINGZHE54/myvm/pkg/utils"

func GetVal(val interface{}) int32 {
	switch v := val.(type) {
	case int32:
		return v
	default:
		utils.Log("expecting int32, got: %T", v)
		return 0
	}
}

func GetLong(val interface{}) int64 {
	switch v := val.(type) {
	case int64:
		return v
	default:
		utils.Log("expecting int64, got: %T", v)
		return 0
	}
}

func GetFloat(val interface{}) float32 {
	switch v := val.(type) {
	case float32:
		return v
	default:
		utils.Log("expecting float32, got: %T", v)
		return 0
	}
}

func GetDouble(val interface{}) float64 {
	switch v := val.(type) {
	case float64:
		return v
	default:
		utils.Log("expecting float64, got: %T", v)
		return 0
	}
}

func GetRef(val interface{}) *Object {
	switch v := val.(type) {
	case *Object:
		return v
	default:
		utils.Log("expecting *Object, got: %T", v)
		return nil
	}
}