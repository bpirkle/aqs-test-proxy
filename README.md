# aqs-test-proxy

This is just toy code for me to play around with something. This README is a bit whimsical, as this code isn't to be taken seriously.

The reason I wrote this Really Bad code was that I was trying to run a test suite against a non-public service by executing the test suite locally through an ssh tunnel. But because the API endpoints provided by the service through the tunnel gave a certificate error. I could get around this in curl by specifying the "insecure" flag, but the test suite had no such option. So I wrote this proxy to make happily make the insecure call and forward the response to the test suite.

Everything in here is extremely specific to the environment I wrote it for.

This is a terrible idea for a lot of reasons, including at least:
* certificates are there for a reason, and bypassing this is usually bad
* the service was publicly inaccessible for a reason, and hitting it from outside the network was bad
* there are probably dozens of existing tools that would have done this for me, without the need to write sketch custom code

So please don't use this code as any kind of pattern. I mostly did it this way to see if it'd work, and as a quick one-off to answer some pressing questions. It was never intended to be part of any ongoing workflow. 

And FWIW, I did clear this with the appropriate folks in my organization before I did it.

Instructions:

* confirm you have working ssh to mwmaint1002.eqiad.wmnet
* add the following to your /etc/hosts file (unless you're on Windows, in which case - why are you on Windows?): "127.0.0.1 staging.svc.eqiad.wmnet"
* clone the test suite from https://gitlab.wikimedia.org/repos/generated-data-platform/aqs/aqs_tests
* hack the utilities/base_uri.py file in the test suite to use this base url: "uri = "http://localhost:8086"
* create the ssh tunnel like this: "ssh -N mwmaint1002.eqiad.wmnet -L 4972:staging.svc.eqiad.wmnet:4972"
* make and run this proxy
* run the test suite like this (to execute just the unique-devices subset of the tests): behave -D env=prod --tags @aqs_tests.unique_devices > test-output.txt
