mkdir petcare-system
cd petcare-system

mkdir petcare-api
mkdir petcare-scheduler
mkdir petcare-common

# Initialize go.mod files
cd petcare-api
go mod init github.com/example/petcare-api

cd ../petcare-scheduler
go mod init github.com/example/petcare-scheduler

cd ../petcare-common
go mod init github.com/example/petcare-common

# Create go.work in root directory
cd ..
go work init
go work use ./petcare-api ./petcare-scheduler ./petcare-common 