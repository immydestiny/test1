#!/usr/bin/python
# -*- coding: UTF-8 -*-

import rpm
import os
import json


def listdir(path,list_name):
    for file in os.listdir(path):
        file_path = os.path.join(path, file)
        if os.path.isdir(file_path):
            listdir(file_path, list_name)
        else:
            list_name.append(file_path)
    return list_name


def get_rpminfo(rpm_file):
    rpminfo = {}
    ts = rpm.TransactionSet()
    ts.setVSFlags(rpm._RPMVSF_NOSIGNATURES |rpm._RPMVSF_NODIGESTS)
    rpmhdr = ts.hdrFromFdno(rpm_file)
    rpminfo["Name"] = rpmhdr["NAME"]
    rpminfo["Version"] = rpmhdr["VERSION"]
    rpminfo["Release"] = rpmhdr["RELEASE"]
    rpminfo["Epoch"] = rpmhdr["Epoch"]
    rpminfo["Arch"] = rpmhdr["Arch"]
    #print(rpminfo)
    return rpminfo
  
  
if __name__=='__main__':
    file_list=[]
    rpm_list = []
    path='/root/rpms'
    listdir(path,file_list)
    for rpm_file in file_list:
        #print(rpm_file)
        a = get_rpminfo(rpm_file)
        #print rpminfo
        print a
        #rpm_list.append(rpminfo)
        rpm_list.append(dict(a))
    print(rpm_list)

    #print(rpminfo)
    jsonObj = json.dumps(rpm_list)
    with open('test_data.json','w') as json_file:
        json_file.write(jsonObj)
