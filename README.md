# OTP Generator CLI 🛠️🔐

A simple Command Line Interface (CLI) tool to generate One-Time Password (OTP) codes based on the TOTP (Time-based One-Time Password) algorithm. This project is written in Go (Golang) and allows you to securely generate OTP codes from a shared secret.


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

1. **Clone the repository**:
    ```bash
    git clone https://github.com/yourusername/otp-generator-cli.git
    cd otp-generator-cli
    ```

2. **Build the project using Make**:
    ```bash
    make build
    ```

3. **Run the executable**:
    After building, you will have the `otp-generator-cli` binary ready to use.

## Usage 📖

To generate an OTP code, you will need to provide a secret key (in base32 encoding) as input. Here's how to use the tool:

### Basic Command
```bash
./otp-generator-cli --secret="your-secret-here"
```

### Example:
```bash
./otp-generator-cli --secret="JBSWY3DPEHPK3PXP"
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
    git clone https://github.com/yourusername/otp-generator-cli.git
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

## License 📄

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

### Key Elements: 📊

- **Project Overview**: Brief description of the tool and the technology used (Go, Makefile).
- **Installation**: Clear instructions on cloning the repo, building, and running the project.
- **Usage**: Examples of commands and options.
- **Development**: How to contribute and run tests.


## Contribution 🧙

I am using the package [OTP](https://github.com/pquerna/otp) from [Paul Querna](https://github.com/pquerna). I just created my version (for studying and working with).