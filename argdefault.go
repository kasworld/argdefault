// Copyright 2015,2016,2017,2018,2019 SeukWon Kang (kasworld@gmail.com)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package argdefault

import (
	"flag"
	"fmt"
	"reflect"
	"strconv"
)

// SetDefault set inObj default value if field is zero value
func SetZeroFieldToDefault(inObj interface{}) interface{} {
	structValue := reflect.ValueOf(inObj).Elem()
	numField := structValue.NumField()
	for i := 0; i < numField; i++ {
		structField := structValue.Field(i)
		fieldTag := structValue.Type().Field(i).Tag
		defaultVal, defaultExist := fieldTag.Lookup("default")
		// is field has some value, not overwrite with default
		if defaultExist && structField.IsZero() {
			err := setFieldDefault(structField, defaultVal)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
	}
	return inObj
}

func setFieldDefault(structField reflect.Value, defaultVal string) error {
	switch structField.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := strconv.Atoi(defaultVal)
		if err != nil {
			return fmt.Errorf("invalid int field %v %v", defaultVal, err)
		} else {
			structField.SetInt(int64(v))
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		v, err := strconv.Atoi(defaultVal)
		if err != nil {
			return fmt.Errorf("invalid int field %v %v", defaultVal, err)
		} else {
			structField.SetUint(uint64(v))
		}

	case reflect.Float64, reflect.Float32:
		v, err := strconv.ParseFloat(defaultVal, 64)
		if err != nil {
			return fmt.Errorf("invalid int field %v %v", defaultVal, err)
		} else {
			structField.SetFloat(v)
		}

	case reflect.Bool:
		v, err := strconv.ParseBool(defaultVal)
		if err != nil {
			return fmt.Errorf("invalid int field %v %v", defaultVal, err)
		} else {
			structField.SetBool(v)
		}

	case reflect.String:
		structField.SetString(defaultVal)

	default:
		return fmt.Errorf("unprocessed %v", structField)
	}
	return nil
}

// ArgStatue remember arg default map[fieldname]value
type ArgStatue struct {
	intArg    map[string]*int
	uintArg   map[string]*uint
	floatArg  map[string]*float64
	boolArg   map[string]*bool
	stringArg map[string]*string
}

// AddArgs set flag by inobj val default
func AddArgsWith(defaultObj interface{}) *ArgStatue {
	as := &ArgStatue{
		intArg:    make(map[string]*int),
		uintArg:   make(map[string]*uint),
		floatArg:  make(map[string]*float64),
		boolArg:   make(map[string]*bool),
		stringArg: make(map[string]*string),
	}
	structValue := reflect.ValueOf(defaultObj).Elem()
	for i := 0; i < structValue.NumField(); i++ {
		fieldName := structValue.Type().Field(i).Name
		structField := structValue.Field(i)
		fieldTag := structValue.Type().Field(i).Tag

		argname, argexist := fieldTag.Lookup("argname")
		// is field has some value, not overwrite with default
		if argexist {
			if argname == "" {
				argname = fieldName
			}
			switch structField.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				var argv int
				as.intArg[fieldName] = &argv
				flag.IntVar(&argv, argname, int(structField.Int()), fieldName)
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				var argv uint
				as.uintArg[fieldName] = &argv
				flag.UintVar(&argv, argname, uint(structField.Uint()), fieldName)
			case reflect.Float64, reflect.Float32:
				var argv float64
				as.floatArg[fieldName] = &argv
				flag.Float64Var(&argv, argname, structField.Float(), fieldName)
			case reflect.Bool:
				var argv bool
				as.boolArg[fieldName] = &argv
				flag.BoolVar(&argv, argname, structField.Bool(), fieldName)
			case reflect.String:
				var argv string
				as.stringArg[fieldName] = &argv
				flag.StringVar(&argv, argname, structField.String(), fieldName)
			default:
			}
		}
	}
	return as
}

// ApplyArgs return inObj with flag value
func (as *ArgStatue) ApplyArgsTo(inObj interface{}) interface{} {
	destStructValue := reflect.ValueOf(inObj).Elem()
	for i := 0; i < destStructValue.NumField(); i++ {
		fieldName := destStructValue.Type().Field(i).Name
		structField := destStructValue.Field(i)
		fieldTag := destStructValue.Type().Field(i).Tag
		destObjField := destStructValue.Field(i)

		_, argexist := fieldTag.Lookup("argname")
		if argexist {
			switch structField.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				if v, exist := as.intArg[fieldName]; exist && *v != int(structField.Int()) {
					destObjField.SetInt(int64(*v))
				}
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
				if v, exist := as.uintArg[fieldName]; exist && *v != uint(structField.Uint()) {
					destObjField.SetUint(uint64(*v))
				}
			case reflect.Float64, reflect.Float32:
				if v, exist := as.floatArg[fieldName]; exist && *v != structField.Float() {
					destObjField.SetFloat(*v)
				}
			case reflect.Bool:
				if v, exist := as.boolArg[fieldName]; exist && *v != structField.Bool() {
					destObjField.SetBool(*v)
				}
			case reflect.String:
				if v, exist := as.stringArg[fieldName]; exist && *v != structField.String() {
					destObjField.SetString(*v)
				}
			default:
			}
		}
	}
	return inObj
}
