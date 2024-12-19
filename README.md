# Delta Encoding example

Project to create custom Delta Encoding. The idea is to create a video streaming with delta encoding.

# Encoding


## VCDIFF

![VCDIFF](./vcdiff_example.png) 

## Read and Write

How to extract frames and use then in memory?

Solutions:
  1. [FFMPEG](https://stackoverflow.com/questions/10957412/fastest-way-to-extract-frames-using-ffmpeg) 
  2. [FFMPEG with Golang Wrapper](https://github.com/u2takey/ffmpeg-go?tab=readme-ov-file#task-frame-from-video) 

## Configurations

### 1. "Byte rate"

The amount of change that can be sent in a request in bytes

# Communication

## TCP/UDP Socket

## USB (Optional)

# References

1. [Wikipedia](https://en.wikipedia.org/wiki/Delta_encoding)
2. [GoUSB](https://github.com/google/gousb)
3. [RFC3284 VCDIFF](https://datatracker.ietf.org/doc/html/rfc3284#section-3) and it's [article](https://www.cs.brandeis.edu/~dilant/cs175/%5BSiying-Dong%5D.pdf)
