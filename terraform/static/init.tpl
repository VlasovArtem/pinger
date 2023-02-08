#!/bin/bash
echo "Updating yum and installing git and wget"
sudo yum update -y && sudo yum install -y git wget
echo "Downloading and installing Go"
wget https://dl.google.com/go/go1.19.linux-amd64.tar.gz
tar -xzf go1.19.linux-amd64.tar.gz
rm -rf go1.19.linux-amd64.tar.gz
sudo mv go /usr/local
echo "Exporting Go variables"
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH
echo "Cloning and building pinger"
cd /usr/ec2-user
git clone https://github.com/VlasovArtem/pinger
cd pinger
go build
mv pinger /usr/local/bin
rm -rf /usr/ec2-user/pinger