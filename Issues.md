# Issues

## Template File

1. json or yaml -> protobuf map
2. 

## Must todo

1. 定义数据结构 OK
2. 添加组合的业务逻辑 OPT

## How to do

1. API in/out

```json
get /v1/note/1

{"id":1,"note":"123213"}

post /v1/hote -body {"note":"123213"}
{"id":2,"note":"123213"}

put /v1/hote/2 -body {"id":2, "note":"123214"}
{"id":2,"note":"123214"}

delete /v1/hote/2 
--
```

2. create db schema 
# code -> model -> db
```go
[]fields{

}
```
# db -> model -> code

3. db -> common api -> 如果有文件不生成, 否则生成模版

- biz
- service

```bash
make db2api model=note template=../db2api/onlydb.tpl
+ server x
+ repo @
+ usercase @
+ biz x
+ entity @
+ protobuffer

make db2api model=note template=../db2api/onlydb.tpl force
# 会重新生成这些
```

4. api -> doc

```bash
make doc

+ swagger/ui
```


