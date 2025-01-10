# Excel Protector

A Go program to **read an existing Excel file**, protect the workbook and sheets with a password, and save the output as a protected Excel file.  
Uses the [`excelize`](https://github.com/xuri/excelize) library for Excel manipulation.

---

## 📦 Features
- ✅ Protects the **workbook** with a password.  
- ✅ Protects **all sheets** with a password.  
- ✅ Prevents unauthorized editing, formatting, and modifications.  
- ✅ Cross-compilation support for **multiple platforms**.  

---

## 🚀 Requirements
- Go 1.20+  
- Excelize library (`go get github.com/xuri/excelize/v2`)  

---

## 📥 Installation
Clone this repository and install the dependencies:
```bash
git clone <repository-url>
cd excel-protector
go mod tidy
```

---

## 🧑‍💻 Usage
### **Run directly using Go:**
```bash
go run main.go <input.xlsx> <output.xlsx> <password>
```

**Example:**
```bash
go run main.go input.xlsx protected_output.xlsx mypassword123
```

---

### **Build and Run Executable:**
```bash
go build -o excel-protector main.go
./excel-protector input.xlsx protected_output.xlsx mypassword123
```

---

## 🔧 Cross-Compilation (Optional)
You can use the provided **`build.sh`** script to cross-compile for different platforms.

### **Usage:**
```bash
./build.sh <platform> <architecture>
```

**Examples:**
```bash
./build.sh linux amd64
./build.sh windows amd64
./build.sh linux arm
```

---

## 📂 Project Structure
```plaintext
.
├── main.go               # Main source code
├── build.sh              # Cross-compilation script
├── README.md             # Project documentation
├── go.mod                # Go module dependencies
├── go.sum                # Go checksum file
├── build/                # Output directory for binaries
```

---

## ✅ License
This project is licensed under the MIT License.

---

## 🎯 Author
- **Your Name**  
- 📧 Contact: [your.email@example.com](mailto:your.email@example.com)  
- 🌐 GitHub: [your-github-profile](https://github.com/your-github-profile)