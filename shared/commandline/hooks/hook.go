package hooks

import "github.com/loewensteinph/fthelper/shared/maps"

// Hook action
type Hook func(config maps.Mapper) error
