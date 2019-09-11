# Website
This example website is an absolute minimum-effort example of how to get data from the database. It can be tested out by opening *index.html* in your browser. More information on how to read data from firestore can be found [here](https://firebase.google.com/docs/firestore/quickstart).

In the *script.js* file there is some credentials that depend on the firebase project used. The information for your database can be found somewhere in the firebase console (under add app -> web maybe?).

```
firebase.initializeApp({
    apiKey: "AIzaSyDLdfkyv0R3BH2nxvclU-SQC--OEvj5gzM",
    authDomain: "hackathon-ntnu.firebaseapp.com",
    databaseURL: "https://hackathon-ntnu.firebaseio.com",
    projectId: "hackathon-ntnu",
    storageBucket: "hackathon-ntnu.appspot.com",
    messagingSenderId: "38635872033",
    appId: "1:38635872033:web:91b586ace71404f27ea860"
});
```