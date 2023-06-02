#!/bin/bash

# Script para ejecutar backend y frontend de Quest en paralelo

cd .quest/app
go run . &

cd ../frontend
npm start
