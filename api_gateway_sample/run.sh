#!/bin/bash

# Start backend server
cd backend
go run main.go &
BACKEND_PID=$!

# Wait for backend to start
sleep 2

# Start gateway server
cd ../gateway
go run main.go &
GATEWAY_PID=$!

echo "Servers started. Press Ctrl+C to stop."

# Wait for Ctrl+C
trap "kill $BACKEND_PID $GATEWAY_PID; exit" INT
wait 