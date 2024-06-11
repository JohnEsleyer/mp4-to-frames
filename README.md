# mp4-to-frames
This is a simple CLI tool written in Go that extracts frames from an .mp4 file and saves them as individual images in a specified output directory. It also includes a GUI mode for ease of use.


### Features
- Extract frames from an .mp4 file using a CLI or GUI interface.
- Specify an output directory for the extracted frames.
- Automatically creates the output directory if it does not exist.

### Requirements
- Go
- ffmpeg installed and available in your system's PATH.

### Installation
```
go build -o mp4-to-frames
```

### Usage 
To use the CLI, run the following command:
```
./mp4-to-frames <input-file> [output-directory]
```
To use the GUI, run the following command:
```
./mp4-to-frames gui
```

### Screenshots
![image](https://github.com/JohnEsleyer/mp4-to-frames/assets/66754038/6d96c094-325f-40b9-a288-057e662d75b9)


