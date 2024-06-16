package commonwealth

// Standard subject/topics
const (
	Status           = "status"
	Database         = "database"
	Query            = "query"
	CookieDatabase   = "cookiedb"
	CompanyIndex     = 0
	ApplicationIndex = 1
	CommandIndex     = 2
)

// NatsTopics are topics that will be handled by each service
type NatsTopics struct {
	Name   string `yaml:"name"`   // the name of the topic to subscribe to (subject)
	Create bool   `yaml:"create"` // the create verb if needed or not ignored if always
	Read   bool   `yaml:"read"`   // the read verb if needed or not ignored if always
	Update bool   `yaml:"update"` // the update verb if needed or not ignored if always
	Delete bool   `yaml:"delete"` // the delete verb if needed or not ignored if always
	Always bool   `yaml:"always"` // the always verb that subscribes to the given name only
}
