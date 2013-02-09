package main

import (
    "fmt"
    "sync"
    "sort"
    "math/rand"
)

const (
    CELLCOUNT = 30000
)

type Experiment struct {
    individuals   []Individual
    generation    int
    fitness_goal  int
}

func CreateExperiment(goal_string string, individual_count, dna_length, fitness_goal int) *Experiment {
    exp := Experiment{generation: 0,
                      fitness_goal: fitness_goal,
                      individuals: make([]Individual, individual_count)}

    for i := range exp.individuals {
        exp.individuals[i] = generateRandomIndividual(dna_length)
        exp.individuals[i].goal = goal_string
    }

    return &exp
}

func (exp *Experiment) Start() {
    fmt.Println("Starting experiment...")
    best_fitness := 0
    for best_fitness < exp.fitness_goal {
        best_fitness = exp.runIteration()
        //if exp.generation % 100 == 0 {
        fmt.Println("Generation finished: ", exp.generation)
        fmt.Println(" -> Best fitness score: ", best_fitness)
        //}
        if exp.generation % 10 == 0 {
            fmt.Println(" -> Output of best performing individual:")
            fmt.Println(Interpret(exp.individuals[0].dna, CELLCOUNT), "\n\n")

            fmt.Println(" -> DNA of best performing individual:")
            fmt.Println(exp.individuals[0].dna, "\n\n")
        }
    }

    // TODO Write beter output, like time elapsed etc.
    fmt.Println("Experiment finished!")
    fmt.Println("Reached target fitness score.")
    fmt.Println(" -> Final BrainFuck DNA:\n")
    fmt.Println(exp.individuals[0].dna, "\n\n")
    fmt.Println(" -> Final output:\n")
    fmt.Println(Interpret(exp.individuals[0].dna, CELLCOUNT), "\n\n")
}

func (exp *Experiment) runIteration() (top_fitness int) {
    exp.generation++

    // Test fitness of all individuals.
    var wg sync.WaitGroup
    for i := range exp.individuals {
        wg.Add(1)
        go func(i int) {
            exp.individuals[i].calculateFitness()
            wg.Done()
        }(i)
    }
    wg.Wait()

    // Sum to normalize fitness-scores.
    // TODO This should probably be done with channels to avoid a few passes.
    fitness_sum := 0
    for _, ind := range exp.individuals {
        fitness_sum += ind.fitness
    }


    // Sort individuals by fitness.
    sort.Sort(Individuals(exp.individuals))

    top_fitness = exp.individuals[0].fitness

    new_gen := make([]Individual, len(exp.individuals))

    // The top ten percent of the previous generation is let through to the next unharmed.
    copy(new_gen, exp.individuals[:(len(exp.individuals)/10)])
    new_gen_ptr := len(exp.individuals)/10

    // FIXME
    // Here be dragons.
    // Extremely ugly code written at 3 am after a case of beer.
    // BADLY needs rewriting, as there is probably no logic to this, read up!
    // Mating of 80% of population, that is len(exp.individuals)*0.4 matings.
    for i := 0; i < int(float64(len(exp.individuals))*0.4); i++ {
        var parent1 Individual
        var par1_index int
        var parent2 Individual
        // TODO Need seed?
        x := rand.Float64()
        sum := 0.0
        for j := range exp.individuals {
            sum += float64(exp.individuals[j].fitness) / float64(fitness_sum)
            if sum >= x {
                parent1 = exp.individuals[j]
                par1_index = j
                break
            }
        }
        x = rand.Float64()
        sum = 0
        for j := range exp.individuals {
            sum += float64(exp.individuals[j].fitness) / float64(fitness_sum)
            if sum >= x && j != par1_index {
                parent2 = exp.individuals[j]
                break
            }
        }

        child1, child2 := twoPointCrossover(parent1, parent2)
        child1.mutateDna()
        child2.mutateDna()
        new_gen[new_gen_ptr] = child1
        new_gen_ptr++
        new_gen[new_gen_ptr] = child2
        new_gen_ptr++
        //new_gen[new_gen_ptr++], new_gen[new_gen_ptr++] = child1, child2
    }

    // Fill the last spots with new blood.
    for new_gen_ptr < len(new_gen) {
        new_gen[new_gen_ptr] = generateRandomIndividual(128) //FIXME Don't inline this
        new_gen_ptr++
    }

    exp.individuals = new_gen
    return
}
