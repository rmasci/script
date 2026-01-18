# Spinner Feature - With Custom Types
The script pipeline now includes a configurable spinner option that you can apply to any stage of your pipeline. When enabled, it displays a loading animation on stderr while data flows through, without interfering with the actual output.
You can choose from **91 different spinner types** to customize the appearance.
## Usage
### Basic Usage (Default Spinner)
Add a spinner to any point in your pipeline:
```go
package main
import "github.com/rmasci/script"
func main() {
    // Display spinner with default animation (type 9: / - \)
    script.File("data.txt").
        Spinner("Reading file...").
        Stdout()
}
```
### Custom Spinner Types
Pass a spinner type number (0-90) as the second parameter:
```go
// Type 0: Arrows (← ↖ ↑ ↗ → ↘ ↓ ↙)
script.Exec("command").Spinner("Processing...", 0).Stdout()
// Type 1: Bar (▁ ▃ ▄ ▅ ▆ ▇ █ ▇ ▆ ▅ ▄ ▃ ▁)
script.File("data.txt").Spinner("Reading...", 1).Stdout()
// Type 25: Bounce (⠁ ⠂ ⠄ ⠂)
script.ListFiles("*.txt").Spinner("Searching...", 25).Stdout()
// Type 14: Dots (⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏)
script.Exec("grep pattern /data").Spinner("Searching...", 14).Stdout()
```
## Available Spinner Types
There are 91 spinner types available (0-90). Here are some popular ones:
| Type | Animation | Description |
|------|-----------|-------------|
| 0 | ← ↖ ↑ ↗ → ↘ ↓ ↙ | Arrows |
| 1 | ▁ ▃ ▄ ▅ ▆ ▇ █ | Bar |
| 2 | ▖ ▘ ▝ ▗ | Box corner |
| 3 | ┤ ┘ ┴ └ ├ ┌ ┬ ┐ | Box line |
| 4 | ◢ ◣ ◤ ◥ | Triangle |
| 8 | . o O @ * | Bounce dot |
| 9 | / - \ | Simple (default) |
| 14 | ⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏ | Dots |
| 25 | ⠁ ⠂ ⠄ ⠂ | Braille bounce ||| 26 | ⠈ ⠐ ⠠ ⠰ ⠸ ⠐ ⠈ | Braille circles |
## Chained Pipelines with Spinner Types
Use different spinner types at different stages:
```go
script.ListFiles("*.txt").
    Spinner("Finding files...", 0).           // Arrows
    ExecForEach("process {{.}}").
    Spinner("Processing files...", 1).        // Bar
    Match("success").
    Spinner("Filtering results...", 25).      // Bounce
    Stdout()
```
## How It Works
1. **User Control**: Only appears when you explicitly call `.Spinner(msg)` or `.Spinner(msg, type)`
2. **Optional Type**: Spinner type defaults to 9 (simple / - \) if not specified
3. **Non-blocking**: The spinner runs on stderr independently
4. **Data Pass-through**: All data is passed through unchanged
5. **Automatic Cleanup**: Stops automatically when that pipe stage completes
6. **Pipeline Agnostic**: Works with any source (files, commands, data, etc.)
## The Spinner Method
```go
// Spinner wraps the pipe with a loading spinner
// msg: The message to display next to the spinner
// spinnerType: (Optional) The spinner type number (0-90). Default: 9
// Returns: The pipe for chaining
func (p *Pipe) Spinner(msg string, spinnerType ...int) *Pipe
```
## Examples
### Example 1: File Reading with Bar Spinner
```go
script.File("large_file.txt").
    Spinner("Loading file...", 1).  // Type 1: bar
    String()
```
### Example 2: Command with Arrow Spinner
```go
script.Exec("find / -name pattern").
    Spinner("Searching...", 0).     // Type 0: arrows
    Match("important").
    Stdout()
```
### Example 3: Multi-stage Pipeline with Different Spinners
```go
script.ListFiles("*.txt").
    Spinner("Finding files...", 0).
    ExecForEach("process {{.}}").
    Spinner("Processing...", 14).
    Stdout()
```
## Invalid Spinner Types
If you pass an invalid spinner type number (< 0 or > 90), it will automatically fall back to the default type 9:
```go
script.File("data.txt").
    Spinner("Reading...", 999).  // Invalid, falls back to type 9
    Stdout()
```
## When to Use Different Spinner Types
- **Quick, simple operations**: Use default (type 9: / - \)
- **Professional appearance**: Use type 1 (bar) or type 14 (dots)
- **Casual/fun**: Use type 25 (bounce) or type 26 (- **Casua- **ASCII only**: Use types 8, 9, 3 (work in all terminals)
- **Visual variety**: Experiment with different types for different operations
That's it! No configuration, no manual cleanup - just call `.Spinner(msg, type)` and pick your favorite animation! 🎉
