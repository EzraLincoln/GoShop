package main

func get(args ...interface{}) (*interface{}) {
	value, err := args[0], args[1]

	if (err != nil) {
		return &err
	} else {
		return &value
	}
}
