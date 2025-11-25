📌 1. Monolithic Architecture

Monolithic adalah arsitektur di mana seluruh bagian aplikasi berada dalam satu project, satu codebase, dan biasanya satu database. Semua fitur ikut dibuild dan dideploy secara bersamaan.

🔍 Ciri-ciri Monolith

Seluruh kode berada dalam 1 repository / 1 project

Semua fitur saling bergantung

Perubahan kecil → tetap harus redeploy seluruh aplikasi

Cocok untuk aplikasi kecil hingga menengah

Deployment lebih sederhana (cukup sekali deploy)

✅ Kelebihan

Mudah dipahami oleh pemula

Arsitektur simple → cepat dibuat

Debugging lebih mudah (semua di 1 tempat)

❌ Kekurangan

Skalabilitas rendah jika aplikasi semakin besar

Jika 1 modul error → semua aplikasi bisa ikut down

Build & deploy semakin lama seiring kompleksitas meningkat

🔗 2. Microservice Architecture

Microservice adalah arsitektur di mana aplikasi dipecah menjadi beberapa service kecil yang dapat berjalan secara independen dan berkomunikasi melalui API.

🔍 Ciri-ciri Microservice

Setiap service punya project dan repo terpisah

Setiap service bisa memakai database berbeda

Deployment dilakukan per service, tidak saling bergantung

Antar service berkomunikasi via HTTP/REST, gRPC, Messaging, dll

Sangat cocok untuk aplikasi besar yang terus berkembang

✅ Kelebihan

Setiap service dapat dikembangkan tim yang berbeda

Lebih scalable — bisa scale service tertentu saja

Jika satu service error → lainnya tetap berjalan

Teknologi berbeda antar service diperbolehkan

❌ Kekurangan

Arsitektur lebih kompleks

Perlu DevOps matang (CI/CD, container, monitoring)

Debugging & monitoring lebih sulit karena banyak service
