if [ -s ./MaJaJalist ]
  then
    cd MaJaJalist
  else
    mkdir MaJaJalist && cd MaJaJalist
  fi
if [ -x "$(command -v docker)" ]
  then
    echo "docker installed"
  else
    sudo apt-get update && sudo apt-get install -y docker
  fi
if [ -x "$(command -v docker-compose)" ]
  then
    echo "docker-compose installed"
  else
    sudo apt-get update && sudo apt-get install -y docker-compose
  fi
wget https://raw.githubusercontent.com/NPC-GO/MaJaJalist-backend/master/docker/docker-compose.yml -O docker-compose.yml
sudo docker-compose pull && sudo docker-compose up -d && sudo docker container prune -f && sudo docker image prune -f