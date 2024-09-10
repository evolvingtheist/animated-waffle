package main

import (
    "fmt"
    "time"
)

var screenHeight int = 49;
var screenWidth int = 49;

type Direction int

const (
    Up Direction = iota
    Down
    Left
    Right
)
 
type Position struct {
    x int
    y int
}

type Particle struct {
    position Position
    direction Direction
    velocity int
}

func main() {
    screen := InitScreen50D()
    particles := []Particle {
        Particle{ position: Position{ x: 24, y: 24 }, direction: Right, velocity: 2 },
        Particle{ position: Position{ x: 13, y: 24 }, direction: Right, velocity: 2 },
    }

    for true {
        screen = updateScreen(screen, particles)

        for i := 0; i < len(particles); i++ {
            particles[i] = move(particles[i])
        }


        fmt.Println("\n\n\n\n\n");
        Display50DB(screen)
        time.Sleep(160*time.Millisecond)
    }
}

func move(particle Particle, particleID int, particles []Particle) Particle {
    switch particle.direction{
        case Up:
            if particle.position.y + 1 != screenHeight {
                particle.position.y++
            } else {
                particle.position.y++
                particle.direction = Down     
            }
        case Down:
            if particle.position.y - 1 != 0 {
                particle.position.y--
            } else {
                particle.position.y--
                particle.direction = Up    
            }
        case Left:
            if particle.position.x - 1 != 0 {
                particle.position.x--
            } else {
                particle.position.x--
                particle.direction = Right     
            }
        case Right: 
            if particle.position.x + 1 != screenWidth {
                particle.position.x++
            } else {
                particle.position.x++
                particle.direction = Left
            }
    }

    return particle
}

func colide(particle Praticle, particleID int,  particles []Particle) {
    
}

func updateScreen(screen [50][50]string, particles []Particle) [50][50]string {
    screen = InitScreen50D() 

    for i, v := range particles {
        screen[v.position.y][v.position.x] = "@ "
        _ = i
    }

    return screen
}
