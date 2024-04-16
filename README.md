# GoDocker
Docker go AWS

#go to EC2

#create instance->ubuntu and other details

#launch

#open EC2 in browser->Contains git pre intsalled

#clone this repository


#install docker

sudo snap install docker


#change permission for docker.sock

chmod 777 /var/run/docker.sock


#build image

docker build -t awsdocker .


#spin container

docker run -p 4000:4000 -t awsdocker

#Public IPv4 address
 16.170.237.65 |open address in instance for url and port

#URLS

 http://16.170.237.65:4000/tejas/home

 http://16.170.237.65:4000/tejas/exit

 
