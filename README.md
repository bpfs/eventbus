# EventBus

[![Go Reference](https://pkg.go.dev/badge/github.com/klauspost/reedsolomon.svg)](https://pkg.go.dev/github.com/bpfs/eventbus) [![Go](https://github.com/klauspost/reedsolomon/actions/workflows/go.yml/badge.svg)]()

EventBus 是用 Go 编写的一种轻量级事件总线

该项目目前正在积极开发中，处于 Release 状态。

## 要求

[Go](http://golang.org) 1.20 或最新版本.

## 安装

要获取包，请使用标准：

```bash
go get -u github.com/bpfs/eventbus
```

推荐使用 Go 模块。

## 用法

本节假设您具有 Go（golang）编码的基础知识。 

```go
import "github.com/bpfs/eventbus"
```

### 例子

首先，您需要使用 NewDeP2P 创建新的 DeP2P 实例。虽然您可以使用默认的配置运行网络，但我们依然强烈建议您对相关参数进行配置。当然，这很大程度上取决于您的使用场景。不过您放心，文后我们会附上可直接运行的配置供您修改。

```Go
func main() {
	// 创建事件注册器
	registry := NewEventRegistry()

	// 注册示例A的事件总线
	registry.RegisterEvent("example:A", New())

	// 注册示例B的事件总线
	registry.RegisterEvent("example:B", New())

	// 获取事件总线
	Example_A_Bus := registry.GetEventBus("example:A")
	if Example_A_Bus == nil {
		fmt.Error("无法获取示例A的事件总线")
	} else {
		// 使用事件总线...
		fmt.Println("已注册 示例A 的事件总线")
	}

	// 注册
	Example_A_Bus.Subscribe("example:A", func(
		a string, // "example:A" 需要的请求参数
	) error {
		fmt.Printf("打印example:A事件参数:\t%s\n\n",a)
		return nil
	})


	Example_A_Bus.Publish("example:A", "参数a")

	///////////////////////////////////////////

	Example_B_Bus := registry.GetEventBus("example:B")
	if Example_B_Bus == nil {
		fmt.Error("无法获取示例B的事件总线。")
	} else {
		// 使用事件总线...
		fmt.Println("已注册 示例B 的事件总线")
	}

	// 注册
	Example_B_Bus.Subscribe("example:B", func(
		b string, // "example:B" 需要的请求参数
	) error {
		fmt.Printf("打印example:B事件参数:\t%s\n\n",b)
		return nil
	})

	Example_B_Bus.Publish("example:B", "参数b")
}

```

### 文档

该文档正在编写中……
 