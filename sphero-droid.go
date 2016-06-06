package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot/platforms/sphero"
	"github.com/thoj/go-ircevent"
)

// Various global variables - FixMe need to reduce their scope

// IRC Info
var ircServer = "irc.freenode.net:6667"
var roomName = "#sphero-control"

// Robot parameters
var botSpeed uint8
var botRED uint8
var botGRN uint8
var botBLU uint8
var botDirection uint16
var botLocation []int16
var commandLine [10]string
var message string

// Tempo of gobot.Every function
// Defined in Microseconds, sets the loop rate
var botTempo int64 = 1000
var botCmdStackPointer int = 0
var botCmdStack []string

// Flag true to execute command stack
// once stack[pointer] == "" assumes end of command list
// Flag reset false
var execstack bool //= false
var spointer int = 0

// Probably botching this, this enables Sphero to flash in Tempo
var botCycle bool // = false
func parse(commands []string) string {
	switch commands[1] {

	case "location":
		{
			fmt.Println("Sphero location requested")
			//con.Privmsg(roomName, "Sphero is located at")
			// FIXME botLocation = driver.ReadLocator()
			//con.Privmsg(roomName, "Hiding!")
			//gbot.Start()
			message = "Sphero is Hiding!"

		}
	case "tempo":
		{
			// FIXME
			/*
			*  As the tempo is assigned at gbot instantion time
			*  resetting the botTempo value makes no difference
			*
			 */
			fmt.Println("Sphero Set Tempo requested")
			if botTempo, err := strconv.ParseInt(commandLine[2], 10, 64); err == nil {
				fmt.Printf("%T, %v\n", botTempo, botTempo)
			}
			//gbot.Stop()
			//gbot.Start()
		}
	case "move":
		{
			switch commands[2] {
			case "fwd":
				{
					fmt.Println("Sphero forward requested")
					message = "Direction Event received - Sphero is now going Forward"
					botDirection = 0
				}
			case "rht":
				{
					fmt.Println("Sphero right requested")
					message = "Direction Event received - Sphero speed now going Right"
					botDirection = 90
				}
			case "lft":
				{
					fmt.Println("Sphero left requested")
					message = "Direction Event received - Sphero speed now going Left"
					botDirection = 270
				}
			case "bwd":
				{
					fmt.Println("Sphero Backward requested")
					message = "Direction Event received - Sphero is now going Backward "
					botDirection = 180
				}
			}
		}
	case "speed":
		{
			switch commands[2] {
			case "slow":
				{
					fmt.Println("Sphero slow speed requested")
					message = "Speed Event received - Sphero speed now Slow"
					botSpeed = 20
				}
			case "medium":
				{
					fmt.Println("Sphero medium speed requested")
					message = "Speed Event received - Sphero speed now Medium"
					botSpeed = 40
				}
			case "fast":
				{
					fmt.Println("Sphero fast speed requested")
					message = "Speed Event received - Sphero speed now Fast"
					botSpeed = 60
				}
			case "turbo":
				{
					fmt.Println("Sphero turbo speed requested")
					message = "Speed Event received - Sphero speed now Turbo"
					botSpeed = 100
				}
			}

		}

	case "colour":
		{
			switch commands[2] {
			case "red":
				{
					fmt.Println("Sphero colour Red requested")
					message = "Colour Event received - Sphero is now Red"
					botRED = 246
					botGRN = 42
					botBLU = 42
				}
			case "green":
				{
					fmt.Println("Sphero colour Green requested")
					message = "Colour Event received - Sphero is now Green"
					botRED = 43
					botGRN = 232
					botBLU = 38
				}
			case "blue":
				{
					fmt.Println("Sphero colour Blue requested")
					message = "Colour Event received - Sphero is now Blue"
					botRED = 39
					botGRN = 123
					botBLU = 248
				}
			case "purple":
				{
					fmt.Println("Sphero colour Purple requested")
					message = "Colour Event received - Sphero is now Purple"
					botRED = 194
					botGRN = 17
					botBLU = 215
				}
			case "yellow":
				{
					fmt.Println("Sphero colour Yellow requested")
					message = "Colour Event received - Sphero is now Yellow"
					botRED = 244
					botGRN = 244
					botBLU = 45
				}
			case "orange":
				{
					fmt.Println("Sphero colour Orange requested")
					message = "Colour Event received - Sphero is now Orange"
					botRED = 250
					botGRN = 200
					botBLU = 20
				}
			}
		}

	case "start":
		{
			fmt.Println("Sphero start requested")
			message = "Sphero Awakens - Starting Sphero"
			botRED = 50
			botGRN = 50
			botBLU = 50
			botDirection = 0
			botSpeed = 0
		}
	case "stop":
		{
			fmt.Println("Sphero stop requested")
			message = "Sphero Sleeps - Stopping Sphero"
			botRED = 0
			botGRN = 0
			botBLU = 0
			botDirection = 0
			botSpeed = 0
			execstack = false
			spointer = 0
		}
	}
	return message
}

func main() {
	botCmdStack := make([]string, 20, 40)
	gbot := gobot.NewGobot()

	adaptor := sphero.NewSpheroAdaptor("sphero", "/dev/rfcomm5")
	driver := sphero.NewSpheroDriver(adaptor, "sphero")

	work := func() {
		gobot.Every(time.Duration(botTempo)*time.Millisecond, func() {
			driver.Roll((botSpeed), (botDirection))
			if botCycle == true {
				driver.SetRGB((botRED), (botGRN), (botBLU))
				botCycle = false
			} else {
				driver.SetRGB(0, 0, 0)
				botCycle = true
			}
			fmt.Println(driver.ReadLocator())
			if execstack == true {
				commands := strings.Split(botCmdStack[spointer], " ")
				fmt.Println(commands)
				parse(commands)
				//con.Privmsg(roomName, message)
				spointer++
				if botCmdStack[spointer] == "sphero stack loop" {
					spointer = 0
				}
				if botCmdStack[spointer] == "" {
					execstack = false
					spointer = 0
				}
			}
		})

		// Connects to IRC, joins room and awaits commands
		//ircCtrl := func() {
		con := irc.IRC("sphero-control", "Sphero IRC Control Bot")
		err := con.Connect(ircServer)
		if err != nil {
			fmt.Println("Failed connecting")
			return
		}
		con.AddCallback("001", func(e *irc.Event) {
			con.Join(roomName)
		})
		con.AddCallback("JOIN", func(e *irc.Event) {
			con.Privmsg(roomName, "Hello! I am the Sphero IRC Robot Orb.")
			con.Privmsg(roomName, "Type: 'sphero help' for instructions.")
		})
		con.AddCallback("PRIVMSG", func(e *irc.Event) {
			commandLine := strings.Split(e.Message(), " ")
			//	DEBUG: Show IRC input parsing to separate commands
			//	for index, each := range commandLine {
			//		fmt.Printf("Command value [%d] is [%s]\n", index, each)
			//	}

			if commandLine[0] == "sphero" {
				// DEBUG: Announce a command call
				//fmt.Println("Sphero called - enter command")
				switch commandLine[1] {
				case "help":
					{
						fmt.Println("Sphero help requested")
						con.Privmsg(roomName, "Sphero is a small robot orb, which you control by typing commands here in IRC ( Internet Relay Chat)")
						con.Privmsg(roomName, "Format: sphero <command> <parameter>")
						con.Privmsg(roomName, "Command Reference: http://www.ricktimmis.com/projects/sphero-droid.html")
						con.Privmsg(roomName, "speed: slow, medium, fast, turbo")
						con.Privmsg(roomName, "colour: red, green, blue, purple, yellow, orange")
						con.Privmsg(roomName, "move: fwd, lft, rht, bwd")
						con.Privmsg(roomName, "start:, stop:")
						con.Privmsg(roomName, "settings: Show the current parameter values")
						con.Privmsg(roomName, "location: Show the current  location of Sphero")
						con.Privmsg(roomName, "stack: push, pop, loop, show , exec")
					}
				case "stack":
					{
						switch commandLine[2] {
						case "push":
							{
								fmt.Println("Sphero Stack Push requested")
								if commandLine[4] == "" {
									botCmdStack[(botCmdStackPointer)] = "sphero " + commandLine[3]
								} else {
									botCmdStack[(botCmdStackPointer)] = "sphero " + commandLine[3] + " " + commandLine[4]
								}
								botCmdStackPointer++
							}
						case "pop":
							{
								botCmdStackPointer--
								botCmdStack[(botCmdStackPointer)] = ""
								fmt.Println("Sphero stack Pop requested")
							}
						case "show":
							{
								fmt.Println("Sphero stack contents requested")
								for stackpointer, command := range botCmdStack {
									fmt.Printf("Stack Pointer [%d] has [%s]\n", stackpointer, command)
									//msg_string = ("Stack Pointer [%d] has [%s]\n", stackpointer, command)
									//con.Privmsg(roomName, msg_string)
								}
							}
						case "exec":
							{
								fmt.Println("Stack Execution Flag = true")
								execstack = true
							}
						}
					}
				case "halt":
					{
						fmt.Println("Sphero stop requested")
						message = "Sphero Halt - Shutting down Sphero & Exiting"
						botRED = 0
						botGRN = 0
						botBLU = 0
						botDirection = 0
						botSpeed = 0
						gbot.Stop()
						break
					}
				default:
					{
						parse(commandLine)
						con.Privmsg(roomName, message)
					}
				}
			}
		})

	}
	robot := gobot.NewRobot("sphero",
		[]gobot.Connection{adaptor},
		[]gobot.Device{driver},
		work,
	)

	gbot.AddRobot(robot)

	gbot.Start()
}
