package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// AddMediaIntentHandlingConfirmInvocationResponse represents the AddMediaIntentHandlingConfirmInvocationResponse schema from the OpenAPI specification
type AddMediaIntentHandlingConfirmInvocationResponse struct {
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
	Debug string `json:"debug,omitempty"`
}

// PlayMediaIntentResponse represents the PlayMediaIntentResponse schema from the OpenAPI specification
type PlayMediaIntentResponse struct {
	Useractivity UserActivity `json:"userActivity"`
	Class string `json:"class"`
}

// AddMediaMediaDestinationResolutionResult represents the AddMediaMediaDestinationResolutionResult schema from the OpenAPI specification
type AddMediaMediaDestinationResolutionResult struct {
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
}

// BooleanResolutionResult represents the BooleanResolutionResult schema from the OpenAPI specification
type BooleanResolutionResult struct {
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
}

// MediaDestinationLibrary represents the MediaDestinationLibrary schema from the OpenAPI specification
type MediaDestinationLibrary struct {
	Mediadestinationtype string `json:"mediaDestinationType"`
}

// AddMediaIntent represents the AddMediaIntent schema from the OpenAPI specification
type AddMediaIntent struct {
	Class string `json:"class"`
	Identifier string `json:"identifier"`
}

// ProtocolExceptionInvocationResponse represents the ProtocolExceptionInvocationResponse schema from the OpenAPI specification
type ProtocolExceptionInvocationResponse struct {
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
	Debug string `json:"debug,omitempty"`
}

// ExplicitDateComponents represents the ExplicitDateComponents schema from the OpenAPI specification
type ExplicitDateComponents struct {
	Month int `json:"month,omitempty"`
	Nanosecond int `json:"nanosecond,omitempty"`
	Second int `json:"second,omitempty"`
	Timezone string `json:"timeZone,omitempty"`
	Day int `json:"day,omitempty"`
	Era int `json:"era,omitempty"`
	Minute int `json:"minute,omitempty"`
	Calendaridentifier string `json:"calendarIdentifier,omitempty"`
	Hour int `json:"hour,omitempty"`
	Year int `json:"year,omitempty"`
}

// PlayMediaRequest represents the PlayMediaRequest schema from the OpenAPI specification
type PlayMediaRequest struct {
	Useractivity UserActivity `json:"userActivity"`
	Version string `json:"version"`
	Constraints Constraints `json:"constraints"`
}

// MediaDestinationPlaylist represents the MediaDestinationPlaylist schema from the OpenAPI specification
type MediaDestinationPlaylist struct {
	Mediadestinationtype string `json:"mediaDestinationType"`
}

// PlayMediaControlCommandSet represents the PlayMediaControlCommandSet schema from the OpenAPI specification
type PlayMediaControlCommandSet struct {
	Seektoplaybackposition bool `json:"seekToPlaybackPosition,omitempty"`
	Skipforward bool `json:"skipForward,omitempty"`
	Skipbackward bool `json:"skipBackward,omitempty"`
	Nexttrack bool `json:"nextTrack,omitempty"`
	Preferskipbackward bool `json:"preferSkipBackward,omitempty"`
	Previoustrack bool `json:"previousTrack,omitempty"`
	Liketrack bool `json:"likeTrack,omitempty"`
	Bookmarktrack bool `json:"bookmarkTrack,omitempty"`
	Disliketrack bool `json:"dislikeTrack,omitempty"`
	Preferskipforward bool `json:"preferSkipForward,omitempty"`
}

// QueuePlayPointer represents the QueuePlayPointer schema from the OpenAPI specification
type QueuePlayPointer struct {
	Contentidentifier string `json:"contentIdentifier,omitempty"`
	Offsetinmillis int64 `json:"offsetInMillis,omitempty"`
}

// UpdateMediaAffinityIntentHandlingInvocation represents the UpdateMediaAffinityIntentHandlingInvocation schema from the OpenAPI specification
type UpdateMediaAffinityIntentHandlingInvocation struct {
	Method string `json:"method"`
	Params map[string]interface{} `json:"params"`
	Session Session `json:"session,omitempty"`
}

// PlayMediaIntentHandlingResolvePlaybackRepeatModeInvocationResponse represents the PlayMediaIntentHandlingResolvePlaybackRepeatModeInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingResolvePlaybackRepeatModeInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// Queue represents the Queue schema from the OpenAPI specification
type Queue struct {
	Controls QueueControlMapping `json:"controls,omitempty"`
	Prerollseconds float64 `json:"prerollSeconds,omitempty"`
	Nextcontenturl string `json:"nextContentUrl,omitempty"`
	Skipsremaining int `json:"skipsRemaining,omitempty"`
	Contentitemscount int `json:"contentItemsCount,omitempty"`
	Identifier string `json:"identifier"`
	Playpointer QueuePlayPointer `json:"playPointer,omitempty"`
	Version string `json:"version"`
	Content []Content `json:"content"`
	Insertpointer QueueInsertPointer `json:"insertPointer,omitempty"`
	Previouscontenturl string `json:"previousContentUrl,omitempty"`
}

// UpdateMediaAffinityIntentHandlingResolveAffinityTypeInvocationResponse represents the UpdateMediaAffinityIntentHandlingResolveAffinityTypeInvocationResponse schema from the OpenAPI specification
type UpdateMediaAffinityIntentHandlingResolveAffinityTypeInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// PlayerContext represents the PlayerContext schema from the OpenAPI specification
type PlayerContext struct {
	Activityidentifier string `json:"activityIdentifier,omitempty"`
	Contentidentifier string `json:"contentIdentifier,omitempty"`
	Offsetinmillis int64 `json:"offsetInMillis,omitempty"`
	Playbackspeed float64 `json:"playbackSpeed,omitempty"`
	Queueidentifier string `json:"queueIdentifier,omitempty"`
}

// IntentResolutionResult represents the IntentResolutionResult schema from the OpenAPI specification
type IntentResolutionResult struct {
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
}

// UserActivity represents the UserActivity schema from the OpenAPI specification
type UserActivity struct {
	Title string `json:"title,omitempty"`
	Userinfo map[string]interface{} `json:"userInfo,omitempty"`
	Version string `json:"version"`
	Activitytype string `json:"activityType"`
	Persistentidentifier string `json:"persistentIdentifier,omitempty"`
}

// PlayMediaIntent represents the PlayMediaIntent schema from the OpenAPI specification
type PlayMediaIntent struct {
	Class string `json:"class"`
	Identifier string `json:"identifier"`
}

// PlayMediaControl represents the PlayMediaControl schema from the OpenAPI specification
type PlayMediaControl struct {
	Activity PlayMediaControlActivity `json:"activity,omitempty"`
	Commands PlayMediaControlCommandSet `json:"commands,omitempty"`
	Scheme string `json:"scheme"`
}

// MediaSearch represents the MediaSearch schema from the OpenAPI specification
type MediaSearch struct {
	Medianame string `json:"mediaName,omitempty"`
	Moodnames []string `json:"moodNames,omitempty"`
	Releasedate DateComponentsRange `json:"releaseDate,omitempty"`
	Sortorder string `json:"sortOrder,omitempty"`
	Mediaidentifier string `json:"mediaIdentifier,omitempty"`
	Mediatype string `json:"mediaType,omitempty"`
	Reference string `json:"reference,omitempty"`
	Albumname string `json:"albumName,omitempty"`
	Artistname string `json:"artistName,omitempty"`
	Genrenames []string `json:"genreNames,omitempty"`
}

// PlayMediaIntentHandlingResolvePlayShuffledInvocationResponse represents the PlayMediaIntentHandlingResolvePlayShuffledInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingResolvePlayShuffledInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// UpdateActivityResponse represents the UpdateActivityResponse schema from the OpenAPI specification
type UpdateActivityResponse struct {
	Queue Queue `json:"queue,omitempty"`
	Useractivity UserActivity `json:"userActivity,omitempty"`
}

// UpdateMediaAffinityIntent represents the UpdateMediaAffinityIntent schema from the OpenAPI specification
type UpdateMediaAffinityIntent struct {
	Class string `json:"class"`
	Identifier string `json:"identifier"`
}

// AddMediaIntentHandlingResolveMediaDestinationInvocationResponse represents the AddMediaIntentHandlingResolveMediaDestinationInvocationResponse schema from the OpenAPI specification
type AddMediaIntentHandlingResolveMediaDestinationInvocationResponse struct {
	Result map[string]interface{} `json:"result"`
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
}

// DateComponentsRange represents the DateComponentsRange schema from the OpenAPI specification
type DateComponentsRange struct {
	Enddatecomponents interface{} `json:"endDateComponents,omitempty"`
	Startdatecomponents interface{} `json:"startDateComponents,omitempty"`
}

// PlaybackQueueLocationResolutionResult represents the PlaybackQueueLocationResolutionResult schema from the OpenAPI specification
type PlaybackQueueLocationResolutionResult struct {
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
	Class string `json:"class"`
}

// ExecutionMetrics represents the ExecutionMetrics schema from the OpenAPI specification
type ExecutionMetrics struct {
	Completed string `json:"completed,omitempty"`
	Duration float32 `json:"duration,omitempty"`
	Received string `json:"received,omitempty"`
}

// AddMediaMediaItemResolutionResult represents the AddMediaMediaItemResolutionResult schema from the OpenAPI specification
type AddMediaMediaItemResolutionResult struct {
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
}

// PlayMediaControlActivity represents the PlayMediaControlActivity schema from the OpenAPI specification
type PlayMediaControlActivity struct {
	Playelapsed int `json:"playElapsed,omitempty"`
	Playelapsedinterval int `json:"playElapsedInterval,omitempty"`
	Playpaused int `json:"playPaused,omitempty"`
}

// PlayMediaIntentHandlingInvocation represents the PlayMediaIntentHandlingInvocation schema from the OpenAPI specification
type PlayMediaIntentHandlingInvocation struct {
	Method string `json:"method"`
	Params map[string]interface{} `json:"params"`
	Session Session `json:"session,omitempty"`
}

// Content represents the Content schema from the OpenAPI specification
type Content struct {
	Url string `json:"url,omitempty"`
	Attributes ContentAttributes `json:"attributes,omitempty"`
	Control string `json:"control,omitempty"`
	Identifier string `json:"identifier"`
	Islive bool `json:"isLive,omitempty"`
	Playindex int `json:"playIndex,omitempty"`
}

// UpdateMediaAffinityIntentHandlingResolveMediaItemsInvocationResponse represents the UpdateMediaAffinityIntentHandlingResolveMediaItemsInvocationResponse schema from the OpenAPI specification
type UpdateMediaAffinityIntentHandlingResolveMediaItemsInvocationResponse struct {
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
}

// PlayMediaMediaItemResolutionResult represents the PlayMediaMediaItemResolutionResult schema from the OpenAPI specification
type PlayMediaMediaItemResolutionResult struct {
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
}

// UpdateMediaAffinityIntentResponse represents the UpdateMediaAffinityIntentResponse schema from the OpenAPI specification
type UpdateMediaAffinityIntentResponse struct {
	Useractivity UserActivity `json:"userActivity"`
	Class string `json:"class"`
}

// PlayMediaIntentHandlingResolvePlaybackQueueLocationInvocationResponse represents the PlayMediaIntentHandlingResolvePlaybackQueueLocationInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingResolvePlaybackQueueLocationInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// IntentResponse represents the IntentResponse schema from the OpenAPI specification
type IntentResponse struct {
	Useractivity UserActivity `json:"userActivity"`
	Class string `json:"class"`
}

// UpdateMediaAffinityIntentHandlingHandleInvocationResponse represents the UpdateMediaAffinityIntentHandlingHandleInvocationResponse schema from the OpenAPI specification
type UpdateMediaAffinityIntentHandlingHandleInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// PlayMediaIntentHandlingResolveResumePlaybackInvocationResponse represents the PlayMediaIntentHandlingResolveResumePlaybackInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingResolveResumePlaybackInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// Constraints represents the Constraints schema from the OpenAPI specification
type Constraints struct {
	Allowexplicitcontent bool `json:"allowExplicitContent,omitempty"`
	Maximumqueuesegmentitemcount int `json:"maximumQueueSegmentItemCount,omitempty"`
	Updateusertasteprofile bool `json:"updateUserTasteProfile,omitempty"`
}

// Session represents the Session schema from the OpenAPI specification
type Session struct {
	Identifier string `json:"identifier"`
	Playercontext PlayerContext `json:"playerContext,omitempty"`
	Requested string `json:"requested"`
	Version string `json:"version"`
	Constraints Constraints `json:"constraints"`
	Deadline string `json:"deadline"`
}

// InvocationResponse represents the InvocationResponse schema from the OpenAPI specification
type InvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// PlaybackRepeatModeResolutionResult represents the PlaybackRepeatModeResolutionResult schema from the OpenAPI specification
type PlaybackRepeatModeResolutionResult struct {
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
}

// ExtensionConfig represents the ExtensionConfig schema from the OpenAPI specification
type ExtensionConfig struct {
	Url string `json:"url,omitempty"`
	Version string `json:"version"`
	Hdr map[string]interface{} `json:"hdr,omitempty"`
	Intent map[string]interface{} `json:"intent"`
	Media map[string]interface{} `json:"media"`
}

// AddMediaIntentResponse represents the AddMediaIntentResponse schema from the OpenAPI specification
type AddMediaIntentResponse struct {
	Useractivity UserActivity `json:"userActivity"`
	Class string `json:"class"`
}

// ContentAttributes represents the ContentAttributes schema from the OpenAPI specification
type ContentAttributes struct {
	Durationinmillis int `json:"durationInMillis,omitempty"`
	Genrenames []string `json:"genreNames,omitempty"`
	Name string `json:"name,omitempty"`
	Tracknumber int `json:"trackNumber,omitempty"`
	Albumname string `json:"albumName,omitempty"`
	Artistname string `json:"artistName,omitempty"`
	Artwork map[string]interface{} `json:"artwork,omitempty"`
	Composername string `json:"composerName,omitempty"`
}

// MediaAffinityTypeResolutionResult represents the MediaAffinityTypeResolutionResult schema from the OpenAPI specification
type MediaAffinityTypeResolutionResult struct {
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
	Class string `json:"class"`
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
}

// PlayMediaIntentHandlingResolveMediaItemsInvocationResponse represents the PlayMediaIntentHandlingResolveMediaItemsInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingResolveMediaItemsInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// UpdateMediaAffinityMediaItemResolutionResult represents the UpdateMediaAffinityMediaItemResolutionResult schema from the OpenAPI specification
type UpdateMediaAffinityMediaItemResolutionResult struct {
	Needsvalue map[string]interface{} `json:"needsValue,omitempty"`
	Notrequired map[string]interface{} `json:"notRequired,omitempty"`
	Unsupported map[string]interface{} `json:"unsupported,omitempty"`
	Class string `json:"class"`
}

// AddMediaIntentHandlingHandleInvocationResponse represents the AddMediaIntentHandlingHandleInvocationResponse schema from the OpenAPI specification
type AddMediaIntentHandlingHandleInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// Invocation represents the Invocation schema from the OpenAPI specification
type Invocation struct {
	Method string `json:"method"`
	Params map[string]interface{} `json:"params"`
	Session Session `json:"session,omitempty"`
}

// PlayMediaIntentHandlingHandleInvocationResponse represents the PlayMediaIntentHandlingHandleInvocationResponse schema from the OpenAPI specification
type PlayMediaIntentHandlingHandleInvocationResponse struct {
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
	Debug string `json:"debug,omitempty"`
}

// ExtensionEndpointConfig represents the ExtensionEndpointConfig schema from the OpenAPI specification
type ExtensionEndpointConfig struct {
	Hdr map[string]interface{} `json:"hdr,omitempty"`
	Url string `json:"url,omitempty"`
}

// AddMediaIntentHandlingInvocation represents the AddMediaIntentHandlingInvocation schema from the OpenAPI specification
type AddMediaIntentHandlingInvocation struct {
	Params map[string]interface{} `json:"params"`
	Session Session `json:"session,omitempty"`
	Method string `json:"method"`
}

// AddMediaIntentHandlingResolveMediaItemsInvocationResponse represents the AddMediaIntentHandlingResolveMediaItemsInvocationResponse schema from the OpenAPI specification
type AddMediaIntentHandlingResolveMediaItemsInvocationResponse struct {
	Debug string `json:"debug,omitempty"`
	Method string `json:"method"`
	Metrics ExecutionMetrics `json:"metrics,omitempty"`
	Result map[string]interface{} `json:"result"`
}

// ProtocolException represents the ProtocolException schema from the OpenAPI specification
type ProtocolException struct {
	Retrywithdelay float32 `json:"retryWithDelay,omitempty"`
	Trace []string `json:"trace,omitempty"`
	Code int64 `json:"code,omitempty"`
	Methodindex int `json:"methodIndex,omitempty"`
	Methodname string `json:"methodName,omitempty"`
	Reason string `json:"reason"`
}

// UpdateActivityRequest represents the UpdateActivityRequest schema from the OpenAPI specification
type UpdateActivityRequest struct {
	Useractivity UserActivity `json:"userActivity"`
	Version string `json:"version"`
	Constraints Constraints `json:"constraints,omitempty"`
	Nowplaying PlayerContext `json:"nowPlaying,omitempty"`
	Previouslyplaying PlayerContext `json:"previouslyPlaying,omitempty"`
	Report string `json:"report"`
	Timestamp string `json:"timestamp"`
}

// QueueControlMapping represents the QueueControlMapping schema from the OpenAPI specification
type QueueControlMapping struct {
	DefaultField PlayMediaControl `json:"default"`
}

// MediaItem represents the MediaItem schema from the OpenAPI specification
type MediaItem struct {
	Artist string `json:"artist,omitempty"`
	Identifier string `json:"identifier"`
	Title string `json:"title,omitempty"`
	TypeField string `json:"type"`
}

// QueueInsertPointer represents the QueueInsertPointer schema from the OpenAPI specification
type QueueInsertPointer struct {
	Afteridentifier string `json:"afterIdentifier,omitempty"`
	Replace bool `json:"replace,omitempty"`
}

// MediaDestination represents the MediaDestination schema from the OpenAPI specification
type MediaDestination struct {
	Mediadestinationtype string `json:"mediaDestinationType"`
}

// Intent represents the Intent schema from the OpenAPI specification
type Intent struct {
	Class string `json:"class"`
	Identifier string `json:"identifier"`
}
