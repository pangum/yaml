package yaml

import (
	"github.com/pangum/pangu"
	"github.com/pangum/yaml/internal"
)

func init() {
	pangu.New().Config().Loader(internal.NewLoader()).Build()
}
