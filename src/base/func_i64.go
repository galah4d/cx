package base

import (
	"fmt"
	"errors"
	"time"
	"math/rand"
	"github.com/skycoin/skycoin/src/cipher/encoder"
)

func addI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.add", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		output := encoder.SerializeAtomic(int64(num1 + num2))

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func subI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.sub", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		output := encoder.SerializeAtomic(int64(num1 - num2))

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func mulI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64mul", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		output := encoder.SerializeAtomic(int64(num1 * num2))

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func divI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.div", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		if num2 == int64(0) {
			return errors.New("divI64: Division by 0")
		}
		
		output := encoder.SerializeAtomic(int64(num1 / num2))

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func modI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.mod", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		if num2 == int64(0) {
			return errors.New("modI64: Division by 0")
		}

		output := encoder.Serialize(int64(num1 % num2))

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func andI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.and", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		output := encoder.Serialize(int64(num1 & num2))

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func orI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.or", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		output := encoder.Serialize(int64(num1 | num2))

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func xorI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.xor", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		output := encoder.Serialize(int64(num1 ^ num2))

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func andNotI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.bitclear", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		output := encoder.Serialize(int64(num1 &^ num2))

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func randI64 (min *CXArgument, max *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.rand", "i64", "i64", min, max); err == nil {
		var minimum int64
		encoder.DeserializeRaw(*min.Value, &minimum)

		var maximum int64
		encoder.DeserializeRaw(*max.Value, &maximum)

		if minimum > maximum {
			return errors.New(fmt.Sprintf("randI64: min must be less than max (%d !< %d)", minimum, maximum))
		}

		rand.Seed(time.Now().UTC().UnixNano())
		output := encoder.SerializeAtomic(int32(rand.Intn(int(maximum - minimum)) + int(minimum)))

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func readI64A (arr *CXArgument, idx *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("[]i64.read", "[]i64", "i32", arr, idx); err == nil {
		var index int32
		encoder.DeserializeRaw(*idx.Value, &index)

		var size int32
		encoder.DeserializeAtomic((*arr.Value)[0:4], &size)

		if index < 0 {
			return errors.New(fmt.Sprintf("readI64A: negative index %d", index))
		}

		if index >= size {
			return errors.New(fmt.Sprintf("readI64A: index %d exceeds array of length %d", index, size))
		}

		var value int64
		encoder.DeserializeRaw((*arr.Value)[(index+1)*4:(index+2)*4], &value)
		output := encoder.Serialize(value)

		assignOutput(&output, "i64", expr, call)
		return nil
	} else {
		return err
	}
}

func writeI64A (arr *CXArgument, idx *CXArgument, val *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkThreeTypes("[]i64.write", "[]i64", "i32", "i64", arr, idx, val); err == nil {
		var index int32
		encoder.DeserializeRaw(*idx.Value, &index)

		var size int32
		encoder.DeserializeAtomic((*arr.Value)[0:4], &size)

		if index < 0 {
			return errors.New(fmt.Sprintf("writeI64A: negative index %d", index))
		}

		if index >= size {
			return errors.New(fmt.Sprintf("writeI64A: index %d exceeds array of length %d", index, size))
		}

		i := (int(index)+1)*4
		for c := 0; c < 4; c++ {
			(*arr.Value)[i + c] = (*val.Value)[c]
		}
		
		return nil
	} else {
		return err
	}
}

func lenI64A (arr *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkType("[]i64.len", "[]i64", arr); err == nil {
		var array []int64
		encoder.DeserializeRaw(*arr.Value, &array)

		output := encoder.SerializeAtomic(int32(len(array)))

		assignOutput(&output, "i32", expr, call)
		return nil
	} else {
		return err
	}
}

func ltI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.lt", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		var val []byte

		if num1 < num2 {
			val = encoder.Serialize(int32(1))
		} else {
			val = encoder.Serialize(int32(0))
		}

		assignOutput(&val, "bool", expr, call)
		return nil
	} else {
		return err
	}
}

func gtI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.gt", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		var val []byte

		if num1 > num2 {
			val = encoder.Serialize(int32(1))
		} else {
			val = encoder.Serialize(int32(0))
		}

		assignOutput(&val, "bool", expr, call)
		return nil
	} else {
		return err
	}
}

func eqI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.eq", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		var val []byte

		if num1 == num2 {
			val = encoder.Serialize(int32(1))
		} else {
			val = encoder.Serialize(int32(0))
		}

		assignOutput(&val, "bool", expr, call)
		return nil
	} else {
		return err
	}
}

func lteqI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.lteq", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		var val []byte

		if num1 <= num2 {
			val = encoder.Serialize(int32(1))
		} else {
			val = encoder.Serialize(int32(0))
		}

		assignOutput(&val, "bool", expr, call)
		return nil
	} else {
		return err
	}
}

func gteqI64 (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("i64.gteq", "i64", "i64", arg1, arg2); err == nil {
		var num1 int64
		var num2 int64
		encoder.DeserializeRaw(*arg1.Value, &num1)
		encoder.DeserializeRaw(*arg2.Value, &num2)

		var val []byte

		if num1 >= num2 {
			val = encoder.Serialize(int32(1))
		} else {
			val = encoder.Serialize(int32(0))
		}

		assignOutput(&val, "bool", expr, call)
		return nil
	} else {
		return err
	}
}

func concatI64A (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("[]i64.concat", "[]i64", "[]i64", arg1, arg2); err == nil {
		var slice1 []int64
		var slice2 []int64
		encoder.DeserializeRaw(*arg1.Value, &slice1)
		encoder.DeserializeRaw(*arg2.Value, &slice2)

		output := append(slice1, slice2...)
		sOutput := encoder.Serialize(output)

		assignOutput(&sOutput, "[]i64", expr, call)
		return nil
	} else {
		return err
	}
}

func appendI64A (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("[]i64.append", "[]i64", "i64", arg1, arg2); err == nil {
		var slice []int64
		var literal int64
		encoder.DeserializeRaw(*arg1.Value, &slice)
		encoder.DeserializeRaw(*arg2.Value, &literal)

		output := append(slice, literal)
		sOutput := encoder.Serialize(output)

		assignOutput(&sOutput, "[]i64", expr, call)
		return nil
	} else {
		return err
	}
}

func copyI64A (arg1 *CXArgument, arg2 *CXArgument, expr *CXExpression, call *CXCall) error {
	if err := checkTwoTypes("[]i64.copy", "[]i64", "[]i64", arg1, arg2); err == nil {
		var slice1 []int32
		var slice2 []int32
		encoder.DeserializeRaw(*arg1.Value, &slice1)
		encoder.DeserializeRaw(*arg2.Value, &slice2)

		copy(slice1, slice2)
		sOutput := encoder.Serialize(slice1)

		*arg1.Value = sOutput
		return nil
	} else {
		return err
	}
}