# 🛡️ DockShield - Secure Your Docker Images with Confidence

**DockShield** is a lightweight, blazing-fast, cross-platform CLI tool that scans Docker images for security vulnerabilities using [Trivy](https://github.com/aquasecurity/trivy) under the hood. It's designed for developers and DevOps engineers who care about **security-first shipping**.

---

## 🚀 Features

- 🔍 Scan any Docker image for known vulnerabilities (powered by Trivy)
- 🎯 Focused output for **HIGH** and **CRITICAL** severity CVEs
- 📊 Beautifully formatted **table output** for CLI readability
- 📁 Save results as **pretty-printed JSON**
- 💻 Cross-platform support (Linux, macOS, Windows)
- 🔧 Built with Go, Cobra CLI, Logrus Logger, and Tablewriter
- 🌐 Open-source and developer-friendly

---

## 📥 Download & Installation

### 📦 Pre-built Binaries

Head to the [Releases](https://github.com/ShivamJha2436/DockShield/releases) section and download the binary for your platform:

| Platform   | Architecture     | File                         |
|------------|------------------|------------------------------|
| Linux      | x86_64           | `dockshield`                 |
| Windows    | x86_64           | `dockshield.exe`             |
| macOS      | Intel (x86_64)   | `dockshield`                 |
| macOS      | Apple Silicon    | `dockshield-arm`             |

> 📌 Tip: Move the binary to a directory in your `$PATH` for global access.

### 🛠 Make Executable (Linux/macOS)

```bash
chmod +x dockshield
```

---

## 🧪 How to Use

### 🔍 Scan a Docker Image

```bash
./dockshield scan alpine:latest
```

### 📝 Save Output as JSON

```bash
./dockshield scan alpine:latest -o report.json
```

### 📊 Sample Output

```
📦 Image Scan Report

🔍 Target: alpine:3.12

+-----------+----------------+------------+-----------------------------------------+
| SEVERITY  | CVE ID         | PACKAGE    | TITLE                                   |
+-----------+----------------+------------+-----------------------------------------+
| CRITICAL  | CVE-2022-37434 | zlib       | heap-based buffer over-read and overflow|
+-----------+----------------+------------+-----------------------------------------+
```

## 🔄 Contributing

We welcome contributions! If you want to suggest a feature, improve the CLI, or fix bugs:

1. Fork the repo
2. Create a new branch (`git checkout -b feat/your-feature`)
3. Commit your changes (`git commit -m 'feat: add new feature'`)
4. Push to your branch (`git push origin feat/your-feature`)
5. Open a Pull Request

---

## 📌 Roadmap

- [x] Scan Docker images using Trivy
- [x] Output in table and JSON format
- [ ] Auto-update support
- [ ] GitHub Action integration
- [ ] HTML or Markdown report generation
- [ ] Severity threshold flags

---

## 📜 License

This project is licensed under the [MIT License](LICENSE).

---

## ❤️ Acknowledgements

Thanks to the [Aqua Security](https://github.com/aquasecurity) team for [Trivy](https://github.com/aquasecurity/trivy), an amazing open-source vulnerability scanner that powers this tool.

---

> Author [@ShivamJha2436](https://github.com/ShivamJha2436)
