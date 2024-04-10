package nex

import (
	"github.com/PretendoNetwork/minecraft-wiiu/globals"
	"github.com/PretendoNetwork/nex-go/v2/constants"
	"github.com/PretendoNetwork/nex-go/v2/types"
	commonticketgranting "github.com/PretendoNetwork/nex-protocols-common-go/v2/ticket-granting"
	ticketgranting "github.com/PretendoNetwork/nex-protocols-go/v2/ticket-granting"
	"os"
	"strconv"
)

func registerCommonAuthenticationServerProtocols() {
	ticketGrantingProtocol := ticketgranting.NewProtocol()
	globals.AuthenticationEndpoint.RegisterServiceProtocol(ticketGrantingProtocol)
	commonTicketGrantingProtocol := commonticketgranting.NewCommonProtocol(ticketGrantingProtocol)

	port, _ := strconv.Atoi(os.Getenv("PN_MINECRAFT_SECURE_SERVER_PORT"))

	secureStationURL := types.NewStationURL("")
	secureStationURL.SetURLType(constants.StationURLPRUDP)
	secureStationURL.SetAddress(os.Getenv("PN_MINECRAFT_SECURE_SERVER_HOST"))
	secureStationURL.SetPortNumber(uint16(port))
	secureStationURL.SetConnectionID(1)
	secureStationURL.SetPrincipalID(types.NewPID(2))
	secureStationURL.SetStreamID(1)
	secureStationURL.SetStreamType(constants.StreamTypeRVSecure)
	secureStationURL.SetType(uint8(constants.StationURLFlagPublic))

	commonTicketGrantingProtocol.SecureStationURL = secureStationURL
	// todo this build string is wrong
	commonTicketGrantingProtocol.BuildName = types.NewString("branch:origin/project/ctr-egd build:3_10_22_2008_0")
	// server account?

	//ticketGrantingProtocol := ticket_granting.NewCommonTicketGrantingProtocol(globals.AuthenticationServer)
	//
	//port, _ := strconv.Atoi(os.Getenv("PN_MINECRAFT_SECURE_SERVER_PORT"))
	//
	//secureStationURL := nex.NewStationURL("")
	//secureStationURL.SetScheme("prudps")
	//secureStationURL.SetAddress(os.Getenv("PN_MINECRAFT_SECURE_SERVER_HOST"))
	//secureStationURL.SetPort(uint32(port))
	//secureStationURL.SetCID(1)
	//secureStationURL.SetPID(2)
	//secureStationURL.SetSID(1)
	//secureStationURL.SetStream(10)
	//secureStationURL.SetType(2)
	//
	//ticketGrantingProtocol.SetSecureStationURL(secureStationURL)
	////ticketGrantingProtocol.SetBuildName("branch:origin/project/ctr-egd build:3_10_22_2008_0")
	//
	//globals.AuthenticationServer.SetPasswordFromPIDFunction(globals.PasswordFromPID)
}
