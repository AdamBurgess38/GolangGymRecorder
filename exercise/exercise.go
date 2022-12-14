package exercise

import (
	"fmt"
	"strconv"
)

type UsersExercise struct{
	Exercises map[string]Exercise
}

type Exercise struct {
	ID int
    Iterations []Iteration
}

type userInput struct {
	Reps []float64
	Weights []float64
	Sets int
	Weight float64
	Date string
	Note string
}

type Iteration struct {
	Reps []float64
	Weights []float64
	Variances []float64
	ID int
	Sets int
	Weight float64
	Date string
	Note string
	TotalWeight float64
	AverageWeight float64
	AverageRep float64
	AverageWeightRepTotal float64
}

type StatsFormat int64

const (
	StandardStats StatsFormat = iota
	AverageOverall        = 1
	SimpleStats        = 2
	MostRecent        = 3
)

func ViewAnExercise(ue *UsersExercise, requestedExercise string, t StatsFormat) string{
	entry, ok := ue.Exercises[requestedExercise];	
	if !ok {
		fmt.Println("Invalid excerise")
		return "";
	}
    switch t {
    case StandardStats:
        return fetchStandardStats(entry)
    case AverageOverall:
        return fetchAverageOverall(entry)
    case SimpleStats:
        return fetchSimpleStats(entry)
	case MostRecent:
		return fetchMostRecent(entry)
    }
	return "Invalid stats type"
}

func boltOnSeperator() string{
	return "-----------------------------------------------------------\n";
}

func fetchStandardStats(ex Exercise) string{
	returnString := ""
	for _ , iter := range ex.Iterations {
		returnString += instanceConverter(iter)
		returnString += boltOnSeperator();
	}
	return returnString;
}

func instanceConverter(iter Iteration) string {
	return "ID: " + strconv.Itoa(iter.ID) + "\n" + 
	"Date: " + iter.Date + "\n" +
	"Planned Weight: " + strconv.FormatFloat(iter.Weight, 'f', 2, 64) + "\n" +
	"Number of Sets: " + strconv.Itoa(iter.Sets) + "\n" +
	"Weights per set: " + arrayToString(iter.Weights) + "\n" +
	"Reps per set: " + arrayToString(iter.Reps) + "\n" +
	"Average Weight: " + strconv.FormatFloat(iter.AverageWeight, 'f', 2, 64)+ "\n" +
	"Average reps: " + strconv.FormatFloat(iter.AverageRep, 'f', 2, 64) + "\n" +
	"Average total: " + strconv.FormatFloat(iter.AverageWeightRepTotal, 'f', 2, 64)+ "\n" +
	"Total Weight pushed: " + strconv.FormatFloat(iter.TotalWeight, 'f',2,64) + "\n"+
	"Note: " + iter.Note +"\n";
}

func fetchAverageOverall(ex Exercise) string{
	returnString := "";
	for _ , iter := range ex.Iterations {
		returnString += "Date: " + iter.Date + " Average Weight: " + strconv.FormatFloat(iter.AverageWeight, 'f', 2, 64) +" Average Rep: " + strconv.FormatFloat(iter.AverageRep, 'f', 2, 64) +"\n"
	}
	return returnString;
}

func fetchSimpleStats(ex Exercise) string{
	returnString := "";
	for _ , iter := range ex.Iterations {
		returnString += "Date: " + iter.Date + " Weights: " + arrayToString(iter.Weights) +" Reps: " + arrayToString(iter.Reps) +"\n"
	}
	return returnString;
}

func fetchMostRecent(ex Exercise) string{
	returnString := ""
	latestIndex := len(ex.Iterations);
	iter := ex.Iterations[latestIndex-1]
	
	returnString += instanceConverter(iter)
	returnString += boltOnSeperator();
	
	return returnString;
}

func arrayToString(array []float64) string{
	returnString := ""
	for _ , x := range array{
		returnString += strconv.FormatFloat(x, 'f', 2, 64) +","
	}
	return returnString[0:len(returnString)-1];

}

func UserTempIteration(reps, weights []float64, sets int, weight float64, date string, note string) *userInput{
	return &(
		userInput{
			Reps: reps, 
			Weights: weights, 
			Sets : sets, 
			Weight: weight,
			Date: date,
			Note: note,
			});
}

func NewIteration(reps []float64, weights []float64, variances []float64, ID int, sets int, weight float64, date string, note string, totalWeight float64, averageRep float64, averageWeight float64, averageWeightRepTotal float64) *Iteration{
	return &(
		Iteration{
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
			AverageWeight: averageWeight,
			AverageWeightRepTotal: averageWeightRepTotal,
			});
}

func initialiseExcerise(ue *UsersExercise,name string){
	ue.Exercises[name] = Exercise{};
}

func DeleteEntireExercise(ue *UsersExercise, name string){
	delete(ue.Exercises, name)
}


func addIteration(ue *UsersExercise, name string, x Iteration){
	fmt.Println(x)
	entry, ok := ue.Exercises[name];
	if ok {
		entry.Iterations = append(entry.Iterations,x)
	}	

	ue.Exercises[name] = entry;
}

func Map[T, V any](ts []T, fn func(T) V) []V {
    result := make([]V, len(ts))
    for i, t := range ts {
        result[i] = fn(t)
    }
    return result
}

func UserRequestNewIteration(ue *UsersExercise, name string, x userInput) bool{
	if(len(x.Reps) != len(x.Weights)){
		return false;
	}
	entry, ok := ue.Exercises[name];

	if !ok {
		initialiseExcerise(ue, name)
	}	

	var newID int = len(entry.Iterations);
	var foundID bool = false

	for !foundID {
		foundID = true;
		for _, x := range entry.Iterations{
			if(x.ID == newID){
				newID++
				foundID = false;
				break;
			}
		}
	}
	var totalWeight float64 = 0
	var totalWeightRep float64
	var totalReps float64 = 0
	for i , w := range x.Weights{
		totalWeightRep += w * x.Reps[i]
		totalReps += x.Reps[i]
		totalWeight += w
	}


	addIteration(ue, name, *NewIteration(
		x.Reps, x.Weights, Map(x.Weights, func(item float64) float64 { return item - x.Weight }), newID, 
		x.Sets, 
		x.Weight, 
		x.Date, 
		x.Note, totalWeightRep, totalReps/float64(len(x.Reps)), totalWeight/float64(len(x.Weights)), totalWeightRep/float64(len(x.Weights))))

		fmt.Println(ViewAnExercise(ue,"Barbell Bench" , MostRecent))

	return true;
}