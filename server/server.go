package server

import (
	"github.com/shiki-tak/shiki-web/server/infrastructure"
)

func Init(port string) {
	infrastructure.Init(port)
}
