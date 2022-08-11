# **راهنمای راه اندازی دیتابیس و استفاده از migration ها**

## **راه اندازی**

ابتدا PostgreSQL را نصب کنید.

از نصب _pgAdmin_ در مراحل نصب دیتابیس اطمینان حاصل کنید.
شاخه ```bin``` از مسیر نصب ````postgres```` را به متغیر محیطی **```PATH```** اضافه کنید. (مثلا ```C:\Program Files\PostgreSQL\14\bin```)


مسیر فایل اجرایی ابزار _goMigrate_ را نیز به **```PATH```** اضافه کنید.  
به مسیر ```infrastructure\datastore``` رفته و دستور  
_```psql -U postgres -d postgres -a -f .\create_db.sql```_  
را برای تولید دیتابیس و یوزر مورد نظر اجرا کنید.  
در مسیر ```infrastructure\datastore\v1``` دستور زیر را برای _**up**_ کردن    جداول اجرا کنید:

_```migrate -database postgres://postgres:postgres@localhost:5432/darolfonoon?sslmode=disable -path . up```_  

---

## **ساخت و استفاده از migrationها**
برای ایجاد یک تغییر جدید در دیتابیس ابتدا مطمئن شوید که  آخرین تغییرات برنچ main را دریافت کرده اید. سپس در مسیر ```infrastructure\datastore``` برای ساخت فایل های _**up**_ و _**down**_ از دستور زیر استفاده کنید: (عبارت داخل [ ] نام migration شما است.)  
‍‍‍‍‍```migrate create -ext sql -dir . -seq [migration-name]```  
سپس اسکریپت های تغییرات مورد نظر خود را در فایل _**up**_ بنویسید و در فایل _**down**_ عملیاتی که آن تغییرات را بر می گرداند را بنویسید.

مطمئن شوید که قبل از ادغام برنچ خود با برنچ main شماره مایگریشن شما تکراری نباشد زیرا مایگریشن ها به ترتیب از آخرین مایگریشنی که در نوبت قبل اجرا شده بوده ادامه پیدا می کنند. در صورت وجود مایگریشن تکراری، عدد ابتدای مایگریشن خود را به یکی بیشتر از آخرین عدد موجود تغییر دهید.

### **استفاده**
برای اجرای مایگریشن ها و اعمال تغییرات جدید(آپگرید کردن دیتابیس) در مسیر ```infrastructure\datastore\v1``` دستور زیر را اجرا کنید:  
_```migrate -database postgres://postgres:postgres@localhost:5432/darolfonoon?sslmode=disable -path . up```_   
و برای بازگرداندن تغییرات(داونگرید کردن دیتابیس) از دستور زیر استفاده کنید:
_```migrate -database postgres://postgres:postgres@localhost:5432/darolfonoon?sslmode=disable -path . down```_ 