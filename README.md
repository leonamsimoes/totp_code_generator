# OTP Generator CLI 🛠️🔐

A simple Command Line Interface (CLI) tool to generate One-Time Password (OTP) codes based on the TOTP (Time-based One-Time Password) algorithm. This project is written in Go (Golang) and allows you to securely generate OTP codes from a shared secret.


#### Note 📋: _I realize this project might seem over-engineered, but I'm using it as a sandbox to experiment with testing methodologies, which will ultimately be a valuable addition to my portfolio._.


_(TOTP and other features coming soon)_ 🚀


## Features ✨

- Generate OTP codes using a secret key
- TOTP standard implementation
- Lightweight, fast, and easy to use
- Built with Go, and uses Makefile for easy build management

## Requirements 📋

- Go 1.18 or higher
- Git
- Make (for building the project)
  
## Installation 🔨

1. **Using go install**:
    ```bash
    go install github.com/leonamsimoes/totp_code_generator@latest
    ```

2. **Clone the repository**:
    ```bash
    git clone https://github.com/leonamsimoes/totp_code_generator.git
    cd totp_code_generator
    ```

    2a. **Build the project using Make**:
    ```bash
    make build
    ```

    2b. **Run the executable**:
    After building, you will have the `totp_code_generator` binary ready to use.

## Usage 📖

To generate an OTP code, you will need to provide a secret key (in base32 encoding) as input. Here's how to use the tool:

### Basic Command
```bash
./totp_code_generator --secret="SECRETVALUEHERE" --account test@gmail.com --issuer outlook --duration 60 --length 10
```

### Help argument
```bash
./totp_code_generator --help
```

### Secret argument
```bash
./totp_code_generator --secret SECRETVALUEHERE
```

### Account argument
```bash
./totp_code_generator --account test@gmail.com
```

### Issuer argument
```bash
./totp_code_generator --issuer outlook
```

### Duration argument
```bash
./totp_code_generator --duration 60
```

### Length argument
```bash
./totp_code_generator --length 10
```

### Example:
```bash
============================================
============================================
BEGIN
Code: 1617198347
END
============================================
============================================
```

This will generate the OTP code for the provided secret using the current time-based window.

### Options 🛠️

- `--secret` (required): The base32-encoded secret key used to generate the OTP.
- `--window` (optional): The time window (in seconds) to adjust the OTP's validity. Default is 30 seconds.

## Example Output ⚡

```bash
OTP: 123456
```

The OTP code will be displayed on the terminal.

## Development 🚀

1. **Clone the repository**:
    ```bash
    git clone https://github.com/leonamsimoes/otp-generator-cli.git
    cd otp-generator-cli
    ```

2. **Install dependencies** (if any):
    ```bash
    make install
    ```

3. **Run tests**:
    ```bash
    make test
    ```

4. **Build the project**:
    ```bash
    make build
    ```

## Contributing 🤝

If you'd like to contribute to this project, feel free to fork the repository, create a feature branch, and submit a pull request.

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature-name`).
3. Commit your changes (`git commit -am 'Add new feature'`).
4. Push to the branch (`git push origin feature-name`).
5. Open a pull request.

Please ensure that your changes adhere to the coding standards and include proper tests.

### **Especial thanks to github.com/pquerna!** 🤝

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Key Elements:

- **Project Overview**: Brief description of the tool and the technology used (Go, Makefile).
- **Installation**: Clear instructions on cloning the repo, building, and running the project.
- **Usage**: Examples of commands and options.
- **Development**: How to contribute and run tests.
- **Contributing**: Instructions on contributing to the project.
