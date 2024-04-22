package helpers

import (
	"regexp"
	"strings"
)
func CheckValidEmailLength(Email string) (string, bool){
	if len(Email) > 64{
		return "Length of the Email should be less then 64" , false
	}
	if len(Email) > 255{
		return "Length of the Email should be less then 255 characters", false
	}
	return "" , true
}

func ValidateEmail(Email string) (string , bool){
	if !strings.Contains(Email,"@"){
		return "Email should contain '@'", false
	}
	Match , _ := regexp.MatchString(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`, Email);{
		if !Match{
			return "Entered Email Address is not Valid", false
		}
	}
	return "" , true
}