package main

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type Course struct {
	CourseName string  `json:"course_name"`
	Score      float32 `json:"score"`
}

type CustomUint64 uint64

type Student struct {
	Name     string       `json:"name"`
	Age      CustomUint64 `json:"age"`
	Gender   bool         `json:"gender"`
	AvgScore float64      `json:"avg_score"`
	Course   []Course     `json:"course"`
}

func (t *CustomUint64) Error(err error) {

}

func (t *CustomUint64) UnmarshalJSON(b []byte) error {
	var data interface{}
	err := json.Unmarshal(b, &data)
	if err != nil {
		return err
	}
	switch data.(type) {
	case string:
		d, err := strconv.Atoi(data.(string))
		if err != nil {
			*t = 0
		}
		*t = CustomUint64(d)
	case float64 :
		*t = CustomUint64(data.(float64))
	case bool:
		if data.(bool) {
			*t = 1
		} else {
			*t = 0
		}
	default:
		*t = 0
	}
	return nil
}

func main() {
	//c:=[]Course{}
	//student := Student{
	//	Name:     "vega",
	//	Age:      27,
	//	Gender:   true,
	//	AvgScore: 99.99,
	//	Course:   []Course{{
	//		CourseName: "chinese",
	//		Score:      100,
	//	},{
	//		CourseName: "English",
	//		Score: 99,
	//	}},
	//}
	//marshalResult,err:=json.Marshal(student)
	//if err != nil {
	//	log.Error(err)
	//	return
	//}
	testData := make([]string, 3)
	testData[0] = `{"name":"vega","age":123,"gender":true,"avg_score":99.99,"course":[{"course_name":"chinese","score":100},{"course_name":"English","score":99}]}`
	var unMarshalResult Student
	//m := make(map[string]interface{})
	err := json.Unmarshal([]byte(testData[0]), &unMarshalResult)
	if err != nil {
		fmt.Println(err)
	}
	//for k, v := range m {
	//fmt.Printf("key: %s,val: %v,typeOfVal: %s\n", k, v, reflect.TypeOf(v).Kind())
	//}
	fmt.Println(unMarshalResult.Age)
	//fmt.Println(m)

}
