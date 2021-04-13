# logtime
A simple experiment with time.Time, channels, and wait groups.
Logtime is a program that a potential employer asks me to write.
 I guess they're tired of FizzBuzz ;-)

 The initial definition was to write a program that will:
 1. display's 'tick' every second.
 2. display 'tock' every minute.
 3. display 'bong' every hour.
 4. And to stop running after 3 hours.
 5. And send the source code to them within 30 minutes.

 So I begin writing, ...
 I'll need a ticker channel, and a for loop reading the ticker,
  and I'll log messages based on a modulus of time.Duration.... sounds simple enough.
 I can do that in 30 minutes, right? heck maybe in 10 minutes!

 Then the potential employer says: Oh, and one more thing......
 we want to be able to change the 'tick', 'tock' or 'bong' messages
 to other messages while the program is running.

 Now that's a little more complicated.
 It will require the messages to be read as input into the program.

 Should I read in a file, if so I would need to read in the file on
 every iteration of the ticker loop.
 Reading in a file in every iteration, no matter if the file has changed or not, seems like a waste of CPU cycles.

 Or I could use a Go routine to read stdin, while the ticker runs in main?
 That seems messy, typing in a new message into stdin, while the program is writing tick\n tick\n tick.... to stdout.
 Also there are three different messages, I'll have to include
 which message to replace, second, minute or hour.

 Like sec=tick or hr=tock, then parse the input to see which msg to update.


 Which is the quickest to write?

 Quick should not be the question, but is part of the test.

 The correct way to approach this is a combination of both.
 1. A Go routine, but to read from stdin.
 2. Read from a file, but not in every iteration of the ticker loop.

I will use a Go routine to watch for file system changes, specifically to watch one file. The file where the messages (tick,tock,bong) are stored, a config file. When changes are made, the go routine will send a msg to the main routine via channel. When the main routine gets the msg, it will read the config file, and update the messages.

The config file will be a json file, which will map to a struct, will read and unmarshal it. 

I've used fsnotify to watch file system activity before, so I've thought it all thru.... 20 minutes left to go.

At 5 minutes left to go, I call the potential employer, and tell them that I'm almost done, but realistically it's gonna be another 15-20 minutes. They reply: That's OK, everyone has taken more or less about an hour. We thought it could be done in 30 minutes, but that's ok, sent us your code when you are done.

And it took me an hour to get it working, the code was messy, but it worked perfect.

They said my code was confusing, to complicated, too many hard-coded literals. So I didn't get the job. :-(

    Since the race is over.... I've cleaned up the code, considered the critique, it much easier to understand, except for some log messages there are no hard-coded literals. It's design and functionality is like a Swiss-Watch.