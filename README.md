# 🧠 Go Reloaded

**Go Reloaded** is a text transformation and auto-correction tool written in **Go**.  
It reads a text file, applies various formatting and editing rules, and outputs a corrected version.  
The goal of this project is to practice file handling, string manipulation, and text processing using only the Go standard library.

---

## 🚀 Features

- 🔢 **Number Conversion**  
  - Converts hexadecimal numbers followed by `(hex)` to their decimal equivalent.  
    - Example: `1E (hex)` → `30`  
  - Converts binary numbers followed by `(bin)` to their decimal equivalent.  
    - Example: `10 (bin)` → `2`

- 🔠 **Case Formatting**  
  - `(up)` → Converts the previous word to uppercase.  
  - `(low)` → Converts the previous word to lowercase.  
  - `(cap)` → Capitalizes the previous word.  
  - Supports numbered formats like `(up, 2)` or `(low, 3)` to apply changes to multiple previous words.

- ✍️ **Punctuation Correction**  
  - Ensures `. , ! ? : ;` are correctly placed next to words with proper spacing.  
  - Maintains grouped punctuations such as `...` or `!?` properly.  

- 🗣️ **Quotation Formatting**  
  - Fixes spacing around single quotes `' '` ensuring they wrap around words properly.  
  - Example:  
    `"I am ' awesome '"` → `"I am 'awesome'"`

- 📖 **Grammar Fix**  
  - Replaces `a` with `an` when followed by a word starting with a vowel or `h`.  
    - Example: `a amazing story` → `an amazing story`

---

## 📂 Project Structure

```bash
go-reloaded/
│
├── main.go # Entry point of the program
├── ressources/ # Directory for helper and utility files
│ ├── atoan.go # Functions for changing a to an if a vowel is next
│ └── capitalize.go # Functions for capitalizing or changing case of words
│ └── convert.go # Functions for converting numbers (hex, bin) to decimal
│ └── punctuations.go # Functions for correcting punctuation spacing and placement
│ └── quots.go # Functions for handling formatting of quoted text
├── sample.txt # Example input file
└── result.txt # Example output file
```

## ⚙️ Usage

1. **Use the already given sample.txt file or prepare an input file**
Example:
```txt
    it (cap) was the best of times, it was the worst of times (up), it was the age of wisdom.
```

2. **Run the program**
```bash
go run . sample.txt result.txt
```

3. **Check the output**
```bash
cat result.txt
```
Example output:
```txt
It was the best of times, it was the worst of TIMES, it was the age of wisdom.
```

## 💡 Example Transformations

```markdown
| Input                      | Output                      |
| -------------------------- | --------------------------- |
| 1E (hex)                   | 30                          |
| 10 (bin)                   | 2                           |
| I was shouting (low)       | I was shouting              |
| this is so cool (up, 3)    | THIS IS SO COOL             |
| There was a amazing view   | There was an amazing view   |
| I am ' awesome '           | I am 'awesome'              |
```

## 🧰 Technologies Used

- Language: Go (Golang)
- Packages: Standard library only (os, fmt, strings, etc.)

## Licensing

This project is open-source and available under the [MIT License](LICENSE).

## 👤 Author
**Abderrahmane Fethi**

Junior Full-Stack Developer | Passionate about clean code and problem-solving

🌍 [LinkedIn](https://www.linkedin.com/in/abderrahmane-fethi)