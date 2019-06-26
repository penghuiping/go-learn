package exception

import (
	"errors"
	"fmt"
	"testing"
)

//自定义错误
type SelfDefined0Error struct {
	code int
	msg  string
}

//自定义错误
type SelfDefined1Error struct {
	code int
	msg  string
}

//实现error接口
func (err *SelfDefined0Error) Error() string {
	return err.msg + ",code:" + fmt.Sprint(err.code)
}

//实现error接口
func (err *SelfDefined1Error) Error() string {
	return err.msg + ",code:" + fmt.Sprint(err.code)
}

func errorFunc(value int) (string, error) {
	var err error;
	var result string;
	if value < 0 {
		//定义一个错误的方法,这种错误不会中断程序，只是一个错误提示
		err = fmt.Errorf("value比0小");
	} else if value == 0 {
		//定义一个错误的方法,这种错误不会中断程序，只是一个错误提示
		err = errors.New("value等于0")
	} else {
		result = "value大于0,value是" + fmt.Sprint(value);
	}

	return result, err;
}

func error1Func(value int) (string, error) {
	var err error;
	var result string;
	if value < 0 {
		//定义一个错误的方法,这种错误不会中断程序，只是一个错误提示
		err = &SelfDefined0Error{1, "value小于0"}
	} else if value == 0 {
		//定义一个错误的方法,这种错误不会中断程序，只是一个错误提示
		err = &SelfDefined1Error{1, "value大于0"}
	} else {
		result = "value大于0,value是" + fmt.Sprint(value);
	}

	return result, err;
}

func TestErrorExample(in *testing.T) {
	var err error;
	var a string;
	//a, err = errorFunc(1);
	a, err = errorFunc(-1);
	//a, err = errorFunc(0);
	if err == nil {
		fmt.Println(a)
	} else {
		fmt.Println(err);
	}
}

func TestErrorExample1(in *testing.T) {
	var err error;
	var a string;
	//a, err = error1Func(-1);
	a, err = error1Func(0);
	//a, err = error1Func(1);
	if e, ok := err.(*SelfDefined0Error); ok {
		fmt.Println("这是一个SelfDefined0Error类型")
		fmt.Println(e)
		fmt.Println(a)
	}

	if e, ok := err.(*SelfDefined1Error); ok {
		fmt.Println("这是一个SelfDefined1Error类型")
		fmt.Println(e)
		fmt.Println(a)
	}
}
