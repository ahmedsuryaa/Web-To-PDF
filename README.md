# Web to PDF OpenAgents Plugin

This repository contains a Web to PDF plugin utilizing [URLBox API](https://urlbox.com/url-to-pdf).Built for the OpenAgents platform, written in Go using the Extism Go PDK. This plugin converts a given website URL into a PDF document.


## Compilation and Running:
1. Install TinyGo (if not already installed):

- Follow the instructions at [TinyGo installation guide](https://tinygo.org/getting-started/install/) to install TinyGo.
  
2. Compile the Plugin:
```
tinygo build -o pdf.wasm -target wasi main.go
```
3. Run the Plugin Using Extism CLI:
```
extism call pdf.wasm run --input 'https://some-url.com' --wasi --allow-host "api.urlbox.io"
```
Replace 'https://some-url.com' with the URL you want to render to PDF. The output should now display the rendered URL from the response.

## Running Example 
Command :
```
$ extism call pdf.wasm run --input "https://urlbox.com/url-to-pdf" --wasi --allow-host "api.urlbox.io"
```
Response :
```
PDF URL: https://renders.urlbox.io/urlbox1/renders/664900eb3984217b455b229f/2024/5/18/570be4bc-6dd4-4b33-bfba-4b3d005d215c.pdf
```
