package customStruct

import (
	"encoding/json"
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"strconv"
	"strings"
	"testing"
)

func BenchmarkCustomUint64_UnmarshalJSON(b *testing.B) {
	testData := make([]string, 7)
	testData[0] = `{"age":123,"a":"a"}`
	testData[1] = `{"age":"12"}`
	testData[2] = `{"age":false}`
	testData[3] = `{"age":true}`
	testData[4] = `{"age":null}`
	testData[5] = `{"age":60.5}`
	testData[6] = `{"age":"60.5"}`
	type D struct {
		//Age CustomUint64 `json:"age" validate:"gt=0,lt=101"`
		Age CustomUint64 `json:"age" validate:"has=123#12,lt=200"`
		//Age int `json:"age" validate:"gte=0,lte=90"`
	}
	var validate *validator.Validate
	uni := ut.New(zh.New())
	trans, _ := uni.GetTranslator("zh")
	validate = validator.New()
	err := validate.RegisterValidation("has", ValidateCustomUint64)
	if err != nil {
		b.Error(err)
	}
	//validate := validator.New()
	//验证器注册翻译器
	err = zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < 1; i++ {
		for index, test := range testData {
			data := D{}
			err := json.Unmarshal([]byte(test), &data)
			if err != nil {
				b.Error(err)
			}
			err = validate.Struct(data)
			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					b.Error(err.Translate(trans)) //Age必须大于18
				}
				//fmt.Println(err.Translate(trans))//Age必须大于18
			}
			b.Logf("index: %d,val: %v", index, data)
		}
	}
}
func ValidateCustomUint64(fl validator.FieldLevel) bool {
	val := fl.Field().Uint()
	param := fl.Param()
	params := strings.Split(param, "#")
	for _, p := range params {
		if p == strconv.FormatUint(val, 10) {
			return true
		}
	}
	return false
}
