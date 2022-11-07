package domain

type WebsocketPayload struct {
	Nickname string                  `json:"nickname,omitempty"`
	GameID   string                  `json:"gameId,omitempty"`
	Game     Game                    `json:"game"`
	Value    Move                    `json:"value,omitempty"`
	Command  WebsocketPayloadCommand `json:"command,omitempty"`
	Message  string                  `json:"message"`
}

type WebsocketPayloadCommand string

const (
	WebsocketPayloadCommandConnect WebsocketPayloadCommand = "connect"
	WebsocketPayloadCommandPlay    WebsocketPayloadCommand = "play"
	WebsocketPayloadCommandOK      WebsocketPayloadCommand = "ok"
	WebsocketPayloadCommandError   WebsocketPayloadCommand = "error"
	WebsocketPayloadCommandClose   WebsocketPayloadCommand = "close"
	WebsocketPayloadCommandPause   WebsocketPayloadCommand = "pause"
	WebsocketPayloadCommandUnpause WebsocketPayloadCommand = "unpause"
)
