package minecraft

// ServerInfo represents info about a minecraft server
type ServerInfo struct {
	Version     Ver     `json:"version"`
	Players     Players `json:"players"`
	Description Chat    `json:"description"`
	Icon        string  `json:"favicon"`
}

// Chat is a retarded struct that contains chat messages, in all their retarded glory
type Chat struct {
	Text          string `json:"text"`
	Bold          bool   `json:"bold"`
	Italic        bool   `json:"italic"`
	Underlined    bool   `json:"underlined"`
	Strikethrough bool   `json:"strikethrough"`
	Obfuscated    bool   `json:"obfuscated"`
	Color         string `json:"color"`
	Extra         []Chat `json:"extra"`
}

// Ver contains info about the version of the server
type Ver struct {
	Name     string `json:"name"`     // 1.16.4 Paper
	Protocol int    `json:"protocol"` // <protocol version>
}

// Players contains info about the players on a server
type Players struct {
	MaxPlayers int      `json:"max"`
	CurPlayers int      `json:"online"`
	Sample     []Player `json:"sample"`
}

// Player represents a player on the server
type Player struct {
	Name string `json:"name"`
	UUID string `json:"id"`
}
