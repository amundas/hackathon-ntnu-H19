# Database
The database is used to store data coming from the gateway, and as the place were the end application can get this data. This example uses firebase firestore. 

* Create a new project at the [firebase console](https://console.firebase.google.com/).
* Go to database -> create database and select "start in locked mode"
* Find the "rules" pane in the database page, and change it to the rules below
```
rules_version = '2';
service cloud.firestore {
  match /databases/{database}/documents {
    match /{document=**} {
      allow read: if true;
      allow write: if false;
    }
  }
}
```
This set of rules will allow anyone to read the content of the database, which is not normally what you want, but for the purposes of this example it is good enough.