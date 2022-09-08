package module

// InitMongoDB ...
func InitMongoDB() {
	err := ConnectDB("mongodb://127.0.0.1:27017", "", "", "test", "", "")

	if err != nil {
		panic(err)
	}

}
