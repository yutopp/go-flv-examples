# go-flv-examples

Example programs which using [go-flv](https://github.com/yutopp/go-flv) library.

## How to build

```
make dep-init # install dependencies, dep is required
make
```

Programs are generated under `dist` directory.

## Programs

### cmd/dump

`dump` outputs flv header data and tag data.

#### Usage

```
Usage of dist/dump:
  -i string
    	Input file
```

#### Example

```
> dist/dump -i /tmp/flv_h263_mp3.flv

2018/06/01 19:29:00 Open file: /tmp/flv_h263_mp3.flv
2018/06/01 19:29:00 Header: &{Version:1 Flags:2 DataOffset:9}
2018/06/01 19:29:00 Tag: {TagType:18 Timestamp:0 StreamID:0 Data:<nil>}
2018/06/01 19:29:00 Tag: {TagType:8 Timestamp:0 StreamID:0 Data:0xc42000a160}
2018/06/01 19:29:00 Tag: {TagType:9 Timestamp:0 StreamID:0 Data:0xc42000a1c0}
...
```
(TODO: fix Data section to show internal payload...)

## License

[Boost Software License - Version 1.0](./LICENSE_1_0.txt)
