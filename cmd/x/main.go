// Copyright 2021 The X Programming Language.
// Use of this source code is governed by a BSD 3-Clause
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/the-xlang/x/documenter"
	"github.com/the-xlang/x/parser"
	"github.com/the-xlang/x/pkg/x"
	"github.com/the-xlang/x/pkg/x/xset"
	"github.com/the-xlang/x/pkg/xio"
	"github.com/the-xlang/x/pkg/xlog"
)

func help(cmd string) {
	if cmd != "" {
		println("This module can only be used as single!")
		return
	}
	helpContent := [][]string{
		{"help", "Show help."},
		{"version", "Show version."},
		{"init", "Initialize new project here."},
		{"doc", "Documentize X source code."},
	}
	maxlen := len(helpContent[0][0])
	for _, part := range helpContent {
		length := len(part[0])
		if length > maxlen {
			maxlen = length
		}
	}
	var sb strings.Builder
	const space = 5 // Space of between command name and description.
	for _, part := range helpContent {
		sb.WriteString(part[0])
		sb.WriteString(strings.Repeat(" ", (maxlen-len(part[0]))+space))
		sb.WriteString(part[1])
		sb.WriteByte('\n')
	}
	println(sb.String()[:sb.Len()-1])
}

func version(cmd string) {
	if cmd != "" {
		println("This module can only be used as single!")
		return
	}
	println("The X Programming Language\n" + x.Version)
}

func initProject(cmd string) {
	if cmd != "" {
		println("This module can only be used as single!")
		return
	}
	content := []byte(`{
  "cxx_out_dir": "./dist/",
  "cxx_out_name": "x.cxx",
  "out_name": "main"
}`)
	err := ioutil.WriteFile(x.SettingsFile, content, 0666)
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	println("Initialized project.")
}

func doc(cmd string) {
	cmd = strings.TrimSpace(cmd)
	paths := strings.SplitAfterN(cmd, " ", -1)
	for _, path := range paths {
		path = strings.TrimSpace(path)
		info := compile(path, false, true)
		if info == nil {
			continue
		}
		if printlogs(info.Logs) {
			fmt.Println(path+":",
				"documentation couldn't generated because X source code has an errors")
			continue
		}
		docjson, err := documenter.Documentize(info.Parser)
		if err != nil {
			fmt.Println("Error:", err.Error())
			continue
		}
		path = filepath.Join(x.XSet.CxxOutDir, path+x.DocExtension)
		writeOutput(path, docjson)
	}
}

func processCommand(namespace, cmd string) bool {
	switch namespace {
	case "help":
		help(cmd)
	case "version":
		version(cmd)
	case "init":
		initProject(cmd)
	case "doc":
		doc(cmd)
	default:
		return false
	}
	return true
}

func init() {
	x.ExecutablePath = filepath.Dir(os.Args[0])
	x.StdlibPath = filepath.Join(x.ExecutablePath, x.StdlibName)

	// Not started with arguments.
	// Here is "2" but "os.Args" always have one element for store working directory.
	if len(os.Args) < 2 {
		os.Exit(0)
	}
	var sb strings.Builder
	for _, arg := range os.Args[1:] {
		sb.WriteString(" " + arg)
	}
	os.Args[0] = sb.String()[1:]
	arg := os.Args[0]
	index := strings.Index(arg, " ")
	if index == -1 {
		index = len(arg)
	}
	loadXSet()
	if processCommand(arg[:index], arg[index:]) {
		os.Exit(0)
	}
}

func loadXSet() {
	// File check.
	info, err := os.Stat(x.SettingsFile)
	if err != nil || info.IsDir() {
		println(`X settings file ("` + x.SettingsFile + `") is not found!`)
		os.Exit(0)
	}
	bytes, err := os.ReadFile(x.SettingsFile)
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	x.XSet, err = xset.Load(bytes)
	if err != nil {
		println("X settings has errors;")
		println(err.Error())
		os.Exit(0)
	}
}

// printlogs prints logs and returns true
// if logs has error, false if not.
func printlogs(logs []xlog.CompilerLog) bool {
	var str strings.Builder
	err := false
	for _, log := range logs {
		switch log.Type {
		case xlog.FlatError:
			err = true
			str.WriteString("ERROR: ")
			str.WriteString(log.Message)
		case xlog.FlatWarning:
			str.WriteString("WARNING: ")
			str.WriteString(log.Message)
		case xlog.Error:
			err = true
			str.WriteString("ERROR: ")
			str.WriteString(log.Path)
			str.WriteByte(':')
			str.WriteString(fmt.Sprint(log.Row))
			str.WriteByte(':')
			str.WriteString(fmt.Sprint(log.Column))
			str.WriteByte(' ')
			str.WriteString(log.Message)
		case xlog.Warning:
			str.WriteString("WARNING: ")
			str.WriteString(log.Path)
			str.WriteByte(':')
			str.WriteString(fmt.Sprint(log.Row))
			str.WriteByte(':')
			str.WriteString(fmt.Sprint(log.Column))
			str.WriteByte(' ')
			str.WriteString(log.Message)
		}
		str.WriteByte('\n')
	}
	print(str.String())
	return err
}

func appendStandard(code *string) {
	year, month, day := time.Now().Date()
	hour, min, _ := time.Now().Clock()
	timeString := fmt.Sprintf("%d/%d/%d %d.%d (DD/MM/YYYY) (HH.MM)",
		day, month, year, hour, min)
	*code = `// Auto generated by X compiler.
// X compiler version: ` + x.Version + `
// Date:               ` + timeString + `

// region X_STANDARD_IMPORTS
#include <iostream>
#include <string>
#include <functional>
#include <vector>
#include <codecvt>
#include <locale>
#include <type_traits>
// endregion X_STANDARD_IMPORTS

// region X_CXX_API
// region X_BUILTIN_VALUES
#define nil nullptr
// endregion X_BUILTIN_VALUES

// region X_MISC
class exception: public std::exception {
private:
  std::basic_string<char> _buffer;
public:
  exception(const char *_Str)      { this->_buffer = _Str; }
  const char *what() const throw() { return this->_buffer.c_str(); }
};

#define XALLOC(_Alloc) new(std::nothrow) _Alloc
#define XTHROW(_Msg) throw exception(_Msg)

template <typename _Enum_t, typename _Index_t, typename _Item_t>
static inline void foreach(const _Enum_t _Enum,
                           const std::function<void(_Index_t, _Item_t)> _Body) {
  _Index_t _index{0};
  for (auto _item: _Enum) { _Body(_index++, _item); }
}

template <typename _Enum_t, typename _Index_t>
static inline void foreach(const _Enum_t _Enum,
                           const std::function<void(_Index_t)> _Body) {
  _Index_t _index{0};
  for (auto begin = _Enum.begin(), end = _Enum.end(); begin < end; ++begin)
  { _Body(_index++); }
}
// endregion X_MISC

// region X_BUILTIN_TYPES
typedef int8_t   i8;
typedef int16_t  i16;
typedef int32_t  i32;
typedef int64_t  i64;
typedef ssize_t  ssize;
typedef uint8_t  u8;
typedef uint16_t u16;
typedef uint32_t u32;
typedef uint64_t u64;
typedef size_t   size;
typedef float    f32;
typedef double   f64;
typedef wchar_t  rune;

class str: public std::basic_string<rune> {
public:
// region CONSTRUCTOR
  str(void): str(L"")                                        { }
  str(const rune* _Str)                                      { this->assign(_Str); }
  str(const std::basic_string<rune> _Src): str(_Src.c_str()) { }
// endregion CONSTRUCTOR
};
// endregion X_BUILTIN_TYPES

// region X_STRUCTURES
template<typename _Item_t>
class array {
public:
// region FIELDS
  std::vector<_Item_t> _buffer;
// endregion FIELDS

// region CONSTRUCTORS
  array<_Item_t>(void) noexcept                                                     { this->_buffer = { }; }
  array<_Item_t>(const std::vector<_Item_t>& _Src) noexcept                         { this->_buffer = _Src; }
  array<_Item_t>(const std::nullptr_t) noexcept: array<_Item_t>()                   { }
  array<_Item_t>(const array<_Item_t>& _Src) noexcept: array<_Item_t>(_Src._buffer) { }

  array<_Item_t>(const str _Str) {
    if (std::is_same<_Item_t, rune>::value) {
      this->_buffer = std::vector<_Item_t>(_Str.begin(), _Str.end());
      return;
    }
    if (std::is_same<_Item_t, u8>::value) {
      std::wstring_convert<std::codecvt_utf8_utf16<rune>> _conv;
      std::string _bytes = _conv.to_bytes(_Str);
      this->_buffer = std::vector<_Item_t>(_bytes.begin(), _bytes.end());
      return;
    }
  }
// endregion CONSTRUCTORS

// region DESTRUCTOR
  ~array<_Item_t>(void) noexcept { this->_buffer.clear(); }
// endregion DESTRUCTOR

// region FOREACH_SUPPORT
  typedef _Item_t       *iterator;
  typedef const _Item_t *const_iterator;
  iterator begin(void) noexcept             { return &this->_buffer[0]; }
  const_iterator begin(void) const noexcept { return &this->_buffer[0]; }
  iterator end(void) noexcept               { return &this->_buffer[this->_buffer.size()]; }
  const_iterator end(void) const noexcept   { return &this->_buffer[this->_buffer.size()]; }
// endregion FOREACH_SUPPORT

// region OPERATOR_OVERFLOWS
  operator str(void) const noexcept {
    if (std::is_same<_Item_t, rune>::value) { return str(std::basic_string<rune>(this->begin(), this->end())); }
    if (std::is_same<_Item_t, u8>::value) {
      std::wstring_convert<std::codecvt_utf8_utf16<rune>> _conv;
      const std::string _bytes(this->begin(), this->end());
      return str(_conv.from_bytes(_bytes));
    }
  }

  bool operator==(const array<_Item_t> &_Src) const noexcept {
    const size _length = this->_buffer.size();
    const size _Src_length = _Src._buffer.size();
    if (_length != _Src_length) { return false; }
    for (size _index = 0; _index < _length; ++_index)
    { if (this->_buffer[_index] != _Src._buffer[_index]) { return false; } }
    return true;
  }

  bool operator==(const std::nullptr_t) const noexcept       { return this->_buffer.empty(); }
  bool operator!=(const array<_Item_t> &_Src) const noexcept { return !(*this == _Src); }
  bool operator!=(const std::nullptr_t) const noexcept       { return !this->_buffer.empty(); }
  _Item_t& operator[](const size _Index)                     { return this->_buffer[_Index]; }

  friend std::wostream& operator<<(std::wostream &_Stream,
                                   const array<_Item_t> &_Src) {
    _Stream << L"[";
    const size _length = _Src._buffer.size();
    for (size _index = 0; _index < _length;) {
      _Stream << _Src._buffer[_index++];
      if (_index < _length) { _Stream << L", "; }
    }
    _Stream << L"]";
    return _Stream;
  }
// endregion OPERATOR_OVERFLOWS
};
// endregion X_STRUCTURES

// region X_BUILTIN_FUNCTIONS
template <typename _Obj_t>
static inline void _out(_Obj_t _Obj) { std::wcout << _Obj; }

template <typename _Obj_t>
static inline void _outln(_Obj_t _Obj) {
  _out<_Obj_t>(_Obj);
  std::wcout << std::endl;
}
// endregion X_BUILTIN_FUNCTIONS
// endregion X_CXX_API

// region TRANSPILED_X_CODE
` + *code + `
// endregion TRANSPILED_X_CODE

// region X_ENTRY_POINT
int main() {
// region X_ENTRY_POINT_STANDARD_CODES
  std::setlocale(LC_ALL, "");
// endregion X_ENTRY_POINT_STANDARD_CODES
  _main();

// region X_ENTRY_POINT_END_STANDARD_CODES
  return EXIT_SUCCESS;
// endregion X_ENTRY_POINT_END_STANDARD_CODES
}
// endregion X_ENTRY_POINT`
}

func writeOutput(path, content string) {
	err := os.MkdirAll(x.XSet.CxxOutDir, 0777)
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
	bytes := []byte(content)
	err = ioutil.WriteFile(path, bytes, 0666)
	if err != nil {
		println(err.Error())
		os.Exit(0)
	}
}

func compile(path string, main, justDefs bool) *parser.ParseInfo {
	info := new(parser.ParseInfo)

	// Check standard library.
	inf, err := os.Stat(x.StdlibPath)
	if err != nil || !inf.IsDir() {
		info.Logs = append(info.Logs, xlog.CompilerLog{
			Type:    xlog.FlatError,
			Message: "standard library directory not found",
		})
		return info
	}

	f, err := xio.Openfx(path)
	if err != nil {
		println(err.Error())
		return nil
	}
	routines := new(sync.WaitGroup)
	info.File = f
	info.Routines = routines
	routines.Add(1)
	go info.ParseAsync(main, justDefs)
	routines.Wait()
	return info
}

func main() {
	filePath := os.Args[0]
	info := compile(filePath, true, false)
	if printlogs(info.Logs) {
		os.Exit(0)
	}
	cxx := info.Parser.Cxx()
	appendStandard(&cxx)
	path := filepath.Join(x.XSet.CxxOutDir, x.XSet.CxxOutName)
	writeOutput(path, cxx)
}
