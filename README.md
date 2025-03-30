# 🎥 Go + HLS Video Streaming Server

This project demonstrates how to build a **video streaming server using Go and HLS (HTTP Live Streaming)**, along with a simple HTML client interface.  

## 📁 Project Structure

- `server/` – Go server to serve HLS content
- `client/` – HTML client that plays the video using the `<video>` tag
- `output.m3u8` + `.ts` files – generated video chunks in HLS format

## 🚀 Quick Start

> ⚠️ Before starting, you need to chunk your video using **FFmpeg**.

### 1. ⚙️ Generate HLS segments

Use the following command to convert your video (`input.mp4`) into HLS format:

```bash
ffmpeg -i input.mp4 -bsf:v h264_mp4toannexb -codec copy -hls_list_size 0 output.m3u8
```

This will generate:
- `output.m3u8` – the HLS manifest
- `output0.ts`, `output1.ts`, ... – video segments

Make sure these files are placed in the appropriate public/static directory the server can access.

---

### 2. 🖥️ Run the server

Navigate to the `server` folder and start the Go server:

```bash
go run main.go
```

The server will start at: [http://localhost:8080](http://localhost:8080)

---

### 3. 🌐 Open the client

Open the `client/index.html` file in your browser.

> For best compatibility, use Safari (which supports HLS natively), or include [hls.js](https://github.com/video-dev/hls.js) in the client for browsers like Chrome and Firefox.

---

## 💡 Use Cases

This project can be used as a boilerplate for:
- Simple HLS video streaming servers
- Local testing of video content
- Learning projects for streaming, FFmpeg, and Go backend development

---

## 📦 Dependencies

- [Go](https://golang.org/)
- [FFmpeg](https://ffmpeg.org/)
- [VideoJS](https://videojs.com/)
- HTML5 Video

---
