# - Python - #

import json

import sys
sys.path.insert(0, './libs/sssa-python')
from SSSA import sssa

fi = open('./tmp/Python3_Create_Input.json', 'r')
input = fi.read()
args = json.loads(input)
fo = open('./tmp/Python3_Create_Output.json', 'w')
sss = sssa()
fo.write(json.dumps(sss.create(int(args[0]), int(args[1]), args[2].encode('utf-8'))))
fo.close()
