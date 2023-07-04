# Ethscanner CLI
Simple command line tool to scan ethereum network's incoming or outgoing tranasctions

## List of available commands:
* ``` block ``` Returns last parsed block number

* ``` sub ``` Subscribes address to track it's transactions (program will ask to enter address)

* ``` tx ``` Retreives list of address's transactions since subscription block (program will ask to enter address)

* ``` exit ``` Exits application

## List of simplifications
* No external libraries has been used, following **KISS** and **YAGNI** principle for simple **MVP**
* Program parses block starting from the latest one, not from genesys block
* Program consumes memory a lot because of in-memory storage, don't forget to turn it off after you're done
* Program doesn't validate address format, so it's up to user to enter valid ethereum address
* Program intended for personal non-production usage, so error handling and recovery simplified
* Parser's module exposes minimal interface as it was given in requirements
* Ethereum client, Storage and Daemon are internal details of module, don't expect separate modules for them
* Program was written to to do what's been required, with minimal effort following **YAGNI** principle
* Command line interface is just for the demo purposes, so it's not intended to follow any of Golang's CLI guides


## Requirements

https://trustwallet.notion.site/Backend-Homework-Tx-Parser-abd431fca950427db75d73d90a0244a8

![image](https://github.com/54k/ethscanner/assets/3462493/3631fd43-4c9c-47b0-927a-39fd527a4869)
