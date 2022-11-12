package main

import (
	"encoding/json"
	"fmt"
	"os"
	"packages/exercise"
	"packages/inputters"
	"strings"
)

func initialiseExcerise(ue *exercise.UsersExercise,name string){
	ue.Exercises[name] = exercise.Exercise{};
}

func addIteration(ue *exercise.UsersExercise, name string, x exercise.Iteration){
	fmt.Println(x)
	entry, ok := ue.Exercises[name];
	if ok {
		entry.Iterations = append(entry.Iterations,x)
	}	
	if !ok{
		initialiseExcerise(ue, name)
		addIteration(ue, name, x);
		return;
	}
	ue.Exercises[name] = entry;
}

func saveUser(ue *exercise.UsersExercise){
	file, _ := json.MarshalIndent(ue, "", " ")
	err := os.WriteFile(username+".json", file, 0644)
    if err == nil{
		fmt.Printf("Succesful save for user %s\n", username);
	}
}

func newUser() *exercise.UsersExercise{
	username = inputters.FetchString("Please enter what you would like your username to be");
	var ue = &exercise.UsersExercise{Exercises: map[string]exercise.Exercise{}};
	return ue;
}

func loadUser() * exercise.UsersExercise{
	
	var ue = &exercise.UsersExercise{Exercises: map[string]exercise.Exercise{}};
	var data []byte;
	for{
		username = inputters.FetchString("Please enter your username:");
		fmt.Println(username+".json")
		dat, err := os.ReadFile(username+".json")
		if err != nil{
			fmt.Printf("Not a valid username, please try again\n")
			continue;
		}
		data = dat;
		break;
	}
	
	json.Unmarshal([]byte(data), &ue)
	backupfile()
	return ue;
}

func newLine(){
	fmt.Printf("================================================================================\n")
}

func backupfile(){
	dat, _ := os.ReadFile(username+".json")
	dateFileFormat := strings.Replace(date, "/", ":", 2);
	err := os.WriteFile(username+"-BACKUP:" + dateFileFormat+ ".json", dat, 0644)
	if err != nil{
		fmt.Printf("Unsuccesful backup of: %s" ,username)
	}
}


func startUp() * exercise.UsersExercise {
	fmt.Printf("Welcome!\n")
	integerInput := inputters.FetchInteger("[1] Log in [2] Create Account\n",2)
	fmt.Println(integerInput)
	if(integerInput ==1){
		return loadUser();
	}
	return newUser();
}

func newIteration(reps []float32, weights []float32, variances []float32, ID int, sets []int, weight float32, date string, note string, totalWeight float32, averageRep float32, averageWeightRepTotal float32) *exercise.Iteration{
	return &(
		exercise.Iteration{
			Reps: reps, 
			Weights: weights, 
			Variances: variances,
			ID : ID, 
			Sets : sets, 
			Weight: weight,
			Date: date,
			Note: note, 
			TotalWeight: totalWeight,
			AverageRep: averageRep,
			AverageWeightRepTotal: averageWeightRepTotal,
			});
}


var username string;
var userInfo *exercise.UsersExercise;
var currentExercise string;
var date string;

func main() {
    fmt.Println("hello world")
    test := inputters.FetchInteger("hello wolrd!",2)
    fmt.Println(test)

    userInfo = startUp()
    fmt.Println(userInfo)
}