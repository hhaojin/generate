package Commands

import "flag"

/*
1，初始化CommandSet command map ，
2，加载每一个文件的init 方法注册到map 里面，
3，set.Each ,执行每个命令的Init()方法，设置flag.parse
4，flag.Parse()
5，执行真正的命令
6，存在问题，加入命令行的 init 方法比 CommandSet的init方法先执行，会报错
*/
func Parse() {
	set := NewCommandSet()
	set.Each()
	flag.Parse()
	set.Run()
}
