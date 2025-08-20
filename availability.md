# Availability assessment
  
Two key questions can answer at any moment:

- Is the equipment the customer wants to hire available in a certain timeframe?
- Where in the planning are equipment shortages (more planned than the stock)?

## Database

Together with this assessment you get a database with the current equipment planning of a company. The database consists of two tables:

```
equipment: This table contains the stock items the company has. Each record represents a type of equipment.
- id         (int)      Primary key
- name       (char)     The name of the equipment
- stock      (int)      How much the company has of a certain equipment

planning: This table contains assignments for projects. During these records equipment is not available for other entries.
- id         (int)      Primary key
- equipment  (int)      Refers to the equipment table
- quantity   (int)      How many items are planned in this timeframe
- start      (time)     When the equipment goes out
- end        (time)     When does the equipment come back
 ```
## Assessment

The goal of this assessment is to write Go code that is able to do answer the two questions from the introduction. 
These questions are formulated as two methods in the attached planning_service file. The goal of this 
assessment is to implement those methods:

`func (s *PlanningService) IsAvailable(equipment, quantity int64, startAt, endAt time.Time) (bool, error)` 
This method should check if the `quantity` asked for is available in the timeframe passed or not.

`func (s *PlanningService) GetShortages(startAt, endAt time.Time) (map[int64]int64, error)`
This method should find all shortages in timeframe `start, end`. An item is short if the number of items planned at
the same moment exceeds the stock (stock field in equipment table). The shortage in a given time timeframe for one 
equipment item is defined as stock minus the maximum concurrent planned items in that timeframe.

Also note you need implement your own queries in the `src/data/postgres/equipment_client.go` & `src/data/postgres/planning_client.go` files.
### Example

In the example (not in the provided database) below there are 4 planning entries, all for the same equipment item. We assume the equipment has a
stock of 9. That means, at most 9 can be planned at the same time without having shortages.

```
Equipment
id  | stock | name  
----+-------+---------
100 | 9     | Speaker
```

```
Planning
id         | equipment | start | end | quantity
-----------+-----------|-------+-----+-----------
Planning 1 | 100       | 1     | 5   | 4
Planning 2 | 100       | 3     | 7   | 5
Planning 3 | 100       | 5     | 8   | 3
Planning 4 | 100       | 3     | 9   | 2
```

That results in this timeline:

```
        Day: 0    1    2    3    4    5    6    7    8    9    10 ...
-------------+----+----+----+----+----+----+----+----+----+----+-----
 Planning 1: |    |---------4---------|
 Planning 2: |              |---------5---------|
 Planning 3: |                        |-------3------|
 Planning 4: |              |------------2----------------|
-------------+----+----+----+----+----+----+----+----+----+----+-----
Sum planned: | 0  |    4    |    11   |    10   | 5  | 2  | 0  | 0    
  Available: | 9  |    5    |    -2   |    -1   | 4 | 7  | 9  | 9           
```  

*isAvailable*
The method `isAvailable` returns false when at any moment in timeframe `start, end` it is not possible to plan an
additional `quantity` items without getting shortages. That means at any moment in `start, end` the total available 
items is at least `quantity`.

*getShortages*
This method returns the amount of equipment items that are short within the given timeframe. We assume the equipment 
id is 100 and there are no other equipments in the database (notice in the example how the quantity is always a negative
number, since only shortages are returned). 

Here are some examples:
```
getShortages(0,1): {}          //nothing is planned in this timeframe (day 0 to day 1)
getShortages(2,4): {"100":-2}  //maximum is 11 planned in timeframe day 2 to day 4, while we have only 9
getShortages(7,9): {}          //the planning in this timeframe (day 7 to day 9) does never exceed 9
```

## Deliverables

This package contains a runnable docker environment. To run the code (assume Docker is running):

```
# run the environment
docker-compose up

```

After running these commands, these grpc server will available:

### Todo

1.) Implement the two missing methods in `src/service/planning_service.go`, you can add 
helper methods or make any change you like.
2.) Implement your queries in `src/data/postgres/equipment_client.go` & `src/data/postgres/planning_client.go`.
3.) You can send the whole code base as a zip package. Make sure it runs by just calling `docker-compose up`.

## Remarks

- If anything is unclear, just ask!
- Write clear code: comments, good indentation, clear names of methods, variables, classes 
- Focus on the code that is asked for: implementation of the two missing methods.
- Input validation is not needed since it is only executed in a controlled environment.
- Think about how you do it, considering that a real database would have way more records (efficiency).
- If you want to make changes to the database or see how we can improve things there, just do it and describe it