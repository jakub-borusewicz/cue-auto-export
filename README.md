# cue-auto-export

A [pre-commit](https://pre-commit.com/) hook to automatically export [cue](https://cuelang.org/) files.

### Using with pre-commit

To run hook via pre-commit, add the following to your `.pre-commit-config.yaml`:

```yaml
repos:
  - repo: https://github.com/jakub-borusewicz/cue-auto-export
    rev: v1.0.2
    hooks:
      - id: cue-auto-export
        exclude: (?x)^(
            cue.mod/.* |
            config/.*
          )$
```

In the root of your project, create directory `cue.mod` and `module.cue` file in it, with following content:

```cue
module: "my.app"
language: {
   version: "v0.9.2"
}
```

Now, you can add directory that you can treat as a single source of truth for your configuration files in the project, for example `config`, with file `config/consts.cue`:
```cue
package consts

foo: "bar"
```

Then you can refer to values defined in `config` in other configuration files outside of this directory:

`mylinterconfig.json.cue`
```cue
import consts "my.app/config:consts"

foo: consts.foo
```

Running `pre-commit run --all-files` will result in creation of new file:

`mylinterconfig.json`
```json
{
  "foo": "bar"
}
```