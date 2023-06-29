package main

func main() {
	listDrivers()
	if db, err := openDatabase(); err == nil {
		db.Close()
	} else {
		panic(err)
	}

	openPGDatabase()
}
