package function

import (
	"path"
	"reflect"
	"runtime"
)

// New constructs and returns a pointer to a Value struct. If the given value is
// not a function nil is returned
func New(a any) *Value {
	v := reflect.ValueOf(a)
	if v.Kind() != reflect.Func {
		return nil
	}
	return &Value{v}
}

// Value represents a Go function
type Value struct {
	v reflect.Value
}

// Call invokes this function with the given arguments and returns the result of
// the invocation. The return value can be one of the following:
//   - nil, if it does not return a value
//   - the single return value of this function, if it only returns one value
//   - a slice containing the multiple return values of this function, if it
//     returns multiple values
//
// Call panics if the number or types of arguments provided do not match this
// function's signature
func (v Value) Call(arguments ...any) any {
	var vv []reflect.Value
	for i := range arguments {
		i := arguments[i]
		vv = append(
			vv,
			reflect.ValueOf(i),
		)
	}
	return interfacesOf(v.v.Call(vv)...)
}

// Name returns the name of this function
func (v Value) Name() string {
	s := runtime.
		FuncForPC(v.v.Pointer()).
		Name()
	return path.Base(s)
}

func interfacesOf(vv ...reflect.Value) any {
	var ii []any
	for i := range vv {
		v := vv[i]
		ii = append(ii, v.Interface())
	}
	if len(ii) == 1 {
		return ii[0]
	}
	return ii
}
