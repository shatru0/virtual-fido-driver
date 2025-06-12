//go:build darwin

package virtual_fido

import (
	"github.com/shatru0/virtual-fido-driver/ctap"
	"github.com/shatru0/virtual-fido-driver/ctap_hid"
	"github.com/shatru0/virtual-fido-driver/mac"
	"github.com/shatru0/virtual-fido-driver/u2f"
)

/*
 * Mac client requires installation of Mac USBDriver, which implements a virtual USB device.
 */
func startClient(client FIDOClient) {
	ctapServer := ctap.NewCTAPServer(client)
	u2fServer := u2f.NewU2FServer(client)
	ctapHIDServer := ctap_hid.NewCTAPHIDServer(ctapServer, u2fServer)
	mac.Start(ctapHIDServer)
}
