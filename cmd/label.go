package cmd

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/fatih/color"
	"github.com/kangaechu/rssfilter/rssfilter"
	"github.com/spf13/cobra"
	"golang.org/x/term"
)

var labelRSSJSON string
var labelAll bool

// labelCmd represents the label command
var labelCmd = &cobra.Command{
	Use:   "label",
	Short: "Interactively label entries without reputation",
	Long:  `Interactively label entries without reputation by displaying title, description, and link, and allowing users to assign Good (1) or Bad (0) ratings.`,
	Run: func(_ *cobra.Command, _ []string) {
		// Load JSON
		storageJSON := rssfilter.StorageJSON{FileName: labelRSSJSON}
		rss, err := storageJSON.Load()
		if err != nil {
			log.Fatal("failed to load RSS JSON. json:", labelRSSJSON, ", err: ", err)
		}

		// Filter entries based on --all flag
		var targetEntries []*rssfilter.RSSEntry
		if rss.Entries != nil {
			for i := range *rss.Entries {
				entry := &(*rss.Entries)[i]
				if labelAll {
					// Include all entries when --all is specified
					targetEntries = append(targetEntries, entry)
				} else if entry.LabeledBy != "manual" {
					// Include all entries except manually labeled ones by default
					targetEntries = append(targetEntries, entry)
				}
			}
		}

		if len(targetEntries) == 0 {
			if labelAll {
				fmt.Println("No entries found.")
			} else {
				fmt.Println("No entries to label found.")
			}
			return
		}

		if labelAll {
			fmt.Printf("Found %d entries.\n\n", len(targetEntries))
		} else {
			fmt.Printf("Found %d entries to label.\n\n", len(targetEntries))
		}

		// Process entries in reverse order (newest first)
		for i := len(targetEntries) - 1; i >= 0; i-- {
			entry := targetEntries[i]
			index := len(targetEntries) - i

			// Clear screen before displaying entry
			clearScreen()

			// Display entry with color formatting
			displayEntry(entry, index, len(targetEntries))

			// Prompt for input
			for {
				fmt.Print("Enter: 1 (Good) / 0 (Bad) / b (open in Browser) / q (quit and save): ")

				char, err := readSingleChar()
				if err != nil {
					fmt.Printf("\nError reading input: %v\n", err)
					goto save
				}

				// Echo the character and newline
				fmt.Printf("%c\n", char)

				switch char {
				case '1':
					entry.Reputation = "Good"
					entry.LabeledBy = "manual"
					fmt.Println("✓ Labeled as Good")
					fmt.Println()
					goto nextEntry
				case '0':
					entry.Reputation = "Bad"
					entry.LabeledBy = "manual"
					fmt.Println("✓ Labeled as Bad")
					fmt.Println()
					goto nextEntry
				case 'b', 'B':
					err := openBrowser(entry.Link)
					if err != nil {
						fmt.Printf("Failed to open browser: %v\n", err)
					} else {
						fmt.Println("Opened in browser.")
					}
					// Continue the loop to ask for label again
				case 'q', 'Q':
					fmt.Println("\nSaving and exiting...")
					goto save
				default:
					fmt.Println("Invalid input. Please enter 1, 0, b, or q.")
				}
			}
		nextEntry:
		}

	save:
		// Save JSON
		err = storageJSON.StoreUnique(rss)
		if err != nil {
			log.Fatal("failed to store JSON. json: ", labelRSSJSON, ", err: ", err)
		}
		fmt.Println("Successfully saved labels to", labelRSSJSON)
	},
}

// clearScreen clears the terminal screen
func clearScreen() {
	// ANSI escape code to clear screen and move cursor to top-left
	fmt.Print("\033[H\033[2J")
}

// displayEntry displays an entry with color formatting using fatih/color
func displayEntry(entry *rssfilter.RSSEntry, index, total int) {
	// Color functions
	cyan := color.New(color.FgCyan).SprintFunc()
	blue := color.New(color.FgBlue).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()
	bold := color.New(color.Bold).SprintFunc()

	// Header with progress
	fmt.Printf("=== %s ===\n", yellow(fmt.Sprintf("Entry %d of %d", index, total)))

	// Title in cyan
	fmt.Printf("%s %s\n", bold("Title:"), cyan(entry.Title))

	// Description in default color
	fmt.Printf("%s %s\n", bold("Description:"), entry.Description)

	// Link in blue
	fmt.Printf("%s %s\n\n", bold("Link:"), blue(entry.Link))

	// Show current reputation if it exists
	if entry.Reputation != "" {
		labeledByStr := ""
		if entry.LabeledBy != "" {
			labeledByStr = fmt.Sprintf(" (%s)", entry.LabeledBy)
		}

		reputationColor := green
		if entry.Reputation == "Bad" {
			reputationColor = red
		}

		fmt.Printf("%s %s%s\n\n",
			bold("Current reputation:"),
			reputationColor(entry.Reputation),
			labeledByStr,
		)
	}
}

// openBrowser opens the specified URL in the default browser
func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return cmd.Start()
}

// readSingleChar reads a single character from stdin without requiring Enter
func readSingleChar() (byte, error) {
	// Get the file descriptor for stdin
	fd := int(os.Stdin.Fd())

	// Check if stdin is a terminal
	if !term.IsTerminal(fd) {
		// Not a terminal (e.g., piped input), read normally
		b := make([]byte, 1)
		_, err := os.Stdin.Read(b)
		return b[0], err
	}

	// Save the current terminal state
	oldState, err := term.MakeRaw(fd)
	if err != nil {
		return 0, err
	}
	defer func() {
		_ = term.Restore(fd, oldState)
	}()

	// Read a single byte
	b := make([]byte, 1)
	_, err = os.Stdin.Read(b)
	if err != nil {
		return 0, err
	}

	return b[0], nil
}

func init() {
	rootCmd.AddCommand(labelCmd)
	labelCmd.Flags().StringVarP(&labelRSSJSON, "feed", "f", "", "feed JSON file name")
	err := labelCmd.MarkFlagRequired("feed")
	if err != nil {
		log.Fatal("specify feed JSON file name. err: ", err)
	}
	labelCmd.Flags().BoolVarP(&labelAll, "all", "a", false, "label all entries including those with existing reputation")
}
