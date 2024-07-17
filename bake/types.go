package bake

type File struct {
	Groups  map[string]Group  `json:"group"`
	Targets map[string]Target `json:"target"`
}

type Group struct {
	Targets []string `json:"targets"`
}

type Target struct {
	Args             map[string]string `json:"args,omitempty"`
	Context          string            `json:"context"`
	Contexts         map[string]string `json:"contexts,omitempty"`
	CacheFrom        []string          `json:"cache-from,omitempty"`
	CacheTo          []string          `json:"cache-to,omitempty"`
	Dockerfile       string            `json:"dockerfile"`
	DockerfileInline string            `json:"dockerfile-inline,omitempty"`
	Outputs          []string          `json:"output,omitempty"`
	Tags             []string          `json:"tags,omitempty"`
	Target           string            `json:"target,omitempty"`
}
