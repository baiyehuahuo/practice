package utils

import (
	"AStream-go/consts"
	"encoding/json"
	"os"
	"sync"
)

var (
	jsonHandle      map[string]interface{}
	jsonHandleMutex = sync.RWMutex{}
)

func init() {
	jsonHandle = make(map[string]interface{})
	jsonHandle["playback_info"] = map[string]interface{}{
		"start_time":                 nil,
		"end_time":                   nil,
		"initial_buffering_duration": nil,
		"interruptions": map[string]interface{}{
			"count":          0,
			"events":         new([][]float64),
			"total_duration": 0,
		},
		"up_shifts":   0,
		"down_shifts": 0,
	}
}

func SetJsonHandleValue(key string, val interface{}) {
	jsonHandleMutex.Lock()
	defer jsonHandleMutex.Unlock()
	jsonHandle[key] = val
}

func SetJsonHandleSecondValue(key, key2 string, val interface{}) {
	jsonHandleMutex.Lock()
	defer jsonHandleMutex.Unlock()
	second, ok := jsonHandle[key]
	if !ok {
		jsonHandle[key] = map[string]interface{}{key2: val}
		return
	}
	hash, ok := second.(map[string]interface{})
	if !ok {
		Warnf("%s: SetJsonHandleSecondValue jsonHandle %s is not a map", consts.UtilError, key)
		return
	}
	hash[key2] = val
}

func GetJsonHandleValue(key string) (ans interface{}) {
	jsonHandleMutex.RLock()
	defer jsonHandleMutex.RUnlock()
	ans = jsonHandle[key]
	return ans
}

func GetJsonHandleSecondValue(key, key2 string) (ans interface{}) {
	jsonHandleMutex.Lock()
	defer jsonHandleMutex.Unlock()
	second, ok := jsonHandle[key]
	if !ok {
		Warnf("%s: GetJsonHandleSecondValue jsonHandle %s is not exist", consts.UtilError, key)
		return nil
	}
	hash, ok := second.(map[string]interface{})
	if !ok {
		Warnf("%s: SetJsonHandleSecondValue jsonHandle %s is not a map", consts.UtilError, key)
		return
	}
	ans = hash[key2]
	return ans
}

func SaveJsonHandle(savePath string) {
	jsonHandleMutex.Lock()
	defer jsonHandleMutex.Unlock()
	file, err := os.OpenFile(savePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		Warnf("%s: SaveJsonHandle open file failed: %s", consts.UtilError, err.Error())
		return
	}
	// 创建编码器
	encoder := json.NewEncoder(file)
	err = encoder.Encode(jsonHandle)
	if err != nil {
		Warnf("%s: SaveJsonHandle encode data failed: %s", consts.UtilError, err.Error())
	}
}
