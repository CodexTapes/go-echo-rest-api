# Base Image
FROM mongo:latest

RUN apt-get update -y
RUN apt-get remove mongodb-org -y
RUN apt-get install mongodb -y
RUN mkdir -p ./data/db
RUN service mongodb start -y

# expose MongoDB default port
EXPOSE 27017