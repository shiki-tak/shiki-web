package server

import (
	"github.com/shiki-tak/shiki-web/app/server/infrastructure"
)

func Init(port string) {
	infrastructure.Init(port)
}
