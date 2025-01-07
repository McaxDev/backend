#!/bin/bash
docker build -t mcaxdev/auth . && \
docker-compose -f /srv/axo/auth/docker-compose.yml up
