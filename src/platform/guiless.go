// +build !windows,!macos

package platform

import (
	"github.com/nkanaev/yarr/src/server"
)

func Start(s *server.Server) {
	s.Start()
}
