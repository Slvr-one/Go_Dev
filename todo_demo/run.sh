#!/bin/bash

docker build -t todo .

docker run -itd --name todo_app -p 9000:9000 todo
# docker run -itd --name todo_app -p 9000:9000 -p 3306:3306 todo
# docker rm -f todo_app