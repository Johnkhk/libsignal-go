package prekey

import (
	"github.com/Johnkhk/libsignal-go/protocol/address"
	"github.com/Johnkhk/libsignal-go/protocol/curve"
	"github.com/Johnkhk/libsignal-go/protocol/identity"
)

// Bundle represents a pre-key bundle as defined by the X3DH protocol.
//
// See https://signal.org/docs/specifications/x3dh/#sending-the-initial-message for more information.
type Bundle struct {
	RegistrationID        uint32
	DeviceID              address.DeviceID
	PreKeyID              *ID
	PreKeyPublic          curve.PublicKey
	SignedPreKeyID        ID
	SignedPreKeyPublic    curve.PublicKey
	SignedPreKeySignature []byte
	IdentityKey           identity.Key
}
