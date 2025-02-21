#!/bin/bash
# Run the API service in the background
go run ./petcare-api &
# Run the scheduler service
go run ./petcare-scheduler 