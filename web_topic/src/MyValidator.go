package src

import (
	"gopkg.in/go-playground/validator.v8"
	. "practice_project/web_topic/src/model"
	"reflect"
	"regexp"
)

func TopicUrl(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	_, ok := topStruct.Interface().(*Topics)
	_, oks := topStruct.Interface().(*Topics)
	if ok || oks {
		// fmt.Println(field.String())
		if matched, _ := regexp.MatchString(`^\w{4,10}$`, field.String()); matched {
			return true
		}
		// 正则里， ^开头，$结尾
	}

	return false
}

func TopicsValidate(v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value, field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string) bool {

	topics, ok := topStruct.Interface().(*TopicArray)
	if ok && topics.TopicListSize == len(topics.TopicList) {
		return true
	}

	return false
}
