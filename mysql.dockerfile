# Use the official MySQL base image
FROM mysql:latest

# Set environment variables for MySQL
# These values can be overridden in your docker-compose.yml if needed
ENV MYSQL_ROOT_PASSWORD=rootpassword
ENV MYSQL_DATABASE=testdb
ENV MYSQL_USER=testuser
ENV MYSQL_PASSWORD=testpassword

# Optional: Copy initialization scripts to the container
# Files in /docker-entrypoint-initdb.d will be executed during container startup
#COPY ./initdb /docker-entrypoint-initdb.d/

# Expose MySQL port
EXPOSE 3306

# Default command to start MySQL
CMD ["mysqld"]

