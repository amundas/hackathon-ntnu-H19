
/*  
    Initialize firebase with the credentials for the database. 
    This data is not considered a secret, since security is handled elsewhere
*/
firebase.initializeApp({
    apiKey: "AIzaSyDLdfkyv0R3BH2nxvclU-SQC--OEvj5gzM",
    authDomain: "hackathon-ntnu.firebaseapp.com",
    databaseURL: "https://hackathon-ntnu.firebaseio.com",
    projectId: "hackathon-ntnu",
    storageBucket: "hackathon-ntnu.appspot.com",
    messagingSenderId: "38635872033",
    appId: "1:38635872033:web:91b586ace71404f27ea860"
});

var db = firebase.firestore();
const rssiOneMeter = -61; // Value got from Nordic's example code
rssiRef = document.getElementById("rssi");
distanceRef = document.getElementById("distance");
lastUpdateRef = document.getElementById("lastUpdate");

db.collection("events0").orderBy('timestamp', 'desc').limit(1).onSnapshot(collection => {
    // Since we used limit(1), the array is of length 1
    const data = collection.docs[0].data();
    console.log(data);
    /* 
        rssi to distance formula from https://iotandelectronics.wordpress.com/2016/10/07/how-to-calculate-distance-from-the-rssi-value-of-the-ble-beacon/
        This is far from accurate, and you are encouraged to experiment yourself!
    */
    const distanceEstimate = 10**((rssiOneMeter - data.rssi)/(10 * 2))
    distanceRef.innerHTML = "Distance: " + distanceEstimate.toFixed(1) + "m";
    rssiRef.innerHTML = "RSSI: " + data.rssi;
    lastUpdateRef.innerHTML = "Last update: " + (new Date(data.timestamp)).toLocaleString('en-GB', { timeZone: 'Europe/Oslo' });

})