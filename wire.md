# wire

## 安装

```go
go install github.com/google/wire/cmd/wire@latest
```

## 基础

Wire has two core concepts: providers and injectors.
Wire有两个核心概念：提供者和注入者。

### 提供者

The primary mechanism in Wire is the **provider**: a function that can produce a
value. These functions are ordinary Go code.

Wire中的主要机制是**提供者**：一个可以产生一个值的函数。这些函数是普通的Go代码。