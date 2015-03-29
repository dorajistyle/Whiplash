// Copyright 2015 JoongSeob Vito Kim. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.
package game

import (
"os"
"fmt"
"bufio"
"time"
"os/exec"
"strconv"
)

var (
indent = "    "
input = "> "
BPM float64
username, bpmString string
CPS, margin time.Duration
bpmStringLength int
skipCheckingRoD bool // Rushing or Dragging
)

func IndentedPrintf(str string, arg... interface{}) {
        fmt.Printf(indent+str, arg...)
}

func IndentedPrint(str string) {
        fmt.Print(indent+str)
}

func IndentedPrintln(str string) {
        fmt.Println(indent+str)
}

func intro() {
    whatIsYourName := "What's your name?"
    IndentedPrintln(whatIsYourName)
    IndentedPrint(input)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    username = scanner.Text()
    introAnswer := "I'm here for a reason."
    intro := `
    `+username+`, you just gotta listen
    to the greats, then.
    Richard Stallman, Ken Thompson...
    You know, Rob Pike
    became Owl...
    ...because Thompson threw a topic
    at his head.
    See what I'm saying?
    Listen, the key is to just relax.
    Don't worry about the numbers.
    Don't worry about what
    other guys are debuging.
    You're here for a reason.
    You believe that, right?
    Type it. '`+introAnswer+`'(Ignore quotes around the sentence.)`
        fmt.Println(intro)
    for{
      IndentedPrint(input)
      scanner.Scan()
      userIntroAnswer := scanner.Text()
      if userIntroAnswer == introAnswer {
        IndentedPrintln("Cool. All right, man. Have fun.")
        break
      }
      IndentedPrintln("Wrong answer. Type it again.")
    }
    fmt.Println("")
}
func CheckCharacter(keyIn string, index int) bool{
    if keyIn != string(bpmString[index]) {
      fmt.Println("")
      IndentedPrintln("Am I to understand that you cannot read tempo?")
      IndentedPrintln("Can you even read alphabet? What's that?")
      time.Sleep(2 * time.Second)
      IndentedPrintln("Again.")
      time.Sleep(time.Second)
      HelloWorld()
      return false
    }
    return true
}

func RushingOrDragging(elapsed time.Duration) bool{
   var yesOrNo, rushingOrDragging string
   if elapsed > CPS+margin || elapsed < CPS-margin {
          fmt.Println("")
          IndentedPrintln("Not quite my tempo.")
          IndentedPrintln("Were you rushing or dragging? [rushing/dragging]")
          scanner := bufio.NewScanner(os.Stdin)
          scanner.Scan()
          line := scanner.Text()
          if elapsed > CPS+margin {
              rushingOrDragging = "dragging"
          } else {
              rushingOrDragging = "rushing"
          }

          if line == rushingOrDragging {
              yesOrNo = "Yes"
          } else {
              yesOrNo = "No" 
          }
          fmt.Println("")
          IndentedPrintf(yesOrNo+". You were "+rushingOrDragging+". It should be %s but you did %s", CPS, elapsed)
          
          HelloWorld()
          return false
        }
    return true
}


func HelloWorld() {
    IndentedPrintln("Here we go. (Wait 4 count and type '"+bpmString+"')")
    time.Sleep(CPS)
    IndentedPrintln("Five,")
    time.Sleep(CPS)
    IndentedPrintln("Six,")
    time.Sleep(CPS)
    IndentedPrintln("and...")
    time.Sleep(CPS)
    IndentedPrintln("Type!")
    var totalDuration time.Duration
    IndentedPrint(input)
    i := 0
    for {
        var keybuf [1]byte
        cmd := exec.Command("/bin/stty", "-F", "/dev/tty", "-icanon", "min", "1")
        cmd.Run()
        start := time.Now() 
        _, _ = os.Stdin.Read(keybuf[0:1])
        isCorrect := CheckCharacter(string(keybuf[0]),i)
        if !isCorrect {
          break;
        }
        elapsed := time.Since(start)
        if !skipCheckingRoD {
            isCorrect = RushingOrDragging(elapsed)
            if !isCorrect {
                break;
            }
        }
        totalDuration = totalDuration + elapsed
        i = i+1
        if i == bpmStringLength {
           fmt.Println("")
           fmt.Printf("It should be %s. And you did %s\n", CPS*(time.Duration(bpmStringLength)), totalDuration)
          break
        }
    }
}

func Whiplash(bpm float64, text string, skipIntro bool, skipChecking bool, level string) {
    username = "Zen"
    BPM = bpm
    bpmString = text
    bpmStringLength = len(bpmString)
    skipCheckingRoD = skipChecking
    CPS = time.Duration(1/(BPM/60)*1000) * time.Millisecond // character per second
    margin = CPS/10 // margin
    switch level {
      case "easy":
        margin = margin * 3
      case "normal":
        margin = margin * 2
      case "guru":
        margin = CPS/100
    }

//    fmt.Printf("text len skipintro skipChecking : %s, %d, %t %t\n",bpmString,bpmStringLength,skipIntro, skipChecking)
//    fmt.Printf("BPM CPS margin : %d, %s, %s\n",BPM,CPS,margin)

    if !skipIntro {
        intro()
    }

  bpmStr := strconv.FormatFloat(BPM, 'G', 'G', 64)
    IndentedPrintln(username+", just do your best.")
    IndentedPrintln("Type '"+bpmString+"'("+strconv.Itoa(bpmStringLength)+" characters). "+bpmStr+" BPM.")
    IndentedPrintln("Ready? (Push the return key when you ready.)")
    IndentedPrint(input)
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    HelloWorld()
}
