# DISH - DIscordSHell
## Shell or command and control? 
So this project is an interesting one. It started off as a fun way to control my PC through Discord. Like I could make the PC say "boobies" from across the room, without having to set up SSH. Then I realised, damn, this is a security nightmare. Then I realised DAMN this is a security nightmare. So it kind of became a command and control concept. I haven't really checked, but this feels like the kind of thing that is done to death, so nothing exciting. 

## Setting it up
This project is VERY early on. But at the moment all you will need is 
- A valid discord bot token. In a file called "token" 
- The discord python library
- A discord server, with the discord tokens bot added to it 

## Oooh, I could use this for illegal shit. 
Yes probably. Please don't though. 

## What operating systems has this been tested with
- MacOs (This is what i am using to build most of it)
- Windows 
- Ubuntu (will probably be compatible with most distros)

## Things that need to get done 
- Need to make a "deploy" script 
    - So it's unlikely that a target would have discord.py pre-installed. And setting up Pip to do that would be a ball ache. So i need to make a script, that the user can run when the code is on their own machine, to get it all set up.  
- Need to add a nicer way of sending files
    - At the moment it's just downloading any file you send. This doesn't seem like a great way of doing things. 
- Need to add a way to exfiltrate files 
    - When you put files in, you probably want to get files out. 
    - would make sense to have this for files and folders. 


## Big ol' issues
- If someone was to see they were infected, they could just take over the other nodes by changing what messages their node sent

## Cool future plans 

### Mother ship 
So this would change how the bot works, instead of the user interacting with Discord to send commands, the user would interact with a console application to give commands. Discord would just be the middle man. There would be two types of bots; nodes, motherships. 

Advantages:
- The mothership could store information about the nodes, meaning that the user could do some qurires without having to send commands to the node. 
- The mothership could be used to schedule when commands are sent
- The mothership could keep track of nodes coming and going 
- The discord chat would look cool, and random (would probably encode the commands)
- would mean a comprimised node couldn't control over nodes (as a mothership wouldn't take commands from a node)

Disadvantages: 
- Work 
- added complexity makes most projects bad 
- People are scared of the command line
