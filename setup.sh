#!/bin/bash

curl -sLO https://go.dev/dl/go1.18.4.linux-amd64.tar.gz
sudo rm -rf /usr/local/go && sudo tar -xzf go1.18.4.linux-amd64.tar.gz -C /usr/local
echo "export PATH=$PATH:/usr/local/go/bin" >> /home/vagrant/.bashrc
go get github.com/rabbitmq/amqp091-go
