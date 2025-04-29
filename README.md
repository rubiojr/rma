# Risor Module Assistant

Use AI to create Risor Go modules.

The Risor Code Assistant is a repository containing tools, examples, and resources to enhance Large Language Models' (LLMs) ability to write Risor modules. It provides context, reference implementations, and best practices to help AI models generate more accurate, efficient, and idiomatic Risor programs.

## Using the context

> [!WARNING]
> The context.txt file is rather large (~53K tokens). Keep in mind that it may burn your LLM credits quite fast.

Code editors like Zed can be given a context file to provide the necessary information for the assistant to understand the problem and generate the appropriate Risor code.

With a prompt like this and `context.txt`:

```
Wrap github.com/tkrajina/gpxgo in a Risor module.

tkrajina/gpxgo sources are in the src folder.

Wrap only the following types and methods:

* GPX
  * Parse
  * ParseBytes
  * ParseFile
  * GetTrackPointsNo
  * GetGpxInfo

A GPX should also contain a list of tracks. Each track should contain a list of segments and each segment a list of points.
```

Claude 3.7 was able to generate a Risor module and documentation like https://github.com/rubiojr/gpx-risor, that required only some small tweaks to build and work.
