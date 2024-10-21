package main

import (
	"flag"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/types/pluginpb"

	"kwai/tools/generator"
)

func main() {
	flag.Parse()
	var flags flag.FlagSet
	protogen.Options{ParamFunc: flags.Set}.Run(
		func(plugin *protogen.Plugin) error {
			plugin.SupportedFeatures = uint64(pluginpb.CodeGeneratorResponse_FEATURE_PROTO3_OPTIONAL)
			for _, f := range plugin.Files {
				if !f.Generate {
					continue
				}
				generator.GenerateFile(plugin, f)
			}

			return nil
		},
	)
}
