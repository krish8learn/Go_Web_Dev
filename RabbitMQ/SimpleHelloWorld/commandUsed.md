# Command used #

docker run command -->
sudo docker run -d --hostname my-rabbit --name test-rabbit -p 15672:15672 -p 5672:5672 rabbitmq:3-management
test rabbitmq running or not -->
curl -v 'http://localhost:15672'
close rabbitmq docker -->
sudo docker stop test-rabbit
