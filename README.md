# GreedyGame-Assignment

## Introduction

Creating a simple in-memory key-value datastore that performs operations on it based on certain
commands and uses REST API for communication.

## Getting Started

Followed clean architecture and dependency-driven api. 
Directory cmd - contains command-line files main.go
Directory pkg - contains packages of four layers 1. Delivery(routes and handlers), 2. Use case (Business logic), 3. Repository (Data structures and data storing), 4. Domain (entities, models)


## Usage

Added .env for specifying environments, git ignored the env so add env before running the server (add PORT=:XXXX in env)
Added Makefile for easy command execution (make run - for starting server) instead of execute go run cmd/main.go from the root directory

## API Documentation

Have five API

1. Method: POST  Endpoint: baseurl/key/set  eg: localhost:8080/key/set
   Input Json
   {
    "Key": "sample_key",
    "Value": {
        "Value": "sample_value",
        "Expiry": 40,
        "Condition": "XX"
    }
   }
   here expiry and condition are optional inputs: give 0 for expiry and "" for the condition when no value is provided.

   Response
   success      - status code: 201  { "Value stored for key": "2"}
   json error   - status code: 422  {"error parsing json": "err"}
   invalid key  - status code: 400  {"error": "enter a valid key"}
   server error - status code: 400  {"error from server": "err"}

2. Method: GET Endpoint: baseurl/key/get
   Input Json
   {
    "Key": "4"
   }

   Response  
   success      - status code: 301  {"Value for key ": "s", "is": {"Value": "sample_value", "Expiry": "", "Condition": "XX"}}
   key not found- status code: 400  {"error from server": "Key not found"}
   json error   - status code: 422  {"error parsing json": "err"}
   invalid key  - status code: 400  {"error": "enter a valid key"}
   server error - status code: 400  {"error from server": "err"}

3. Method POST Endpoint: baseurl/que/push
   Input Json
   {
    "Key": "1",
    "Value": ["1", "2", "3"]
   }

   Response
   success      - status code: 201  {"Value stored for key ": "1"}
   json error   - status code: 422  {"error parsing json": "err"}
   invalid key  - status code: 400  {"error": "enter a valid key"}
   server error - status code: 400  {"error from server": "err"}

4. Method GET Endpoint: baseurl/que/pop
   Input Json
   {
    "Key": "1"
   }

   Response
   success      - status code: 302  {"Value for que ": "1", "is": "3"}
   key not found- status code: 400  {"error from server": "Key not found"}
   json error   - status code: 422  {"error parsing json": "err"}
   invalid key  - status code: 400  {"error": "enter a valid key"}
   server error - status code: 400  {"error from server": "err"}

5. Method GET Endpoint: baseurl/que/bpop
   Input Json
   {
    "Key": "1",
    "Time": 2
   }

   Response 
   success      - status code: 302  {"Value for que": "1","is": "2","waited seconds": 2}
   key not found- status code: 400  {"error from server": "Key not found"}
   json error   - status code: 422  {"error parsing json": "err"}
   invalid key  - status code: 400  {"error": "enter a valid key"}
   server error - status code: 400  {"error from server": "err"}


   
