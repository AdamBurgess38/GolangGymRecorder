package main

import (
	"fmt"
	"packages/exercise"
	"packages/inputters"
	"packages/startup"
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

    userInfo = startup.StartUp()
    fmt.Println(userInfo)
	startup.SaveUser(userInfo)
}


/*Functions I need to make:
	Log in
	Create new Account
        //Move this to a different package! 
	Add new excerise instance
	Views:
		Most recent
		Average overall
		Standard
		Simple
	Delete:
		One instance
		All instances
	View all exercise names
	Save

	Type in an array of doubles only
	Type in a number only
	Type in a string only

	Mapping to the map....therefore numbers for each of the excerises. 
		Do this with a function and on this new version of my JSON we can actually log the IDs too....? 
*/
