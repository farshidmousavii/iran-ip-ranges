# iran-ip-ranges — Iran IPv4/IPv6 Prefix Lists & Fetcher

**فارسی** | [English](#english)

---

## فارسی

### این پروژه چیه؟

ترافیک ایرانی را به‌طور خودکار از مسیر VPN یا پروکسی یا روتر  با استفاده از لیست‌های به‌روز IPv4/IPv6 ایران خارج کنید .

مناسب برای سرورهای VPN، روترها، فایروال‌ها و سرورهای پروکسی.

این پروژه هر ۶ ساعت Prefixهای اعلام شده ایران را از RIPE Stat دریافت، merge و normalize می‌کند و خروجی‌های آماده استفاده در قالب‌های مختلف تولید می‌کند.

> این پروژه صرفاً بر اساس داده‌های RIPE Stat کار می‌کند و دقت کامل Geolocation یا Routing را تضمین نمی‌کند.

---

## دانلود مستقیم فایل‌ها

تمامی فایل‌ها در پوشه `dist/` در دسترس هستند:

- **Plain text:** `dist/raw/ipv4.txt`, `dist/raw/ipv6.txt`
- **Clash / Mihomo:** `dist/clash/iran.yaml`
- **Sing-box:** `dist/sing-box/iran.json`
- **Xray:** `dist/xray/iran.json`
- **NFTables ipset:** `dist/firewall/iran.ipset`
- **NFTables config:** `dist/firewall/iran.nft`
- **OpenWRT:** `dist/openwrt/iran.sh`
- **MikroTik RouterOS:** `dist/routeros/ipv4.rsc`, `dist/routeros/ipv6.rsc`

> بسته کامل همه فایل‌ها به‌صورت روزانه در **GitHub Releases** منتشر می‌شود.
> [دانلود آخرین Release](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest)

### لینک‌های دانلود مستقیم

| فایل | لینک |
|------|------|
| `raw/ipv4.txt` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv4.txt) |
| `raw/ipv6.txt` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv6.txt) |
| `routeros/ipv4.rsc` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv4.rsc) |
| `routeros/ipv6.rsc` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv6.rsc) |
| `clash/iran.yaml` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/clash/iran.yaml) |
| `sing-box/iran.json` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/sing-box/iran.json) |
| `xray/iran.json` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/xray/iran.json) |
| `firewall/iran.ipset` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.ipset) |
| `firewall/iran.nft` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.nft) |
| `openwrt/iran.sh` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/openwrt/iran.sh) |
| `checksums.txt` | [Download](https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/checksums.txt) |
| `iran-ip-ranges.zip` | [دانلود از Releases](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest) |

### دانلود باینری (بدون نیاز به Go)

باینری‌های آماده برای سیستم‌عامل‌های مختلف در هر Release منتشر می‌شوند:

| سیستم‌عامل | دانلود |
|-----------|--------|
| Linux | [`iran-ip-linux-amd64`](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest/download/iran-ip-linux-amd64) |
| macOS (Apple Silicon) | [`iran-ip-darwin-arm64`](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest/download/iran-ip-darwin-arm64) |
| Windows | [`iran-ip-windows-amd64.exe`](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest/download/iran-ip-windows-amd64.exe) |

```bash
# لینوکس / مک
chmod +x iran-ip-linux-amd64
./iran-ip-linux-amd64

# ویندوز ( PowerShell )
.\iran-ip-windows-amd64.exe
```

> نیازی به نصب Go یا Docker نیست. فقط باینری را دانلود و اجرا کنید.

---

## Quick Start

### دانلود مستقیم با curl

```bash
# IPv4
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv4.txt

# IPv6
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv6.txt

# Clash / Mihomo
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/clash/iran.yaml

# Sing-box
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/sing-box/iran.json

# Xray
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/xray/iran.json

# NFTables ipset
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.ipset

# NFTables config
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.nft

# OpenWRT script
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/openwrt/iran.sh

# MikroTik
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv4.rsc
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv6.rsc
```

### راهنمای تنظیمات آماده

فایل‌های مثال آماده در پوشه `examples/` قرار دارند — هر کدام شامل تنظیمات کامل برای یک ابزار خاص:

| ابزار | مسیر مثال |
|------|-----------|
| Clash / Mihomo | [`examples/clash/config.yaml`](examples/clash/config.yaml) |
| Sing-box | [`examples/sing-box/config.json`](examples/sing-box/config.json) |
| Xray | [`examples/xray/config.json`](examples/xray/config.json) |
| NFTables + ipset | [`examples/nftables/rules.nft`](examples/nftables/rules.nft) |
| OpenWRT (firewall) | [`examples/openwrt/firewall-config.sh`](examples/openwrt/firewall-config.sh) |
| MikroTik RouterOS | [`examples/mikrotik/import-script.rsc`](examples/mikrotik/import-script.rsc) |
| Split Tunnel (iptables) | [`examples/split-tunnel.sh`](examples/split-tunnel.sh) |

---

## دو روش استفاده

### ۱. دانلود مستقیم از GitHub

به وسیله GitHub Actions هر ۶ ساعت فایل‌ها به‌روزرسانی می‌شوند.

**سه روش دانلود:**
- **فایل‌های تکی** — هر فایل به‌صورت جداگانه از پوشه `dist/` در دسترس است
- **بسته کامل (zip)** — همه فایل‌ها به‌صورت یکجا در [GitHub Releases](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest) هر روز منتشر می‌شوند
- **باینری آماده** — فایل اجرایی مخصوص سیستم‌عامل خود را از Releases دانلود کنید (بدون نیاز به Go)

### ۲. اجرای وب سرور (Self-hosted)

وب سرور را روی سرور خود اجرا کنید. دو روش:

**روش الف — دانلود باینری (ساده‌تر):**

```bash
# لینوکس amd64
wget https://github.com/farshidmousavii/iran-ip-ranges/releases/latest/download/iran-ip-linux-amd64
chmod +x iran-ip-linux-amd64
./iran-ip-linux-amd64
```

**روش ب — داکر:**

از داکر هم می‌توانید استفاده کنید (دستورات در بخش Docker پایین).

ویژگی‌های وب سرور:

- دریافت خودکار Prefixها در startup
- ارائه فایل‌ها از طریق HTTP
- به‌روزرسانی دوره‌ای در پس‌زمینه
- استفاده از cache روی دیسک هنگام قطعی اینترنت
- استفاده از Health Check داخلی
- دارای Graceful Shutdown

---

## نصب

### گزینه ۱ — باینری (ساده‌ترین روش)

```bash
# لینوکس amd64
wget https://github.com/farshidmousavii/iran-ip-ranges/releases/latest/download/iran-ip-linux-amd64
chmod +x iran-ip-linux-amd64
./iran-ip-linux-amd64
```

### گزینه ۲ — با Go

```bash
git clone https://github.com/farshidmousavii/iran-ip-ranges.git
cd iran-ip-ranges

go run ./cmd/
```

---

## Docker

### Docker Compose (پیشنهادی)

```bash
docker compose up -d
```

فایل‌های تولید شده در پوشه `dist/` در دسترس خواهند بود.

### Docker Manual

```bash
docker build -t iran-ip-ranges .

docker run -d \
  --name iran-ip-ranges \
  -p 8080:8080 \
  iran-ip-ranges
```

کانتینر با کاربر non-root اجرا می‌شود و دارای HEALTHCHECK داخلی است.

---

## Flags

| Flag          | Default | Description                           |
| ------------- | ------- | ------------------------------------- |
| `-addr`       | `:8080` | آدرس وب سرور                          |
| `-refresh`    | `6h`    | فاصله به‌روزرسانی خودکار              |
| `-fetch-only` | `false` | فقط دریافت فایل‌ها بدون اجرای وب سرور |

---

## نحوه اجرا

### حالت پیش‌فرض

```bash
go run ./cmd/
```

### پورت دلخواه

```bash
go run ./cmd/ -addr :9090
```

### به‌روزرسانی هر ۱ ساعت

```bash
go run ./cmd/ -refresh 1h
```

### فقط دریافت فایل‌ها

```bash
go run ./cmd/ -fetch-only
```

---

## Web Endpoints

| Endpoint                    | Description                     |
| --------------------------- | ------------------------------- |
| `GET /health`               | وضعیت سلامت سرویس               |
| `GET /ipv4.txt`             | نمایش لیست IPv4 (backward)      |
| `GET /ipv6.txt`             | نمایش لیست IPv6 (backward)      |
| `GET /ipv4.rsc`             | دانلود RouterOS IPv4 (backward) |
| `GET /ipv6.rsc`             | دانلود RouterOS IPv6 (backward) |
| `GET /routeros/ipv4.rsc`    | دانلود اسکریپت RouterOS IPv4    |
| `GET /routeros/ipv6.rsc`    | دانلود اسکریپت RouterOS IPv6    |
| `GET /clash/iran.yaml`      | دانلود rule-provider Clash      |
| `GET /sing-box/iran.json`   | دانلود rule-set Sing-box        |
| `GET /xray/iran.json`       | دانلود routing rules Xray       |
| `GET /firewall/iran.ipset`  | دانلود اسکریپت ipset restore    |
| `GET /firewall/iran.nft`    | دانلود کانفیگ nftables          |
| `GET /openwrt/iran.sh`      | دانلود اسکریپت OpenWRT          |
| `GET /raw/iran.json`        | دانلود JSON خام                 |
| `GET /raw/iran.yaml`        | دانلود YAML خام                 |
| `GET /checksums.txt`        | دانلود SHA256 checksums         |

تمام endpointها دارای:

```text
Cache-Control: public, max-age=21600
```

هستند.

---

## Health Check

اندپوینت `/health` خروجی JSON برمی‌گرداند:

### در حال initialization

```json
{ "status": "initializing" }
```

HTTP Status:

```text
503
```

### خطا در آخرین دریافت

```json
{ "status": "stale", "last_fetch": "...", "last_error": "..." }
```

HTTP Status:

```text
503
```

### وضعیت سالم

```json
{ "status": "ok", "last_fetch": "..." }
```

HTTP Status:

```text
200
```

---

## منبع داده ها

داده ها از منابع زیر دریافت میگردند :

- RIPE Stat API
- Country Resource List (IR)

---

## مجوز

MIT

---

# English

## What is this?

Automatically route Iranian traffic outside your VPN or proxy using continuously updated Iran CIDR ranges.

Designed for VPN servers, routers, firewalls, and proxy servers.

This project fetches announced IP prefixes for Iran from RIPE Stat every 6 hours, merges and normalizes them, and generates ready-to-use output files in multiple formats.

> This project relies on RIPE Stat data and does not guarantee perfect geolocation or routing accuracy.

---

## Direct Downloads

All output files are available under the `dist/` directory.

> A complete archive of all files is published daily as **GitHub Releases**.
> [Download latest release](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest).

### Pre-built Binaries (no Go required)

Ready-to-run binaries for every platform are published in each release:

| Platform | Download |
|----------|----------|
| Linux | [`iran-ip-linux-amd64`](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest/download/iran-ip-linux-amd64) |
| macOS (Apple Silicon) | [`iran-ip-darwin-arm64`](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest/download/iran-ip-darwin-arm64) |
| Windows | [`iran-ip-windows-amd64.exe`](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest/download/iran-ip-windows-amd64.exe) |

```bash
# Linux / macOS
chmod +x iran-ip-linux-amd64
./iran-ip-linux-amd64

# Windows (PowerShell)
.\iran-ip-windows-amd64.exe
```

> No Go or Docker installation required. Just download and run.

### Clash / Mihomo

- `clash/iran.yaml`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/clash/iran.yaml`

### Sing-box

- `sing-box/iran.json`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/sing-box/iran.json`

### Xray

- `xray/iran.json`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/xray/iran.json`

### NFTables

- `firewall/iran.ipset`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.ipset`

- `firewall/iran.nft`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.nft`

### OpenWRT

- `openwrt/iran.sh`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/openwrt/iran.sh`

### MikroTik RouterOS

- `routeros/ipv4.rsc`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv4.rsc`

- `routeros/ipv6.rsc`
- `https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv6.rsc`

---

## Quick Start

### Download with curl

```bash
# IPv4
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv4.txt

# IPv6
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/raw/ipv6.txt

# Clash / Mihomo
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/clash/iran.yaml

# Sing-box
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/sing-box/iran.json

# Xray
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/xray/iran.json

# NFTables ipset
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.ipset

# NFTables config
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/firewall/iran.nft

# OpenWRT script
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/openwrt/iran.sh

# MikroTik
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv4.rsc
curl -O https://raw.githubusercontent.com/farshidmousavii/iran-ip-ranges/main/dist/routeros/ipv6.rsc
```

### Example Configurations

Ready-to-use example configs are available in the `examples/` directory:

| Tool | Example Path |
|------|-------------|
| Clash / Mihomo | [`examples/clash/config.yaml`](examples/clash/config.yaml) |
| Sing-box | [`examples/sing-box/config.json`](examples/sing-box/config.json) |
| Xray | [`examples/xray/config.json`](examples/xray/config.json) |
| NFTables + ipset | [`examples/nftables/rules.nft`](examples/nftables/rules.nft) |
| OpenWRT (firewall) | [`examples/openwrt/firewall-config.sh`](examples/openwrt/firewall-config.sh) |
| MikroTik RouterOS | [`examples/mikrotik/import-script.rsc`](examples/mikrotik/import-script.rsc) |
| Split Tunnel (iptables) | [`examples/split-tunnel.sh`](examples/split-tunnel.sh) |

---

## Two Usage Modes

### 1. Download from GitHub

GitHub Actions automatically refreshes the files every 6 hours.

**Three download options:**
- **Individual files** — each file available directly from the `dist/` folder
- **Full archive (zip)** — all files bundled in daily [GitHub Releases](https://github.com/farshidmousavii/iran-ip-ranges/releases/latest)
- **Pre-built binary** — download the executable for your platform from Releases (no Go required)

### 2. Self-hosted Web Server

Run the web server on your own machine. Two methods:

**Option A — Download binary (simpler):**

```bash
# Linux amd64
wget https://github.com/farshidmousavii/iran-ip-ranges/releases/latest/download/iran-ip-linux-amd64
chmod +x iran-ip-linux-amd64
./iran-ip-linux-amd64
```

**Option B — Docker:**

You can also use Docker (see Docker section below).

Features:

- Fetches prefixes on startup
- Serves files via HTTP
- Background auto-refresh
- Disk cache fallback
- Built-in health checks
- Graceful shutdown

---

## Installation

### Option 1 — Binary (easiest)

```bash
# Linux amd64
wget https://github.com/farshidmousavii/iran-ip-ranges/releases/latest/download/iran-ip-linux-amd64
chmod +x iran-ip-linux-amd64
./iran-ip-linux-amd64
```

### Option 2 — With Go

```bash
git clone https://github.com/farshidmousavii/iran-ip-ranges.git
cd iran-ip-ranges

go run ./cmd/
```

---

## Docker

### Docker Compose (recommended)

```bash
docker compose up -d
```

Generated files will be available in the `dist/` directory.

### Docker

```bash
docker build -t iran-ip-ranges .

docker run -d \
  --name iran-ip-ranges \
  -p 8080:8080 \
  iran-ip-ranges
```

Container runs as non-root user and includes a built-in health check.

---

## Flags

| Flag          | Default | Description               |
| ------------- | ------- | ------------------------- |
| `-addr`       | `:8080` | Web server listen address |
| `-refresh`    | `6h`    | Auto-refresh interval     |
| `-fetch-only` | `false` | Fetch files and exit      |

---

## Usage

### Default Mode

```bash
go run ./cmd/
```

### Custom Listen Address

```bash
go run ./cmd/ -addr :9090
```

### Refresh Every Hour

```bash
go run ./cmd/ -refresh 1h
```

### Fetch-only Mode

```bash
go run ./cmd/ -fetch-only
```

---

## Web Endpoints

| Endpoint                    | Description                       |
| --------------------------- | --------------------------------- |
| `GET /health`               | Service health endpoint           |
| `GET /ipv4.txt`             | View IPv4 list (backward compat)  |
| `GET /ipv6.txt`             | View IPv6 list (backward compat)  |
| `GET /ipv4.rsc`             | Download RouterOS (backward)      |
| `GET /ipv6.rsc`             | Download RouterOS (backward)      |
| `GET /routeros/ipv4.rsc`    | Download RouterOS IPv4 script     |
| `GET /routeros/ipv6.rsc`    | Download RouterOS IPv6 script     |
| `GET /clash/iran.yaml`      | Download Clash rule-provider      |
| `GET /sing-box/iran.json`   | Download Sing-box rule-set        |
| `GET /xray/iran.json`       | Download Xray routing rules       |
| `GET /firewall/iran.ipset`  | Download ipset restore script     |
| `GET /firewall/iran.nft`    | Download nftables config          |
| `GET /openwrt/iran.sh`      | Download OpenWRT script           |
| `GET /raw/iran.json`        | Download raw JSON list            |
| `GET /raw/iran.yaml`        | Download raw YAML list            |
| `GET /checksums.txt`        | Download SHA256 checksums         |

All file endpoints include:

```text
Cache-Control: public, max-age=21600
```

---

## Health Check

The `/health` endpoint returns JSON:

### Initializing

```json
{ "status": "initializing" }
```

HTTP Status:

```text
503
```

### Last Fetch Failed

```json
{ "status": "stale", "last_fetch": "...", "last_error": "..." }
```

HTTP Status:

```text
503
```

### Healthy

```json
{ "status": "ok", "last_fetch": "..." }
```

HTTP Status:

```text
200
```

---

## Data Source

Data is fetched from:

- RIPE Stat API
- Country Resource List (IR)

---

## License

MIT
