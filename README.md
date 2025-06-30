# Virtual FIDO

Virtual FIDO is a virtual USB device that implements the FIDO2/U2F protocol (like a YubiKey) to support 2FA and WebAuthN. Please note that this software is still in beta and under active development, so APIs may be subject to change.

## Features

-   Support for both Windows and Linux through USB/IP 
-   Connect using both U2F and FIDO2 protocols for both normal 2FA and WebAuthN
-   Store credentials in an encrypted format with a passphrase
-   Store credential data anywhere (example provided: a local file)
-   Generic approval mechanism for credential creation and login (example provided: terminal-based)

