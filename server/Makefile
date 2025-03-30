run:
	go build main.go
	main.exe

convert:
	ffmpeg -i input.mp4 -bsf:v h264_mp4toannexb -codec copy -hls_list_size 0 output.m3u8