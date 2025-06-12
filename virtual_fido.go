package virtual_fido

import (
	"io"

	"github.com/shatru0/virtual-fido-driver/ctap"
	"github.com/shatru0/virtual-fido-driver/u2f"
	"github.com/shatru0/virtual-fido-driver/util"
)

type FIDOClient interface {
	u2f.U2FClient
	ctap.CTAPClient
}

func Start(client FIDOClient) {
	// Calls either the Mac or USB/IP client, based on system
	startClient(client)
}

func SetLogLevel(level util.LogLevel) {
	util.SetLogLevel(level)
}

func SetLogOutput(out io.Writer) {
	util.SetLogOutput(out)
}
