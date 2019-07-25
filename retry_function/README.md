# retry function

有时候我们需要一些重试函数，当函数调用失败时，可以进行重试操作。超出一定的重试次数后，函数返回错误。

示例中有2个函数：

- alwaysFailed：该函数总是失败，尝试3次后，返回错误。
- succInThirdCall：该函数前2次会失败，第3次会成功。



###  retry函数的使用

入参：

- attempts int
  - 尝试的最大次数
- f func() error
  - 实际的函数



```go
if err := doRetry(3, f); err != nil {
  log.Println("try do best, but failed")
} else {
  log.Println("success")
}
```

### 运行

```bash
go run main.go
```

