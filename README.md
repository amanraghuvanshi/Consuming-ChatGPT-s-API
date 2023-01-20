# Consuming-ChatGPT-s-API
Consuming the ChatAPI
THis is a CLI tool that interacts directly with the ChatGPT SERVERS.

This tools will again take the API key but instead of using joho/godotenv we will be using the Viper and Cobra from spf13. 

Configiure your go mod file using > go mod init <filename>
and then set it up using 
> go get  github.com/PullRequestInc/go-gpt3 <br>
> go get github.com/spf13/viper <br>
> go get github.com/spf13/cobra <br>


After performing these stuffs we will start working on the function were we would be taking the input from the user for the search request

In the main function, we would start by setting up the enviroment for the API key and would be storing them in a variable and using the viper package by spf13, we would retrieve them

Then we will set the context and create a client for the gpt3 for connecting the endpoint of API, And using the Cobra package we would be commanding the controls over the API

We would be running a command that will have the command line arguments function for taking input from the command line.
Using the scanner function of BUF.IO we would be taking inputs. 
A variable would be declared in case if the user wants to quit
And along with it conditions will be applied

We would also define a function that will response and it will take a client, context, and quest that is searched. And using the CompletionStreamwWithEngine function of client from the PullRequestInc's Repo we would be framing the response
