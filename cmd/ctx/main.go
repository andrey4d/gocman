/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */

package main

import (
	"godman/cmd/ctx/context"
	"godman/cmd/ctx/somepkg"
	"log"
)

var Context *context.Context
var Context_v2 *context.Context_v2

func init() {
	log.Println("Init context")
	Context = context.NewCtx()
	param := context.NewParam("path", "/usr/home")
	Context.AddParam(*param)

	Context_v2 = context.NewContext_v2()
	Context_v2.Add("temp", "/temp")

}

func main() {

	log.Println("Main() -->", Context.GetParam("path"))

	somepkg.AnyFunctionUseCtx(Context)

	log.Println("Main()_v2 -->", Context_v2.Get("temp"))
	somepkg.AnyFunctionUseCtxV2(Context_v2)
	log.Println("Main()_v2 -->", Context_v2.Get("temp"))

}
