/*
To get this program to work you need two files

1. token.txt (A file containing your discord bots token)


*/

//https://discord.com/api/oauth2/authorize?client_id= 758403936784482315&scope=bot&permissions=1

const Discord = require('discord.js')
const fs = require('fs');
const client = new Discord.Client()

//Allows exectuing programs on the server
const { exec } = require("child_process");



//https://discord.com/oauth2/authorize?client_id=717442260131774485&scope=bot


var token = fs.readFileSync('token',{encoding:'utf8', flag:'r'});
var theUserIndex = null
token = token.replace(/(\r\n|\n|\r)/gm, ""); //Removes the newline from the token file
client.login(token)




//Runs neofetch on the server. The program must be installed 
function runCommand(msg){

    var command = msg.content.substring(6);

    exec(command, (error, stdout, stderr) => {
        // if (error) {
        //     //console.log(`error: ${error.message}`);
        //     msg.reply("Could not run neofetch")
        //     return;
        // }
        if (stderr) {
            msg.reply(`There was an error: ${stderr}`)
            console.log(`stderr: ${stderr}`);
            return;
        }
        msg.reply(`\n${stdout}`);
    });

}







//When the client connects
client.on('ready', () => {

  console.log(`logged in as ${client.user.tag}!`)

  console.log(client.user.id)

  // Does this  work? Well at some point it did 


  //Sets the bots username and activity
  client.user.setUsername('DISH');
  //client.user.setActivity(':-)');
  client.user.setActivity('your shit', { type: 'pwning' });

  

})




//Runs when the message is read
client.on('message', msg => {

    //Don't talk to yourself 
    if (msg.author.id == "758403936784482315"){ //bot id
        console.log("It's me")
        return 0; 
    }

  
    // Commands (Soon this will be a different thing )
    if (msg.author.id === "326743504443146241" ){  // My ID.
        var arg = null



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
