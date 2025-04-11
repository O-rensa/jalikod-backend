run: run-bffe

kill:
	taskkill /F /IM jalikod-bffe.exe

run-bffe:
	start "" ./bffe/bin/jalikod-bffe.exe