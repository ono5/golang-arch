# Walk - Building Graphical Windows Applications

Walk is a Windows GUI toolkit for the Go programming language to enable us to build native desktop GUI application for windows using Go.

Thw Walk project is one of the lodest GUIL toolkits for Go.

The name stands for Windows Application Library Kit.

https://github.com/lxn/walk

# Manifest file
The manifest file is required to to tell Windows runtime that we are using the Common Controls framework version 6.0.0.0, which is required by the Walk APIs.

# Build
The walk applications require a manifest file that will be embedded in the executable we are building.

To do this, we need to download the rsrc tool from github.com/akavel/rsrc, which will embed the required metadata.

```
go get github.com/akavel/rsrc

rsrc.exe -manifest hello.exe.manifest

go build -ldflags="-H windowsgui"
```

# Run
The hello world app that was built in the last step can be executed in tow ways:
either by running it from the command line, or by clicking the icon the file manager.