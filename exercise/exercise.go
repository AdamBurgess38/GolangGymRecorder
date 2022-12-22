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
	return "Date: " + iter.Date + "\n" +
	"Planned Weight: " + strconv.FormatFloat(iter.Weight, 'f', 2, 64) + "\n" +
	"Number of Sets: " + strconv.Itoa(iter.Sets) + "\n" +
	"Weights per set: " + arrayToString(iter.Weights) + "\n" +
	"Reps per set: " + arrayToString(iter.Reps) + "\n" +
	"Average Weight: " + strconv.FormatFloat(iter.AverageWeight, 'f', 2, 64)+ "\n" +
	"Average reps: " + strconv.FormatFloat(iter.AverageRep, 'f', 2, 64) + "\n" +
	"Average total: " + strconv.FormatFloat(iter.AverageWeightRepTotal, 'f', 2, 64)+ "\n" +
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

func AddIteration(ue *UsersExercise, name string, x Iteration){
	fmt.Println(x)
	entry, ok := ue.Exercises[name];
	if ok {
		entry.Iterations = append(entry.Iterations,x)
	}	
	if !ok{
		initialiseExcerise(ue, name)
		AddIteration(ue, name, x);
		return;
	}
	ue.Exercises[name] = entry;
}