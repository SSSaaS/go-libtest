# - Python - #

import json

import sys
sys.path.insert(0, './libs/sssa-python')
from SSSA import sssa

fi = open('./tmp/Python3_Combine_Input.json', 'r')
input = fi.read()
args = json.loads(input)
fo = open('./tmp/Python3_Combine_Output.json', 'w')
sss = sssa()
for index,a in enumerate(args):
	args[index] = a.encode('utf-8')

fo.write(json.dumps(sss.combine(args)))
fo.close()
