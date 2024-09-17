# Uy vazifasi: Go va sqlc yordamida API integratsiyasiga ega task management system yaratish

## Maqsad
Siz `Go` va `sqlc` yordamida `task management system` yaratishingiz kerak bo'ladi. Tizim vazifalar va ularning mualliflarini boshqaradi, va `RESTful API`-lar orqali `CRUD` operatsiyalarini hamda "`author-task`" tayinlashni amalga oshiradi.

## Talablar
1. **Loyihani Sozlash**: 
    - `sqlc` ni o'rnating va `sqlc.yaml` konfiguratsiya faylini sozlang.

2. **Ma'lumotlar Bazasi Shemasini Belgilash**: 
    - SQL shemasini (`schema.sql`) yarating va quyidagi jadvallarni belgilang:
        - `tasks`: Vazifa ma'lumotlarini saqlaydi.
        - `authors`: Muallif ma'lumotlarini saqlaydi.
        - `task_authors`: Vazifalar va mualliflar o'rtasidagi ko'p-ko'pga munosabatlarni boshqaradi.

3. **SQL So'rovlarini Yozish:**
    - SQL so'rovlari (`queries.sql`) faylini yarating va quyidagilarni yozing:
        - Vazifalarni yaratish, o'qish, yangilash va o'chirish.
        - Mualliflarni yaratish, o'qish va ro'yxatini olish.
        - Mualliflarni vazifalarga tayinlash.

4. **Go Kodini Yaratish:**
    - `SQL` shemasi va so'rovlardan `Go` kodini yaratish uchun `sqlc` ni ishga tushiring

5. **API Endpointlarni Amalga Oshirish:**
    - `POST /tasks`: Yangi vazifani yaratish.
    - `GET /tasks`: Barcha vazifalarni ro'yxatlash.
    - `GET /tasks/{id}`: ID bo'yicha vazifani olish.
    - `PUT /tasks/{id}`: ID bo'yicha vazifani yangilash.
    - `DELETE /tasks/{id}`: ID bo'yicha vazifani o'chirish.
    - `POST /authors`: Yangi muallifni yaratish.
    - `GET /authors`: Barcha mualliflarni ro'yxatlash.
    - `GET /authors/{id}`: ID bo'yicha muallifni olish.
    - `POST /tasks/{task_id}/authors/{author_id}`: Muallifni vazifaga tayinlash.





 















