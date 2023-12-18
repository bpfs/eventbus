package eventbus

import (
	"fmt"
	"testing"
)

func TestEventRegistry(t *testing.T) {
	// 创建事件注册器
	registry := NewEventRegistry()

	// 注册切片上传事件总线
	registry.RegisterEvent("file:SliceUpload", New())

	// 注册切片转发事件总线
	registry.RegisterEvent("file:SliceTransfer", New())

	// 获取事件总线
	sliceUploadBus := registry.GetEventBus("file:SliceUpload")
	if sliceUploadBus == nil {
		t.Error("Failed to get the slice upload event bus.")
	} else {
		// 使用事件总线...
		fmt.Println("Slice upload event bus registered.")
	}

	sliceTransferBus := registry.GetEventBus("sliceTransfer")
	if sliceTransferBus == nil {
		t.Error("无法获取切片传输事件总线。")
	} else {
		// 使用事件总线...
		fmt.Println("Slice transfer event bus registered.")
	}
}
