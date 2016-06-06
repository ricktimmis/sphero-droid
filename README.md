# sphero-droid
An implementation of gobot to control the Sphero Robot Orb via Internet Relay Chat IRC

Sphero-Droid is an implementation of gobot.io and go-ircevent, ( both can be found on GitHub.com ) which uses
Internet Relay Chat as a communication channel and programming interface..

Overview
--------
The project implements the gobot libraries, coupled with go-ircevent.

The software connects to Sphero via Bluetooth, and logs into a definable IRC channel.
In my implementation this is irc.freenode.net #sphero-control.
It provides a variety of commands for controlling movement, and the lighting colours of the droid.
It can be controlled by others logged into IRC, and makes for fun times if it is being controlled
remotely by others who can watch its actions via webcam.

Commands and Controls
---------------------
Control Sphero-Droid by passing him commands via IRC.
When available, you can find him logged into IRC at irc.freenode.net in channel #sphero-control.

Format: sphero command [parameter] 
            
            help
            speed  [slow, medium, fast, turbo]
            colour [red, green, blue, purple, yellow, orange]
            move   [fwd, lft, rht, bwd]
            start
            stop
            settings
            location
            stack  [push, pop, loop, show, exec]
            
Sphero posts responses, in channel, to the commands you send him

Command Stack
-------------
The most interesting and programatically useful feature of Sphero-Droid is the command stack.
The stack enables you to prime Sphero with a sequence of commands to execute.

They are called using the "Stack" command here are what the commands do :-

push  :- push the following command onto the stack.
pop   :- pop the last entry from the stack.
loop  :- Reset the stack pointer to 0 and cycle through the commands again. usually you push this to the stack as the last entry
show  :- List the current stack entries
exec  :- Execute the stack. This iterates through the command set one instruction per Tempo cycle

The Tempo cycle, is the pulse of Sphero. Each cycle lasts a specific number of Milliseconds, and is defined by the internal
botTempo variable as a number ( default 1000 ) Milliseconds.

Example
-------
        sphero stack push speed medium
        sphero stack push colour red
        sphero stack push move fwd
        sphero stack push move fwd
        sphero stack push move fwd
        sphero stack push colour green
        sphero stack push stop
        sphero stack exec

In the above example we set the speed of Sphero to medium, and switch him to red.
Sphero does not begin to move until the first move command is executed, which wouldbe 3 seconds after "sphero stack exec"
command is issued. he then moves forward for a further 2 seconds, changes colour to green whilst still moving forward for 
a further 1 second, and finally Sphero stops. The stop command sets Sphero's speed, direction and lights all back
to zero ( i.e off ).

With just this simple command set it is possible to devise simple games and challenges, here are some examples

1.) Camelion  :-  Using coloured paper on the arena floor, the challenge is to get Sphero to Stop on each Spot whilst
                  disguising himself in the correct colour.
                  
2.) Labyrinth :-  Set up a Labyrinth using cardboard tubes, Sphero must negotiate it, knocking down as few tubes as possible.

3.) Skittles  :-  With a defined time limit, you must control Sphero to knock down as many cardboard tubes as possible.

I hope you like Sphero-droid, and get the opportunity to play not only with him but also the code. Please change extend modify,
pull requests and feature ideas, bugs etc..are all welcome too :-D Enjoy !
