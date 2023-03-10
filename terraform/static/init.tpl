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
git clone https://github.com/VlasovArtem/pinger
cd pinger
go build
mv pinger /usr/local/bin/pinger
echo "Installing Nginx"
amazon-linux-extras install nginx1 -y
cp /home/ec2-user/pinger/nginx/pinger.conf /etc/nginx/conf.d/pinger.conf
systemctl enable nginx && systemctl start nginx
echo "Installing Systemctl"
cp /home/ec2-user/pinger/systemd/pinger.service /etc/systemd/system/pinger.service
rm -rf /home/ec2-user/pinger
