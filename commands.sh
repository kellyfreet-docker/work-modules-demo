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

# Create v1 and v2 directories for common module
mv petcare-common petcare-common-v1
cp -r petcare-common-v1 petcare-common-v2

# Update v2 with some changes
cd petcare-common-v2
# We'll modify the module name and add some new functionality 

# Restructure common directory
cd petcare-common
mkdir v1 v2
mv pet.go utils.go v1/
cp v1/pet.go v1/utils.go v2/ 