package dtype

import (
	"fmt"
	"reflect"
	"strconv"
	"runtime"
	"regexp"
)

func Int(val string) int {
	value, _ := strconv.Atoi(val)
	return value
}

func Str(val int) string {
	return strconv.FormatInt(int64(val), 10)
}

func ShowType(val interface{}) {
	fmt.Println(reflect.TypeOf(val))
}

// func InterfaceToOther(value interface{}){
// 		switch value.(type) {
// 		case string:
// 		 // 将interface转为string字符串类型
// 		 op, ok := value.(string)
// 		 fmt.Println(op, ok)
// 		case int32:
// 		 // 将interface转为int32类型
// 		 op, ok := value.(int32)
// 		 fmt.Println(op, ok)
// 		case int64:
// 		 // 将interface转为int64类型
// 		 op, ok := value.(int64)
// 		 fmt.Println(op, ok)
// 		case User:
// 		 // 将interface转为User struct类型，并使用其Name对象
// 		 op, ok := value.(User)
// 		 fmt.Println(op.Name, ok)
// 		case []int:
// 		 // 将interface转为切片类型
// 		 op := make([]int, 0)
// 		 op = value.([]int)
// 		 fmt.Println(op)
// 		default:
// 		 fmt.Println("unknown")
		
// }


/*****************************************************************
*  @brief  GetMethodName 获取外层方法的名称
*  @param  me None
*  @retval None
******************************************************************/
func GetMethodName(me interface{}) string {

	str:=  runtime.FuncForPC(reflect.ValueOf(me).Pointer()).Name()
	rg := regexp.MustCompile(`\.(.*)`)
	name := rg.FindAllStringSubmatch(str,-1)[0][1]
	
    return name



}
