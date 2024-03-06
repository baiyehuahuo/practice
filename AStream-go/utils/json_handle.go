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
			"events":         []interface{}{},
			"total_duration": 0.0,
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

func getLatestMap(keys []string) map[string]interface{} {
	jsonHandleMutex.Lock()
	defer jsonHandleMutex.Unlock()
	var (
		cur, next map[string]interface{}
		v         interface{}
		ok        bool
		keyLength = len(keys)
	)
	cur = jsonHandle
	for i := 0; i < keyLength-1; i++ {
		v, ok = cur[keys[i]]
		if !ok {
			next = map[string]interface{}{}
			cur[keys[i]] = next
			cur = next
			continue
		} else {
			next, ok = v.(map[string]interface{})
			if !ok {
				Errorf("%s: getLatestMap jsonHandle %s is not a map", consts.UtilError, keys[i])
				return nil
			}
			cur = next
		}
	}
	return cur
}

func SetJsonHandleMultiValue(keys []string, val interface{}) {
	latestMap := getLatestMap(keys)
	if latestMap == nil {
		return
	}
	latestMap[keys[len(keys)-1]] = val
}

func SetJsonHandleMultiValueIntIncrease(keys []string) {
	latestMap := getLatestMap(keys)
	if latestMap == nil {
		return
	}
	latestKey := keys[len(keys)-1]

	x := latestMap[latestKey].(int) // needn't judge may be need zero
	latestMap[latestKey] = x + 1
}

func SetJsonHandleMultiValueFloatAdd(keys []string, addVal float64) {
	latestMap := getLatestMap(keys)
	if latestMap == nil {
		return
	}
	latestKey := keys[len(keys)-1]

	x := latestMap[latestKey].(float64) // needn't judge may be need zero
	latestMap[latestKey] = x + addVal
}

func SetJsonHandleMultiValueSliceAppend(keys []string, appendVal interface{}) {
	latestMap := getLatestMap(keys)
	if latestMap == nil {
		return
	}
	latestKey := keys[len(keys)-1]

	x := latestMap[latestKey].([]interface{})
	latestMap[latestKey] = append(x, appendVal)
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
		Warnf("%s: GetJsonHandleSecondValue jsonHandle %s is not a map", consts.UtilError, key)
		return
	}
	ans = hash[key2]
	return ans
}

func GetJsonHandleMultiValue(keys []string) (ans interface{}) {
	jsonHandleMutex.Lock()
	defer jsonHandleMutex.Unlock()
	var hash = jsonHandle
	var ok bool
	for _, key := range keys[:len(keys)-1] {
		hash, ok = hash[key].(map[string]interface{})
		if !ok {
			Warnf("%s: GetJsonHandleMultiValue jsonHandle %s is not a map", consts.UtilError, key)
			return
		}
	}
	ans = hash[keys[len(keys)-1]]
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
