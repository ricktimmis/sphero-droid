
Obviously you'll be needing a Sphero Robot ball. I have the version 2.0

I'm afraid I'm a Linux only dude, these instructions should work for Debian based distro's
I am using Kubuntu 15.10 at the time of writing.

You'll need Bluetooth enabled, and active

Pair your Sphero with Bluetooth, check out the pairing instructions that come with Sphero
the default pairing password is 1234

I used the Blueman-Applet 
  $> apt-get install blueman
  
Once paired blueman will tell you which rfcomm serial port has been assigned to Sphero e.g /dev/rfcomm1
Change the adaptor assignment in main() to suit your settings.

  adaptor := sphero.NewSpheroAdaptor("sphero", "/dev/rfcomm5")
  
Logon to irc.freenode.net #sphero-control

Build and Install sphero-droid, and run it.

At the moment you need sudo to give you read/write priviledges to /dev/rfcomm resources

Sphero should logon to IRC after a few seconds....
