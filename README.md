# LibTest

Generic library function invoker.

## Usage
```
LIBTEST - Tool to call arbitrary dynamic library function.

Usage: ./libtest library type function [optional inputs]
Types: int_fn_string, fn_string, string_fn, int_fn
```

Arguments:
- library - dynamic library to call
- type - format of the function called (add _ suffix to append a new line)
- function - name of function to call
- inputs - input(s) to send to function

## Examples
Print a string with libc puts()
```
./libtest libc.so.6 fn_string puts "This string is printed by puts"
This string is printed by puts
```

Convert a string to an integer
```
$ ./libtest libc.so.6 int_fn_string_ atoi "007"
7
$ ./libtest libc.so.6 int_fn_string_ strlen "007"
3
$ ./libtest libc.so.6 char_fn_char_ tolower T
t
```

Call the `curl_version()` in libcurl
```
$ ./libtest /lib64/libcurl.so.4 string_fn_ curl_version
libcurl/7.29.0 NSS/3.53.1 zlib/1.2.7 libidn/1.28 libssh2/1.8.0
```


