# Blockchain using Pow consensus algorithm 
## Steps to run :
Please use .env file according to .env.example

1. ```go mod init github.com/<Name>/<ProjName>```

2. ```go mod tidy```

- Use the Makefile and add the relevant name and project name
3. ```make build```
4. ```make run```

The server runs on localhost 3000
- Check the router file which provides all the mux routes relevant to creating nodes, blocks, checking and printing chain

- Also to run more than 1 node, change the port numbers (same machine different nodes)
