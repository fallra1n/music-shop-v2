## We will use docker to run the database
        docker run -it --name some-postgres -e POSTGRES_PASSWORD=pass -e POSTGRES_USER=user -e POSTGRES_DB=db -p 5432:5432 --rm postgres