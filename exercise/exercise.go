package exercise

import "fmt"

type UsersExercise struct{
	Exercises map[string]Exercise
}

type Exercise struct {
	ID int
    Iterations []Iteration
}

type Iteration struct {
	Reps []float32
	Weights []float32
	Variances []float32
	ID int
	Sets int
	Weight float32
	Date string
	Note string
	TotalWeight float32
	AverageWeight float32
	AverageRep float32
	AverageWeightRepTotal float32
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
	if ok {
		fmt.Println("Excercise exists")
	}	
	if !ok {
		fmt.Println("Invalid")
		return "";
	}
    switch t {
    case StandardStats:
        return fetchStandardStats(entry)
    case AverageOverall:
        return fetchStandardStats(entry)
    case SimpleStats:
        return fetchStandardStats(entry)
	case MostRecent:
		return fetchStandardStats(entry)
    }
	return fetchStandardStats(entry)
}

func boltOnSeperator() string{
	return "-----------------------------------------------------------\n";
}

func fetchStandardStats(ex Exercise) string{
	returnString := ""
	// for _ , iter := range ex.Iterations {
	// 	returnString += "Date: " +iter.Date + "\n" +
	// 					"Planned Weight: " + iter.Weight + "\n" +
	// 					"Number of Sets: " + iter.Sets + "\n" +
	// 					"Weights per set: " + iter.Weights + "\n" +
	// 					"Reps per set: " + iter.Reps + "\n" +
	// 					"Average Weight: " + iter.
	// 	returnString += boltOnSeperator();
	// }

	return returnString;
}

func NewIteration(reps []float32, weights []float32, variances []float32, ID int, sets int, weight float32, date string, note string, totalWeight float32, averageRep float32, averageWeight float32, averageWeightRepTotal float32) *Iteration{
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