package main
import ("logger/app")

func main(){
	err := app.Setup()
	if err != nil {
		panic(err)
	}
}