# Gateway
The example gateway is based on a raspberry pi, and some code written in go. There is also an example using python, but that code was just an early test before I decided to change to go. Both codes should also work on laptops with bluetooth capabilities. The programs is required to be run with root acess. A bunch of dependencies needs to be installed, but that should be relatively self-explanatory.

## Connecting a raspberry pi to eduroam
[This guide](https://autottblog.wordpress.com/raspberry-pi-arduino/connecting-raspberry-pi-to-eduroam/) shows how to connect a raspberry pi to eduroam on NTNU. 

## Giving the gateway access to the database
In order to give the gateway acess to the database, we can use a service account. This can be done in [cloud console](https://console.cloud.google.com/). Set the service account role to "Firebase Admin", and create and download a key in JSON format. The key is the file referenced in the gateway code (raspi-admin.json). When the raspberry pi uses a service account, it has full acess to the firestore project, and it even ignores the database rules that is set. In other words; the service account file is a secret that should not be made available to others.