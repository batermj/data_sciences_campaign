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

# Deal with bibliography file in the format same as
# ./bibliography/KDD/2019/Research_Track_Papers_Full_List.txt
def check_line_format_kdd_2019(lines=None):
    line_no = 1
    for idx in range(len(lines)//2):
        if( re.match('Authors: ', lines[2*idx:(2*idx+2)][1]) == None):
            print(f"line_no = {line_no}")
            print(lines[2 * idx:(2 * idx + 2)][1])
            print(re.match('Authors: ', lines[2 * idx:(2 * idx + 2)][1]))
            print(re.split('^Authors: ', lines[2 * idx:(2 * idx + 2)][1]))
            break

        line_no = line_no + 1

    if line_no == len(lines):
        print(f"line_no = {line_no} != the lines of input file!")
        exit(0)

    return(None)

# Deal with bibliography file in the format same as
# ./bibliography/KDD/2017/Research_Track_Papers_Full_List.txt
def check_line_format_kdd_2017(lines=None):
    line_no = 1
    for idx in range(len(lines)//2):
        if( re.match('Author\(s\):', lines[2*idx:(2*idx+2)][1]) == None):
            print(f"line_no = {line_no}")
            print(lines[2 * idx:(2 * idx + 2)])
            print(lines[2 * idx:(2 * idx + 2)][1])
            print(re.match('Author\(s\):', lines[2 * idx:(2 * idx + 2)][1]))
            print(re.split('^Author\(s\):', lines[2 * idx:(2 * idx + 2)][1]))
            exit(0)

        line_no = line_no + 1

    if line_no == len(lines):
        print(f"line_no = {line_no} != the lines of input file!")
        exit(0)

    return(None)

# Deal with bibliography file in the format same as
# ./bibliography/KDD/2012/Research_Track_Full_lisit.txt
def check_line_format_kdd_2012(lines=None):
    line_no = 1
    for idx in range(len(lines)//4):
        if( (re.match('Title:', lines[4*idx:(4*idx+4)][1]) == None) & (re.match('Author\(s\):', lines[4*idx:(4*idx+4)][1]) == None) ):
            print(f"line_no = {line_no}")
            print(lines[4*idx:(4*idx+4)])
            print(lines[4*idx:(4*idx+4)][1])
            print(re.match('Title:', lines[4*idx:(4*idx+4)][1]))
            print(re.match('Author\(s\):', lines[4*idx:(4*idx+4)][2]))
            print(re.split('^Title:', lines[4*idx:(4*idx+4)][1]))
            print(re.split('^Author\(s\):', lines[4*idx:(4*idx+4)][2]))
            exit(0)

        line_no = line_no + 1

    if line_no*4 != len(lines)+1:
        print(f"line_no = {line_no} != the lines of input file, {len(lines)}!")
        exit(0)

    return(None)


# Deal with bibliography file in the format same as ./bibliography/bibliography.py
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

def sort_bibliography_kdd_2019(input_file_path, output_file_path, pub_year=None):
    f = open(input_file_path)
    lines = f.readlines()
    f.close()

    check_line_format_kdd(lines=lines)

    with open(output_file_path, "a") as result_file:
        line_no = 1
        for idx in range(len(lines)//2):
            print(f"line_no = {line_no}")
            print(lines[2*idx:(2*idx+2)])
            print(re.match('Authors: ', lines[2*idx:(2*idx+2)][1]))
            print(re.split('^Authors: ', lines[2*idx:(2*idx+2)][1]))
            merged_line = f"+ [{pub_year[2:]}], " + re.split('^Authors: ', lines[2*idx:(2*idx+2)][1])[1].replace("\n","") + ", " + lines[2*idx:(2*idx+2)][0].replace("\n",f", KDD{pub_year}\n")
            result_file.write(merged_line)
            line_no = line_no + 1


def sort_bibliography_kdd_2018(input_file_path, output_file_path, pub_year=None):
    f = open(input_file_path)
    lines = f.readlines()
    f.close()

    check_line_format_kdd(lines=lines)

    with open(output_file_path, "a") as result_file:
        line_no = 1
        for idx in range(len(lines)//2):
            print(f"line_no = {line_no}")
            print(lines[2*idx:(2*idx+2)])
            merged_line = f"+ [{pub_year[2:]}], " + lines[2*idx:(2*idx+2)][1].replace("\n","") + ", " + lines[2*idx:(2*idx+2)][0].replace("\n",f", KDD{pub_year}\n")
            result_file.write(merged_line)
            line_no = line_no + 1

# uses for 2017, 2016
def sort_bibliography_kdd_2017(input_file_path, output_file_path, pub_year=None):
    f = open(input_file_path)
    lines = f.readlines()
    f.close()

    check_line_format_kdd_2017(lines=lines)

    with open(output_file_path, "a") as result_file:
        line_no = 1
        for idx in range(len(lines)//2):
            print(f"line_no = {line_no}")
            print(lines[2*idx:(2*idx+2)])
            print(re.match('Author\(s\):', lines[2*idx:(2*idx+2)][1]))
            print(re.split('^Author\(s\):', lines[2*idx:(2*idx+2)][1]))
            merged_line = f"+ [{pub_year[2:]}], " + re.split('^Author\(s\):', lines[2*idx:(2*idx+2)][1])[1].replace("\n","") + ", " + lines[2*idx:(2*idx+2)][0].replace("\n",f", KDD{pub_year}\n")
            result_file.write(merged_line)
            line_no = line_no + 1

# uses for 2012
def sort_bibliography_kdd_2012(input_file_path, output_file_path, pub_year=None):
    f = open(input_file_path)
    lines = f.readlines()
    f.close()

    check_line_format_kdd_2012(lines=lines)

    with open(output_file_path, "a") as result_file:
        line_no = 1
        for idx in range(len(lines)//4):
            print(f"line_no = {line_no}")
            print(lines[4*idx:(4*idx+4)])
            print(re.match('Title:', lines[4*idx:(4*idx+4)][1]))
            print(re.match('Author\(s\):', lines[4*idx:(4*idx+4)][2]))
            print(re.split('^Title:', lines[4*idx:(4*idx+4)][1]))
            print(re.split('^Author\(s\):', lines[4*idx:(4*idx+4)][2]))

            merged_line = f"+ [{pub_year[2:]}], " + re.split('^Author\(s\):', lines[4*idx:(4*idx+4)][2])[1].replace("\n","") + ", " + re.split('^Title:', lines[4*idx:(4*idx+4)][1])[1].replace("\n",f", KDD{pub_year}\n")
            result_file.write(merged_line)
            line_no = line_no + 1

def load_and_sort_bibliography():
    input_filename = "bibliography.txt"
    output_filename = "sorted_"+str(time.strftime('%Y%m%d%H%m%S', time.localtime())) + input_filename
    sort_bibliography(input_file_path=input_filename,output_file_path=output_filename)

def load_and_sort_bibliography_from_kdd():
    pub_year = "2012"
    # for 2016~2019
    #input_filenames = [f"bibliography/KDD/{pub_year}/Research_Track_Papers_Oral_Full_List", f"bibliography/KDD/{pub_year}/Research_Track_Papers_Poster_Full_List",f"bibliography/KDD/{pub_year}/Applied_Data_Science_Track_Papers_Oral_Full_List",f"bibliography/KDD/{pub_year}/Applied_Data_Science_Track_Papers_Poster_Full_List"]
    input_filenames_2012 = [f"bibliography/KDD/{pub_year}/Research_Track_Full_lisit", f"bibliography/KDD/{pub_year}/Industrial_and_Government_Track_Full_List"]

    for input_filename in input_filenames_2012:
        output_filename = input_filename + str(time.strftime('%Y%m%d%H%m%S', time.localtime())) + ".txt"
        input_filename = input_filename + ".txt"
        sort_bibliography_kdd_2012(input_file_path=input_filename,output_file_path=output_filename,pub_year=pub_year)

if __name__ == "__main__":
    load_and_sort_bibliography_from_kdd()
