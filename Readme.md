This is a trip manager application.

1. This is going be an API only application.

2. It will have a Trip model

    Trip
        name:string
        start_time:date_time
        end_time:date_time
        location:string
        people_ids: [] #list of all people in the trip


3. A trip will have many people. So the People model will look like so.
    Person
        name:string
        phone_number:int
        password:string

4. A trip will have many expenses. 
    Expense
        trip_id:int
        amount:int
        description:string
        label:string #This we will hard code (Travel,Lodging,Dining,Toll,..)

5. On completing a trip (setting a trip as done) It will divide the amount by number of 
   people in the trip and conclude.
