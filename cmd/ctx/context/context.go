/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package context

type param struct {
	key   string
	value string
}

func NewParam(key, value string) *param {
	return &param{
		key:   key,
		value: value,
	}
}

type params []param

type Context struct {
	params params
	val    map[string]string
}

func NewCtx() *Context {
	ctx := &Context{
		params: params{},
	}
	return ctx
}

func (c *Context) AddParam(param param) {
	c.params = append(c.params, param)
}

func (c *Context) GetParam(key string) string {
	for _, entry := range c.params {
		if entry.key == key {
			return key
		}
	}
	return ""
}
