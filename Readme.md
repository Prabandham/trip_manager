This is a trip manager application.

1. This is going be an API only application.

2. It will have a Trip model
    Trip
        name:string
        start_time:date_time
        end_time:date_time
        location:string

3. A trip will have many people. So the People model will look like so.
    Person
        name:string
        phone_number:int
        password:string

4. The link between people and Trips will be in a join table called trip_people
    TripPeople
        trip_id:int
        person_id:int

5. A trip will have many expenses. 
    Expense
        trip_id:int
        amount:int
        description:string
        label:string #This we will hard code (Travel,Lodging,Dining,Toll,..)

6. A trip will have many Moments (Images and Videos and audio files)
    Moment
        trip_id int
        person_id int
        file_type string
        file_size int
        file_location int

7. On completing a trip (setting a trip as done) It will divide the amount by number of 
   people in the trip and conclude.



Development dependancies
    mysql
    go (> 1.6)
    go get -u (from root to install all dependancies)

Folder Structure

  `Readme.md` // This file
  `main.go`   // The main app file (This also is responsible for loading routes and setting up logger)
  `endpoints/`
    // This is the API endpoints.
  `models/`
    //This folder is going to serve as the model for the backend.
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
    
