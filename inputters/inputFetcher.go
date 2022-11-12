package inputters

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func FetchInteger(statement string, intRange int) int{
	valueFound := false;
	var returnValue int;
	for(!valueFound) {
		fmt.Printf("%s",statement);
		userInput  := bufio.NewReader(os.Stdin)
		userVal, err := userInput.ReadString('\n')
		if err != nil {
			fmt.Printf("Enter a value between 0 and %d\n" , intRange);
			continue;
		}

		input := strings.TrimSpace(userVal)
		returnValue, err = strconv.Atoi(input)
		if err != nil {
			fmt.Printf("Enter a value between 0 and %d\n" , intRange);
			continue;
		}

		if(returnValue <= 0 || returnValue > intRange){
			fmt.Printf("Enter a value between 0 and %d\n" , intRange);
			continue;
		}
		fmt.Println(returnValue)
		valueFound = true;
	}
	return returnValue;
}

func FetchString(statement string) string{
	var returnValue string;
	var err error;
	for {
		fmt.Printf("%s",statement);
		userInput  := bufio.NewReader(os.Stdin)
		returnValue, err = userInput.ReadString('\n')
		if err != nil {
			fmt.Printf("Please enter a valid string\n");
			continue;
		}
		break;

	}
	fmt.Printf("Here" + returnValue)
	return returnValue[0:len(returnValue)-1];
}