cp Dockerfile .. &&
cd .. &&
docker build -t mcaxdev/games . &&
rm Dockerfile &&
docker-compose -f /srv/axo/games/docker-compose.yml up -d &&
docker logs -f axo-games
