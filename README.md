# LibTest

Generic library function invoker.

## Usage
```
LIBTEST - Tool to call arbitrary dynamic library function.

Usage: ./libtest64 library type function [optional inputs]
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
./libtest64 libc.so.6 fn_string puts "This string is printed by puts"
This string is printed by puts
```

Convert a string to an integer
```
$ ./libtest64 libc.so.6 int_fn_string_ atoi "007"
7
$ ./libtest64 libc.so.6 int_fn_string_ strlen "007"
3
$ ./libtest64 libc.so.6 char_fn_char_ tolower T
t
```

Call the `curl_version()` in libcurl
```
$ ./libtest64 /lib64/libcurl.so.4 string_fn_ curl_version
libcurl/7.29.0 NSS/3.53.1 zlib/1.2.7 libidn/1.28 libssh2/1.8.0
```


How to get versions of libcurl in McAfee installations:
```
LD_LIBRARY_PATH=/opt/McAfee/runtime/2.0/lib:/opt/McAfee/ens/esp/lib libtest64 /opt/McAfee/ens/esp/lib/libcurl.so string_fn_ curl_version

libcurl/8.4.0 OpenSSL/1.0.2zi-fips
```

```
LD_LIBRARY_PATH=/opt/McAfee/ens/runtime/3.0/lib:/opt/McAfee/agent/lib/tools libtest32 /opt/McAfee/agent/lib/tools/libcurl.so string_fn_ curl_version

libcurl/7.80.0 OpenSSL/1.0.2zg-fips zlib/1.2.13
```
