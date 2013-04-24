#  GoLang clickstart (click and go ! get it? nevermind...)

This shows off the google <a href="http://golang.org">go language</a> on CloudBees. 
This starter app includes a simple wiki app (in Go) and a Jenkins build that runs tests and deploys if the tests pass.

<img src="https://d3ko533tu1ozfq.cloudfront.net/clickstart/googlego.png">

See <a href="http://golang.org/doc/articles/wiki/">here</a> for the article about building the app in Go.

Press the button to build, test and deploy this instantly:

<a href="https://grandcentral.cloudbees.com/?CB_clickstart=https://raw.github.com/michaelneale/golang-clickstart/master/clickstart.json"><img src="https://d3ko533tu1ozfq.cloudfront.net/clickstart/deployInstantly.png"/></a>

# To run manually locally

0. Install Google Go
1. Clone this repo
2. cd app
3. export PORT=8080 && go run

# Procfile

This also uses the CloudBees procfile support - the app is packaged up as a binary with a Procfile in the root specifying the launch application.

Fork this and make it your own!

