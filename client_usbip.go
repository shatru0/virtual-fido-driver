//go:build linux || windows

package virtual_fido

import (
	"github.com/shatru0/virtual-fido-driver/ctap"
	"github.com/shatru0/virtual-fido-driver/ctap_hid"
	"github.com/shatru0/virtual-fido-driver/u2f"
	"github.com/shatru0/virtual-fido-driver/usb"
	"github.com/shatru0/virtual-fido-driver/usbip"
)

func startClient(client FIDOClient) {
	ctapServer := ctap.NewCTAPServer(client)
	u2fServer := u2f.NewU2FServer(client)
	ctapHIDServer := ctap_hid.NewCTAPHIDServer(ctapServer, u2fServer)
	usbDevice := usb.NewUSBDevice(ctapHIDServer)
	server := usbip.NewUSBIPServer([]usbip.USBIPDevice{usbDevice})
	server.Start()
}
