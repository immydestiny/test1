# -*- coding: utf-8 -*-
import json
import csv

'''
@2021.11
@在windows 环境 python3.6 成功将json 文件转换为csv 文件
'''
def csv_json():
    # 1.分别 读，创建文件
    #json_fp = open("1.1.1-report.json", "r")
    #csv_fp = open("08csv.csv", "w")

    # 2.提出表头和表的内容 这个代码提取的内容有点少，分层比较简单，需要改进
    #data_list = json.load(json_fp)
    #print("原数据类型为：" + str(type(data_list)))

    # 读取json文件
    with open("1.1-report.json",encoding='utf-8') as json_obj:
        data = json.load(json_obj)
        print("3原数据类型为：" + str(type(data)))
        data_dict = data[0]
        print("4数据类型为：" + str(type(data_dict)))
        # print(type(json_dict))


    #字典的形式读取数据
    #获取json文件中Vulnerabilities 键对应的所有值,为list
    vul_list = data_dict.get("Vulnerabilities")
    print("5数据类型为：" + str(type(vul_list)))
    #获取Vulnerabilities list 中的字典数据
    vul_data = vul_list[0]
    print("6数据类型为：" + str(type(vul_data)))


    # 3.csv 写入器 需要添加newline='' 删除多余空行 添加encoding='utf-8' 解决UnicodeEncodeError
    with open("1.1-new.csv", "w", newline='',enc;oding='utf-8') as csv_fp:
        writer = csv.writer(csv_fp)
        # 4.写入表头 手动写入，因为之前 vul_data.keys()获取的会少fixedversion,可以自己选择顺序
        writer.writerow(["VulnerabilityID","PkgName","InstalledVersion","FixedVersion","Severity","SeveritySource",
                         "PrimaryURL","Title","Description","CweIDs","CVSS","References","PublishedDate",
                         "LastModifiedDate"])
        # 5.写入内容
        for data in vul_list:
            #print("the data type " + str(type(data)))
            writer.writerow([data.get('VulnerabilityID',''),data.get('PkgName'),data.get('InstalledVersion'),
                             data.get('FixedVersion', ''),data.get('Severity'),data.get('SeveritySource'),data.get('PrimaryURL'),data.get('Title'),
                             data.get('Description'),data.get('CweIDs'),data.get('CVSS'),data.get('References'),
                             data.get('PublishedDate'),data.get('LastModifiedDate')])


if __name__ == "__main__":
    csv_json()
