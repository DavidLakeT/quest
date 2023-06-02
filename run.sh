#!/bin/bash

# Script para ejecutar backend y frontend de Quest en paralelo

cd app
go run . &

cd ../frontend
npm start
