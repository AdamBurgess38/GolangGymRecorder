package main

import (
	"fmt"
	"packages/exercise"
	"packages/inputters"
	"packages/sort"
	"packages/startup"
	"strings"
	"time"
)

// type StatsFormat int64

// const (
// 	StandardStats StatsFormat = iota
// 	AverageOverall        = 1
// 	SimpleStats        = 2
// 	MostRecent        = 3
// )

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

func userInputNewExercise(){
	currentExercise = inputters.FetchString("State the name of the exercise")
	fmt.Println("Adding iteration of " , currentExercise);
	weight := inputters.FetchDouble("Please state the weight you worked at");
	constantWeight := inputters.FetchBoolean("Was the weight constant throughout?");
	constantReps := inputters.FetchBoolean("Were the reps constant throughout?");
	var reps []float64
	var weights []float64
	var sets int
	if(constantWeight && constantReps){
		rep := inputters.FetchDouble("And what was this value?");
		sets := inputters.FetchInteger("For how many sets?",1000);
		for sets != 0{
			reps = append(reps, rep);
			weights = append(weights, weight)
			sets --;
		}
		
	}
	if(!constantWeight && constantReps){
		rep := inputters.FetchDouble("And what was this value?");
		weights = inputters.FetchArray("Please state the weight throughout the sets")
		for x := 0; x < len(weights); x++{
			reps = append(reps, rep)
		}
	}
	if(constantWeight && !constantReps){
		reps = inputters.FetchArray("Please state the reps throughout the sets")
		for x := 0; x < len(weights); x++{
			weights = append(weights, weight)
		}
	}
	if(!constantWeight && !constantReps){
		reps = inputters.FetchArray("Please state the reps throughout the sets")
		weights = inputters.FetchArray("Please state the weight throughout the sets")
	}

	requestNote := inputters.FetchBoolean("Would you like to leave a note?");
	var note string = ""
	if(requestNote){
		note += inputters.FetchString("Please enter what you would like the note to be?")
	}

	var daysAgo = 0
	altTime := inputters.FetchBoolean("Did you perform this exercise today?");
	if(!altTime){
		daysAgo = inputters.FetchInteger("How many days ago did you perform this exercise?", 365);
	}
	if(exercise.UserRequestNewIteration(userInfo, currentExercise, *exercise.UserTempIteration(reps, weights, sets, weight, 
		strings.Replace((time.Now().Local().AddDate(0, 0, -daysAgo)).Format("01-02-2006"),"/", ":", 2),note))){
			fmt.Println("Instance of " , currentExercise , " has been successfully added")
			return;
		}

	fmt.Println("Instance of " , currentExercise , " has been unsuccessfully added due to the number of reps and weights not alligning")
	
}

func main() {

	// currentTime := strings.Replace(time.Now().Local().Add(time.Hour*24).Format("01-02-2006"),"/", ":", 2);
    // fmt.Println("Current Time in String: ", currentTime)
    
	// fmt.Println(currentTime)
    // fmt.Println("hello world")
    // test := inputters.FetchInteger("hello wolrd!",2)
    // fmt.Println(test)

    userInfo = startup.StartUp()
	
	userInputNewExercise()
    // // fmt.Println(userInfo)
	// startup.SaveUser(userInfo)
	// exercise.AddIteration(userInfo, "ZZZZZZZZZZZZZ", *exercise.NewIteration([]float64{1,2,3,4,5}, 
	// 	[]float64{1,2,3,4,5}, []float64{1,2,3,4,5},
	// 	5,
	// 	 5, 
	// 	 4.5,
	// 	 "Hello", 
	// 	 "Hello there!", 
	// 	 6.9, 
	// 	 7.1,
	// 1.0,
	// 	1.1))
	
	// printAllExceriseNames()
	// initaliseListOfKeys()
	// printAllExceriseNames()
	// fmt.Println(inputters.FetchArray("dsfghibdshjfbdfshjfgbdhjsbfdhsj"))
	
	// // fmt.Println(exercise.ViewAnExercise(userInfo,"Barbell Bench" , exercise.AverageOverall))
	// fmt.Println(inputters.FetchBoolean("Yes or no test"))	
	
}


/*Functions I need to make:
	Log in --> DONE
	Create new Account --> DONE
        //Move this to a different package!  --> Done
	Add new excerise instance 
		--> Need to get inputs first. --> Done
	Views: --> Put this into a packge tbh where you just pass in what you want. --> DONE
		Most recent
		Average overall
		Standard
		Simple
	Delete:
		One instance 
		All instances --> Done
	View all exercise names --> DONE
	Save --> DONE
	Type in an array of doubles only --> DOne
	Type in a number only --> Done
	Type in a string only --> Done
		Typing, have it all use the same function? 

	Mapping to the map....therefore numbers for each of the excerises. 
		Do this with a function and on this new version of my JSON we can actually log the IDs too....? 
*/
