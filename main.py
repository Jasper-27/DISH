# imports

import os
import subprocess
import discord

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


    if message.content.startswith("DISH GETFILE"):
        await message.channel.send ("Getting file")
        filePath = message.content[13:]

        try: 
            await message.channel.send(file=discord.File(filePath))
        except: 
            await message.channel.send("Error retrieving file: " + filePath)
        return 

    if message.content.startswith("DISH!"):
        command  = message.content[6:]

        print("Command: " + command)

        if command.startswith("cd"):
            path = command.split("cd")[1]
            print(path)
            # os.chdir(os.path.join(os.getcwd(), path)
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


