## soul_project_api

# install go 
# install postgres

# DATABASE_SETUP

sudo su postgres
createdb soul_api
sudo -u 'user_name' psql soul_api
psql
ALTER USER 'user_name' WITH PASSWORD 'password'
#update password in db.go
\c soul_api
