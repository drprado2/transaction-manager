# Transaction Manager

Project made to study some concepts in the go language

## Project Structure

The structure of the project was inspired by the ideas demonstrated here https://github.com/golang-standards/project-layout

## Migrations

The project uses migrations to version the database and automate the application of scripts https://github.com/golang-migrate/migrate

## Configurations

In the development environment you can change the settings using the **configs/config.json** file,   
the application replaces these values ​​with environment variables if there is a variable configured with the same name. 

## Documentation

The API is documented with Open API Specification using Swagger, you can access it through [http://localhost:9000/swagger/index.html](http://localhost:9000/swagger/index.html)

## Running
To execute the entire project in the most practical way it is necessary to have Docker and Docker Compose installed and execute the following commands:

    git clone https://github.com/drprado2/transaction-manager.git
    cd transaction-manager/build
    docker-compose up -d
  
After the command is executed you will have Postgres running on port 5432 and the application running on port 9000, then you can access http://localhost:9000/swagger/index.html
  
  If you prefer to run directly in your IDE or CLI the main method is in the **cmd/transaction-manager/transaction-manager.go** file
