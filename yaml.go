package yaml

import (
	"github.com/pangum/yaml/internal"
)

func New() *internal.Loader {
	return internal.NewLoader()
}
