# go-code-generate

#### 介绍
一个go代码生成工具，可以生成，三层架构代码，接口实现类，数据库模型


#### 安装
```
go get gitee.com/hhaojin/generate
```

#### 使用说明

1.  查看版本
```
generate -v

```

2.  数据库模型生成
```
generate db -m TableName -d app/models
会连接当前目录下app.yaml里面配置的数据库，查找表名并生成相应的实体在 app/models/目录下

-m 指定表名称
-d 可选，生成文件的目录，默认app/models目录下
注：
需要在项目根目录下添加文件app.yaml，内容：
db:
  driver: mysql
  dsn: root:123456@tcp(127.0.0.1:3306)/study?charset=utf8&parseTime=True&loc=Local
```

3.  接口实现类生成
```
generate service -c -i app/api/ICourseService -d app/service/course
会去查找 app/api/ICourseService.go文件，然后生成在 app/service/course 目录下

-c 生成文件
-i 所实现的接口
-d 可选，生成文件的目录，默认app/service目录下

```

4.  三层架构代码生成
```
generate lib -c app/lib/course -i ICourseService -r /test -m POST
生成三层架构代码到 app/lib/course 目录下

-c 可选，指定生成的目录名，默认app/lib 下面
-i 指定接口名称
-r 路由uri
-m 可选，请求方式，默认GET

```
