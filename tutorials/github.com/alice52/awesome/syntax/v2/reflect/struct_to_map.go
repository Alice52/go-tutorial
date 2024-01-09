package reflect

import (
	"errors"
	"reflect"
	"strconv"
)

var (
	// ErrUnsupportedType 表示数据结构中包含不支持的类型。查看 reflectionToInterface 获取支持的类型
	ErrUnsupportedType = errors.New("unsupported type")
)

// StructTagName 是结构体标签的默认名称
const StructTagName = "structtomap"

// Decode 将 Struct 转换成 map[string]interface{}
func Decode(obj interface{}) (map[string]interface{}, error) {
	return DecodeWithTagName(obj, StructTagName)
}

// DecodeWithTagName 支持自定义结构体标签名称
func DecodeWithTagName(obj interface{}, tagName string) (map[string]interface{}, error) {
	v := reflect.ValueOf(obj)
	// 解多层引用
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	// 只支持 Struct
	if v.Kind() != reflect.Struct {
		return nil, ErrUnsupportedType
	}
	if result, err := reflectionToInterface(v, tagName); err != nil {
		return nil, err
	} else {
		return result.(map[string]interface{}), nil
	}
}

// reflectionToInterface 将反射对象转成 interface{}
func reflectionToInterface(v reflect.Value, tagName string) (interface{}, error) {
	// 如果 v 是零值，那么返回相应类型的零值
	if v.IsZero() {
		return reflect.Zero(v.Type()).Interface(), nil
	}
	// 解多层引用
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// 处理基本类型
	switch val := v.Interface().(type) {
	// 将复数转成字符串，其它基本类型直接返回
	case complex64:
		return strconv.FormatComplex(complex128(val), 'g', -1, 64), nil
	case complex128:
		return strconv.FormatComplex(val, 'g', -1, 128), nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr:
		return val, nil
	case float32, float64:
		return val, nil
	case bool:
		return val, nil
	case string:
		return val, nil
	}

	// 处理结构体、切片/数组、映射、接口
	switch v.Kind() {
	case reflect.Struct:
		return structToMap(v, tagName)
	case reflect.Slice, reflect.Array:
		return sliceToSlice(v, tagName)
	case reflect.Map:
		return mapToMap(v, tagName)
	case reflect.Interface:
		if v.IsNil() {
			return nil, nil
		}
		return reflectionToInterface(v.Elem(), tagName)
	}

	return nil, ErrUnsupportedType
}

// mapToMap 将映射转成 map[string]interface{}
func mapToMap(v reflect.Value, tagName string) (map[string]interface{}, error) {
	m := make(map[string]interface{}, 0)
	iter := v.MapRange()
	for iter.Next() {
		key := iter.Key()
		if key.Kind() != reflect.String {
			return nil, ErrUnsupportedType
		}
		if val, err := reflectionToInterface(iter.Value(), tagName); err != nil {
			return nil, err
		} else {
			m[key.Interface().(string)] = val
		}
	}
	return m, nil
}

// sliceToSlice 将切片转成切片
func sliceToSlice(v reflect.Value, tagName string) (interface{}, error) {
	s := make([]interface{}, 0)
	for i := 0; i < v.Len(); i++ {
		if ele, err := reflectionToInterface(v.Index(i), tagName); err != nil {
			return nil, err
		} else {
			s = append(s, ele)
		}
	}
	return s, nil
}

// structToMap 将 Struct 转成 map[string]interface
func structToMap(v reflect.Value, tagName string) (map[string]interface{}, error) {
	t := v.Type()
	m := make(map[string]interface{})
	for i := 0; i < v.NumField(); i++ {
		structField := t.Field(i)
		name := structField.Name
		if tagVal := structField.Tag.Get(tagName); len(tagVal) > 0 {
			name = tagVal
		}
		if val, err := reflectionToInterface(v.Field(i), tagName); err != nil {
			return nil, err
		} else {
			m[name] = val
		}
	}
	return m, nil
}
