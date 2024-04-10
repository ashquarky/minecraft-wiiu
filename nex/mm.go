package nex

//func GetPlayingSession(err error, packet nex.PacketInterface, callID uint32, listPID []uint32) uint32 {
//	rmcResponseStream := nex.NewStreamOut(globals.SecureServer)
//	rmcResponseStream.WriteUInt32LE(0)
//
//	rmcResponseBody := rmcResponseStream.Bytes()
//	fmt.Println(hex.EncodeToString(rmcResponseBody))
//
//	rmcResponse := nex.NewRMCResponse(matchmake_extension.ProtocolID, callID)
//	rmcResponse.SetSuccess(matchmake_extension.MethodGetPlayingSession, rmcResponseBody)
//
//	rmcResponseBytes := rmcResponse.Bytes()
//
//	responsePacket, _ := nex.NewPacketV1(packet.Sender(), nil)
//
//	responsePacket.SetVersion(1)
//	responsePacket.SetSource(0xA1)
//	responsePacket.SetDestination(0xAF)
//	responsePacket.SetType(nex.DataPacket)
//	responsePacket.SetPayload(rmcResponseBytes)
//
//	responsePacket.AddFlag(nex.FlagNeedsAck)
//	responsePacket.AddFlag(nex.FlagReliable)
//
//	globals.SecureServer.Send(responsePacket)
//
//	return 0
//}
