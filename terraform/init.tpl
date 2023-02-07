#!/bin/bash
cd /usr/ec2-user
echo "Updating yum and installing git and wget"
sudo yum update -y && sudo yum install -y git wget
echo "Installing Nginx"
sudo amazon-linux-extras install nginx1 -y
echo "Downloading and installing Go"
wget https://dl.google.com/go/go1.19.linux-amd64.tar.gz
tar -xzf go1.19.linux-amd64.tar.gz
sudo mv go /usr/local
echo "Exporting Go variables"
echo "export GOROOT=/usr/local/go" >> .bash_profile
echo "PATH=$GOROOT/bin:$PATH" >> .bash_profile
echo "export TELEGRAM_BOT_ENABLING_TOKEN=${bot_enabling_token}" >> .bash_profile
echo "export TELEGRAM_BOT_TOKEN=${telegram_bot_token}" >> .bash_profile
source .bash_profile
#echo "Start nginx"
#sudo systemctl enable nginx && sudo systemctl start nginx
#echo "Cloning and building pinger"
#git clone https://github.com/VlasovArtem/pinger
#cd pinger
#go build