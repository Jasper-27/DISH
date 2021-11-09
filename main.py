# imports

import os
import subprocess
import discord

import uuid #used for checking unique hosts 

id = (hex(uuid.getnode()))

#Getting the users token
f = open("token", "r")
token = f.read()
f.close()

os.chdir(os.path.expanduser('~'))


client = discord.Client()

@client.event
async def on_ready():
    print(f'{client.user} has connected to Discord!')

@client.event
async def on_message(message):
    
    # Bit off fluff to see the message in the terminal 
    print("\n=======   ")
    print("Message: ")
    print(message.content)
    print("++++++++ \n")


    # Checks if this is the node the command is meant for 
    idString = "[" + id + "]! "
    if message.content.startswith(idString): 

        command = message.content.split(idString)[1]
        print(message.content)
        print(command)

        ### This bit is repeated code, and needs to be fixed 

        # just running "cd" doesn't change the wd or the pythin script
        if command.startswith("cd"):
            path = command.split("cd ")[1]
            os.chdir(path)
            return 
        

        # I wan't to find a better way of doing this. It has some issues 
        result = subprocess.check_output(command, shell=True, text=True)

        if result == "": 
            await message.channel.send("Command produced no output")
            return 
        
        
        await message.channel.send(result)
        return




    # Get file by path. (sends the file to the chat)
    if message.content.startswith("DISH GETFILE"):
        await message.channel.send ("Getting file")
        filePath = message.content[13:]

        try: 
            await message.channel.send(file=discord.File(filePath))
        except: 
            await message.channel.send("Error retrieving file: " + filePath)
        return 

    # Report in to the sever, so the user can see what nodes are online 
    if message.content == ("dish report"):

        reportString = "=========" + "\n" + "ID: " + str(id) + "\n" + str(os.uname()) + "\n" + "========="

        await message.channel.send(reportString) 



    # Standard running of commands on all nodes 
    if message.content.startswith("!"):
        command  = message.content[2:]

        print("Command: " + command)


        # just running "cd" doesn't change the wd or the pythin script
        if command.startswith("cd"):
            path = command.split("cd ")[1]
            os.chdir(path)
            return 
            

        # I wan't to find a better way of doing this. It has some issues 
        result = subprocess.check_output(command, shell=True, text=True)

        if result == "": 
            await message.channel.send("Command produced no output")
            return 
        
        
        await message.channel.send(result)
        return

       



    # no need to d anything with the bots own messages 
    if message.author == client.user:
        return

    if message.content == 'test':
        response = "Hello World"
        await message.channel.send(response)

   
    # Downloads any files that are sent, and stores them in the downloads folder 
    if str(message.attachments) != "[]": # Checks if there is an attachment on the message
        print(message)
        split_v1 = str(message.attachments).split("filename='")[1]
        filename = str(split_v1).split("' ")[0]
  
        await message.attachments[0].save(fp="downloads/{}".format(filename)) # saves the file
        await message.channel.send("File recieved")


    
# run the bot 
client.run(token)


