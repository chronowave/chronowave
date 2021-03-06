## ChronoWave Search

Distributed *schema agnostic* search & analytics solution for append only machine generated data. It is built on self compressed full text search engine, capable of handling ad hoc query and machine learning work load. Focusing on getting application faster to the market for an agile business while lowering the cost of development & operation.

ChronoWave can be used for:

   * distributed tracing
   * logging and log analytics
   * infrastructure and application performance monitoring
   * machine learning on IOT data
   * data archiving
   * utf-8 encoded multi-byte text document search, ie. fluent search on Chinese, Korean and Japanese characters.

#### Design Goal: Towards Simple, Small and Performant
ChronoWave batches input data stream into segments (or blocks) and transforms the semi structured data into columnar formats, followed by creating self compressed index to produce a succinct index data structure. 

*Simple* to develop, use and operate. Machine generated data can be in different shape and size. ChronoWave simplifies development cycle with its support of schema agnostic semi structured data. Operation is also fairly easy. ChronoWave is engineered on succinct self-compressed index data structure, the basis of shared nothing architecture that can scale out on load.

ChronoWave uses only single copy of machine generated data to support major use cases, like Analytics, AI and Real Time Monitoring. The size of self-compressed index is only a *fraction* of the original. ChronoWave requires *only* the index to filter/extract information or restore entire data set with *SSQL*, Semi Structured Query Language.

ChronoWave transforms semi structured data into columnar formats, leverages modern CPU vector instructions to execute queries and full text search.

### ChronoWave In Action
   * [ChronoWave as Jaeger storage backend](https://github.com/chronowave/opentelemetry)
   
### SSQL
   * [Introduction](https://github.com/chronowave/chronowave/wiki/Semi-Structured-Query-Language)

### command line example

1. build command line executable
```shell script
cd cmd/waverider
go build
```

sample.json is a distributed trace data generated by HotRod app captured in Jaeger. Only partial data listed. 

```json
{
   "tags" : [
      {
         "type" : "string",
         "key" : "http.url",
         "value" : "http://0.0.0.0:8083/route?dropoff=577%2C322&pickup=516%2C208"
      }
   ],
  "startTime" : 1601613777130370,
  "spanID" : "2ef6e3c30af421ea",
   "traceID" : "464382d9a88849ff"
}
```

2. construct index: /startTime as time partition *required*, /traceID, /spanID will be used latter as K/V usage (optional)
```shell script
./waverider index ./testdata/sample.json -d data -t '/startTime' -k '/traceID' -k '/spanID'
```

3. query data: *timeframe* is a SSQL keyword that tells ChronoWave searches data between the required time range.
ChronoWave supports partial words and wild card full text search.
```shell script
./waverider query -d data 'find $log where [$log /logs][/startTime timeframe(1601613777130350, 1801613777130470)] [/tags [/key contain("http.url")] [/value contain("dropoff*pickup")]]'
```

4. key/value lookup: *key* is a SSQL keyword that tells ChronoWave lookups by the key. The JSON path of the key must be provided at the time of building index.
```shell script
./waverider query -d data 'find $a where [$a /process][/traceID key("464382d9a88849ff")]'
```

### Version
| feature | Open Source | Community Edition |
| ------- | ----------- | ----------------- |
| search & analytics | :white_check_mark: | :white_check_mark: |
| single node / embedded | :white_check_mark: | :white_check_mark: |
| multi-nodes cluster | :negative_squared_cross_mark: | :white_check_mark: |
| label index selection | :negative_squared_cross_mark: | :white_check_mark: |
| SIMD / AVX512 | :negative_squared_cross_mark: | :white_check_mark: |
| License | Apache v2.0 | coming soon |

### Questions or Suggestions
Please comment on [Gitter](https://gitter.im/chronowave/community)