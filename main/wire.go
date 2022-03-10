//go:build wireinject
// +build wireinject

package main

import (
	"FitnessSpace/provider"
	"github.com/google/wire"
)

func InitDB() *provider.Database {
	panic(wire.Build(provider.NewDB))
}
