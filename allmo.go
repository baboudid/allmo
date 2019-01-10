// Show all the packages in your project (including sub-packages' sub-packages)
package allmo

import (
	"encoding/json"
	"fmt"
	"os/exec"
	"time"
)

//show all modules.
func Mo() []string {
	result := []string{}
	RestMap = make(map[string]int)
	tt(varS)
	for { // 此处一直阻塞，直到 start和end在一秒的时间区域内都是相等的时候然后退出。
		// Blocked here until exit and end are equal in the time zone of one second and then exit.
		start := len(RestMap)
		time.Sleep(time.Second)
		end := len(RestMap)
		if start == end {
			fmt.Println(end)
			break
		}
	}
	for k, _ := range RestMap {
		result = append(result, k)
	}
	return result

}
func tt(s string) {
	re := new(Result)
	cmd := exec.Command("/usr/local/bin/go", "list", "-e", "-json", s)
	data, err := cmd.Output()
	if err != nil {
		fmt.Print(err)
	}
	json.Unmarshal(data, re)
	for _, v := range re.Imports {
		sy.Lock()
		RestMap[v]++
		sy.Unlock() // 如果在lock前调用unlock那么会发生error错误If you call unlock before lock, an error will occur.
		go tt(v)
	}

}
