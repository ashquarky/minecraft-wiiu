package nex

import (
	matchmakingtypes "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
)

func cleanupSearchMatchmakeSessionHandler(matchmakeSession *matchmakingtypes.MatchmakeSession) {
	//matchmakeSession.Attributes[2] = 0
	//matchmakeSession.MatchmakeParam = matchmakingtypes.NewMatchmakeParam()
	//matchmakeSession.ApplicationData = make([]byte, 0)
	//fmt.Println(matchmakeSession.String())
}

// func cleanupMatchmakeSessionSearchCriteriaHandler(searchCriteria []*match_making_types.MatchmakeSessionSearchCriteria){
// }

func gameSpecificMatchmakeSessionSearchCriteriaChecksHandler(searchCriteria *matchmakingtypes.MatchmakeSessionSearchCriteria, matchmakeSession *matchmakingtypes.MatchmakeSession) bool {
	return true
}

func createReportDBRecordHandler(pid uint32, reportID uint32, reportData []byte) error {
	return nil
}

func registerCommonSecureServerProtocols() {
	//secureProtocol := secureconnection.NewProtocol()
	//commonsecureconnection.NewCommonProtocol(secureProtocol)
	//
	//matchmakingProtocol := matchmaking.NewProtocol()
	//
	//commonSecureconnectionProtocol := secureconnection.NewCommonSecureConnectionProtocol(globals.SecureServer)
	//commonSecureconnectionProtocol.CreateReportDBRecord(createReportDBRecordHandler)
	//commonmatchmaking.NewCommonMatchMakingProtocol(globals.SecureServer)
	//commonmatchmakingext.NewCommonMatchMakingExtProtocol(globals.SecureServer)
	//commonMatchmakeExtensionProtocol := commonmatchmakeextension.NewCommonMatchmakeExtensionProtocol(globals.SecureServer)
	//commonMatchmakeExtensionProtocol.CleanupSearchMatchmakeSession(cleanupSearchMatchmakeSessionHandler)
	//// 	commonMatchmakeExtensionProtocol.CleanupMatchmakeSessionSearchCriteria(cleanupMatchmakeSessionSearchCriteriaHandler)
	//commonMatchmakeExtensionProtocol.GameSpecificMatchmakeSessionSearchCriteriaChecks(gameSpecificMatchmakeSessionSearchCriteriaChecksHandler)
	//commonMatchmakeExtensionProtocol.DefaultProtocol.GetPlayingSession(GetPlayingSession)
	////commonMatchmakeExtensionProtocol.DefaultProtocol.AutoMatchmakeWithParam_Postpone(autoMatchmakeWithParam_Postpone)
	//nattraversal.NewCommonNATTraversalProtocol(globals.SecureServer)
	//
	//globals.SecureServer
}
