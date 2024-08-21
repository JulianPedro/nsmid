package nvidiasmi

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func GetNvidiaSMIOutput(query string) (string, string, error) {
	cmd := exec.Command("nvidia-smi", fmt.Sprintf("--query-gpu=%s", query), "--format=csv")
	output, err := cmd.Output()
	if err != nil {
		return "", "", err
	}
	return string(output), query, nil
}

func ParseNvidiaSMIOutput(output string, query string, data *NvidiaSMIInfo) {
	resultValue := strings.Split(output, "\n")
	keys := strings.Split(query, ",")
	values := strings.Split(resultValue[1], ", ")

	dataMap := make(map[string]string)
	for i, key := range keys {
		dataMap[key] = values[i]
	}

	fmt.Printf("dataMap: %v\n", dataMap)

	if index, err := strconv.Atoi(dataMap["index"]); err == nil {
		data.StaticNvidiaSMIInfo.Index = uint8(index)
	}

	if name, ok := dataMap["name"]; ok {
		data.StaticNvidiaSMIInfo.Name = name
	}

	if utilizationGPU, err := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(dataMap["utilization.gpu"], " %"))); err == nil {
		data.DynamicNvidiaSMIInfo.Utilization.GPU = uint8(utilizationGPU)
	}

	if utilizationMemory, err := strconv.Atoi(strings.TrimSpace(strings.TrimSuffix(dataMap["utilization.memory"], " %"))); err == nil {
		data.DynamicNvidiaSMIInfo.Utilization.Memory = uint8(utilizationMemory)
	}

	if powerDraw, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(dataMap["power.draw"], " W")), 32); err == nil {
		data.DynamicNvidiaSMIInfo.Power.Draw = float32(powerDraw)
	}

	if temperatureGPU, err := strconv.ParseFloat(dataMap["temperature.gpu"], 32); err == nil {
		data.DynamicNvidiaSMIInfo.Temperature.GPU = float32(temperatureGPU)
	}

	if fanSpeed, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(dataMap["fan.speed"], " %")), 32); err == nil {
		data.DynamicNvidiaSMIInfo.Fan.Speed = uint32(fanSpeed)
	}

	if memoryTotal, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(dataMap["memory.total"], " MiB")), 32); err == nil {
		data.DynamicNvidiaSMIInfo.Memory.Total = float32(memoryTotal)
	}

	if memoryFree, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(dataMap["memory.free"], " MiB")), 32); err == nil {
		data.DynamicNvidiaSMIInfo.Memory.Free = float32(memoryFree)
	}

	if memoryUsed, err := strconv.ParseFloat(strings.TrimSpace(strings.TrimSuffix(dataMap["memory.used"], " MiB")), 32); err == nil {
		data.DynamicNvidiaSMIInfo.Memory.Used = float32(memoryUsed)
	}
}

func HandleNvidiaSMI() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			command := DynamicInfoCommand
			if MemoryStoreData.staticIsEmpty() {
				command = AllInfoCommand
			}
			output, query, err := GetNvidiaSMIOutput(command)
			if err != nil {
				fmt.Printf("Error %s\n", err)
			}
			mu.Lock()
			ParseNvidiaSMIOutput(output, query, &MemoryStoreData.NvidiaSMIInfo)
			mu.Unlock()
			UpdateChan <- struct{}{}
		}
	}
}

func GetData() *NvidiaSMIInfo {
	mu.RLock()
	defer mu.RUnlock()
	return &MemoryStoreData.NvidiaSMIInfo
}
