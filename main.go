package main

import (
    "fmt"
    "os"
)

const (
    Up string = "UP"
    Down string = "DOWN"
    Left string = "LEFT"
    Right string = "RIGHT"
    maxX int = 30
    maxY int = 20
)

var board [][]int

type lumicycle struct {
    lastDir string
    pos coord
}

type coord struct {
    x int
    y int
}

func move(dir string) string {
    var lastDir string
    switch dir {
    case Up:
        fmt.Println(Up)
        lastDir = Up
    case Down:
        fmt.Println(Down)
        lastDir = Down
    case Left:
        fmt.Println(Left)
        lastDir = Left
    case Right:
        fmt.Println(Right)
        lastDir = Right
    }
    return lastDir
}

func computeNextPos(pos coord, dir string) coord {
    var newPos coord
    fmt.Fprintf(os.Stderr, "\t\tCompute next position from (%d,%d) with direction %s\n", pos.x, pos.y, dir)
    switch dir {
    case Up:
        newPos = coord{pos.x, pos.y - 1}
    case Down:
        newPos = coord{pos.x, pos.y + 1}
    case Left:
        newPos = coord{pos.x - 1, pos.y}
    case Right:
        newPos = coord{pos.x + 1, pos.y}
    }
    
    fmt.Fprintf(os.Stderr, "\t\tNext position = (%d,%d)\n", newPos.x, newPos.y)
    return newPos
}

func isPosFree(pos coord) bool {
    isFree := board[pos.x][pos.y] == -1
    fmt.Fprintf(os.Stderr, "\t\tOccupancy of pos (%d,%d) = %d ==> free = %t\n", pos.x, pos.y, board[pos.x][pos.y], isFree)
    return isFree
}

func isPosInsideBoard(pos coord) bool {
    fmt.Fprintf(os.Stderr, "\t\tIs pos (%d,%d) in the board ?\n", pos.x, pos.y)
    if pos.x < 0 || pos.x > 29 || pos.y < 0 || pos.y > 19 {
        fmt.Fprintf(os.Stderr, "\t\t- Outside !!\n")
        return false
    }
        fmt.Fprintf(os.Stderr, "\t\t- Inside\n")
    return true
}

func turn(currentPos coord, lastDir string) string {
    fmt.Fprintf(os.Stderr, "Turn from %s", lastDir)
    newDir := lastDir
    switch lastDir {
    case Up:
        newDir = Right
    case Down:
        newDir = Left
    case Left:
        newDir = Up
    case Right:
        newDir = Down
    }
    fmt.Fprintf(os.Stderr, " to %s\n", newDir)
    return newDir
}

func shouldTurn(nextPos coord) bool {
    if isPosInsideBoard(nextPos) {
        if isPosFree(nextPos) {
            //  if next pos is in the board and free, do not turn
            return false
        }
        //  if next pos is in the board but is not free free, turn
        return true
    }
    //  if next pos is outside the board, turn
    return true
}

func chooseDirection(currentPos coord, lastDir string) string {
    fmt.Fprintf(os.Stderr, "Choose Direction\n")
    fmt.Fprintf(os.Stderr, "\t- current direction: %s\n", lastDir)
    fmt.Fprintf(os.Stderr, "\t- current position: (%d,%d)\n", currentPos.x, currentPos.y)

    // if no reason to turn, keep same direction
    newDir := lastDir

    // compute next position if we do not change direction
    nextPos := computeNextPos(currentPos, newDir)

    // while next position is out of board or next position is already used, turn
    for ;shouldTurn(nextPos); {
        newDir = turn(currentPos, newDir)
        nextPos = computeNextPos(currentPos, newDir)
    }



    // check border
/*
    if pos.y == 0 {
        if pos.x < 29 {
            dir = Right
        } else {
            dir = Down
        }
    } else if pos.y == 19 {
        if pos.x > 0 {
            dir = Left
        } else {
            dir = Up
        }
    } else if pos.x == 0 {
        if pos.y > 0 {
            dir = Up
        } else {
            dir = Right
        }
    } else if pos.x == 29 {
        if pos.y < 19 {
            dir = Down
        } else {
            dir = Left
        }
    }
    */

    return newDir
}

func initGame() string{
    // init board
    board = make([][]int, maxX)
    for x := 0 ; x < maxX ; x++ {
        board[x] = make([]int, maxY)
        for y := 0 ; y < maxY; y++ {
            board[x][y] = -1
        }
    }
    // init direction
    return Up
}

func main() {
    var lastDir, newDir string
    lastDir = initGame()
    myCurrentPos := coord{}

    /**
     * Auto-generated code below aims at helping you parse
     * the standard input according to the problem statement.
    **/

    for {
        // N: total number of players (2 to 4).
        // P: your player number (0 to 3).
        var N, P int
        fmt.Scan(&N, &P)

        for i := 0; i < N; i++ {
            // X0: starting X coordinate of lightcycle (or -1)
            // Y0: starting Y coordinate of lightcycle (or -1)
            // X1: starting X coordinate of lightcycle (can be the same as X0 if you play before this player)
            // Y1: starting Y coordinate of lightcycle (can be the same as Y0 if you play before this player)
            var X0, Y0, X1, Y1 int
            fmt.Scan(&X0, &Y0, &X1, &Y1)
            board[X1][Y1] = i
            if i == P {
                myCurrentPos =coord{X1, Y1}
            }
        }

        newDir = chooseDirection(myCurrentPos, lastDir)
        lastDir = move(newDir)
    }
}

