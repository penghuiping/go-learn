/**
必须要在main包里的main方法才是程序入口
*/
package main

import "serialize"

func main() {
	//basic_type.DoActionBase()
	//basic_type.DoActionStrings()
	//basic_type.DoActionTime()

	//collection.ArrayExample()
	//collection.MapExample()
	//collection.SliceExample()
	//
	//_io.FileIOExample()
	//_io.OSArgsExample()
	//_io.StdioExample()
	//
	//object_function.FuncExample()
	//object_function.InterfaceExample()
	//object_function.ReflectExample()
	//object_function.StructExample()

	serialize.JsonExample()

	//tea := object_function.Teacher{30,"小王"}
	//fmt.Println(tea)

	/**
	thread
	 */
	//thread.GoRoutineExample()
	//thread.LockExample()
	//thread.ChannelExample()

	/**
	exception
	 */
	//exception.ErrorExample();
	//exception.ErrorExample1();
	//exception.PanicExample()

	/**
	net
	 */
	 //server := false
	 //if server {
		// nets.TcpServerExample()
	 //}else {
		// nets.TcpClientExample()
	 //}

}
