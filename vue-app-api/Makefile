# .PHONY: build run clean start stop restart
# SHELL=cmd.exe
# DSN=host=localhost port=5432 user=postgres password=password dbname=vue_app_api sslmode=disable timezone=UTC connect_timeout=5
# BINARY_NAME=vueapi.exe

# ## build: builds all binaries
# build:
# 	@go build -o ${BINARY_NAME} ./cmd/api
# 	@echo back end built!

# run: build
# 	@echo Starting back end...
# 	cmd /C "set DSN=${DSN} && .\${BINARY_NAME}"
# 	@echo back end started!

# clean:
# 	@echo Cleaning...
# 	@DEL ${BINARY_NAME}
# 	@go clean
# 	@echo Cleaned!

# start: run

# stop:
# 	@echo "Stopping the back end..."
# 	@taskkill /IM ${BINARY_NAME} /F
# 	@echo Stopped back end

# restart: stop start


# DSN="host=localhost port=5432 user=postgres password=password dbname=vueapi sslmode=disable timezone=UTC connect_timeout=5"
# BINARY_NAME=vueapi
# ENV=development

# ## build: Build binary
# build:
# 	@echo "Building back end..."
# 	go build -o ${BINARY_NAME} ./cmd/api/
# 	@echo "Binary built!"

# ## run: builds and runs the application
# run: build
# 	@echo "Starting back end..."
# 	@env DSN=${DSN} ENV=${ENV} ./${BINARY_NAME} &
# 	@echo "Back end started!"

# ## clean: runs go clean and deletes binaries
# clean:
# 	@echo "Cleaning..."
# 	@go clean
# 	@rm ${BINARY_NAME}
# 	@echo "Cleaned!"

# ## start: an alias to run
# start: run

# ## stop: stops the running application
# stop:
# 	@echo "Stopping back end..."
# 	@-pkill -SIGTERM -f "./${BINARY_NAME}"
# 	@echo "Stopped back end!"

# ## restart: stops and starts the running application
# restart: stop start


.PHONY: build run clean start stop restart

SHELL=cmd.exe

DSN=host=localhost port=5432 user=postgres password=password dbname=vue_app_api sslmode=disable timezone=UTC connect_timeout=5
ENV=development
BINARY_NAME=vueapi.exe

## build: Build binary
build:
	@echo Building back end...
	@go build -o ${BINARY_NAME} ./cmd/api/
	@echo Binary built!

## run: builds and runs the application
run: stop build
	@echo Starting back end...
	cmd /C "set DSN=${DSN} && set ENV=${ENV} && ${BINARY_NAME}"
	@echo Back end started!

## clean: runs go clean and deletes binaries
clean:
	@echo Cleaning...
	@go clean
	@del ${BINARY_NAME} >nul 2>&1
	@echo Cleaned!

## start: alias to run
start: run

## stop: stops the running application
stop:
	@echo Stopping back end...
	-@taskkill /IM ${BINARY_NAME} /F >nul 2>&1 || echo No running process found
	@echo Stopped back end!

## restart: stops and starts the application
restart: stop run