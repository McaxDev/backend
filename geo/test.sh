docker build -t mcaxdev/geo . &&
docker-compose -f /srv/axo/geo/docker-compose.yml up -d &&
docker logs -f axo-geo
