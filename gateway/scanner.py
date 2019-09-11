#!/usr/bin/env python3
from bluepy.btle import Scanner, DefaultDelegate, BTLEManagementError, BTLEDisconnectError
import firebase_admin
from firebase_admin import credentials
from firebase_admin import firestore
from datetime import datetime
import sys
import os
import time

if sys.version_info[0] < 3:
    raise Exception("Must be using Python 3")

cred = credentials.Certificate('raspi-admin.json')
firebase_admin.initialize_app(cred)
db = firestore.client()
doc_ref = db.collection(u'events0')


# if failed scanend, try sudo hciconfig hci0 down && sudo hciconfig hci0 up
class ScanDelegate(DefaultDelegate):
    def __init__(self):
        DefaultDelegate.__init__(self)

    def handleDiscovery(self, dev, isNewDev, isNewData):
        for (adtype, desc, value) in dev.getScanData():
            if adtype == 255:
                if value[:6] == "590002" and value[-10:-2] == "deadc0de":
                    print(dev.rssi)
                    """
                    timestamp = int(datetime.utcnow().strftime("%s")) * 1000 
                    doc_ref.add({
                        u'raw': str(value),
                        u'timestamp': timestamp,
                        u'uuid': value[8:40],
                        u'app_value': value[-10:-2],
                        u'last_byte': value[-2:],
                        u'rssi': dev.rssi
                    })
                    """
                    

scanner = Scanner().withDelegate(ScanDelegate())

while True:
    scanner.scan(1)
"""
    # Had some errors, this code restarts hci0 and then the script when errors occur 
    except BTLEManagementError as error:
        print(error)
        print("Applying fix")
        os.system('sudo hciconfig hci0 down && sudo hciconfig hci0 up')
        print("restarting")
        os.execv('./scanner.py', ['scanner']) 
    except BTLEDisconnectError as error:
        print(error)
        print("Applying fix")
        os.system('sudo hciconfig hci0 down && sudo hciconfig hci0 up')
        print("restarting")
        os.execv('./scanner.py', ['scanner']) 

"""

