# Tables
A Restful Api interface written in go, displayed by means of HTML tables

## Deploying Requests to Table Interface 
Deploying Requests With Curl: To append data to the above table, it is reccomended

-GET Requests
To recieve data from stored variables divulged in the table, use the following curl command in your terminal  window. 
``curl http://localhost:8000/events``

-POST Requests
To send data to stored variables, and therefore protracting the table, use the following curl command in your terminal window. 
``curl -X POST -H 'Content-type:application/json' -d '{"ID":"IDHERE", "Title":"TITLEHERE", "Description":"DESHERE"}' http://localhost:8000/event``
                     
-PATCH (Update) Requests
To retract data from the table,by means of ID,use the following curl command in your terminal window.
``curl -X PATCH -H 'Content-type:application/json' -d '{"ID":"IDHERE", "Title":"TITLEHERE", "Description":"DESHERE"}' http://127.0.0.1:8000/{id}``

-DELETE Requests
To retract data from the table,by means of ID, use the following curl command in your terminal window. 
``curl -X DELETE -H '{"Content-type":"application/json"} -d '{"ID":"IDHERE", "Title":"TITLEHERE", "Description":"DESHERE"}' http://127.0.0.1:8000/{id}``
