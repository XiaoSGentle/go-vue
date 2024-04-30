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
