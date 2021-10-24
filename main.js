

const Discord = require('discord.js')
const fs = require('fs');
const client = new Discord.Client()

//Allows exectuing programs on the server
const { exec } = require("child_process");


var token = fs.readFileSync('token',{encoding:'utf8', flag:'r'});
var theUserIndex = null
token = token.replace(/(\r\n|\n|\r)/gm, ""); //Removes the newline from the token file
client.login(token)



// run a command s
function runCommand(msg){

    var command = msg.content.substring(6);

    exec(command, (error, stdout, stderr) => {
        if (stderr) {
            msg.reply(`There was an error: ${stderr}`)
            console.log(`stderr: ${stderr}`);
            return;
        }
        console.log(`${stdout}`)
        msg.channel.send(`\n${stdout}`)
    });

}


//When the client connects
client.on('ready', () => {

  console.log(`logged in as ${client.user.tag}!`)

  console.log("details: ")
  console.log(client.user.id)

  //Sets the bots username and activity
  client.user.setUsername('DISH');
  //client.user.setActivity(':-)');
  client.user.setActivity('YOUR MACHINES', { type: 'WATCHING' });

  

})




//Runs when the message is read
client.on('message', msg => {

    //Don't talk to yourself 
    if (msg.author.id == "758403936784482315"){ //bot id
        console.log("It's me")
        return 0; 
    }

  
    // Commands (Soon this will be a different thing )
    if (msg.author.id === "326743504443146241" ){  // My ID. This stops people running commands, on my bot

        if(msg.content.substring(0,5) == "DISH:"){

            runCommand(msg)

        }
    }


})



//Says when it disconnects
client.on("disconnected", function () {

	console.log("Disconnected!"); // send message that bot has disconnected.
	process.exit(1); //exit node.js with an error

});
