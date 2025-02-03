package internal

type Model struct {
	ModelLoaded      bool   `json:"modelLoaded"`
	ModelName        string `json:"modelName"`
	ModelID          string `json:"modelID"`
	VtsModelName     string `json:"vtsModelName"`
	VtsModelIconName string `json:"vtsModelIconName"`
}

type ModelPosition struct {
	PositionX float32 `json:"position_x"`
	PositionY float32 `json:"position_y"`
	Rotation  float32 `json:"rotation"`
	Size      float32 `json:"size"`
}

type HotKey struct {
	Name             string   `json:"name"`
	Type             string   `json:"type"`
	Description      string   `json:"description"`
	File             string   `json:"file"`
	HotkeyID         string   `json:"hotkeyID"`
	KeyCombination   []string `json:"keyCombination"`
	OnScreenButtonID int      `json:"onScreenButtonID"`
}

type UsedHotKey struct {
	Name string `json:"name"`
	Id   string `json:"id"`
}

type Parameters struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Expression struct {
	Name                       string       `json:"name"`
	File                       string       `json:"file"`
	Active                     bool         `json:"active"`
	DeactivateWhenKeyIsLetGo   bool         `json:"deactivateWhenKeyIsLetGo"`
	AutoDeactivateAfterSeconds bool         `json:"autoDeactivateAfterSeconds"`
	SecondsRemaining           uint16       `json:"secondsRemaining"`
	UsedInHotkeys              []UsedHotKey `json:"used_in_hotkeys"`
	Parameters                 []Parameters `json:"parameters"`
}

type Data struct {
	Active                      bool          `json:"active,omitempty"`
	VTubeStudioVersion          string        `json:"vTubeStudioVersion,omitempty"`
	CurrentSessionAuthenticated bool          `json:"currentSessionAuthenticated,omitempty"`
	Port                        uint16        `json:"port,omitempty"`
	InstanceID                  string        `json:"instanceID,omitempty"`
	WindowTitle                 string        `json:"windowTitle,omitempty"`
	PluginName                  string        `json:"pluginName,omitempty"`
	PluginDeveloper             string        `json:"pluginDeveloper,omitempty"`
	PluginIcon                  string        `json:"pluginIcon,omitempty"`
	AuthenticationToken         string        `json:"authenticationToken,omitempty"`
	ErrorID                     uint16        `json:"errorID,omitempty"`
	Message                     string        `json:"message,omitempty"`
	Reason                      string        `json:"reason,omitempty"`
	Uptime                      uint16        `json:"uptime,omitempty"`
	Framerate                   int           `json:"framerate,omitempty"`
	AllowedPlugins              uint16        `json:"allowedPlugins,omitempty"`
	ConnectedPlugins            uint16        `json:"connectedPlugins,omitempty"`
	StartedWithSteam            bool          `json:"startedWithSteam,omitempty"`
	WindowWidth                 uint16        `json:"windowWidth,omitempty"`
	WindowHeight                uint16        `json:"windowHeight,omitempty"`
	WindowIsFullscreen          bool          `json:"windowIsFullscreen,omitempty"`
	Models                      string        `json:"models,omitempty"`
	Backgrounds                 string        `json:"backgrounds,omitempty"`
	Items                       string        `json:"items,omitempty"`
	Config                      string        `json:"config,omitempty"`
	Logs                        string        `json:"logs,omitempty"`
	Backup                      string        `json:"backup,omitempty"`
	ModelLoaded                 bool          `json:"modelLoaded,omitempty"`
	ModelName                   string        `json:"modelName,omitempty"`
	ModelID                     string        `json:"modelID,omitempty"`
	VtsModelName                string        `json:"vtsModelName,omitempty"`
	VtsModelIconName            string        `json:"vtsModelIconName,omitempty"`
	Live2DModelName             string        `json:"live2DModelName,omitempty"`
	ModelLoadTime               uint16        `json:"modelLoadTime,omitempty"`
	TimeSinceModelLoaded        uint32        `json:"timeSinceModelLoaded,omitempty"`
	NumberOfLive2DParameters    uint16        `json:"numberOfLive2DParameters,omitempty"`
	NumberOfLive2DArtmeshes     uint16        `json:"numberOfLive2DArtmeshes,omitempty"`
	HasPhysicsFile              bool          `json:"hasPhysicsFile,omitempty"`
	NumberOfTextures            uint16        `json:"numberOfTextures,omitempty"`
	TextureResolution           uint16        `json:"textureResolution,omitempty"`
	ModelPosition               ModelPosition `json:"modelPosition,omitempty"`
	NumberOfModels              uint16        `json:"numberOfModels,omitempty"`
	AvailableModels             []Model       `json:"availableModels,omitempty"`
	TimeInSeconds               float32       `json:"timeInSeconds,omitempty"`
	ValuesAreRelativeToModel    bool          `json:"valuesAreRelativeToModel,omitempty"`
	PositionX                   float32       `json:"position_x,omitempty"`
	PositionY                   float32       `json:"position_y,omitempty"`
	Rotation                    float32       `json:"rotation,omitempty"`
	Size                        float32       `json:"size,omitempty"`
	Live2DItemFileName          string        `json:"live2DItemFileName,omitempty"`
	AvailableHotkeys            []HotKey      `json:"available_hotkeys,omitempty"`
	HotkeyID                    string        `json:"hotkeyID,omitempty"`
	ItemInstanceID              string        `json:"itemInstanceID,omitempty"`
	Details                     bool          `json:"details,omitempty"`
	ExpressionFile              string        `json:"expressionFile,omitempty"`
	Expressions                 []Expression  `json:"expressions,omitempty"`
	FadeTime                    float32       `json:"fadeTime,omitempty"`
	// Only finish until expression
	// Have not started ArtMeshes with color selection
	// Need to reorg this struct later, to be respond to request
	// for now I'm only using expression anyways
}

type Message struct {
	ApiName     string `json:"apiName"`
	ApiVersion  string `json:"apiVersion"`
	Timestamp   int64  `json:"timestamp,omitempty"`
	RequestID   string `json:"requestID"`
	MessageType string `json:"messageType"`
	Data        Data   `json:"data,omitempty"`
}

const (
	Status = "APIStateRequest"
)

type RequestMessageBase struct {
	ApiName     string `json:"apiName"`
	ApiVersion  string `json:"apiVersion"`
	RequestID   string `json:"requestID"`
	MessageType string `json:"messageType"`
}

type RespondMessageBase struct {
	ApiName     string `json:"apiName"`
	ApiVersion  string `json:"apiVersion"`
	Timestamp   int64  `json:"timestamp"`
	RequestID   string `json:"requestID"`
	MessageType string `json:"messageType"`
}

type StatusRequest struct {
	RequestMessageBase
}

type StatusResponseData struct {
	Active                      bool   `json:"active"`
	VtubeStudioVersion          string `json:"vTubeStudioVersion"`
	CurrentSessionAuthenticated bool   `json:"currentSessionAuthenticated"`
}

type StatusResponse struct {
	RespondMessageBase
	StatusResponseData StatusResponseData `json:"data"`
}

type AuthenticationRequestData struct {
	PluginName      string `json:"pluginName"`
	PluginDeveloper string `json:"pluginDeveloper"`
	PluginIcon      string `json:"pluginIcon"`
}
type AuthenticationRequest struct {
	RequestMessageBase
	AuthenticationRequestData AuthenticationRequestData `json:"authenticationRequestData"`
}

type AuthenticationResponseData struct {
	Authenticated bool   `json:"authenticated"`
	Reason        string `json:"reason"`
}
type AuthenticationResponse struct {
	RespondMessageBase
	AuthenticationResponseData AuthenticationResponseData `json:"data"`
}

type StatisticsRequest struct {
	RespondMessageBase
}
type StatisticsRespondData struct {
	Uptime              uint32  `json:"uptime"`
	FrameRate           float32 `json:"frameRate"`
	VtuberStudioVersion string  `json:"vtuberStudioVersion"`
	AllowedPlugins      uint16  `json:"allowedPlugins"`
	ConnectedPlugins    uint16  `json:"connectedPlugins"`
	StartedWithSteam    bool    `json:"startedWithSteam"`
	WindowWidth         uint16  `json:"windowWidth"`
	WindowHeight        uint16  `json:"windowHeight"`
	WindowIsFullscreen  bool    `json:"windowIsFullscreen"`
}

type StatisticsResponse struct {
	RespondMessageBase
	StatisticsRequestData StatisticsRequest `json:"data"`
}

type VTSFolderInfoRequest struct {
	RequestMessageBase
}

type VTSFolderInfoResponseData struct {
	Model      string `json:"model"`
	Background string `json:"background"`
	Items      string `json:"items"`
	Config     string `json:"config"`
	Logs       string `json:"logs"`
	Backup     string `json:"backup"`
}

type VTSFolderInfoResponse struct {
	RespondMessageBase
	VTSFolderInfoResponseData VTSFolderInfoResponseData `json:"data"`
}

type CurrentModelRequest struct {
	RequestMessageBase
}
type CurrentModelResponseData struct {
	ModelLoaded              bool          `json:"modelLoaded"`
	ModelName                string        `json:"modelName"`
	ModelID                  string        `json:"modelID"`
	VtsModelName             string        `json:"vtsModelName"`
	VtsModelIconName         string        `json:"vtsModelIconName"`
	Live2DModelName          string        `json:"live2DModelName"`
	ModelLoadTime            uint32        `json:"modelLoadTime"`
	TimeSinceModelLoaded     uint32        `json:"timeSinceModelLoaded"`
	NumberOfLive2DParameters uint16        `json:"numberOfLive2DParameters"`
	NumberOfLive2DArtmeshes  uint16        `json:"numberOfLive2DArtmeshes"`
	HasPhysicsFile           bool          `json:"hasPhysicsFile"`
	NumberOfTextures         uint16        `json:"numberOfTextures"`
	TextureResolution        uint16        `json:"textureResolution"`
	ModelPosition            ModelPosition `json:"modelPosition"`
}

type CurrentModelResponse struct {
	RespondMessageBase
	CurrentModelResponseData CurrentModelResponseData `json:"data"`
}

type AvailableModelsRequest struct {
	RequestMessageBase
}

type AvailableModelsResponseData struct {
	NumberOfModels  uint16 `json:"numberOfModels"`
	AvailableModels Model  `json:"availableModels"`
}

type AvailableModelsResponse struct {
	RespondMessageBase
	AvailableModelsResponseData AvailableModelsResponseData `json:"data"`
}

type ModelLoadedRequestData struct {
	ModelID string `json:"modelID"`
}

type ModelLoadedRequest struct {
	RequestMessageBase
	ModelLoadedRequestData ModelLoadedRequestData `json:"data"`
}

type ModelLoadedResponseData struct {
	ModelID string `json:"modelID"`
}

type ModelLoadedResponse struct {
	RespondMessageBase
	ModelLoadedResponseData ModelLoadedResponseData `json:"data"`
}

type MoveModelRequestData struct {
	TimeInSeconds            float32 `json:"timeInSeconds"`
	ValuesAreRelativeToModel bool    `json:"valuesAreRelativeToModel"`
	PositionX                float32 `json:"positionX"`
	PositionY                float32 `json:"positionY"`
	Rotation                 float32 `json:"rotation"`
	Size                     float32 `json:"size"`
}

type MoveModelRespondData struct {
}

type MoveModelRequest struct {
	RequestMessageBase
	MoveModelRequestData MoveModelRequestData `json:"data"`
}

type MoveModelResponse struct {
	RespondMessageBase
	ModelLoadedResponseData ModelLoadedResponseData `json:"data"`
}

type HotkeyInCurrentModelRequestData struct {
	ModelID            string `json:"modelID"`
	Live2DItemFileName string `json:"live2DItemFileName"`
}

type HotkeysInCurrentModelRequest struct {
	RequestMessageBase
}

type HotkeysInCurrentModelResponseData struct {
	ModelLoaded      bool     `json:"modelLoaded"`
	ModelName        string   `json:"modelName"`
	ModelID          string   `json:"modelID"`
	AvailableHotKeys []HotKey `json:"availableHotKeys"`
}

type HotkeysInCurrentModelResponse struct {
	RespondMessageBase
	HotkeysInCurrentModelResponseData HotkeysInCurrentModelResponseData `json:"data"`
}

type HotkeyTriggerRequestData struct {
	HotkeyID       string `json:"hotkeyID"`
	ItemInstanceID string `json:"itemInstanceID"`
}

type HotkeyTriggerRequest struct {
	RequestMessageBase
	HotkeyTriggerRequestData HotkeyTriggerRequestData `json:"data"`
}

type HotkeyTriggerResponseData struct {
	HotkeyID string `json:"hotkeyID"`
}

type HotkeyTriggerResponse struct {
	RespondMessageBase
	HotkeyTriggerResponseData HotkeyTriggerResponseData `json:"data"`
}

type ExpressionStateRequestData struct {
	Details    bool   `json:"details"`
	Expression string `json:"expression"`
}
type ExpressionStateRequest struct {
	RequestMessageBase
	ExpressionStateRequestData ExpressionStateRequestData `json:"data"`
}

type ExpressionStateResponseData struct {
	ModelLoaded bool         `json:"modelLoaded"`
	ModelName   string       `json:"modelName"`
	ModelID     string       `json:"modelID"`
	Expressions []Expression `json:"expressions"`
}

type ExpressionStateResponse struct {
	RespondMessageBase
	ExpressionStateResponseData ExpressionStateResponseData `json:"data"`
}

type ExpressionActivationRequestData struct {
	ExpressionFile string  `json:"expressionFile"`
	FadeTime       float32 `json:"fadeTime"`
	Active         bool    `json:"active"`
}

type ExpressionActivationRequest struct {
	RequestMessageBase
	ExpressionActivationRequestData ExpressionActivationRequestData `json:"data"`
}

type ExpressionActivationResponseData struct{}

type ExpressionActivationResponse struct {
	RespondMessageBase
	ExpressionActivationResponseData ExpressionActivationResponseData `json:"data"`
}

type ArtMeshListRequest struct {
	RequestMessageBase
}

type ArtMeshListResponseData struct {
	ModelLoaded          bool     `json:"modelLoaded"`
	NumberOfArtMeshNames uint16   `json:"numberOfArtMeshNames"`
	NumberOfArtMeshTags  uint16   `json:"numberOfArtMeshTags"`
	ArtMeshNames         []string `json:"artMeshNames"`
	ArtMeshTags          []string `json:"artMeshTags"`
}

type ArtMeshListResponse struct {
	RespondMessageBase
	ArtMeshListResponseData ArtMeshListResponseData `json:"data"`
}

type ColorTint struct {
	ColorR                    uint8 `json:"colorR"`
	ColorG                    uint8 `json:"colorG"`
	ColorB                    uint8 `json:"colorB"`
	ColorA                    uint8 `json:"colorA"`
	MixWithSceneLightingColor uint8 `json:"mixWithSceneLightingColor"`
}

type ArtMeshMatcher struct {
	TintAll       bool     `json:"tintAll"`
	ArtMeshNumber []uint8  `json:"artMeshNumber"`
	NameExact     []string `json:"nameExact"`
	TagExact      []string `json:"tagExact"`
	TagContains   []string `json:"tagContains"`
}

type ColorTintRequestData struct {
	ColorTint      ColorTint      `json:"colorTint"`
	ArtMeshMatcher ArtMeshMatcher `json:"artMeshMatcher"`
}

type ColorTintRequest struct {
	RequestMessageBase
	ColorTintRequestData ColorTintRequestData `json:"data"`
}

type ColorTintResponseData struct {
	MatchedArtMeshes uint8 `json:"matchedArtMeshes"`
}

type ColorTintResponse struct {
	RespondMessageBase
	ColorTintResponseData ColorTintResponseData `json:"data"`
}

type SceneColorOverlayInfoRequest struct {
	RequestMessageBase
}

type CapturePart struct {
	Active bool  `json:"active"`
	ColorR uint8 `json:"colorR"`
	ColorG uint8 `json:"colorG"`
	ColorB uint8 `json:"colorB"`
}

type SceneColorOverlayInfoResponseData struct {
	Active            bool        `json:"active"`
	ItemsIncluded     bool        `json:"itemsIncluded"`
	IsWindowCapture   bool        `json:"isWindowCapture"`
	BaseBrightness    uint16      `json:"baseBrightness"`
	ColorBoost        uint16      `json:"colorBoost"`
	Smoothing         uint16      `json:"smoothing"`
	ColorOverlayR     uint16      `json:"colorOverlayR"`
	ColorOverlayG     uint16      `json:"colorOverlayG"`
	ColorOverlayB     uint16      `json:"colorOverlayB"`
	LeftCapturePart   CapturePart `json:"leftCapturePart"`
	MiddleCapturePart CapturePart `json:"middleCapturePart"`
	RightCapturePart  CapturePart `json:"rightCapturePart"`
}

type SceneColorOverlayInfoResponse struct {
	RespondMessageBase
	SceneColorOverlayInfoResponse SceneColorOverlayInfoResponseData `json:"data"`
}

type FaceFoundRequest struct {
	RequestMessageBase
}

type FaceFoundResponseData struct {
	Found bool `json:"found"`
}

type FaceFoundResponse struct {
	RespondMessageBase
	FaceFoundResponseData FaceFoundResponseData `json:"data"`
}

type InputParameterListRequest struct {
	RequestMessageBase
}

type Parameter struct {
	Name         string `json:"name"`
	AddedBy      string `json:"addedBy"`
	Value        int    `json:"value"`
	Min          int    `json:"min"`
	Max          int    `json:"max"`
	DefaultValue int    `json:"defaultValue"`
}

type InputParameterListResponseData struct {
	ModelLoaded       bool        `json:"modelLoaded"`
	ModelName         string      `json:"modelName"`
	ModelID           string      `json:"modelID"`
	CustomParameters  []Parameter `json:"customParameters"`
	DefaultParameters []Parameter `json:"defaultParameters"`
}

type InputParameterListResponse struct {
	RespondMessageBase
	InputParameterListResponseData InputParameterListResponseData `json:"data"`
}

type ParameterValueRequestData struct {
	Name string `json:"name"`
}

type ParameterValueRequest struct {
	RequestMessageBase
	ParameterValueRequestData ParameterValueRequestData `json:"data"`
}

type ParameterValueResponseData struct {
	Name         string `json:"name"`
	AddedBy      string `json:"addedBy"`
	Value        int    `json:"value"`
	Min          int    `json:"min"`
	Max          int    `json:"max"`
	DefaultValue int    `json:"defaultValue"`
}

type ParameterValueResponse struct {
	RespondMessageBase
	ParameterValueResponseData ParameterValueResponseData `json:"data"`
}

type Live2DParameterListRequest struct {
	RequestMessageBase
}

type Live2DParameterListResponseData struct {
	ModelLoaded bool        `json:"modelLoaded"`
	ModelName   string      `json:"modelName"`
	ModelID     string      `json:"modelID"`
	Parameters  []Parameter `json:"parameters"`
}

type Live2DParameterListResponse struct {
	RespondMessageBase
	Live2DParameterListResponseData Live2DParameterListResponseData `json:"data"`
}

type ParameterCreationRequestData struct {
	ParameterName string `json:"parameterName"`
	Explanation   string `json:"explanation"`
	Min           int    `json:"min"`
	Max           int    `json:"max"`
	DefaultValue  int    `json:"defaultValue"`
}

type ParameterCreationRequest struct {
	RequestMessageBase
	ParameterCreationRequestData ParameterCreationRequestData `json:"data"`
}

type ParameterCreationResponseData struct {
	ParameterName string `json:"parameterName"`
}

type ParameterCreationResponse struct {
	RespondMessageBase
	ParameterCreationResponseData ParameterCreationResponseData `json:"data"`
}

type ParameterDeletionRequestData struct {
	ParameterName string `json:"parameterName"`
}

type ParameterDeletionRequest struct {
	RequestMessageBase
	ParameterDeletionRequestData ParameterDeletionRequestData `json:"data"`
}

type ParameterDeletionResponseData struct {
	ParameterName string `json:"parameterName"`
}

type ParameterDeletionResponse struct {
	RespondMessageBase
	ParameterDeletionResponseData ParameterDeletionResponseData `json:"data"`
}

type ParameterValue struct {
	Id    string  `json:"id"`
	Value float32 `json:"value"`
}

type InjectParameterDataRequestData struct {
	FaceFound       bool             `json:"faceFound"`
	Model           string           `json:"model"`
	ParameterValues []ParameterValue `json:"parameterValues"`
}

type InjectParameterDataRequest struct {
	RequestMessageBase
	InjectParameterDataRequestData InjectParameterDataRequestData `json:"data"`
}

type InjectParameterDataResponseData struct {
}

type InjectParameterDataResponse struct {
	RespondMessageBase
	InjectParameterDataResponseData InjectParameterDataResponseData `json:"data"`
}

type GetCurrentModelPhysicsRequest struct {
	RequestMessageBase
}

type PhysicsGroup struct {
	GroupID            string  `json:"groupID"`
	GroupName          string  `json:"groupName"`
	StrengthMultiplier float32 `json:"strengthMultiplier"`
	WindMultiplier     float32 `json:"windMultiplier"`
}

type GetCurrentModelPhysicsResponseData struct {
	ModelLoaded                  bool           `json:"modelLoaded"`
	ModelName                    string         `json:"modelName"`
	ModelID                      string         `json:"modelID"`
	ModelHasPhysics              bool           `json:"modelHasPhysics"`
	PhysicsSwitchedOn            bool           `json:"physicsSwitchedOn"`
	UsingLegacyPhysics           bool           `json:"usingLegacyPhysics"`
	PhysicsFPSSettings           int            `json:"physicsFPSSettings"`
	BaseStrength                 int            `json:"baseStrength"`
	BaseWind                     int            `json:"baseWind"`
	ApiPhysicsOverrideActive     bool           `json:"apiPhysicsOverrideActive"`
	ApiPhysicsOverridePluginName string         `json:"apiPhysicsOverridePluginName"`
	PhysicsGroups                []PhysicsGroup `json:"physicsGroups"`
}

type SetCurrentModelPhysicsRequestData struct {
	Id              string  `json:"id"`
	Value           float32 `json:"value"`
	SetBaseValue    bool    `json:"setBaseValue"`
	OverrideSeconds int     `json:"overrideSeconds"`
}

type SetCurrentModelPhysicsRequest struct {
	RequestMessageBase
	SetCurrentModelPhysicsRequestData SetCurrentModelPhysicsRequestData `json:"data"`
}

type SetCurrentModelPhysicsResponseData struct{}

type SetCurrentModelPhysicsResponse struct {
	RespondMessageBase
	SetCurrentModelPhysicsResponseData SetCurrentModelPhysicsResponseData `json:"data"`
}

type NDIConfigRequestData struct {
	SetNewConfig        bool `json:"setNewConfig"`
	NdiActive           bool `json:"ndiActive"`
	UseNDI5             bool `json:"useNDI5"`
	UseCustomResolution bool `json:"useCustomResolution"`
	CustomWidthNDI      int  `json:"customWidthNDI"`
	CustomHeightNDI     int  `json:"customHeightNDI"`
}

type NDIConfigRequest struct {
	RequestMessageBase
	NDIConfigRequestData NDIConfigRequestData `json:"data"`
}

type NDIConfigResponseData struct {
	SetNewConfig        bool `json:"setNewConfig"`
	NdiActive           bool `json:"ndiActive"`
	UseNDI5             bool `json:"useNDI5"`
	UseCustomResolution bool `json:"useCustomResolution"`
	CustomWidthNDI      int  `json:"customWidthNDI"`
	CustomHeightNDI     int  `json:"customHeightNDI"`
}

type NDIConfigResponse struct {
	RespondMessageBase
	NDIConfigResponseData NDIConfigResponseData `json:"data"`
}

type ItemListRequestData struct {
	IncludeAvailableSpots       bool   `json:"includeAvailableSpots"`
	IncludeItemInstancesInScene bool   `json:"includeItemInstancesInScene"`
	IncludeAvailableItemFiles   bool   `json:"includeAvailableItemFiles"`
	OnlyItemsWithFileName       string `json:"onlyItemsWithFileName"`
	OnlyItemsWithInstanceID     string `json:"onlyItemsWithInstanceID"`
}

type ItemListRequest struct {
	RequestMessageBase
	ItemListRequestData ItemListRequestData `json:"data"`
}

type Item struct {
	FileName        string  `json:"fileName"`
	InstanceID      string  `json:"instanceID"`
	Order           int     `json:"order"`
	Type            string  `json:"type"`
	Censored        bool    `json:"censored"`
	Flipped         bool    `json:"flipped"`
	Locked          bool    `json:"locked"`
	Smoothing       bool    `json:"smoothing"`
	Framerate       float32 `json:"framerate"`
	FrameCount      int     `json:"frameCount"`
	CurrentFrame    int     `json:"currentFrame"`
	PinnedToModel   bool    `json:"pinnedToModel"`
	PinnedModelID   string  `json:"pinnedModelID"`
	PinnedArtMeshID string  `json:"pinnedArtMeshID"`
	GroupName       string  `json:"groupName"`
	SceneName       string  `json:"sceneName"`
	FromWorkshop    bool    `json:"fromWorkshop"`
}

type ItemFile struct {
	FileName    string `json:"fileName"`
	Type        string `json:"type"`
	LoadedCount int    `json:"loadedCount"`
}

type ItemListResponseData struct {
	ItemsInSceneCount      uint8      `json:"itemsInSceneCount"`
	TotalItemsAllowedCount uint8      `json:"totalItemsAllowedCount"`
	CanLoadItemsRightNow   bool       `json:"canLoadItemsRightNow"`
	AvailableSpots         []int      `json:"availableSpots"`
	ItemInstancesInScene   []Item     `json:"itemInstancesInScene"`
	AvailableItemFiles     []ItemFile `json:"availableItemFiles"`
}

type ItemListResponse struct {
	RespondMessageBase
	ItemListResponseData ItemListResponseData `json:"data"`
}

type ItemLoadRequestData struct {
	FileName                              string  `json:"fileName"`
	PositionX                             float32 `json:"positionX"`
	PositionY                             float32 `json:"positionY"`
	Size                                  float32 `json:"size"`
	Rotation                              int     `json:"rotation"`
	FadeTime                              float32 `json:"fadeTime"`
	Order                                 int     `json:"order"`
	FailIfOrderTaken                      bool    `json:"failIfOrderTaken"`
	Smoothing                             int     `json:"smoothing"`
	Censored                              bool    `json:"censored"`
	Flipped                               bool    `json:"flipped"`
	Locked                                bool    `json:"locked"`
	UnloadWhenPluginDisconnects           bool    `json:"unloadWhenPluginDisconnects"`
	CustomDataBase64                      string  `json:"customDataBase64"`
	CustomDataAskUserFirst                bool    `json:"customDataAskUserFirst"`
	CustomDataSkipAskingUserIfWhitelisted bool    `json:"customDataSkipAskingUserIfWhitelisted"`
	CustomDataAskTimer                    int     `json:"customDataAskTimer"`
}

type ItemLoadRequest struct {
	RequestMessageBase
	ItemLoadRequestData ItemLoadRequestData `json:"data"`
}

type ItemLoadResponseData struct {
	InstanceID string `json:"instanceID"`
	FileName   string `json:"fileName"`
}

type ItemLoadResponse struct {
	RespondMessageBase
	ItemLoadResponseData ItemLoadResponseData `json:"data"`
}

type ItemUnloadRequestData struct {
	UnloadAllInScenes                             bool     `json:"unloadAllInScenes"`
	UnloadAllLoadedByThisPlugin                   bool     `json:"unloadAllLoadedByThisPlugin"`
	AllowUnloadingITemsLoadedByUserOrOtherPlugins bool     `json:"allowUnloadingITemsLoadedByUserOrOtherPlugins"`
	InstanceIDs                                   []string `json:"instanceIDs"`
	FileNames                                     []string `json:"fileNames"`
}

type ItemUnloadRequest struct {
	RequestMessageBase
	ItemUnloadRequestData ItemUnloadRequestData `json:"data"`
}

type UnloadItem struct {
	InstanceID string `json:"instanceID"`
	FileName   string `json:"fileName"`
}

type ItemUnloadResponseData struct {
	UnloadItems []UnloadItem `json:"unloadItems"`
}

type ItemUnloadResponse struct {
	RespondMessageBase
	ItemUnloadResponseData ItemUnloadResponseData `json:"data"`
}

type ItemAnimationControlRequestData struct {
	ItemInstanceID        string `json:"itemInstanceID"`
	FrameRate             int    `json:"frameRate"`
	Frame                 int    `json:"frame"`
	Brightness            int    `json:"brightness"`
	Opacity               int    `json:"opacity"`
	SetAutoStopFrames     bool   `json:"setAutoStopFrames"`
	AutoStopFrames        []int  `json:"autoStopFrames"`
	SetAnimationPlayState bool   `json:"setAnimationPlayState"`
	AnimationPlayState    bool   `json:"animationPlayState"`
}

type ItemAnimationControlRequest struct {
	RequestMessageBase
	ItemAnimationControlRequestData ItemAnimationControlRequestData `json:"data"`
}

type ItemAnimationControlResponseData struct {
	Frame            int  `json:"frame"`
	AnimationPlaying bool `json:"animationPlaying"`
}

type ItemAnimationControlResponse struct {
	RespondMessageBase
	ItemAnimationControlResponseData ItemAnimationControlResponseData `json:"data"`
}

type ItemToMove struct {
	ItemInstanceID string  `json:"itemInstanceID"`
	TimeInSeconds  int     `json:"timeInSeconds"`
	FadeMode       string  `json:"fadeMode"`
	PositionX      float32 `json:"positionX"`
	PositionY      float32 `json:"positionY"`
	Size           float32 `json:"size"`
	Rotation       int     `json:"rotation"`
	Order          int     `json:"order"`
	SetFilp        bool    `json:"setFilp"`
	Flip           bool    `json:"flip"`
	UserCanStop    bool    `json:"userCanStop"`
}

type ItemMoveRequestData struct {
	ItemsToMove []ItemToMove `json:"itemsToMove"`
}

type ItemMoveRequest struct {
	RequestMessageBase
	ItemMoveRequestData ItemMoveRequestData `json:"data"`
}

type MovedItem struct {
	ItemInstanceID string `json:"itemInstanceID"`
	Success        bool   `json:"success"`
	ErrorID        int    `json:"errorId"`
}

type ItemMoveResponseData struct {
	MovedItems []MovedItem `json:"movedItems"`
}

type ItemMoveResponse struct {
	RespondMessageBase
	ItemMoveResponseData ItemMoveResponseData `json:"data"`
}

type ArtMeshSelectionRequestData struct {
	TextOverride          string   `json:"textOverride"`
	HelpOverride          string   `json:"helpOverride"`
	RequestedArtMeshCount int      `json:"requestedArtMeshCount"`
	ActiveArtMeshes       []string `json:"activeArtMeshes"`
}

type ArtMeshSelectionRequest struct {
	RequestMessageBase
	ArtMeshSelectionRequestData ArtMeshSelectionRequestData `json:"data"`
}

type ArtMeshSelectionResponseData struct {
	Success           bool     `json:"success"`
	ActiveArtMeshes   []string `json:"activeArtMeshes"`
	InactiveArtMeshes []string `json:"inactiveArtMeshes"`
}

type ArtMeshSelectionResponse struct {
	RespondMessageBase
	ArtMeshSelectionResponseData ArtMeshSelectionResponseData `json:"data"`
}

type PinInfo struct {
	ModelID       string  `json:"modelID"`
	ArtMeshID     string  `json:"artMeshID"`
	Angle         float32 `json:"angle"`
	Size          float32 `json:"size"`
	VertexID1     int     `json:"vertexID1"`
	VertexID2     int     `json:"vertexID2"`
	VertexID3     int     `json:"vertexID3"`
	VertexWeight1 float32 `json:"vertexWeight1"`
	VertexWeight2 float32 `json:"vertexWeight2"`
	VertexWeight3 float32 `json:"vertexWeight3"`
}

type ItemPinRequestData struct {
	Pin             bool    `json:"pin"`
	ItemInstanceID  string  `json:"itemInstanceID"`
	AngleRelativeTo string  `json:"angleRelativeTo"`
	SizeRelativeTo  string  `json:"sizeRelativeTo"`
	VertexPinType   string  `json:"vertexPinType"`
	PinInfo         PinInfo `json:"pinInfo"`
}

type ItemPinRequest struct {
	RequestMessageBase
	ItemPinRequestData ItemPinRequestData `json:"data"`
}
type ItemPinRequestResponseData struct {
	IsPinned       bool   `json:"IsPinned"`
	ItemInstanceID string `json:"itemInstanceID"`
	ItemFileName   string `json:"itemFileName"`
}

type ItemPinResponse struct {
	RespondMessageBase
	ItemAnimationControlResponseData ItemAnimationControlResponseData `json:"data"`
}

type PostProcessingListRequestData struct {
	FillPostProcessingPresetArray  bool     `json:"fillPostProcessingPresetArray"`
	FillPostProcessingEffectsArray bool     `json:"fillPostProcessingEffectsArray"`
	EffectIDFilter                 []string `json:"effectIDFilter"`
}

type PostProcessingListRequest struct {
	RequestMessageBase
	PostProcessingListRequestData PostProcessingListRequestData `json:"data"`
}

type ConfigEntry struct {
	InternalID       string  `json:"internalID"`
	EnumID           string  `json:"enumID"`
	Explanation      string  `json:"explanation"`
	Type             string  `json:"type"`
	ActivationConfig bool    `json:"activationConfig"`
	FloatValue       float64 `json:"floatValue"`
	FloatMin         float32 `json:"floatMin"`
	FloatMax         float32 `json:"floatMax"`
	FloatDefault     float32 `json:"floatDefault"`
	IntValue         int     `json:"intValue"`
	IntMin           int     `json:"intMin"`
	IntMax           int     `json:"intMax"`
	IntDefault       int     `json:"intDefault"`
	ColorValue       string  `json:"colorValue"`
	ColorDefault     string  `json:"colorDefault"`
	ColorHasAlpha    bool    `json:"colorHasAlpha"`
	BoolValue        bool    `json:"boolValue"`
	BoolDefault      bool    `json:"boolDefault"`
	StringValue      string  `json:"stringValue"`
	StringDefault    string  `json:"stringDefault"`
	SceneItemValue   string  `json:"sceneItemValue"`
	SceneItemDefault string  `json:"sceneItemDefault"`
}

type PostProcessingEffect struct {
	InternalID         string        `json:"internalID"`
	EnumID             string        `json:"enumID"`
	Explanation        string        `json:"explanation"`
	EffectIsActivated  bool          `json:"effectIsActivated"`
	EffectIsRestricted bool          `json:"effectIsRestricted"`
	ConfigEntries      []ConfigEntry `json:"configEntries"`
}

type PostProcessingListResponseData struct {
	PostProcessingSupported                    bool                   `json:"PostProcessingSupported"`
	PostProcessingActive                       bool                   `json:"PostProcessingActive"`
	CanSendPostProcessingUpdateRequestRightNow bool                   `json:"CanSendPostProcessingUpdateRequestRightNow"`
	RestrictedEffectsAllowed                   bool                   `json:"RestrictedEffectsAllowed"`
	PresetIsActive                             bool                   `json:"PresetIsActive"`
	ActivePReset                               string                 `json:"ActivePReset"`
	PresetCount                                int                    `json:"PresetCount"`
	ActiveEffectCount                          int                    `json:"ActiveEffectCount"`
	EffectCountBeforeFilter                    int                    `json:"EffectCountBeforeFilter"`
	ConfigCountBeforeFilter                    int                    `json:"ConfigCountBeforeFilter"`
	EffectCountAfterFilter                     int                    `json:"EffectCountAfterFilter"`
	ConfigCountAfterFilter                     int                    `json:"ConfigCountAfterFilter"`
	PostProcessingEffects                      []PostProcessingEffect `json:"PostProcessingEffects"`
	PostProcessingPresets                      []string               `json:"PostProcessingPresets"`
}

type PostProcessingListResponse struct {
	RespondMessageBase
	PostProcessingListResponseData PostProcessingListResponseData `json:"data"`
}

type PostProcessingValues struct {
	ConfigID    string  `json:"configID"`
	ConfigValue float32 `json:"configValue"`
}

type PostProcessingUpdateRequestData struct {
	PostProcessingOn           bool                   `json:"postProcessingOn"`
	SetPostProcessingPReset    bool                   `json:"setPostProcessingPReset"`
	SetPostProcessingValues    bool                   `json:"setPostProcessingValues"`
	PresetToSet                string                 `json:"presetToSet"`
	PostProcessingFadeTime     float32                `json:"postProcessingFadeTime"`
	SetAllOtherValuesToDefault bool                   `json:"setAllOtherValuesToDefault"`
	UsingRestrictedEffects     bool                   `json:"UsingRestrictedEffects"`
	RandomizeAll               bool                   `json:"randomizeAll"`
	RandomizeAllChaosLevel     float32                `json:"randomizeAllChaosLevel"`
	PostProcessingValues       []PostProcessingValues `json:"postProcessingValues"`
}

type PostProcessingUpdateRequest struct {
	RequestMessageBase
	PostProcessingUpdateRequestData PostProcessingUpdateRequestData `json:"data"`
}

type PostProcessingUpdateResponseData struct {
	PostProcessingActive bool   `json:"PostProcessingActive"`
	PresetIsActive       bool   `json:"PresetIsActive"`
	ActivePReset         string `json:"ActivePReset"`
	ActiveEffectCount    int    `json:"ActiveEffectCount"`
}

type PostProcessingUpdateResponse struct {
	RespondMessageBase
	PostProcessingUpdateResponseData PostProcessingUpdateResponseData `json:"data"`
}
