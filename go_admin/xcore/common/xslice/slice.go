package xslice

import (
	"fmt"
	"strconv"
)

func StringToInt32(strings []string) []int32 {
	int32Slice := make([]int32, len(strings))
	for i, str := range strings {
		num, err := strconv.ParseInt(str, 10, 32)
		if err != nil {
			fmt.Printf("转换失败，无效的字符串: %s\n", str)
			return nil
		}
		int32Slice[i] = int32(num)
	}

	return int32Slice
}
func StringExist(slice []string, target string) bool {
	found := false
	for _, s := range slice {
		if s == target {
			found = true
			break
		}
	}
	return found
}

func StringDuplicate(slc []string) []string {
	var result []string
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
