


import time
import os
import logging
#!/usr/bin/python
fname=('file_%s.txt'%os.getpid())
formatStr='%(asctime)s-[%(name)s]-[%(levelname)s]-[%(filename)s:%(lineno)d] %(message)s'
logging.basicConfig(filename=fname,format=formatStr,level=logging.DEBUG,filemode='a')
logger = logging.getLogger()


i = 0
while True:
	# print("i=%d"%i)
	logger.debug('i=%d'%i)
	time.sleep(1)
	i+=1

    