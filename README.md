# golang-docker-concurrency-exercise
Some Go language exercises in text processing and concurrency with Docker containers

### 3 exercises

* Exercise 1: wf1.go, Dockerfile1

    Go program to print out map showing frequency of occurence of words in a given text file
    Dockerized to run in a container (uses Dockerfile1). Currently hard coded to parse input file input_files/moby-000.txt, can be adjusted to work on any input file

    Assumes this is running on a Linux host with Docker runtime support.

    Tested on ubuntu 14.04 with Docker runtime version 1.12.1 

    Example usage: 

    *docker build --file Dockerfile1 -t wf1_v1 .*

    *docker run wf1_v1* 

* Exercise 2: wf2.go, Dockerfile2, input_files/*
 
    Similar function as above but this time processes multiple input files concurrently using GoRoutines and then aggregates the results from each file. Dockerized to run in a container (uses Dockerfile2). Parses all files present in the input_files/ directory in parallel and then merges the results from processing each individual result.
 
    Assumes this is running on a Linux host with Docker runtime support.

    Tested on ubuntu 14.04 with Docker runtime version 1.12.1 

    Example usage: 

    *docker build --file Dockerfile2 -t wf2_v1 .*

    *docker run wf2_v1* 


* Exercise 3: <<work in progress>>

