//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

func InitializeEvent(phrase string) (Event, error) {
	// 使用したいイニシャライザを渡す
	wire.Build(NewEvent, NewGreeter, NewMessage)
	// 戻り値の定義を満たすためにゼロ値を返却する
	return Event{}, nil
}
