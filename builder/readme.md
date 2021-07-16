# Deploy code

ทำการ run shellscript  โดยเข้าไปที่โฟเดอร์ builder malar และพิมพ์คำสั่ง

```bash
./builder.sh
```

## builder.sh

เป็นไฟล์ที่รวมสคริปการแปลงไฟล์โปรเจ็คให้เป็น .exe แล้วนำขึ้น cloud

เป็นส่วนที่ประกาศให้ .exe ของเรามีโครงสร้างไปตามที่เราประกาศไว้

```
GOOS=linux
GOARCH=amd64
cd .. && git checkout $env && git pull origin $env && go build .
```

เป็นส่วนที่นำไฟล์ของของเราขึ้นไปไว้บน cloud โดย 
```
scp -r -i big.pem ../api-wecode-supplychain ubuntu@ec2-54-174-95-219.compute-1.amazonaws.com:/home/ubuntu/


scp -r -i <key> <file deploy> <url connect cloud>:<ที่ตั้งที่จะวางไฟล์>
```
## checkout.sh

เป็นคำสั่งให้ทำการ kill process  ที่ run ทิ้งไว้บน cloud โดย run แล้วจะเป็นการ connect cloud + kill process 

pidof api-wecode-supplychain มีไว้ดูเลข precess

awk สร้างข้อมูลในรูปแบบ text

xargs ใช้ในการสร้างคำสั่งใหม่จาก ouput ที่ได้ก่อนหน้า (ก็คือเลข process)



```

ssh -i "big.pem" ubuntu@ec2-54-174-95-219.compute-1.amazonaws.com  "pidof api-wecode-supplychain | awk '{print $1}' | xargs kill "

```

# start service

ให้ทำการ connect cloud และ run file script service.sh ที่อยู่บน cloud  อยู่แล้ว

```
./service.sh
```

