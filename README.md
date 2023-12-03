# KaihorBackend

### My First Time using golang ^ _ ^
### Installation need
- golang v 1.16++

### run server ด้วยคำสั่ง
```
go run main.go
```
### JSON Data Spec

- Camp  ข้อมูลของค่าย

```
{
 "campid": string
 "name": string
 "time": string
 "location" : string
 "director" : string
}
```

| ชื่อ field | คำอธิบาย |
| ------ | ------ |
| campid | ครั้งที่ |
| name | ชื่อค่าย |
| time | ช่วงเวลาที่จัดค่าย |
| location | สถานที่ที่จัดค่าย |
| director | ผู้อำนวยการค่าย |

 หากข้อมูลช่องไหนขาดหาย จะคืน string "ไม่ระบุ"

# Service ของ  API


### Method GET
#### รูปเเบบ response ของ api

```
{
 "status": number
 "message": string
 "data": {
     camps: [
        {
         "campid": string
         "name": string
         "time": string
         "location" : string
         "director" : string  
        }
     ]
     error : Error
 }
}
```        
* ``` baseurl/admin/update``` เรียกใช้เมื่อ admin ต้องการ update ข้อมูลจาก google sheet ลง Database
* ``` baseurl/camp```  จะคืนข้อมูล Camp ทั้งหมดที่มีอยู่ใน Database
* ``` baseurl/camp/id/:id```  จะคืนข้อมูล Camp ที่มี id ตรงกับใน params
* ``` baseurl/camp/location/:location```  จะคืนข้อมูล Camp ที่มี location ตรงกับใน params
* ``` baseurl/camp/keyword/:keyword```  จะคืนข้อมูล Camp ที่มี keyword ตรงกับใน params
* ``` baseurl/camp/year/:year```  จะคืนข้อมูล Camp ที่จัดปีตรงกับใน params


  
