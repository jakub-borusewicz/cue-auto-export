package main
import "fmt"
import "os"
import S "strings"
import "path"
import "cuelang.org/go/cue/load"
import "cuelang.org/go/cmd/cue/cmd"
import "cuelang.org/go/cue/cuecontext"
func main() {
  ctx := cuecontext.New()
  files := os.Args[1:]
  for _, file := range files {
    no_cue_filename := S.TrimSuffix(file, ".cue")
    extension := path.Ext(no_cue_filename)
    fmt.Printf("%v\n", extension)
    fmt.Printf("%v\n", no_cue_filename)
    insts := load.Instances([]string{file}, nil)
    v := ctx.BuildInstance(insts[0])
    fmt.Printf("l1%v\n", v)

//     command_string := fmt.Sprintf("export %[1]v --out yaml --outfile %[2]v  --force", file, no_cue_filename)
//     command_string := "export other_example.cue --out yaml --outfile other_example  --force"
    command_string := "export"
    export_command, _ := cmd.New([]string{command_string, "other_example.cue", "--out", "yaml", "--outfile", "other_example", "--force"})
    res := export_command.Execute()
    fmt.Printf("l2%v\n", res)
//     ctx.CompileString()
//     f, err := os.OpenFile(no_cue_filename, os.O_CREATE|os.O_WRONLY, 0644)
//     if err != nil {
//         log.Fatal(err)
//     }
//
//     f.WriteString(v)
//     defer f.Close()

  }

  fmt.Println("Hello, World!" + S.Join(files, " "))
}