**Register User**
----
 Register user name,phone,role

* **URL**

  /user

* **Method:**

  `POST`
  
*  **URL Params**

   None

* **Data Params**

  `name=[string]`
  `phone=[string]`
  `role=[string]` 

* **Success Response:**

 
    **Result:** `{"result":{"ID":1,"CreatedAt":"2020-04-18T22:07:58+07:00","UpdatedAt":"2020-04-18T22:07:58+07:00","DeletedAt":null,"Name":"name 1","Phone":"11111","Role":"user","Password":"jaw9"}}`
 
* **Error Response:**

-

* **Sample Call:**

-

**Login**
----
 Login User to get token

* **URL**

  /login

* **Method:**

  `POST`
  
*  **URL Params**

   None

* **Data Params**

  `phone=[string]` 
  `password=[string]`
   

* **Success Response:**

    **Message:** `logged in` <br> 
    **token:** `{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJOYW1lIjoibmFtZSAzIiwiUGhvbmUiOiIzMzMzMyIsIlJvbGUiOiJhZG1pbiIsIlBhc3N3b3JkIjoiTXBMbSIsIlRpbWUiOiIyMDIwLTA0LTIwVDExOjM1OjE3LjU0MDEyNzkrMDc6MDAifQ.aUwhkatGYwN7MqD2e4TtzVOlJhjJleLkO0IUMFJs-FI"}`
 
* **Error Response:**

 **Message:** `{"message":"Phone not Found!"}` <br>
 OR 
**Message:** `{"message":"failed log in"}` <br>

* **Sample Call:**

-

**Token Info**
----
 Show token payload

* **URL**

  /tokeninfo

* **Method:**

  `POST`
  
*  **URL Params**

   None

* **Data Params**

  `token=[string]` 
  
   

* **Success Response:**

    **result:** `{"result":{"Name":"name 3","Password":"MpLm","Phone":"33333","Role":"admin","Time":"2020-04-19T15:05:10.6128548+07:00"}}`
 
* **Error Response:**

 **result:** `{"result":"token not valid"}` <br>
 OR 
**result:** `{"result":"invalid jwt token"}` <br>

* **Sample Call:**

-

**Show Storage**
----
Show fetching resources + add USD field

* **URL**

  /showstorage

* **Method:**

  `GET`
  
*  **URL Params**

   None

* **Data Params**

  **middleware**
  `Authorization=[string]` 
  
   

* **Success Response:**

    **data:** `{"data":[{"Uuid":"","Komoditas":"","Area_provinsi":"","Area_kota":"","Size":"","Price":"","Tgl_parsed":"0001-01-01T00:00:00Z","Timestamp":"","Usd":0},{"Uuid":"8a23fcab-ef67-48b8-8ba1-7055ea91ea3b","Komoditas":"Penaeus Vannamei","Area_provinsi":"BANTEN","Area_kota":"PANDEGLANG","Size":"40","Price":"72000","Tgl_parsed":"2019-11-11T17:00:00Z","Timestamp":"1573491600","Usd":4,667}]}`
 
* **Error Response:**

 **message:** `"message": "not authorized",` <br>

* **Sample Call:**

-

**Admin Login**
----
role admin access only endpoint for filter result by area_provisi

* **URL**

  /adminpage

* **Method:**

  `GET`
  
*  **URL Params**



* **Data Params**

  **middleware**
  `Authorization=[string]` 
  
   

* **Success Response:**

    **Claims:**`{"Claims":{"Name":"name 3","Password":"MpLm","Phone":"33333","Role":"admin","Time":"2020-04-20T10:55:12.504988+07:00"}}`<br>

    **data:** `{"data":[{"Uuid":"","Komoditas":"","Area_provinsi":"","Area_kota":"","Size":"","Price":"","Tgl_parsed":"0001-01-01T00:00:00Z","Timestamp":"","Usd":0},{"Uuid":"8a23fcab-ef67-48b8-8ba1-7055ea91ea3b","Komoditas":"Penaeus Vannamei","Area_provinsi":"BANTEN","Area_kota":"PANDEGLANG","Size":"40","Price":"72000","Tgl_parsed":"2019-11-11T17:00:00Z","Timestamp":"1573491600","Usd":4,667}]}`
 
* **Error Response:**

 **message:** `"message": "not authorized",` <br>

* **Sample Call:**

-


