#coding:utf-8
from __future__ import division

__author__ = 'bater.makhabel'

import os
import re
import sys
import json
import time
import pickle
import builtins
import itertools

from datetime import datetime

def check_line_format(line,line_no):
    if x is None:
        #print(x)
        #print(f"line_no = {line_no}")
        line_no = line_no + 1
    elif x[0] != '+':
        print(x)
        print(f"line_no = {line_no}")


def sort_bibliography(input_file_path, output_file_path):
    f = open(input_file_path)
    lines = f.readlines()
    f.close()

    line_no = 1
    x_dict = {}
    for x in map(lambda x: (re.split('[\[\]]',x)[1],x), lines):
        #print(x)
        #print(re.findall('\d+', str(x[0])))
        check_format = False
        if check_format:      
            check_line_format(line=x,line_no=line_no)

        line_no = line_no + 1
        new_year_key = re.findall('\d+', str(x[0]))[0]
        if int(new_year_key)<20:
            new_year_key = "20" + new_year_key
        else:
            new_year_key = "19" + new_year_key

        if new_year_key in x_dict.keys():
            x_dict[new_year_key].append(x[1])
        else:
            x_dict[new_year_key] = [x[1]]

    print(f"item count = {len(x_dict)}")
    x_dict_sorted = sorted(x_dict.items(),key=lambda item:item,reverse=True)
    print(f"item count = {len(x_dict_sorted)}")
    for x in x_dict_sorted:
        print(x[0],x)

    with open(output_file_path, "a") as result_file:
        for books_per_year in x_dict_sorted:
            print(f"Year: {books_per_year[0]}")
            for book in books_per_year[1]:
                #print(book)
                result_file.write(book)

if __name__ == "__main__":
    input_filename = "bibliography.txt"
    output_filename = "sorted_"+str(time.strftime('%Y%m%d%H%m%S', time.localtime())) + input_filename
    sort_bibliography(input_file_path=input_filename,output_file_path=output_filename)
