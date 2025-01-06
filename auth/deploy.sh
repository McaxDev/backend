#!/bin/bash
docker build -t mcaxdev/auth . && \
docker-compose up -d
