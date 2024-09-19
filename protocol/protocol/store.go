package protocol

import (
	"github.com/Johnkhk/libsignal-go/protocol/identity"
	"github.com/Johnkhk/libsignal-go/protocol/prekey"
	"github.com/Johnkhk/libsignal-go/protocol/session"
)

type Store interface {
	SessionStore() session.Store
	IdentityStore() identity.Store
	PreKeyStore() prekey.Store
	SignedPreKeyStore() prekey.SignedStore
	GroupStore() session.GroupStore
}
