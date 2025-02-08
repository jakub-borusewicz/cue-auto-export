package main
import "os"
import S "strings"
import "path"
import "cuelang.org/go/cmd/cue/cmd"
import "fmt"

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
  files := os.Args[1:]
  run_for_files(files)
}

func run_for_files(files []string) {
    for _, file := range files {
        fmt.Println("Exporting file", file)
        no_cue_filename := S.TrimSuffix(file, ".cue")
        extension := path.Ext(no_cue_filename)
        out_format, extension_mapped := extension_out_map[extension]
        if !extension_mapped {
           out_format = "text"
        }
        export_command, _ := cmd.New([]string{"export", file, "--out", out_format, "--outfile", no_cue_filename, "--force"})
        result := export_command.Execute()
        if result != nil {
            fmt.Println("Error - could not export file", file)
            panic(result)
        }
    }
}