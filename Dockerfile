FROM scratch

#COPY ./.env ./
ENV DRIVER_NAME="postgres"
ENV DATASOURCE_URL="postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
ENV PORT="8080"

# Copy the binary file
COPY ./manager ./

EXPOSE 8080

#service can be run with following command
#docker run -p 8080:8080 kaktas/task-manager ./manager