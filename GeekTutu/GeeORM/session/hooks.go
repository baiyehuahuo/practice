package session

import (
	"geeorm/log"
	"reflect"
)

const (
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
)

// CallMethod calls the registered hooks
func (s *Session) CallMethod(method string, value interface{}) {
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method) // 得到 Session 的 模型中 名为 method 的结构体方法
	if value != nil {
		fm = reflect.ValueOf(value).MethodByName(method) // 如果有输入 则使用输入类型的名为 method 的结构体方法
	}
	param := []reflect.Value{reflect.ValueOf(s)} // 入参为 session，通过反射传入
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 { // 调用该 method 方法 入参为 session
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}
}
