# <img src="https://github.com/Ykaros/phrame/blob/main/utils/logo.png" width="101" height="100"> 

Phrame: A Simple CLI Tool for Adding Frames to Photos

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

There will be a release version soon..

### Installing
To build locally, first
```
git clone https://github.com/Ykaros/phrame.git
```
and then 
```
cd phrame && task build
```
finally
```
./phrame -h
```
To set the command as global, simply 
```
task global
```

### Running Tests
Use:
```
task test
```
If everything goes well, you should see: `Image successfully saved to: out/test.png ✓`.

## CLI Manual
The fundamental function of this CLI is to add customizable frames to images. To add frame(s) to image(s), use the following command:

```
phrame -q -i test -o out -r 0.1 -c "#FFFFFF"
```

Except the root command, there are two available subcommands: sign and cut. 
### Sign
This subcommand is used to add a signature (watermark) on the frame. I personally am not interested in anything that blinds any part of my photo so this tool is for you if you have same concerns.
The complete command goes like:
```
phrame sign -s [signature] -i [input_path] -o [output_path] -r [border_ratio] -c [frame_color] -x [font_size] -y [font_color]
```

To make it more customizable, it supports loading any font (.ttf) by simply replacing the default font (Inter-Regular) named by "font.ttf" with [whatever font you are fond of](https://fonts.google.com/). 
### CUT
The cut subcommand can be used to divide an image into either four or nine equal parts:
```
phrame cut [image_path] -g [4 or 9]
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details


