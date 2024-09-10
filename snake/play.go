package main

import (
    "fmt"
    "time"
    "math/rand"
)



    type Pos struct {
        x int;
        y int;
    }

    type Direction int

   const(
        Up Direction = iota
        Down
        Right
        Left
    )

    type Snake struct {
        headPosition Pos
        body []Pos
        direction Direction 
    }

    var screenHeight int = 9
    var screenWidth int = 9

func main() {

    screen := InitScreen10D()

    apple := false
    applePosition := Pos{
        x: 0,
        y: 0,
    }

    snake := Snake{
        headPosition: Pos{
            x: 4,
            y: 4,
        },
        body: []Pos{},
        direction: Right,
    }

    screen[snake.headPosition.y][snake.headPosition.x] = "@ "

    //screen[pos.x][pos.y]
    charr := make(chan []byte)
    go ReadKey(charr) 

    for true {

        select {
        case key :=  <-charr:
            snake.direction = turn(key, snake.direction)
        
        default: 
            if apple == false {
                applePosition = placeApple(snake.headPosition, screenHeight, screenWidth)
                screen[applePosition.y][applePosition.x] = "& "
                apple = true
            }

            fmt.Println("\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n")
            Display10B(screen)

            screen, snake = move(screen, snake)
            if snake.headPosition == applePosition && apple == true {
                apple = false;
                 
                snake.body = append([]Pos{snake.headPosition}, snake.body...)
            }

            time.Sleep(160*time.Millisecond)
        }
    }   
}

func move(screen [10][10]string, snake Snake) ([10][10]string, Snake) {
    screen[snake.headPosition.y][snake.headPosition.x] = "  "
    for i := 0; i < len(snake.body); i++ {
        screen[snake.body[i].y][snake.body[i].x] = "  "
    }

    for i := len(snake.body); i > 1; i-- {
        screen[snake.body[i-1].y][snake.body[i-1].x] = "  "
        snake.body[i-1] = snake.body[i-2]
    }
    if len(snake.body) > 0 {
        screen[snake.body[0].y][snake.body[0].x] = "  "
        snake.body[0] = snake.headPosition    
    }

    if  snake.direction == Up {
        snake.headPosition.y --
        if snake.headPosition.y < 0 {
            snake.headPosition.y = screenHeight
        }
    } else if  snake.direction == Down {
        snake.headPosition.y ++
        snake.headPosition.y = snake.headPosition.y%(screenHeight + 1)
    } else if  snake.direction == Left {
        snake.headPosition.x --
        if snake.headPosition.x < 0 {
            snake.headPosition.x = screenWidth
        }
    } else if  snake.direction == Right {
        snake.headPosition.x ++
        snake.headPosition.x = snake.headPosition.x % (screenHeight + 1)
    }
    
//    for i := len(snake.body); i < 1; i-- {
//        snake.body[i-1] = snake.body[i-2]
//    }
//    if len(snake.body) > 0 {
//        snake.body[0] = snake.headPosition    
//    }

    screen[snake.headPosition.y][snake.headPosition.x] = "@ "
    for i := 0; i < len(snake.body); i++ {
        screen[snake.body[i].y][snake.body[i].x] = "@ "
    }
    return screen, snake
}

func turn(key []byte, direction Direction) Direction {
    if string(key) == "w" && direction != Down {
        direction = Up
    } else if string(key) == "s" && direction != Up {
        direction = Down
    } else if string(key) == "a" && direction != Right {
        direction = Left
    } else if string(key) == "d" && direction != Left {
        direction = Right
    } else {
        return direction
    }

   return direction
}

func placeApple(snakePos Pos, screenHeight int, screenWidth int) Pos {//[50][50]string {
    valid := false 
    applePosition := Pos {
        x: 0,
        y: 0,
    }

    for !valid {
        applePosition.y = rand.Intn(screenHeight)
        applePosition.x = rand.Intn(screenWidth)

        if applePosition.y != snakePos.y || applePosition.x != snakePos.x {
            valid = true
        }
    }

    return applePosition
}

