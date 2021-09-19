package diskv_wrapper

import (
	"testing"

	"github.com/peterbourgon/diskv/v3"
)

var (
	disk *diskv.Diskv = Init()
)

func VersionUseTest(t *testing.T) {
	existingVersionUse := VersionUse(disk)

	if len(existingVersionUse) == 0 {
		t.Errorf("Expected has version use")
	}
}

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"

// 	diskv_utils "github.com/mkhuda/diskv_wrapper"
// )

// func wrapper_test() {
// 	var version, key, value string

// 	fmt.Println("::Using Diskv")

// 	currentDisk := diskvInit()

// 	fmt.Println("::::Existing")
// 	existingVersionUse := diskv_utils.VersionUse(currentDisk)
// 	diskv_utils.ReadByVersions(existingVersionUse, currentDisk)
// 	diskv_utils.ReadAllKeys(currentDisk)

// 	fmt.Println("::::version")
// 	fmt.Print("Enter version (v1, v2, v...): ")

// 	// Input version name {v1, v2, v...}
// 	_version, err := fmt.Scanln(&version)
// 	if err != nil {
// 		diskv_utils.JustError(_version, err)
// 		return
// 	}

// 	diskv_utils.WriteVersion(currentDisk, version)

// 	fmt.Println("::::key")
// 	fmt.Print("Enter key: ")

// 	// Input key name
// 	_key, err := fmt.Scanln(&key)
// 	if err != nil {
// 		diskv_utils.JustError(_key, err)
// 		return
// 	}

// 	fmt.Println("::::value")
// 	fmt.Print("Enter the value: ")

// 	// Input value
// 	valueReader := bufio.NewReader(os.Stdin)
// 	valueBytes, _, err := valueReader.ReadLine()
// 	value = string(valueBytes)
// 	if err != nil {
// 		diskv_utils.JustError(valueBytes, err)
// 		return
// 	}

// 	path := strings.Join([]string{version, key}, "/")

// 	currentDisk.WriteString(path, value)

// }
