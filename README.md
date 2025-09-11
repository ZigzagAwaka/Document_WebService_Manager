## Simple Document web service manager

This is a small web service application to manage simple documents.

Written in Go as an exposed RESTful API.


### How it works

The program will start a small server from which users can consult the stored documents, add new ones, or remove some.

The server can be comunicated with by consulting the local address on a browser, or by using one of the example commands listed in [example-commands.txt](https://github.com/ZigzagAwaka/Document_WebService_Manager/blob/main/example-commands.txt) in a command prompt window.

### Notes

The documents are stored in memory for this simple example service. But this system can be replaced by a database or any other storage solution.