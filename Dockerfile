#################################################
# Define a Builder stage and build app inside it
FROM ubuntu:latest
LABEL maintainer="lzha@redhat.com"
RUN apt-get update && apt-get install -y git && apt-get install -y wget
ENTRYPOINT ["git"]
