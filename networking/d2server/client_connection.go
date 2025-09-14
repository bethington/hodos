package d2server

import (
	d2hero "nostos/core/character"
	"nostos/networking/d2client/d2clientconnectiontype"
	"nostos/networking/d2netpacket"
)

// ClientConnection is an interface for abstracting local and remote
// clients.
type ClientConnection interface {
	GetUniqueID() string
	GetConnectionType() d2clientconnectiontype.ClientConnectionType
	SendPacketToClient(packet d2netpacket.NetPacket) error
	GetPlayerState() *d2hero.HeroState
	SetPlayerState(playerState *d2hero.HeroState)
}
