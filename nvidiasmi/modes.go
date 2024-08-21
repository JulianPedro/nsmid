package nvidiasmi

import "fmt"

var (
	StaticInfoCommand  string = "index,name"
	DynamicInfoCommand string = "utilization.gpu,utilization.memory,power.draw,temperature.gpu,memory.total,memory.free,memory.used"
	AllInfoCommand     string = fmt.Sprintf("%s,%s", StaticInfoCommand, DynamicInfoCommand)
)
