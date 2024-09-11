package internal

import (
	"context"
	"encoding/json"
	"path/filepath"
	"strings"

	"github.com/goexl/exception"
	"github.com/goexl/gox/field"
	"github.com/pangum/config"
)

var _ config.Loader = (*Loader)(nil)

type Loader struct{}

func NewLoader() *Loader {
	return new(Loader)
}

func (l *Loader) Local() bool {
	return true
}

func (l *Loader) Load(ctx context.Context, target any) (err error) {
	if path, pok := ctx.Value(config.ContextFilepath).(string); !pok {
		err = exception.New().Message("未指定配置文件路径").Field(field.New("loader", "yaml")).Build()
	} else if bytes, bok := ctx.Value(config.ContextBytes).([]byte); !bok {
		err = exception.New().Message("配置文件无内容").Field(field.New("loader", "yaml")).Build()
	} else {
		err = l.load(&path, &bytes, target)
	}

	return
}

func (l *Loader) load(path *string, bytes *[]byte, target any) (err error) {
	ext := strings.ToLower(filepath.Ext(*path))
	if ".yaml" == ext || ".yml" == ext {
		err = json.Unmarshal(*bytes, target)
	}

	return
}
