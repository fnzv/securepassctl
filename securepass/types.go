package securepass

// APIResponse gives access to the response details
type APIResponse interface {
	ErrorCode() int
	ErrorMessage() string
}

// ApplicationDescriptor describes the basic attributes of
// a Securepass application
type ApplicationDescriptor struct {
	Label            string `json:"label"`
	Realm            string `json:"realm"`
	Group            string `json:"group"`
	Write            bool   `json:"write"`
	AllowNetworkIPv4 string `json:"allow_network_ipv4"`
	AllowNetworkIPv6 string `json:"allow_network_ipv6"`
	Privacy          bool   `json:"privacy"`
}

// UserDescriptor defines the attributes of SecurePass users
type UserDescriptor struct {
	Username string
	Name     string
	Surname  string
	Email    string
	Mobile   string
	Nin      string
	Rfid     string
	Manager  string
	Type     string
	Password bool
	Enabled  bool
	Token    string
}

// Response is the base type for API calls responses
type Response struct {
	APIResponse
	ErrorMsg string
	RC       int
}

// ErrorCode returns the API call's numeric return code
func (r *Response) ErrorCode() int {
	return r.RC
}

// ErrorMessage returns the API call's text message
func (r *Response) ErrorMessage() string {
	return r.ErrorMsg
}

// PingResponse represents the /api/v1/ping call's HTTP response
type PingResponse struct {
	IP        string
	IPVersion int `json:"ip_version"`
	Response
}

// AppAddResponse describes the expected response from the
// /api/v1/apps/add
type AppAddResponse struct {
	AppID     string `json:"APP_ID"`
	AppSecret string `json:"APP_SECRET"`
	Response
}

// AppInfoResponse encapsulates the /api/v1/apps/info call's HTTP response
type AppInfoResponse struct {
	ApplicationDescriptor
	Response
}

// AppListResponse encapsulates the /api/v1/apps HTTP response
type AppListResponse struct {
	AppID []string `json:"APP_ID"`
	Response
}

// LogEntry is a SecurePass application's log entry
type LogEntry struct {
	// SecurePass response is currently broken, this
	// should be a time.Time object.
	Timestamp string
	UUID      string
	Message   string
	Level     int
	App       string
	Realm     string
}

// LogEntriesByTimestamp sorts log entries by timestamp
type LogEntriesByTimestamp []LogEntry

func (l LogEntriesByTimestamp) Less(i, j int) bool { return l[i].Timestamp < l[j].Timestamp }
func (l LogEntriesByTimestamp) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l LogEntriesByTimestamp) Len() int           { return len(l) }

// LogsResponse encapsulates SecurePass application's logs
type LogsResponse struct {
	Logs map[string]LogEntry
	Response
}

// GroupMemberResponse encapsulates whether a group belogs to a member
type GroupMemberResponse struct {
	Member bool
	Response
}

// UserInfoResponse encapsulates the /api/v1/users/info HTTP response
type UserInfoResponse struct {
	UserDescriptor
	Response
}

// UserListResponse encapsulates the /api/v1/users HTTP response
type UserListResponse struct {
	Username []string
	Response
}

// NSSConfig encapsulates the SecurePass's config section '[nss]'
type NSSConfig struct {
	Realm        string `ini:"realm"`
	DefaultGid   int    `ini:"default_gid"`
	DefaultHome  string `ini:"default_home"`
	DefaultShell string `ini:"default_shell"`
}

// SSHConfig encapsulates the SecurePass's config section '[ssh]'
type SSHConfig struct {
	Root string `ini:"root"`
}

// GlobalConfig encapsulates the SecurePass's whole configuration
type GlobalConfig struct {
	SecurePass `ini:"default"`
	NSSConfig  `ini:"nss"`
	SSHConfig  `ini:"ssh"`
}
