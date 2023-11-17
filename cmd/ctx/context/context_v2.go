/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package context

type param_v2 struct {
	key   string
	value string
}

func NewParam_v2(key, value string) *param_v2 {
	return &param_v2{
		key:   key,
		value: value,
	}
}

type params_v2 map[string]string

type Context_v2 struct {
	params_v2 params_v2
}

func NewContext_v2() *Context_v2 {
	return &Context_v2{
		params_v2: make(params_v2),
	}
}

func (c *Context_v2) Add(key, value string) {
	c.params_v2[key] = value
}

func (c *Context_v2) Get(key string) string {
	return c.params_v2[key]
}

func (c *Context_v2) Remove(key string) {
	delete(c.params_v2, key)
}
