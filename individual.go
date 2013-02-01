package main

import (
    "time"
    "math/rand"
    "math"
    "bytes"
)

const (
    INSTRUCTIONS = "<>+-.,[]"
    MUTATION_RATE = 0.01
)

type Individual struct {
    dna         string
    fitness     int
    goal        string
}

func generateRandomIndividual(dna_length int) Individual {
    ind := Individual{fitness:0}

    rand.Seed(time.Now().UTC().UnixNano()) //TODO Verify randomness somewhat
    for i := 0; i < dna_length; i++ {
        ind.dna += string(INSTRUCTIONS[rand.Intn(8)])
    }

    return ind
}

func (ind *Individual) calculateFitness() {
    // TODO Halting problem yo
    output := Interpret(ind.dna, CELLCOUNT)

    fitness := 0
    for i := 0; i < len(output); i++ {
        fitness += 256 - math.Abs(output[i] - ind.goal[i])
    }

    ind.fitness = fitness
}

func (ind *Individual) mutateDna() {
    // FIXME Horribly slow way of doing this. DNA should probably be a []uint8.
    var buffer bytes.Buffer
    for i := range ind.dna {
        if rand.Float64() < MUTATION_RATE {
            buffer.WriteString(string(INSTRUCTIONS[rand.Intn(8)]))
        } else {
            buffer.WriteString(ind.dna[i])
        }
    }

    ind.dna = buffer.String()
}

func twoPointCrossover(parent1, parent2 Individual) (Individual, Individual) {
    child1 := Individual{fitness: 0, goal: parent1.goal}
    child2 := Individual{fitness: 0, goal: parent1.goal}
    // Choose two random crossoverpoints.
    // TODO Do we not need reseed here?
    pos1 := rand.Intn(len(parent1.dna))
    pos2 := pos1 +  rand.Intn(len(parent1.dna) - pos1)

    child1.dna = parent1.dna[0:pos1] + parent2.dna[pos1:pos2] + parent1.dna[pos2:]
    child2.dna = parent2.dna[0:pos1] + parent1.dna[pos1:pos2] + parent2.dna[pos2:]

    return child1, child2
}

// Methods for sorting.
type Individuals []Individual
func (inds Individuals) Len() int { return len(inds) }
func (inds Individuals) Swap(i, j int) { inds[i], inds[j] = inds[j], inds[i] }
func (inds Individuals) Less(i, j int) bool { return inds[i].fitness < inds[j].fitness }
