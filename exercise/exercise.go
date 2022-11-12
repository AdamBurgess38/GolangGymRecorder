package exercise

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
	Sets []int
	Weight float32
	Date string
	Note string
	TotalWeight float32
	AverageRep float32
	AverageWeightRepTotal float32
}