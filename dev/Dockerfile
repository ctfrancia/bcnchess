# Derived from official mysql image (our base image)
FROM mysql:5.7
# Add a database
ENV MYSQL_DATABASE bcnchess
ENV  MYSQL_ROOT_PASSWORD abc123
# Add the content of the sql-scripts/ directory to your image
# All scripts in docker-entrypoint-initdb.d/ are automatically
# executed during container startup
COPY ./db-setup.sql /docker-entrypoint-initdb.d/

