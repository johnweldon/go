package quassel

type QuasselUser struct {
	Id       int
	Username string
	Password string
}

type Sender struct {
	Id   int
	Name string
}

type Network struct {
	Id                      int
	UserId                  int
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
	Id                  int
	UserId              int
	GroupId             int
	NetworkId           int
	Name                string
	CanonicalName       string
	Type                int
	LastSeenMessageId   int
	MarkerLineMessageId int
	Key                 string
	Joined              bool
}

type Backlog struct {
	MessageId int
	Time      int
	BufferId  int
	Type      int
	Flags     int
	SenderId  int
	Message   string
}

type CoreInfo struct {
	Key   string
	Value string
}

type IRCServer struct {
	ServerId   int
	UserId     int
	NetworkId  int
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
	UserId       int
	SettingName  string
	SettingValue []byte
}

type Identity struct {
	Id                      int
	UserId                  int
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
	Id         int
	IdentityId int
	Nick       string
}
