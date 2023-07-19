# memoir

Memoir is a client tool to manage daily tasks.

## Installation

To install this tool, you can download the pre-compiled binary from the 
GitHub releases page: https://github.com/enkodr/memoir/releases or by executing
the following command:

```bash
curl -sL https://raw.githubusercontent.com/enkodr/memoir/main/bin/install | sh - 
```

## Usage

The following parameters are the available:

### init 

Initialises the memoir structure in the `~/.memoir` directory.

### add

Adds a new task to the memoir file.

```bash
memoir add "Title of the task to add"
```

### show

Show all tasks for a specific day. 

If no options are passed, it will show the list of tasks for the current day.

```bash
memoir show
```

You can pass the amount of days to count from today.

```bash
memoir show -1 # it will show yesterday's tasks 
```

### edit 

Edit a task in the daily file.

The file will be open in the editor specified in the `$EDITOR` environment variable. 

### do

Marks a task as completed.

```bash
memoir do 1 # marks the taks with ID 1 as completed
```

### undo

Unmarks a task as completed.

```bash
memoir undo 1 # marks the taks with ID 1 as completed
```

### rm

Removes a task from the list.

```bash
memoir undo 1 # marks the taks with ID 1 as completed
```

### completion

Generate shell completion for memoir.

You can add the following to the `~/.bashrc` file.

```bash
source <(machina completion bash)
```

**Note:** Replace `bash` with `fish` or `zsh` depending on the bash you're using.

## Contributing

If you would like to contribute to this project, please fork the repository and submit a pull request.


## License

This project is licensed under the MIT License.
