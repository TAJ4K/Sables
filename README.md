# Sables
A tool for the visually impaired

## Installation
Visit our website [sables.tech](https://sables.tech) and click on the "Install" button.
Run the downloaded .exe and you're good to go!

## Usage
With the exe running in the background, simply press Alt+H to turn it on
Then, drag your mouse over the object you want to visualize and after a few seconds, you'll see a popup with the object's name.

## Want to build it for yourself?
Download the source code from [GitHub](https://github.com/TAJ4K/Sables/tree/main/client) and download Golang here [https://go.dev/](https://go.dev/)
Edit the .env.example to a .env file with the following content:
```env
AWS_ACCESS_KEY_ID=[your access key]
AWS_SECRET_ACCESS_KEY=[your secret key]
```
Then, run 
```sh
go build -ldflags -H=windowsgui
```
to have the exe open without the cli, or use 
```sh
go build .
```
to have it run with the cli.
