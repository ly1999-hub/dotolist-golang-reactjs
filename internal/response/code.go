package response

import "github.com/thoas/go-funk"

// Code ...
type Code struct {
	Key     string
	Message string
	Code    int
}

var notFoundKey = Code{
	Key:     CommonNotFound,
	Message: "không tìm thấy",
	Code:    -1,
}
var list []Code

// GetByKey give key and receive message + code
func GetByKey(key string) Code {
	item := funk.Find(list, func(item Code) bool {
		return item.Key == key
	})

	if item == nil {
		return notFoundKey
	}
	return item.(Code)
}

func init() {
	list = append(list, common...)
}
