## maui-gen completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	maui-gen completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
maui-gen completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [maui-gen completion](maui-gen_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 25-Jul-2023