# slack-slash-command
Slack slash command service written in Golang

## Prerequisites

### Download the project

```bash
git clone https://github.com/anubhavmishra/slack-slash-command.git
```

### Download and install ngrok

```bash
brew install ngrok/ngrok/ngrok
```

or

```bash
sudo unzip ~/Downloads/ngrok-stable-darwin-amd64.zip -d /usr/local/bin
```

## Usage

### Build the project

Build the project

```bash
cd slack-slash-command
```

```bash
make build-service
```

### Create a Slack app and generate verification token

* Go to [Slack API](https://api.slack.com/) webpage
* Click "Create an app"
* Select "From scratch" option from the "Create an app" modal
* Enter application name and select a workspace for your application
* Click "Create App" button to create your application. This should create your application and take you to the "Basic Information" page of your application.
* Copy the "Verification Token" from the "App Credentials" section. You can find the "App Credentials" section by scrolling down towards the bottom of the "Basic Information" page.

The video below showcases how to create a Slack app and find the verification token.

https://user-images.githubusercontent.com/4024506/158033562-b41e8678-7151-492a-9cf8-47a28a76c204.mov

Set Slack token

```bash
cd slack-slash-command
```

```bash
export SLACK_VERIFICATION_TOKEN="SLACK TOKEN HERE"
```

### Start slash command application

```bash
cd slack-slash-command
```

```bash
./build/slack-slash-command
```

Expected output

```bash
2022/03/12 18:53:59 PORT environment variable not found defaulting to port 8080
2022/03/12 18:53:59 Server listening on port 8080

```

The slash command application is now listening on port 8080

### Start a ngrok tunnel to expose the slash command application

In a new terminal window, run the command below

```bash
ngrok http 8080
```

Expected output

```bash
ngrok by @inconshreveable                                                                                                                                    (Ctrl+C to quit)

Session Status                online
Session Expires               1 hour, 59 minutes
Version                       2.3.40
Region                        United States (us)
Web Interface                 http://127.0.0.1:4040
Forwarding                    http://8dec-49-207-209-126.ngrok.io -> http://localhost:8080
Forwarding                    https://8dec-49-207-209-126.ngrok.io -> http://localhost:8080

Connections                   ttl     opn     rt1     rt5     p50     p90
                              0       0       0.00    0.00    0.00    0.00
```

Copy the `https` address from the output. In this case, it's `https://8dec-49-207-209-126.ngrok.io`


### Create a slash command

* From your application homepage, select "Slash Commands" from the side navigation bar.
* From the "Slash Commands" section, select "Create New Command"
* Enter details for your slash command. First, in the "Command" input box enter `/hello`
* In the "Request URL" input box paste the ngrok `https` address and add `/receive` at the end of it.
For this example, the `https` URL will look like this `https://8dec-49-207-209-126.ngrok.io/receive`
* Enter a short description in the "Short Description" input box
* Click the "Save" button in the lower right corner of the webpage.

The video below showcases how to create a slash command.

https://user-images.githubusercontent.com/4024506/158033572-bac3294d-4b85-43f0-b38f-8bd7ef8d3869.mov

### Install Slack app

* Go to the "Basic Information" page of your application
* In the "Install your app" section click on "Install to Workspace" button
* Click "Allow" to give permission to the app. Application should now be installed.

The video below showcases how to install a Slack app.

https://user-images.githubusercontent.com/4024506/158033577-8f542a28-f304-491f-a648-a42c055af37b.mov

### Use the slash command in Slack

* Go to the Slack channel you want to install the slash command.
* Click on the View all members of this channel" button on the top right corner of the Slack channel
* Select the "Integrations" tab in the modal.
* In the "Apps" section click "Add apps" button.
* Scroll to find the "Slack Slash Command" app and click the "Add" button beside its name. Slack application should now be installed in the channel.
* In the message box of the channel type `/hello` to invoke the slash command. You should see a message "Hello World!" back from the "Slack Slash Command" application.

The video below showcases how to install Slack application in a channel and invoke `/hello` slash command to get a response back from the "Slack Slash Command" application.

https://user-images.githubusercontent.com/4024506/158033583-5e245540-0702-4b43-978d-21873a1c511e.mov
