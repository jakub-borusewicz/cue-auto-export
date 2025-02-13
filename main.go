package main
import "os"
import S "strings"
import "path"
import "cuelang.org/go/cmd/cue/cmd"
import "fmt"
import "strings"
import "github.com/spf13/cobra"
import "regexp"
import cartesian "github.com/schwarmco/go-cartesian-product"


var extension_out_map = map[string]string{
    ".json": "json",
    ".cue": "cue",
    ".yaml": "yaml",
    ".yml": "yaml",
    ".jsonl": "jsonl",
    ".ldjson": "jsonl",
    ".textproto": "textproto",
    ".proto": "proto",
    ".go": "go",
    ".txt": "text",
    "": "text",
}

func main() {
  var parametrizationMapFlag map[string]string
  cmd := &cobra.Command{
     Use: "cue-auto-export",
     Run: func(cmd *cobra.Command, args []string) {
         parametrizationMap := map[string][]string{}
         for k, v := range parametrizationMapFlag {
              parametrizationMap[k] = strings.Split(v, ",")
         }
         // todo pass parametrizationMap to run_for_files and use it
         fmt.Printf("parametrizationMap: %v\n", parametrizationMap)
         run_for_files(args, parametrizationMap)
     },
  }

  cmd.PersistentFlags().StringToStringVarP(&parametrizationMapFlag, "parametrize", "p", nil, "parametrization map string")
  if err := cmd.Execute(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
}

type FileExportParameters struct {
    cue_filename string
    no_cue_filename string
    out_format string
    tags_map map[string]string
}

func cartesianProduct(parametrizationMap map[string][]string) []map[string]string {
    keys := []string{}
    values := [][]interface{}{}
    for k, v := range parametrizationMap {
        keys = append(keys, k)
        var interfaceSlice []interface{}
        for _, value := range v {
            interfaceSlice = append(interfaceSlice, value)
        }
        values = append(values, interfaceSlice)
    }

    c := cartesian.Iter(values...)
    result := []map[string]string{}

    for tuple := range c {
        tuple_map := map[string]string{}
        for i, key := range keys {
            tuple_map[key] = tuple[i].(string)
        }
        result = append(result, tuple_map)
    }

    return result
}


func run_for_files(files []string, parametrizationMap map[string][]string) {
    for _, file := range files {
        fmt.Println("Exporting file", file)
        filename_parse_result := parse_filename(file)
        out_format, extension_mapped := extension_out_map[filename_parse_result.extension]
        if !extension_mapped {
           out_format = "text"
        }
        if len(filename_parse_result.parametrizationVariables) == 0 {
            export_file(FileExportParameters{file, filename_parse_result.no_cue_filename, out_format, map[string]string{}})
        } else {
            fileParamMap := map[string][]string{}
            for _, paramVar := range filename_parse_result.parametrizationVariables {
                paramValueList, is_mapped := parametrizationMap[paramVar]
                if !is_mapped {
                    fmt.Printf("Error - cannot match placeholder {%s} with any provided parameters while exporting file {%s}\n", paramVar, file)
                    panic("")
                }
                fileParamMap[paramVar] = paramValueList
            for _, tags_map := range cartesianProduct(fileParamMap) {
                result_file_name := filename_parse_result.no_cue_filename
                for k, v := range tags_map {
                    result_file_name = S.ReplaceAll(result_file_name, "{" + k + "}", v)

                }
                export_file(FileExportParameters{file, result_file_name, out_format, tags_map})

            }
            }
        }
    }
}


func export_file(params FileExportParameters) {
    command_list := []string{"export", params.cue_filename, "--out", params.out_format, "--outfile", params.no_cue_filename, "--force"}
    for k, v := range params.tags_map {
        command_list = append(command_list, []string{"--inject", k + "=" + v}...)
    }
    export_command, _ := cmd.New(command_list)
    result := export_command.Execute()
    if result != nil {
        fmt.Println("Error - could not export file", params.cue_filename)
        panic(result)
    }
}


type FilenameParseResult struct {
    no_cue_filename string
    extension string
    parametrizationVariables []string
}

func parse_filename(filename string) FilenameParseResult {
    no_cue_filename := S.TrimSuffix(filename, ".cue")
    extension := path.Ext(no_cue_filename)
    r, _ := regexp.Compile("{(.+?)}")
    parametrizationVariablesMatches := r.FindAllStringSubmatch(no_cue_filename, -1)
    parametrizationVariables := []string{}
    for _, match := range parametrizationVariablesMatches {
        parametrizationVariables = append(parametrizationVariables, match[1])
    }
    return FilenameParseResult{no_cue_filename, extension, parametrizationVariables}
}
