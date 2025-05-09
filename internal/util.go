package internal

import (
	"encoding/json"
	"log"
)

var messageMapping = map[string]interface{} {
	"VTubeStudioAPIStateBroadcast": StatusRequest{},
	"AuthenticationTokenResponse" : AuthenticationRequest{},
	"APIError" : AuthenticationRequest{},
	"AuthenticationResponse" : AuthenticationRequest{},
	"StatisticsResponse" : AuthenticationRequest{},
	"VTSFolderInfoResponse" : AuthenticationRequest{},
	"CurrentModelResponse" : AuthenticationRequest{},
	"AvailableModelsResponse" : AuthenticationRequest{},
	"ModelLoadResponse" : AuthenticationRequest{},
	"MoveModelResponse" : AuthenticationRequest{},
	"HotkeysInCurrentModelResponse" : AuthenticationRequest{},
	"HotkeyTriggerResponse" : AuthenticationRequest{},
	"ExpressionStateResponse" : AuthenticationRequest{},
	"ExpressionActivationResponse" : AuthenticationRequest{},
	"ArtMeshListResponse" : AuthenticationRequest{},
	"ColorTintResponse" : AuthenticationRequest{},
	"SceneColorOverlayInfoResponse" : AuthenticationRequest{},
	"FaceFoundResponse" : AuthenticationRequest{},
	"InputParameterListResponse" : AuthenticationRequest{},
	"ParameterValueResponse" : AuthenticationRequest{},
	"Live2DParameterListResponse" : AuthenticationRequest{},
	"ParameterCreationResponse" : AuthenticationRequest{},
	"ParameterDeletionResponse" : AuthenticationRequest{},
	"InjectParameterDataResponse" : AuthenticationRequest{},
	"GetCurrentModelPhysicsResponse" : AuthenticationRequest{},
	"SetCurrentModelPhysicsResponse" : AuthenticationRequest{},
	"NDIConfigResponse" : AuthenticationRequest{},
	"ItemListResponse" : AuthenticationRequest{},
	"ItemLoadResponse" : AuthenticationRequest{},
	"ItemUnloadResponse" : AuthenticationRequest{},
	"ItemAnimationControlResponse" : AuthenticationRequest{},
	"ItemMoveResponse" : AuthenticationRequest{},
	"ArtMeshSelectionResponse" : AuthenticationRequest{},
	"ItemPinResponse" : AuthenticationRequest{},
	"PostProcessingListResponse" : AuthenticationRequest{},
	"PostProcessingUpdateResponse" : AuthenticationRequest{},
}

//https://stackoverflow.com/questions/47911187/golang-elegantly-json-decode-different-structures

/*
Two different idea here, that I can have a metadata wrapper structure and have json raw value to parse it correctlu
Or
I can just have a bunch of if statements and parse it
*/

func ParseByMessageType(message []byte) ([]interface{}, error) {
	response := RespondMessageBase{}
	err := json.Unmarshal(message, &response)

	if err != nil {
		log.Fatalf("json unmarshal error: %v", err)
	}

	switch response.MessageType {
		case:

		case API_VERSION:

	}


	return nil, nil
}
