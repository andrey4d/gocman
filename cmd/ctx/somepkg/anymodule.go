/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package somepkg

import (
	"godman/cmd/ctx/context"
	"log"
)

func AnyFunctionUseCtx(ctx *context.Context) {

	log.Println("AnyFunction() -->", ctx.GetParam("path"))

}

func AnyFunctionUseCtxV2(ctx *context.Context_v2) {

	log.Println("AnyFunction()_v2 -->", ctx.Get("temp"))
	ctx.Remove("temp")

}
