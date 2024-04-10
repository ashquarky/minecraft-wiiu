package nex

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PretendoNetwork/minecraft-wiiu/globals"
	nex "github.com/PretendoNetwork/nex-go/v2"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewPRUDPServer()

	globals.SecureEndpoint = nex.NewPRUDPEndPoint(1)
	globals.SecureEndpoint.IsSecureEndPoint = true
	//globals.SecureEndpoint.ServerAccount = globals.ServerAccount
	globals.SecureServer.BindPRUDPEndPoint(globals.SecureEndpoint)

	globals.SecureServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(3, 10, 0))
	globals.SecureServer.AccessKey = "f1b61c8e"

	//globals.SecureServer = nex.NewServer()
	//globals.SecureServer.SetPRUDPVersion(1)
	//globals.SecureServer.SetPRUDPProtocolMinorVersion(3)
	//globals.SecureServer.SetDefaultNEXVersion(nex.NewNEXVersion(3, 10, 0))
	//globals.SecureServer.SetKerberosPassword(globals.KerberosPassword)
	//globals.SecureServer.SetAccessKey("f1b61c8e")

	globals.Timeline = make(map[uint32][]uint8)

	globals.SecureEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("==Minecraft: Wii U Edition - Secure==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID)
		fmt.Printf("Method ID: %#v\n", request.MethodID)
		fmt.Println("===============")
	})

	//registerCommonSecureServerProtocols()
	//registerSecureServerNEXProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_MINECRAFT_SECURE_SERVER_PORT"))

	globals.SecureServer.Listen(port)
}
