package juju

type EnvironInfoData struct {
	User         string
	Password     string
	EnvironUUID  string                 `yaml:"environ-uuid,omitempty"`
	StateServers []string               `yaml:"state-servers"`
	CACert       string                 `yaml:"ca-cert"`
	Config       map[string]interface{} `yaml:"bootstrap-config,omitempty"`
}
