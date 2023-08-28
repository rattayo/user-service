#! /bin/bash
docker build . -t rattayo/user-service:1.0.0
docker push rattayo/user-service:1.0.0