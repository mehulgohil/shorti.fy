package main

func main() {
	app := Router().InitRouter()
	err := app.Listen(":8080")
	if err != nil {
		panic("unable to start server")
	}
}
