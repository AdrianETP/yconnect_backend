package config

import (
	"log"
	"os"

	"github.com/Davincible/goinsta/v3"
	"github.com/joho/godotenv"
)

var Insta *goinsta.Instagram
func InstaLogin() {

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}
	}
	insta := goinsta.New(os.Getenv("INSTAUSERNAME"),os.Getenv("INSTAPASS"))
    
    err := insta.Login()
    if err != nil{
        panic(err)
    }
    Insta = insta
}
