docker build -t mcaxdev/bbs . &&
docker-compose -f /srv/axo/bbs/docker-compose.yml up -d &&
docker logs -f axo-bbs
