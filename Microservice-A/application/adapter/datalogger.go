package adapter

import "github.com/m4rc0nd35/test-fluid/application/entity"

type AdapterDataLogger interface {
	LogQueue(user entity.User) bool
}
