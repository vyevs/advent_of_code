package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	root := readDirStructure()
	root.calculateTotalSize()
	
	part1(root)
	part2(root)
}


func part1(root dir) {
	sz := sumSmallDirs(&root)
	
	fmt.Printf("total size of all small dirs: %d\n", sz)
}

func sumSmallDirs(d *dir) int {
	var sz int
	if d.totalSize <= 100000 {
		sz += d.totalSize
	}
	
	for _, subD := range d.dirs {
		sz += sumSmallDirs(subD)
	}
	
	return sz
}

func part2(root dir) {
	spaceLeft := 70000000 - root.totalSize
	needSpace := 30000000 - spaceLeft
	
	d := findSmallestToDelete(&root, needSpace, &root)
	fmt.Printf("the smallest directory to delete is %s of size %d\n", d.name, d.totalSize)
}

func findSmallestToDelete(cur *dir, need int, best *dir) *dir {
	if cur.totalSize < best.totalSize && cur.totalSize >= need {
		best = cur
	}
	
	for _, subD := range cur.dirs {
		canD := findSmallestToDelete(subD, need, best)
		
		if canD.totalSize < best.totalSize {
			best = canD
		}
	}
	
	return best
}

type file struct {
	name string
	size int
}

type dir struct {
	name string
	files []file
	dirs []*dir
	parent *dir
	
	totalSize int
}

func readDirStructure() dir {
	plainLines := readLines()
	plainLines = plainLines[1:] // to discard the initial $ cd /
	
	lines := parseLines(plainLines)
	
	return buildDirStructure(lines)
}

func buildDirStructure(lines []line) dir { 
	rootDir := dir {
		name: "/",
	}
	curDir := &rootDir
	_ = curDir
	
	for _, l := range lines {
		curDir = modifyDir(l, curDir)
	}
	
	return rootDir
}

func modifyDir(l line, d *dir) *dir {
	if l.typ == lineTypeCD {
		target := d.findDir(l.dirName)
		if target == nil {
			panic(fmt.Sprintf("dir %s not found in parent dir %+v", l.dirName, d))
		}
		return target
		
	} else if l.typ == lineTypeLS {
		// intentionally do nothing.
		return d
		
	} else if l.typ == lineTypeFileListing {
		d.newFile(l.fileName, l.fileSize)
		return d
		
	} else if l.typ == lineTypeDirListing {
		_ = d.newDir(l.dirName)
		return d
	}
	
	panic(fmt.Sprintf("unrecognized command type for line %+v", l))
}

func (d *dir) findDir(name string) *dir {
	if name == ".." {
		return d.parent
	}
	for _, subD := range d.dirs {
		if subD.name == name {
			return subD
		}
	}
	return nil
}

func (d *dir) newDir(name string) *dir {
	newDir := dir {
		name: name,
		parent: d,
	}
	d.dirs = append(d.dirs, &newDir)
	return &newDir
}

func (d *dir) newFile(name string, size int) {
	f := file {
		name: name,
		size: size,
	}
	d.files = append(d.files, f)
}

type lineType int
const (
	lineTypeCD lineType = iota
	lineTypeLS
	lineTypeFileListing
	lineTypeDirListing
)

func (t lineType) String() string {
	switch t {
		case lineTypeCD: return "cd"
		case lineTypeLS: return "ls"
		case lineTypeFileListing: return "file listing"
		case lineTypeDirListing: return "dir listing"
	}
	panic(fmt.Sprintf("unknown lineType %d", t))
}


type line struct {
	typ lineType
		
	fileName string // for a file listing only
	fileSize int // for a file listing only
	
	dirName string // for cd and dir listing
}

func (l line) String() string {
	switch l.typ {
		case lineTypeCD: return fmt.Sprintf("%s %s", l.typ, l.dirName)
		case lineTypeLS: return l.typ.String()
		case lineTypeFileListing: return fmt.Sprintf("%s %s %d", l.typ, l.fileName, l.fileSize)
		case lineTypeDirListing: return fmt.Sprintf("%s %s", l.typ, l.dirName)
	}
	panic(fmt.Sprintf("unknown lineType %d", l.typ))
}

func parseLines(plainLines []string) []line {
	lines := make([]line, 0, 1024)
	for _, pl := range plainLines {
		lines = append(lines, parseLine(pl))
	}
	return lines
}
	
func parseLine(pl string) line {
	var l line
	
	if pl[0] == '$' {
		cmd := pl[2:4]
		if cmd == "cd" {
			l.typ = lineTypeCD
			l.dirName = pl[5:]
			
		} else if cmd == "ls" {
			l.typ = lineTypeLS
		}
	} else if pl[:3] == "dir" {
		l.typ = lineTypeDirListing
		l.dirName = pl[4:]
	 
	} else {
		l.typ = lineTypeFileListing
		
		pts := strings.Split(pl, " ")
		
		
		l.fileSize, _ = strconv.Atoi(pts[0])
		l.fileName = pts[1]
		
	}
	
	return l
	
}
	

func readLines() []string {
	lines := make([]string, 0, 1024)
	s := bufio.NewScanner(os.Stdin)
	
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	return lines
}

func (d dir) String() string {
	var buf strings.Builder
	return d.buildString(&buf, "")
}

func (d dir) buildString(buf *strings.Builder, prefix string) string {
	dirLine := fmt.Sprintf("%s- %s (dir)\n", prefix, d.name)
	buf.WriteString(dirLine)
	
	prefix += "  "
	
	for _, f := range d.files {
		fileLine := fmt.Sprintf("%s- %s (file, size=%d)\n", prefix, f.name, f.size)
		buf.WriteString(fileLine)
	}
	
	for _, d := range d.dirs {
		d.buildString(buf, prefix)
	}
	
	return buf.String()
}


func (d *dir) calculateTotalSize() {
	var sz int
	
	for _, f := range d.files {
		sz += f.size
	}
	
	for _, subD := range d.dirs {
		subD.calculateTotalSize()
		sz += subD.totalSize
	}
	
	d.totalSize = sz
}