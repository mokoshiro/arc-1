package trackerruntime

type ContainerDBConfig struct {
	Image         string   `json:"image"`
	Command       []string `json:"command"`
	Env           []string `json:"env"`
	ContainerPort string   `json:"container_port"`
	HostPort      string   `json:"host_port"`
}
