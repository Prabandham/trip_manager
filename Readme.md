DEVELOPMENT DEPENDANCIES
------------------------
    `mysql`
    `go (> 1.6)`
    `go get -u (from root to install all dependancies)`

Folder Structure
------------------------

  `Readme.md` 
    // This file  (Documentation)

  `main.go`
    // The main app file (This also is responsible for loading routes and setting up logger)

  `endpoints/`
    // This is the API endpoints.

  `models/`
    // This folder is going to serve as the model for the backend.
    // Also contains the schema for the respective models.

  `db/`
    // The place where DB connection is defined. 


API Documentation
--------------------------------------------------------------------------------------------------------------------------------------------------------------------
    METHOD            IP                           REQUEST PARAMS                    RESPONSE
--------------------------------------------------------------------------------------------------------------------------------------------------------------------
    GET               127.0.0.1:3000/ping          -                                 "Pong !!" (text/html)
    -------------------------------------------------------------------------------------------------------------------------------
    POST              127.0.0.1:3000/register      user_name,phone_number (string)   ->
                                                                                        (success)
                                                                                        200,
                                                                                        application/JSON
                                                                                        { message: "User Saved", code: "XXXXX" }
                                                                                        { message: "User exists", code: "XXXXX" }
                                                                                        (error)
                                                                                        400,
                                                                                        application/JSON
                                                                                        { message: "Bad Request" }
    -------------------------------------------------------------------------------------------------------------------------------
    POST              127.0.0.1:3000/confirm       user_name,phone_number,           -> 
                                                   confirmation_code                   (success)
                                                                                       200,
                                                                                       application/JSON
                                                                                       { message: "Verified user" }
    
