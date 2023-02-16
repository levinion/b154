# base114514

## 介绍
114514恶臭编码

```txt
输入：你好
一次编码：1145145414411154545454144144415444
解码：你好
```

## 下载
使用包
```go
go get github.com/levinion/b154/coding
```

使用命令行程序对文件编解码
```go
go install github.com/levinion/b154@latest
```

## 用法
包的使用
```go
func EncodeToString[T []rune | string | []byte](str T) (r string)
func EncodeToBytes[T []rune | string | []byte](str T) (r []byte)
func DecodeToString[T string | []byte | []rune](str T) (r string)
func DecodeToString[T string | []byte | []rune](str T) (r []byte)
```
命令行程序使用
```sh
# 基本使用
b154 encode filename/dirname
b154 decode filename/dirname

# 指定编解码次数
b154 encode filename -n 6   # 进行6次编码

# 在标准输出显示解码后结果
b154 decode filename -s
```
