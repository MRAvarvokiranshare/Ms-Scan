# Ms-Scan

# 🚀 MrScanr

**MrScanr** یک ابزار شبیه Sherlock است که برای بررسی وجود یک نام کاربری در سایت‌ها و شبکه‌های اجتماعی طراحی شده است.  
ابزاری سبک، سریع و چند زبانه برای موبایل و ترموکس!

---

## 🔍 کاربرد ابزار

- بررسی می‌کند یک **username** در کدام سایت‌ها موجود است و در کدام سایت‌ها نیست.  
- سایت‌های بررسی شده شامل شبکه‌های اجتماعی مهم:
  - Instagram, Twitter, Facebook, YouTube, Telegram, GitHub, Snapchat, Reddit, TikTok, LinkedIn
- مناسب برای OSINT، تحلیل اکانت‌ها و بررسی سریع نام کاربری.

---

## ⚡ ویژگی‌ها

- **چند زبانه:** استفاده از Python, JavaScript و Go  
- **سریع و سبک:** با درخواست‌های HEAD سایت‌ها را بررسی می‌کند  
- **خروجی بصری:** علامت ✅ و ❌ نشان‌دهنده موجود بودن یا نبودن username  
- **قابل اجرا روی موبایل:** با استفاده از Termux  

---

## 🛠️ تکنولوژی‌های استفاده شده

![Python](https://img.shields.io/badge/Python-3776AB?style=for-the-badge&logo=python&logoColor=white)
![JavaScript](https://img.shields.io/badge/JavaScript-F7DF1E?style=for-the-badge&logo=javascript&logoColor=black)
![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)
![Git](https://img.shields.io/badge/Git-F05032?style=for-the-badge&logo=git&logoColor=white)
![Termux](https://img.shields.io/badge/Termux-000000?style=for-the-badge&logo=termux&logoColor=white)

---


## 💻 نحوه اجرا در ترمینال‌ها

<div align="center">

| سیستم عامل | دستور اجرا |
|------------|------------|
| **Linux / macOS** | <pre>cd ~/mrscanr\npython3 scan.py &lt;username&gt;\nnode scan.js &lt;username&gt;\ngo run scan.go &lt;username&gt;</pre> |
| **Termux (Android)** | <pre>cd ~/mrscanr\npkg install python nodejs golang -y\npip install aiohttp rich\npython scan.py &lt;username&gt;\nnode scan.js &lt;username&gt;\ngo run scan.go &lt;username&gt;</pre> |
| **Windows (CMD / Powershell)** | <pre>cd C:\Users\<username>\mrscanr\npython scan.py &lt;username&gt;\nnode scan.js &lt;username&gt;\ngo run scan.go &lt;username&gt;</pre> |

</div>

### 🔹 نکات:
- `<username>` را با نام کاربری که می‌خوای بررسی کنی جایگزین کن.  
- نسخه Python بهتر است **Python 3** باشد (`python3 scan.py`).  
- قبل از اجرای نسخه JS یا Go، Node.js و Go باید نصب شده باشند.  
- در Termux، بهتر است dependencies (`aiohttp`, `rich`) یکبار نصب شوند تا ابزار بدون مشکل اجرا شود.








## 💻 اجرا در ترموکس

```bash
# داخل پوشه پروژه برو
cd ~/mrscanr

# اجرای نسخه Python
python scan.py <username>

# اجرای نسخه JavaScript
node scan.js <username>

# اجرای نسخه Go
go run scan.go <username>


## 💻 نحوه اجرا در ترمینال‌ها

<div align="center">

### 🐧 **Linux**
![Linux](https://img.shields.io/badge/Linux-000000?style=for-the-badge&logo=linux&logoColor=white)
```bash
cd ~/mrscanr
python3 scan.py <username>
node scan.js <username>
go run scan.go <username>



🪟 Windows



cd C:\Users\<username>\mrscanr
python scan.py <username>
node scan.js <username>
go run scan.go <username>



---

🍎 macOS



cd ~/mrscanr
python3 scan.py <username>
node scan.js <username>
go run scan.go <username>

</div>🔹 نکات:

<username> را با نام کاربری که می‌خوای بررسی کنی جایگزین کن.

نسخه Python بهتر است Python 3 باشد.

قبل از اجرای JS یا Go، Node.js و Go نصب شده باشند.

در Termux، dependencies (aiohttp, rich) یکبار نصب شوند.






