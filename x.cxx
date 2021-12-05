// Auto generated by X compiler.
// X compiler version: @developer_beta 0.0.1
// Date:               5/12/2021 14.42 (DD/MM/YYYY) (HH.MM)

#pragma region X_STANDARD_IMPORTS
#include <iostream>
#include <string>
#include <functional>
#include <vector>
#include <locale.h>
#pragma endregion X_STANDARD_IMPORTS

#pragma region X_RUNTIME_FUNCTIONS
inline void throw_exception(const std::wstring message) {
  std::wcout << message << std::endl;
  exit(1);
}
#pragma endregion X_RUNTIME_FUNCTIONS

#pragma region X_BUILTIN_TYPES
typedef int8_t int8;
typedef int16_t int16;
typedef int32_t int32;
typedef int64_t int64;
typedef uint8_t uint8;
typedef uint16_t uint16;
typedef uint32_t uint32;
typedef uint64_t uint64;
typedef float float32;
typedef double float64;
typedef wchar_t rune;

#define function std::function

class str {
public:
#pragma region FIELDS
  std::wstring string;
#pragma endregion FIELDS

#pragma region CONSTRUCTORS
  str(const std::wstring& string) {
    this->string = string;
  }

  str(const rune* string) {
    this->string = string;
  }
#pragma endregion CONSTRUCTORS

#pragma region DESTRUCTOR
  ~str() {
    this->string.clear();
  }
#pragma endregion DESTRUCTOR

#pragma region OPERATOR_OVERFLOWS
  bool operator==(const str& string) {
    return this->string == string.string;
  }

  bool operator!=(const str& string) {
    return !(this->string == string.string);
  }

  str operator+(const str& string) {
    return str(this->string + string.string);
  }

  void operator+=(const str& string) {
    this->string += string.string;
  }

  rune& operator[](const int index) {
    const uint32 length = this->string.length();
    if (index < 0) {
      throw_exception(L"stackoverflow exception:\n index is less than zero");
    } else if (index >= length) {
      throw_exception(L"stackoverflow exception:\nindex overflow " +
        std::to_wstring(index) + L":" + std::to_wstring(length));
    }
    return this->string[index];
  }

  friend std::wostream& operator<<(std::wostream &os, const str& string) {
    os << string.string;
    return os;
  }
#pragma endregion OPERATOR_OVERFLOWS
};
#pragma endregion X_BUILTIN_TYPES

#pragma region X_BUILTIN_VALUES
#define null nullptr
#pragma endregion X_BUILTIN_VALUES

#pragma region X_STRUCTURES
template<typename T>
class array {
public:
#pragma region FIELDS
  std::vector<T> vector;
#pragma endregion FIELDS

#pragma region CONSTRUCTORS
  array() {
    this->vector = {};
  }

  array(std::nullptr_t ) : array() {}

  array(const std::vector<T>& vector) {
    this->vector = vector;
  }
#pragma endregion CONSTRUCTORS

#pragma region DESTRUCTOR
  ~array() {
    this->vector.clear();
  }
#pragma endregion DESTRUCTOR

#pragma region OPERATOR_OVERFLOWS
  bool operator==(const array& array) {
    const uint32 vector_length = this->vector.size();
    const uint32 array_vector_length = array.vector.size();
    if (vector_length != array_vector_length) {
      return false;
    }
    for (int index = 0; index < vector_length; ++index) {
      if (this->vector[index] != array.vector[index]) {
        return false;
      }
    }
    return true;
  }

  bool operator==(std::nullptr_t) {
    return this->vector.empty();
  }

  bool operator!=(const array& array) {
    return !(*this == array);
  }

  bool operator!=(std::nullptr_t) {
    return !this->vector.empty();
  }

  T& operator[](const int index) {
    const uint32 length = this->vector.size();
    if (index < 0) {
      throw_exception(L"stackoverflow exception:\n index is less than zero");
    } else if (index >= length) {
      throw_exception(L"stackoverflow exception:\nindex overflow " +
        std::to_wstring(index) + L":" + std::to_wstring(length));
    }
    return this->vector[index];
  }

  friend std::wostream& operator<<(std::wostream &os, const array<T>& array) {
    os << L"[";
    const uint32 size = array.vector.size();
    for (int index = 0; index < size;) {
      os << array.vector[index++];
      if (index < size) {
        os << L", ";
      }
    }
    os << L"]";
    return os;
  }
#pragma endregion OPERATOR_OVERFLOWS
};
#pragma endregion X_STRUCTURES

#pragma region X_BUILTIN_FUNCTIONS
template<typename any>
inline void _out(any v) {
  std::wcout << v;
}

template<typename any>
inline void _outln(any v) {
  _out(v);
  std::wcout << std::endl;
}
#pragma endregion X_BUILTIN_FUNCTIONS

#pragma region TRANSPILED_X_CODE
#pragma region TYPES
#pragma endregion TYPES

#pragma region PROTOTYPES
void _main();
#pragma endregion PROTOTYPES

#pragma region GLOBAL_VARIABLES
#pragma endregion GLOBAL_VARIABLES

#pragma region GLOBAL_VARIABLES
#pragma endregion GLOBAL_VARIABLES

#pragma region FUNCTIONS

void _main(){
  int32 _a = 5;
  _a <<= 10;
  _outln(_a);
}

#pragma endregion FUNCTIONS
#pragma endregion TRANSPILED_X_CODE

#pragma region X_ENTRY_POINT
int main() {
#pragma region X_ENTRY_POINT_STANDARD_CODES
  setlocale(0x0, "");
#pragma endregion X_ENTRY_POINT_STANDARD_CODES

  _main();

#pragma region X_ENTRY_POINT_END_STANDARD_CODES
  return EXIT_SUCCESS;
#pragma endregion X_ENTRY_POINT_END_STANDARD_CODES
}
#pragma endregion X_ENTRY_POINT
