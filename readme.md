Go Lang Crud Project

Folder structure

    Folder_Name->
                Subfolder_Name


        cmd -> contains server and its confugurations
                server-> file contains port location and corsheader

        internal ->
                    db-> contains the database client connection
                    model-> schema of model which is being used
                    repository-> all Find Save Delete function
                    routes->  all the routes of the api
                    service-> business logic the api

        supertoken-> contains the congifuration file of the supertokens library

        tests -> contains the test files

Steps to Run the projects 1) change the database configurations in .env 2) chagne the supertoken configuration in .env 3) open terminal, run the command ""go run main.go""

API request Flow

    request-> main.go ->cmd/server.go -> internal/routes/routes.go -> internal/service/service.go->
    internal/repository/repository.go -> internal/db/db.go and internal/model/model.go
