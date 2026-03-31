docker exec -it vue-app-api-postgres-1 psql -U postgres -d vue_app_api

docker exec -it <container_name> psql -U <username> -d <database_name>


docker compose up --build --force-recreate