#!/bin/bash

# Script para ejecutar backend y frontend de Quest en paralelo

cd docker
./service.sh up

cd ../app
go run . &

cd ../frontend
npm install
npm start
