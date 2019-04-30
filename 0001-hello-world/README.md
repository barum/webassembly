# webassembly

There are two parts: client and server.
Client = gets loaded into the browser

Server = will serve the files. If you try and use the `index.html` in the client directory from file open or live server, you will get an error.

### Client code

Go into the client directory to create the client code
```
cd client
```

Need to get the Golang `wasm_exec.js` file needed that comes with Golang.

```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

We make a simple `main.go` file and build it. a `main.wasm` file is created to be used in `index.html`file.

```
GOOS=js GOARCH=wasm go build -o main.wasm
 ```

After these steps there should be 4 files in the client folder:

1. main.go
2. index.html
3. wasm_exec.js
4. main.wasm

### Server

From the project root:

```
cd server
```
Now we do a normal:
```
go run main.go
```
Now you can navigate to <a> http://localhost:9999</a>

Open up the developer tools and in the Console you should see "Hello, WebAssembly"

