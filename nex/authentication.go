package nex

import (
	"fmt"
	"os"
	"strconv"

	"github.com/PretendoNetwork/minecraft-wiiu/globals"
	nex "github.com/PretendoNetwork/nex-go/v2"
)

var serverBuildString string

func StartAuthenticationServer() {
	globals.AuthenticationServer = nex.NewPRUDPServer()

	globals.AuthenticationEndpoint = nex.NewPRUDPEndPoint(1)
	// todo: server account..?
	// NEED AccountDetailsByUsername
	globals.AuthenticationServer.BindPRUDPEndPoint(globals.AuthenticationEndpoint)

	globals.AuthenticationServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(3, 10, 0))
	globals.AuthenticationServer.AccessKey = "f1b61c8e"

	//globals.AuthenticationServer.SetPRUDPVersion(1)
	//globals.AuthenticationServer.SetPRUDPProtocolMinorVersion(2)
	//globals.AuthenticationServer.SetDefaultNEXVersion(nex.NewNEXVersion(3, 10, 0))
	//globals.AuthenticationServer.SetKerberosPassword(globals.KerberosPassword)
	//globals.AuthenticationServer.SetAccessKey("f1b61c8e")

	globals.AuthenticationEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("==Minecraft: Wii U Edition - Auth==")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID)
		fmt.Printf("Method ID: %#v\n", request.MethodID)
		fmt.Println("===============")
	})

	registerCommonAuthenticationServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_MINECRAFT_AUTHENTICATION_SERVER_PORT"))

	globals.AuthenticationServer.Listen(port)
}
