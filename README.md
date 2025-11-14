# Core Element Template

This repository is a template for developers to create custom HTML elements for the core web3 framework. It includes a Go backend, a Lit-based custom element, and a full release cycle configuration.

## Getting Started

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/your-username/core-element-template.git
    ```

2.  **Install the dependencies:**
    ```bash
    cd core-element-template
    go mod tidy
    cd ui
    npm install
    ```

3.  **Run the development server:**
    ```bash
    go run ./cmd/demo-cli serve
    ```
    This will start the Go backend and serve the Lit custom element.

## Building the Custom Element

To build the Lit custom element, run the following command:

```bash
cd ui
npm run build
```

This will create a JavaScript file in the `dist` directory that you can use in any HTML page.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the EUPL-1.2 License - see the [LICENSE](LICENSE) file for details.
