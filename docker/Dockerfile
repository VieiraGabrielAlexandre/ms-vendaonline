# Telling to use Docker's golang ready image
FROM golang

EXPOSE 5566

# Set pwd to the go folder
WORKDIR ${GOPATH}

RUN apt-get update -y && \
	apt-get install -y \
	vim
# Start process
ENTRYPOINT ["main.go"]