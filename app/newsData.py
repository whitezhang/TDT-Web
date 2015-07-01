# Usage: Extract content from json separately and store in different files
# Month: 4 - 6
# Input: news_data.json

import json
from datetime import datetime
import os
import re

fin = open("../db/news_data.json")
fout_path = "../db/gms/s_month-"
lines = fin.readlines()

for mon in range(4, 7):
	for line in lines:
		info = json.loads(line)
		# print info["timeStamp"]
		day, month, year = [int(n) for n in info["timeStamp"].split(" ")[0].split("/")]
		time = datetime(year, month, day)
		if time < datetime(2015, mon+1, 1) and time >= datetime(2015, mon, 1):
			if not os.path.exists(fout_path + str(mon)):
				os.makedirs(fout_path + str(mon))
			filename = fout_path + str(mon) +"/" + info["_id"]["$oid"]
			content = info["mainStory"].encode('utf8')
			content = re.sub(r'</?\w+[^>]*>', '', content)
			fout = open(filename, "w")
			fout.write(content)
			fout.close()