package quassel

type User struct {
	ID       int
	Username string
	Password string
}

type Sender struct {
	ID   int
	Name string
}

type Network struct {
	ID                      int
	UserID                  int
	Name                    string
	Identity                int
	EncodingCodec           string
	DecodingCodec           string
	ServerCodec             string
	UseRandomServer         bool
	Perform                 string
	UseAutoIdentify         bool
	AutoIdentifyService     string
	AutoIdentifyPassword    string
	UseSASL                 bool
	SASLAccount             string
	SASLPassword            string
	UseAutoReconnect        bool
	AutoReconnectInterval   int
	AutoReconnectRetries    int
	UnlimitedConnectRetries bool
	RejoinChannels          bool
	Connected               bool
	UserMode                string
	AwayMessage             string
	AttachPerform           string
	DetachPerform           string
}

type Buffer struct {
	ID                  int
	UserID              int
	GroupID             int
	NetworkID           int
	Name                string
	CanonicalName       string
	Type                int
	LastSeenMessageID   int
	MarkerLineMessageID int
	Key                 string
	Joined              bool
}

type Backlog struct {
	MessageID int
	Time      int
	BufferID  int
	Type      int
	Flags     int
	SenderID  int
	Message   string
}

type CoreInfo struct {
	Key   string
	Value string
}

type IRCServer struct {
	ServerID   int
	UserID     int
	NetworkID  int
	Hostname   string
	Port       int
	Password   string
	SSL        bool
	SSLVersion int
	UseProxy   bool
	ProxyType  int
	ProxyHost  string
	ProxyPort  int
	ProxyUser  string
	ProxyPass  string
}

type UserSetting struct {
	UserID       int
	SettingName  string
	SettingValue []byte
}

type Identity struct {
	ID                      int
	UserID                  int
	Name                    string
	RealName                string
	AwayNick                string
	AwayNickEnabled         bool
	AwayReason              string
	AwayReasonEnabled       bool
	AutoAwayEnabled         bool
	AutoAwayTime            int
	AutoAwayReason          string
	AutoAwayReasonEnabled   bool
	DetachAwayTime          int
	DetachAwayReason        string
	DetachAwayReasonEnabled bool
	Ident                   string
	KickReason              string
	PartReason              string
	QuitReason              string
	SSLCert                 []byte
	SSLKey                  []byte
}

type IdentityNick struct {
	ID         int
	IdentityID int
	Nick       string
}
