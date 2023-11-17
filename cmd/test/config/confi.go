/*
 *   Copyright (c) 2023 Andrey Danilov andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package config

type ConfigContext struct {
	temp string
	path string
}

func (cc ConfigContext) SetPath(path string) {
	cc.path = path
}

func (cc ConfigContext) SetTemp(temp string) {
	cc.temp = temp
}

func (cc ConfigContext) GetPath() string {
	return cc.path
}

func (cc ConfigContext) GetTemp() string {
	return cc.temp
}

type Cfg struct {
	*ConfigContext
}

func NewCfg() *Cfg {
	return &Cfg{
		ConfigContext: &ConfigContext{},
	}
}
