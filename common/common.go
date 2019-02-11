package common

import "fmt"

func EXPECT(input interface{}, expect interface{}) {
	ErrMsg := ""
	switch valueInput := input.(type) {
	case byte:
		valueExpect, ok := expect.(byte)
		if !ok {
			ErrMsg = "interface type not same"
		} else if valueExpect != valueInput {
			ErrMsg = fmt.Sprintf("Expected %v but Got %v", expect, input)
		}

	case bool:
		valueExpect, ok := expect.(bool)
		if !ok {
			ErrMsg = "interface type not same"
		} else if valueExpect != valueInput {
			ErrMsg = fmt.Sprintf("Expected %v but Got %v", expect, input)
		}
	case int:
		valueExpect, ok := expect.(int)
		if !ok {
			ErrMsg = "interface type not same"
		} else if valueExpect != valueInput {
			ErrMsg = fmt.Sprintf("Expected %v but Got %v", expect, input)
		}
	case int8:
		valueExpect, ok := expect.(int8)
		if !ok {
			ErrMsg = "interface type not same"
		} else if valueExpect != valueInput {
			ErrMsg = fmt.Sprintf("Expected %v but Got %v", expect, input)
		}
	case int16:
		valueExpect, ok := expect.(int16)
		if !ok {
			ErrMsg = "interface type not same"
		} else if valueExpect != valueInput {
			ErrMsg = fmt.Sprintf("Expected %v but Got %v", expect, input)
		}
	case int32:
		valueExpect, ok := expect.(int32)
		if !ok {
			ErrMsg = "interface type not same"
		} else if valueExpect != valueInput {
			ErrMsg = fmt.Sprintf("Expected %v but Got %v", expect, input)
		}
	case int64:
		valueExpect, ok := expect.(int64)
		if !ok {
			ErrMsg = "interface type not same"
		} else if valueExpect != valueInput {
			ErrMsg = fmt.Sprintf("Expected %v but Got %v", expect, input)
		}
	case string:
		valueExpect, ok := expect.(string)
		if !ok {
			ErrMsg = "interface type not same"
		} else if valueExpect != valueInput {
			ErrMsg = fmt.Sprintf("Expected %v but Got %v", expect, input)
		}
	default:
		ErrMsg = "interface type not support"
	}
	if ErrMsg != "" {
		panic(ErrMsg)
	}
}

