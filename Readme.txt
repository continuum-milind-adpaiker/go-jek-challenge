Project is divided into below layers:
1. cliMain - client program
2. Data Access Layer - Data Access Layer which reads input from and write output to file
3. Service - Starts the game by invoking go routines
4. Model - Data model for the project

Basic error handling like invalid values in input.txt is in place

Exe name - go-jek-challenge/src/climain/battleField

How to run exe:
./battleField --input=../input.txt
By default reads input from src/input.txt. User can supply own absolute path

Output generated at src/output.txt

