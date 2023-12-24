package controllers

import (
	"fmt"
	"net/http"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/oOWinOo/Chawin_agnos_backend/database"
	"github.com/oOWinOo/Chawin_agnos_backend/models"
)

func SubmitPassword(c *gin.Context){
	password := new(models.LogInOut)
	if err := c.ShouldBind(password) ;err != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : err.Error(),
		})
	}
	password.Steps = findSteps(password.Password)
	result := database.Db.Create(password)
	if result.Error != nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error" : result.Error.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"Password Receive": password,
	})
}

func findSteps(password string) uint{
	length := len(password)
	if length <= 3 { // if "...." = 3
		return uint(6 - len(password))
	}else if length == 4{
		if password == "...." || password == "!!!!" {
			return 3
		}
		return 2
	}else if length ==5{
		inTheRow := stepToEditRepeating(password)
		editSteps := 0
		for i := range inTheRow{
			editSteps += stepOfOrderNum(i)*inTheRow[i]
		}
		return uint(max(1,stepsToAdd(password),editSteps))
	}else if length <= 20 {
		inTheRow := stepToEditRepeating(password)
		editSteps := 0
		for i := range inTheRow{
			editSteps += stepOfOrderNum(i)*inTheRow[i]
		}
		return uint(max(stepsToAdd(password),editSteps))
	}else{
		inTheRow := stepToEditRepeating(password)
		editSteps := removeUntilTwenty(inTheRow,length)
		for i := range inTheRow{
			editSteps += stepOfOrderNum(i)*inTheRow[i]
		}
		return uint(max(stepsToAdd(password),editSteps))
		
	}


}

func stepsToAdd(word string) int{ //  Contains at least 1 lowercase letter, at least 1 uppercase letter, and at least 1 digit
	result := 3
	check1,check2,check3 := false,false,false;
	for _,char := range word{
		if unicode.IsLower(char) {
			result -= 1
			check1 = true
		} else if unicode.IsUpper(char) {
			result -= 1
			check2 = true
		} else if unicode.IsDigit(char) {
			result -= 1
			check3 = true
		}

		if check1 && check2 && check3 {
			break
		}
	}
	
	return result

}



func stepToEditRepeating(word string)map[int]int { // Does not contain 3 repeating characters in a row
	result := make(map[int]int)
	count := 1;
	lastChar := rune(0);
	for _,char := range word{
		if lastChar == char{
			count += 1
		}else{
			result[count] += 1;
			count = 1
			lastChar = char
		}
	}
	if count > 1 {
		result[count] += 1;
	}
	fmt.Println(result)
	return result


}

func stepOfOrderNum(num int) int{ // Cost to edit when repeat num = num
	if num < 3{
		return 0
	}else if num <= 5{
		return 1
	}
	y := num - 5;
	result := int((y-1+3)/3) //ceiling divide 3
	return result+1
}

func removeUntilTwenty(inTheRow map[int]int, length int) int{
	newLen := length
	editSteps := 0
	for i := 3 ; i < length ; i += 3{// remove repeat 3,6,9,.. in a row
		for inTheRow[i] > 0 {
			inTheRow[i] -= 1
			inTheRow[i-1] += 1
			editSteps += 1
			newLen -= 1
			if newLen <= 20 {
				break
			}
		}
		if newLen <= 20{
			break
		}
	}
	for i := 4 ; i < length ; i += 3{// remove repeat 4,7,10,.. in a row
		for inTheRow[i] > 0 {
			inTheRow[i] -= 1
			inTheRow[i-1] += 1
			editSteps += 1
			newLen -= 1
			if newLen <= 20 {
				break
			}
			inTheRow[i-1] -= 1// remove repeat 3,6,9,.. in a row
			inTheRow[i-2] += 1
			editSteps += 1
			newLen -= 1
			if newLen <= 20 {
				break
			}
		}
		if newLen <= 20{
			break
		}
	}
	
	var start int = (length-5)/3
	start *= 3 + 5 // Begin At most number of set 5,8,11,..
	for i := start ; i > 5 ; i -= 3{// remove repeat 5,8,11,.. in a row
		for inTheRow[i] > 0 {
			inTheRow[i] -= 1
			inTheRow[i-1] += 1
			editSteps += 1
			newLen -= 1
			if newLen <= 20 {
				break
			}
			inTheRow[i-1] -= 1// remove repeat 4,7,10,.. in a row
			inTheRow[i-2] += 1
			editSteps += 1
			newLen -= 1
			if newLen <= 20 {
				break
			}
			inTheRow[i-2] -= 1// remove repeat 3,6,9,.. in a row
			inTheRow[i-3] += 1
			editSteps += 1
			newLen -= 1
			if newLen <= 20 {
				break
			}
		}
		if newLen <= 20{
			break
		}
	}
	// If remove until dont have any char repeat 3 or more in a row, Just remove anything until 20 
	//(remove most repeat Char in a row (2),remove most repeat Char, remove dot and !, remove most case) that will dont make password weak again
	if newLen > 20{ 
		editSteps += newLen-20
		newLen = 20;
	}
	return editSteps

	
}