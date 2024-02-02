package main

import (
	"bufio"
	"container/list"
	"container/ring"
	"encoding/base64"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"time"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

// Package Error-1
func GetById(id string) error {
	if id == "" {
		return ValidationError
	} else if id != "Rachi" {
		return NotFoundError
	} else {
		fmt.Println("Berhasil")
	}
	// sukses
	return nil
}

// Package Error-2
func HandleError() {
	err := GetById("Rachi")
	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}
}

// Package OS
func Os() {
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("Error", err.Error())
	}
}

// Package Flag
func BenderaFlag() {
	var username *string = flag.String("username", "root", "database username")
	var password *string = flag.String("password", "root", "database password")
	var host *string = flag.String("host", "localhost", "database host")
	var port *int = flag.Int("port", 0, "database port")

	flag.Parse()

	fmt.Println("Username", *username)
	fmt.Println("Password", *password)
	fmt.Println("Host", *host)
	fmt.Println("Port", *port)
}

// Package Strings
func Setring() {
	fmt.Println(strings.Contains("Rizky Ziaul Haq", "Zia"))
	fmt.Println(strings.Split("Rizky Ziaul Haq", " "))
	fmt.Println(strings.ToLower("Rizky Ziaul Haq"))
	fmt.Println(strings.ToUpper("Rizky Ziaul Haq"))
	fmt.Println(strings.Trim("       Rizky Ziaul Haq        ", " "))
	fmt.Println(strings.ReplaceAll("Zia Zia Zia Zia", "Zia", "Rachi"))
}

// Package Strconv (String Conversion)
func StringBool() {
	result, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println("Error", err.Error())
	}

	resultInt, err := strconv.Atoi("1000")

	if err == nil {
		fmt.Println(resultInt)
	} else {
		fmt.Println("Error", err.Error())
	}

	octal := strconv.FormatInt(999, 8)
	fmt.Println(octal)

	var StringINT string = strconv.Itoa(999)
	fmt.Println(StringINT)
}

// Package Time
func waktu() {
	formatter := "2006-01-02 15:04:05"

	value := "2024-01-12 17:16:29"

	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println("Error", err.Error())
	} else {
		fmt.Println(valueTime)
	}
}

// Package double linked list
func doubleList() {
	data := list.New()

	data.PushBack("Rizky")
	data.PushBack("Ziaul")
	data.PushBack("Haq")

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// Package Data Circular List
func circularList() {
	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}

// Package Sort
type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (s UserSlice) Len() int {
	return len(s)
}

func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func (s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func urutanSort() {
	users := []User{
		{"Rizky", 21},
		{"Ziaul", 25},
		{"Haq", 31},
		{"Rachi", 11},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}

// Package regexp (regular expression)
func PkgRegex() {
	var regex *regexp.Regexp = regexp.MustCompile(`r([a-z])y`)

	fmt.Println(regex.MatchString("roy"))
	fmt.Println(regex.MatchString("ray"))
	fmt.Println(regex.MatchString("rIy"))

	fmt.Println(regex.FindAllString("ray riy ruy rIy rey", 10))
}

// Package encoding populer (base64, csv, json)
// base64
func encoDecodeBase64() {
	value := "Rizky Ziaul Haq"

	// encode base64
	var encoded = base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println(encoded)

	// decode base64
	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
	}
}

// csv reader
func PkgCsvReader() {
	csvString := "rizky,ziaul,haq\n" +
		"rachi,ern,yarouwa\n" +
		"lenovo,acer,asus"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}

// csv writer
func PkgCsvWriter() {
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"rizky", "ziaul", "haq"})
	_ = writer.Write([]string{"rachi", "ern", "yarouwa"})
	_ = writer.Write([]string{"lenovo", "acer", "asus"})

	writer.Flush()
}

// Packgae Slices
func PkgSlice() {
	names := []string{"Rizky", "Zia", "Ul", "Haq"}
	values := []int{21, 11, 31, 41, 51, 61, 71, 81, 91}

	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
	fmt.Println(slices.Contains(names, "Ul"))
	fmt.Println(slices.Index(names, "Zia"))
}

// Package Path
func PkgPath() {
	fmt.Println(path.Dir("hello/world.go"))
	fmt.Println(path.Base("hello/world.go"))
	fmt.Println(path.Ext("hello/world.go"))
	fmt.Println(path.IsAbs("hello/world.go"))
	fmt.Println(path.Join("development", "go", "golang-dasar", "main.go"))
}

// Package IO
// Reader
func BufioReader() {
	input := strings.NewReader("this is long string\nwith new line\n")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}

// Writer
func BufioWriter() {
	writer := bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("hello world\n")
	_, _ = writer.WriteString("semangat terus!")
	writer.Flush()
}

// File Manipulation

// Membuat file baru
func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)

	return nil
}

// Membaca file
func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string

	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

// Menambahkan isi file
func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)

	return nil
}
