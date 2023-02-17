// package main

// func main()  {
	
// }

// // // import (
// // // 	// "bufio"
// // // 	// "bytes"
// // // 	"bufio"
// // // 	"bytes"
// // // 	"encoding/json"
// // // 	"fmt"

// // // 	// "fmt"
// // // 	"io"
// // // 	// "io/ioutil"
// // // 	// "os"
// // // 	// "strconv"
// // // 	// "strings"
// // // )


// // type Manager struct {
// //     FullName       string `json:"full_name,omitempty"`
// //     Position       string `json:"position,omitempty"`
// //     Age            int32 `json :"age,omitempty"`
// //     YearsInCompany int32 `json:"years_in_company,omitempty"`
// // }

// // func EncodeManager(manager *Manager) (io.Reader, error) {
// // 	jsonBytes, err := json.Marshal(manager)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	bufio := bytes.NewBuffer(jsonBytes)
// // 	return bufio, nil
// // }