package nvidiasmi

type Power struct {
	Draw float32 `json:"draw"`
}

type Utilization struct {
	GPU    uint8 `json:"gpu"`
	Memory uint8 `json:"memory"`
}

type Fan struct {
	Speed uint32 `json:"speed"`
}

type Temperature struct {
	GPU float32 `json:"gpu"`
}

type Memory struct {
	Total float32 `json:"total"`
	Free  float32 `json:"free"`
	Used  float32 `json:"used"`
}

type Clocks struct {
	CurrentGraphics uint32 `json:"current_graphics"`
	CurrentMemory   uint32 `json:"current_memory"`
}

type StaticNvidiaSMIInfo struct {
	Index uint8  `json:"index"`
	Name  string `json:"name"`
}

type DynamicNvidiaSMIInfo struct {
	Utilization Utilization `json:"utilization"`
	Power       Power       `json:"power"`
	Temperature Temperature `json:"temperature"`
	Fan         Fan         `json:"fan"`
	Memory      Memory      `json:"memory"`
}

type NvidiaSMIInfo struct {
	StaticNvidiaSMIInfo  StaticNvidiaSMIInfo  `json:"static"`
	DynamicNvidiaSMIInfo DynamicNvidiaSMIInfo `json:"dynamic"`
}
