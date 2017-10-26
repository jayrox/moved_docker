package main

import (
	//"bufio"
	"flag"
	"fmt"
	"os"
	//	"path"
	//"io"
	//"log"
	"path/filepath"
	"strings"
	"time"
)

var (
	flagdebug      = flag.Bool("d", true, "show debug output")
	src            = "/mnt/src"
	dst            = "/mnt/dst"
	flagminsize    = flag.Int64("min", 200000000, "minimum file size to include in scan. default is 200MB") // 3MB
	flagmovesample = flag.Bool("ms", false, "move sample files")
	flagTest       = flag.Bool("test", false, "test move but don't actually move files.")
	flagNoLogo     = flag.Bool("nologo", false, "hide the logo, useful for automation logging.")
	rf             mvdFlags
	moveext        = []string{
		".mkv", ".mp4", ".avi", ".m4v", ".divx",
	}
)

func main() {
	flag.Parse()
	// Print the logo :P
	printLogo()

	rf.Min = flagInt(flagminsize)

	// Root folder to scan
	fpSAbs, _ := filepath.Abs(src)
	rf.Dir = fpSAbs

	// Root folder to move to
	fpTAbs, _ := filepath.Abs(dst)
	rf.Target = fpTAbs

	fmt.Printf("Scanning directory: %s\n", rf.Dir)
	fmt.Println("_____________________")

	i := folderWalk(rf.Dir)
	if i < 1 {
		fmt.Println("No movable files found.")
	}
}

func flagString(fs *string) string {
	return fmt.Sprint(*fs)
}

func flagInt(fi *int64) int64 {
	return int64(*fi)
}

func flagBool(fb *bool) bool {
	return bool(*fb)
}

func folderWalk(file string) (i int64) {
	i = 0
	var err = filepath.Walk(file, func(file string, _ os.FileInfo, _ error) error {
		for _, x := range moveext {
			if !flagBool(flagmovesample) && strings.Contains(strings.ToLower(file), "sample") {
				//fmt.Println("Skipping sample.")
				continue
			}
			if filepath.Ext(file) == x {
				var ok bool

				ok = moveable(file)

				if ok == true {
					ok = move(file, rf.Target)

					if ok == false {
						printDebug("Move failed %s\n", "")
					}
				}
				fmt.Println("_________")
			}
		}
		return nil
	})
	check(err)

	return
}

func moveable(file string) bool {
	fmt.Printf("Checking file size: %s\n", file)

	sizeOne := fileSize(file)
	fmt.Println(sizeOne)

	time.Sleep(500 * time.Millisecond)
	sizeTwo := fileSize(file)
	fmt.Println(sizeTwo)

	time.Sleep(600 * time.Millisecond)
	sizeThree := fileSize(file)
	fmt.Println(sizeThree)

	time.Sleep(700 * time.Millisecond)
	sizeFour := fileSize(file)
	fmt.Println(sizeFour)

	time.Sleep(1000 * time.Millisecond)
	sizeFive := fileSize(file)
	fmt.Println(sizeFive)

	if sizeOne == sizeTwo && sizeOne == sizeThree && sizeOne == sizeFour && sizeOne == sizeFive {
		return true
	}

	fmt.Println("Skipping: File sizes don't match.")
	return false
}

func fileSize(path string) int64 {
	file, err := os.Open(path)
	check(err)

	fi, err := file.Stat()
	check(err)

	file.Close()
	return fi.Size()
}

func move(file, destpath string) (ok bool) {
	ok = true

	destpath = strings.Replace(destpath, "\\", "/", -1)
	file = strings.Replace(file, "\\", "/", -1)

	basename := filepath.Base(file)

	fp := filepath.Dir(file)

	name := filepath.Base(fp)
	fmt.Println(name)

	// Prevents /Media/Movies/Movies/
	if "src" == name {
		name = ""
	}

	// Make target directory
	err := os.Mkdir(filepath.Join(destpath, name), os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}

	target := filepath.Join(destpath, name, basename)

	printDebug("Moving: %s\nDestination: %s\n", file, target)

	if *flagTest {
		printDebug("Test mode enabled. Exiting before move.%s\n", "")
		return true
	}

	err = os.Rename(file, target)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	fmt.Printf("CopyFile succeeded\n")

	//err = os.Remove(file)
	//if err != nil {
	//	fmt.Printf("Remove failed %q\n", err)
	//	return false
	//}
	fmt.Printf("Remove succeeded\n")

	return true
}

// Check err
func check(e error) bool {
	if e != nil {
		fmt.Println(e.Error())
		fmt.Println(e)
		//panic(e)
		return false
	}
	return true
}

// Only print debug output if the debug flag is true
func printDebug(format string, vars ...interface{}) {
	if *flagdebug {
		if vars[0] == nil {
			fmt.Println(format)
			return
		}
		fmt.Printf(format, vars...)
	}
}

// Hold flag data
type mvdFlags struct {
	Dir    string
	Target string
	Debug  bool
	Min    int64
}

// Print the logo, obviously
func printLogo() {
	if *flagNoLogo {
		return
	}
	fmt.Println("███╗   ███╗██╗   ██╗██████╗")
	fmt.Println("████╗ ████║██║   ██║██╔══██╗")
	fmt.Println("██╔████╔██║██║   ██║██║  ██║")
	fmt.Println("██║╚██╔╝██║╚██╗ ██╔╝██║  ██║")
	fmt.Println("██║ ╚═╝ ██║ ╚████╔╝ ██████╔╝")
	fmt.Println("╚═╝     ╚═╝  ╚═══╝  ╚═════╝ moved")
	fmt.Println("")
}
