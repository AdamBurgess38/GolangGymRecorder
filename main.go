package main

import (
	"fmt"
	"packages/exercise"
	"packages/sort"
	"packages/startup"
)

type StatsFormat int64

const (
	StandardStats StatsFormat = iota
	AverageOverall        = 1
	SimpleStats        = 2
	MostRecent        = 3
)

func initaliseListOfKeys(){
	var kys []string
	for key, _ := range userInfo.Exercises{
		kys = append(kys, key)
	}
	kys = sort.Sort(kys)
	keys = kys
	initialiseExceriseNamesAndIDs()
}

func printAllExceriseNames(){
	for v, d := range keys{
		fmt.Printf("%d. %s\n", v, d)
	}
}

func initialiseExceriseNamesAndIDs(){
	for v, d := range keys{
		fmt.Printf("%d. %s\n", v, d)
		if entry, ok := userInfo.Exercises[d]; ok {
			entry.ID = v
			userInfo.Exercises[d] = entry
		}
	}
}

// func initialiseExcerise(ue *exercise.UsersExercise,name string){
// 	ue.Exercises[name] = exercise.Exercise{};
// }

// func addIteration(ue *exercise.UsersExercise, name string, x exercise.Iteration){
// 	fmt.Println(x)
// 	entry, ok := ue.Exercises[name];
// 	if ok {
// 		entry.Iterations = append(entry.Iterations,x)
// 	}	
// 	if !ok{
// 		initialiseExcerise(ue, name)
// 		addIteration(ue, name, x);
// 		return;
// 	}
// 	ue.Exercises[name] = entry;
// }

// func newIteration(reps []float32, weights []float32, variances []float32, ID int, sets []int, weight float32, date string, note string, totalWeight float32, averageRep float32, averageWeightRepTotal float32) *exercise.Iteration{
// 	return &(
// 		exercise.Iteration{
// 			Reps: reps, 
// 			Weights: weights, 
// 			Variances: variances,
// 			ID : ID, 
// 			Sets : sets, 
// 			Weight: weight,
// 			Date: date,
// 			Note: note, 
// 			TotalWeight: totalWeight,
// 			AverageRep: averageRep,
// 			AverageWeightRepTotal: averageWeightRepTotal,
// 			});
// }

var userInfo *exercise.UsersExercise;
var currentExercise string;
var keys []string

func main() {
    // fmt.Println("hello world")
    // test := inputters.FetchInteger("hello wolrd!",2)
    // fmt.Println(test)

    userInfo = startup.StartUp()
    // fmt.Println(userInfo)
	startup.SaveUser(userInfo)
	exercise.AddIteration(userInfo, "ZZZZZZZZZZZZZ", *exercise.NewIteration([]float32{1,2,3,4,5}, 
		[]float32{1,2,3,4,5}, []float32{1,2,3,4,5},
		5,
		 5, 
		 4.5,
		 "Hello", 
		 "Hello there!", 
		 6.9, 
		 7.1,
	1.0,
		1.1))
	
	printAllExceriseNames()
	initaliseListOfKeys()
	printAllExceriseNames()
	
	
	fmt.Println(exercise.ViewAnExercise(userInfo,"Barbell Bench" , AverageOverall))
}


/*Functions I need to make:
	Log in --> DONE
	Create new Account --> DONE
        //Move this to a different package!  --> Done
	Add new excerise instance 
		--> Need to get inputs first. 
	Views: --> Put this into a packge tbh where you just pass in what you want. 
		Most recent
		Average overall
		Standard
		Simple
	Delete:
		One instance
		All instances
	View all exercise names --> DONE
	Save --> DONE
	Type in an array of doubles only
	Type in a number only
	Type in a string only

	Mapping to the map....therefore numbers for each of the excerises. 
		Do this with a function and on this new version of my JSON we can actually log the IDs too....? 
*/
