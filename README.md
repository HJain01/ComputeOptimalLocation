# Compute Optimal Location Service

## Description

This is the backend api to an app that find the most optimal destination with multiple starting locations.
You input multiple starting locations and multiple ending locations, and the api will spit out the best destination
using the variance of the distance between each origin and destination

## Running the app

This app runs in a docker container, so you will need Docker and Docker Compose installed on your machine.
There is a bash script named `_up.sh` in `./env/local` that runs the app locally. This will create and run a docker container in the background.
You can use `_down.sh` in `./env/local` to break down all the containers associated with this app (which is currently just the one).